package twigger

import (
	"net/url"
	"strconv"
	"unsafe"
)

func (c *Connection) GetXTweetsByQuery(query string, v url.Values) (Tweets, error) {
	v.Set("count", "100")
	searchResp, err := c.Client.GetSearch(query, v)
	if err != nil {

	}
	tweets := *(*Tweets)(unsafe.Pointer(&searchResp.Statuses))
	return tweets, err
}

func (c *Connection) getResultTypeTweetsByQuery(query string, resultType string) (Tweets, error) {
	v := url.Values{}
	v.Set("result_type", resultType)
	c.InfoLog.Printf("Querying %v tweets with keyword (%v) has started.\n", resultType, query)
	tweets, err := c.GetXTweetsByQuery(query, v)
	if err != nil {
		c.ErrLog.Printf("An error occurred during querying tweets.\n")
		c.ErrLog.Printf("Error message: %v", err)
	} else {
		c.InfoLog.Printf("Querying %v tweets with keyword (%v) has completed successfully.\n", resultType, query)
	}
	return tweets, err
}

func (c *Connection) GetPopularTweetsByQuery(query string) (Tweets, error) {
	return c.getResultTypeTweetsByQuery(query, "popular")
}

func (c *Connection) GetRecentTweetsByQuery(query string) (Tweets, error) {
	return c.getResultTypeTweetsByQuery(query, "recent")
}

func (c *Connection) GetMixedTweetsByQuery(query string) (Tweets, error) {
	return c.getResultTypeTweetsByQuery(query, "mixed")
}

func (c *Connection) GetTopUsersFromQuery(query string) (Users, error) {
	v := url.Values{}
	pageNum := 1
	v.Set("page", strconv.Itoa(pageNum))
	allUsers := make([]User, 0)
	for {
		anacodaUsers, err := c.Client.GetUserSearch(query, v)
		if err != nil {
			return allUsers, err
		}
		if len(anacodaUsers) == 0 {
			break
		}
		curUsers := *(*Users)(unsafe.Pointer(&anacodaUsers))
		allUsers = append(allUsers, curUsers...)
		pageNum++
		v.Set("page", strconv.Itoa(pageNum))
	}
	return allUsers, nil
}
