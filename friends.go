package twigger

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"unsafe"
)

type RelationFunc func(v url.Values) (c anaconda.UserCursor, err error)

type RelationInfo struct {
	Func RelationFunc
	Type string
}

const (
	RelationFriend   = "friend"
	RelationFollower = "follower"
)

func (c *Connection) getNUsersFromValues(n int, info RelationInfo, v url.Values) (Users, error) {
	countValue := v.Get("count")
	if countValue == "" {
		v.Set("count", "200")
	}

	err := c.FillMissingUserValues(v)
	if err != nil {
		c.ErrLog.Printf("During filling missing user values an error occurred.\n Values: %v \nError: %v\n", v, err)
		c.ErrLog.Printf("Retrieval task of %vs of %v cannot be initiated. ***\n", info.Type, v.Get("screen_name"))
	}
	screenName := v.Get("screen_name")

	nextCursor := int64(-1)
	nexCursorStr := fmt.Sprintf("%v", nextCursor)
	v.Set("cursor", nexCursorStr)
	allUsers := make([]User, 0)

	c.InfoLog.Printf("*** Retrieval task of %vs of %v has started. ***\n", info.Type, screenName)
	for {
		userCursor, err := info.Func(v)
		if err != nil {
			c.ErrLog.Printf("While retrieving %vs of %v an error occurred. Error message: %v", info.Type, screenName, err)
			return allUsers, err
		}
		nextCursor = userCursor.Next_cursor
		nextCursorStr := fmt.Sprintf("%v", nextCursor)
		v.Set("cursor", nextCursorStr)
		curUsers := *(*Users)(unsafe.Pointer(&userCursor.Users))
		//c.InfoLog.Printf("Retrieving %vs[%v:%v]", info.Type, oldC )
		oldCount := len(allUsers)
		allUsers = append(allUsers, curUsers...)

		// If n is 0 or a negative number that means the functions should return all users.
		if n > 0 && len(allUsers) >= n {
			allUsers = allUsers[:n]
			break
		}

		newCount := len(allUsers)
		c.InfoLog.Printf("Retrieved %vs[%v:%v]", info.Type, oldCount, newCount)
		if nextCursor == 0 {
			break
		}
	}
	return allUsers, nil
}

// Friend Utility Methods

func (c *Connection) getAllFriendsFromValues(v url.Values) (Users, error) {
	info := RelationInfo{
		Func: c.Client.GetFriendsList,
		Type: RelationFriend,
	}
	return c.getNUsersFromValues(-1, info, v)
}

func (c *Connection) GetAllFriendsFromScreenName(screenName string) (Users, error) {
	v := url.Values{}
	v.Set("screen_name", screenName)
	return c.getAllFriendsFromValues(v)
}

func (c *Connection) GetAllFriendsFromID(userID string) (Users, error) {
	v := url.Values{}
	v.Set("user_id", userID)
	return c.getAllFriendsFromValues(v)
}

func (c *Connection) getNFriendsFromValues(n int, v url.Values) (Users, error) {
	info := RelationInfo{
		Func: c.Client.GetFriendsList,
		Type: RelationFriend,
	}
	return c.getNUsersFromValues(n, info, v)
}

func (c *Connection) GetNFriendsFromScreenName(n int, screenName string) (Users, error) {
	v := url.Values{}
	v.Set("screen_name", screenName)
	return c.getNFriendsFromValues(n, v)
}

func (c *Connection) GetNFriendsFromID(n int, userID string) (Users, error) {
	v := url.Values{}
	v.Set("user_id", userID)
	return c.getNFriendsFromValues(n, v)
}

// Follower Utility methods

func (c *Connection) getAllFollowersFromValues(v url.Values) (Users, error) {
	info := RelationInfo{
		Func: c.Client.GetFollowersList,
		Type: RelationFollower,
	}
	return c.getNUsersFromValues(-1, info, v)
}

func (c *Connection) GetAllFollowersFromScreenName(screenName string) (Users, error) {
	v := url.Values{}
	v.Set("screen_name", screenName)
	return c.getAllFollowersFromValues(v)
}

func (c *Connection) GetAllFollowersFromID(userID string) (Users, error) {
	v := url.Values{}
	v.Set("user_id", userID)
	return c.getAllFollowersFromValues(v)
}

func (c *Connection) getNFollowersFromValues(n int, v url.Values) (Users, error) {
	info := RelationInfo{
		Func: c.Client.GetFollowersList,
		Type: RelationFriend,
	}
	return c.getNUsersFromValues(n, info, v)
}

func (c *Connection) GetNFollowersFromScreenName(n int, screenName string) (Users, error) {
	v := url.Values{}
	v.Set("screen_name", screenName)
	return c.getNFollowersFromValues(n, v)
}

func (c *Connection) GetNFollowersFromID(n int, userID string) (Users, error) {
	v := url.Values{}
	v.Set("user_id", userID)
	return c.getNFollowersFromValues(n, v)
}
