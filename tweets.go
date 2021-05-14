package twigger

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"math"
	"net/url"
	"strconv"
	"unsafe"
)

const(
	EntityTweet        = "tweet"
	EntityFavorite     = "favorite"
	EntityRetweet      = "retweet"
	EntityHomeTimeLine = "home-timeline-tweet"
	EntityMention      = "mention-tweet"
	EntityAPILimit     = 5000 // It is actually 3200 and this constant is set to 5000 to be on the safe side!
)

type EntityFunc func(values url.Values) ([]anaconda.Tweet, error)

type TweetEntity struct{
	Func EntityFunc
	Entity string
}

func (c *Connection) getRecentNEntityFromValues(n int, entity TweetEntity, values url.Values) (Tweets, error){
	maxID := math.MaxInt64 - 1
	maxIDStr := fmt.Sprintf("%v", maxID)
	values.Set("max_id", maxIDStr)

	countVal := values.Get("count")
	if countVal == ""{
		values.Set("count", "200")
	}

	err := c.FillMissingUserValues(values)
	if err != nil{
		c.ErrLog.Printf("During filling missing user values an error occurred.\n Values: %v \nError: %v\n", values, err)
		c.ErrLog.Printf("Retrieval task of %vs of %v cannot be initiated. ***\n",entity.Entity, values.Get("screen_name"))
		return nil, err
	}

	allTweets := make([]Tweet, 0)
	screenName := values.Get("screen_name")
	c.InfoLog.Printf("*** Retrieval task of %vs of %v has started. ***\n",entity.Entity, screenName)

	for{
		tweets, err := entity.Func(values)
		if err != nil{
			c.ErrLog.Printf("While retrieving %vs of %v an error occurred. Error message: %v", entity.Entity, screenName, err)
			return allTweets, err
		}
		if len(tweets) == 0{
			break
		}

		maxID := tweets[len(tweets) -  1].Id - 1
		maxIDStr := fmt.Sprintf("%v", maxID)
		values.Set("max_id", maxIDStr)

		curTweets := *(*Tweets)(unsafe.Pointer(&tweets))
		oldCount := len(allTweets)
		allTweets = append(allTweets, curTweets...)
		newCount := len(allTweets)
		c.InfoLog.Printf("%vs (%v:%v) are retrieved.\n", entity.Entity, oldCount +1, newCount)

		if len(allTweets) >= n {
			allTweets = allTweets[:n]
			break
		}
		if len(curTweets) == 0{
			break
		}
	}
	c.InfoLog.Printf("%v retrieval task for %v has ended. %v most recent %vs obtained.\n",entity.Entity, screenName, len(allTweets), entity.Entity)
	return allTweets, nil
}


func (c *Connection) GetAllRecentTweetsFromScreenName(screenName string)(Tweets, error){
	return c.GetRecentNTweetsFromScreenName(screenName, EntityAPILimit)
}

func (c *Connection) GetRecentNTweetsFromScreenName(screenName string, n int)(Tweets, error){
	entity := TweetEntity{
		Func:   c.Client.GetUserTimeline,
		Entity: EntityTweet,
	}

	values := url.Values{}
	values.Add("screen_name", screenName)
	return c.getRecentNEntityFromValues(n, entity, values)
}

func (c *Connection) GetAllRecentFavoritesFromScreenName(screenName string)(Tweets, error){
	return c.GetRecentNFavoritesFromScreenName(screenName, EntityAPILimit)
}

func (c *Connection) GetRecentNFavoritesFromScreenName(screenName string, n int)(Tweets, error){
	entity := TweetEntity{
		Func:   c.Client.GetFavorites,
		Entity: EntityFavorite,
	}

	values := url.Values{}
	values.Add("screen_name", screenName)
	return c.getRecentNEntityFromValues(n, entity, values)
}

func (c *Connection) GetRecentNMentions(n int)(Tweets, error){
	entity := TweetEntity{
		Func:   c.Client.GetMentionsTimeline,
		Entity: EntityMention,
	}

	values := url.Values{}
	values.Add("screen_name", c.User.ScreenName)
	return c.getRecentNEntityFromValues(n, entity, values)
}

func (c *Connection) GetAllRecentMentions()(Tweets, error){
	return c.GetRecentNMentions(EntityAPILimit)
}

func (c *Connection) GetRecentNMentionsSince(n int, sinceID int64)(Tweets, error){
	entity := TweetEntity{
		Func:   c.Client.GetMentionsTimeline,
		Entity: EntityMention,
	}

	values := url.Values{}
	values.Set("since_id", fmt.Sprintf("%v", sinceID))
	values.Set("screen_name", c.User.ScreenName)

	return c.getRecentNEntityFromValues(n, entity, values)
}

func (c *Connection) GetAllRecentMentionsSince(sinceID int64)(Tweets, error){
	return c.GetRecentNMentionsSince(EntityAPILimit, sinceID)
}

func (c *Connection) GetRecentNHomeTimelineTweets(n int)(Tweets, error){
	entity := TweetEntity{
		Func:   c.Client.GetHomeTimeline,
		Entity: EntityHomeTimeLine,
	}

	values := url.Values{}
	return c.getRecentNEntityFromValues(n, entity, values)
}

func (c *Connection) GetAllRecentNHomeTimelineTweets()(Tweets, error){
	return c.GetRecentNHomeTimelineTweets(EntityAPILimit)
}

func (c *Connection) DeleteTweetFromID(id string) (Tweet, error){
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		c.ErrLog.Printf("While trying to delete tweet with ID:%v and error occurred becuase given tweet ID is not a proper value. Error: %v.\n", id, err)
		return Tweet{}, err
	}
	tweet, err := c.Client.DeleteTweet(idInt, false)
	if err != nil{
		c.ErrLog.Printf("While trying to delete tweet with ID:%v and error occurred. Error: %v.\n", id, err)
		return Tweet{}, err
	}
	c.InfoLog.Printf("Tweet with ID of %v has been deleted successfully.\n", id)
	return Tweet(tweet), nil
}

func (c *Connection) GetSingleTweetFromID(id int64)(Tweet, error){
	tw, err := c.Client.GetTweet(id, nil)
	if err != nil{
		c.ErrLog.Printf("Tweet with ID of %v cannot be retrieved. Error: %v\n",id, err)
	}
	c.InfoLog.Printf("Tweet with ID of %v has been retrieved successfully.\n", id)
	return Tweet(tw), nil
}




