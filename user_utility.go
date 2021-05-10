package twigger

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *Connection)GetScreenNameFromID(id int64) (string, error){
	user, err := c.Client.GetUsersShowById(id, nil)
	if err != nil{
		return "", err
	}
	return user.ScreenName, nil
}

func (c *Connection)GetIDFromScreenName(screenName string)(int64, error){
	user, err := c.Client.GetUsersShow(screenName, nil)
	if err != nil{
		return 0, err
	}
	return user.Id, nil
}

func (c *Connection)CheckUserValueIntegrity(v url.Values) bool{
	id := v.Get("user_id")
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		return false
	}
	screenName := v.Get("screen_name")
	user1, err1 := c.Client.GetUsersShowById(idNum, nil)
	user2, err2 := c.Client.GetUsersShow(screenName, nil)
	if err1 != nil || err2 != nil{
		return false
	}
	if user1.ScreenName == user2.ScreenName{
		return true
	}
	return false
}

func (c *Connection)FillMissingUserValues(v url.Values) error{
	screenName :=  v.Get("screen_name")
	id := v.Get("user_id")
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		idNum = 0 // ParseInt already returns 0 in the case of parsing error
	}

	if screenName == "" && idNum == 0{
		screenName = c.User.ScreenName
		user, err := c.Client.GetUsersShow(screenName, nil)
		if err != nil{
			return err
		}
		idNum = user.Id
	}

	if screenName == "" && idNum !=0 {
		screenName, err = c.GetScreenNameFromID(idNum)
		if err != nil{
			return err
		}
		v.Set("screen_name", screenName)
	}

	if screenName != "" && idNum == 0 {
		idNum, err := c.GetIDFromScreenName(screenName)
		if err != nil{
			return err
		}
		v.Set("user_id", fmt.Sprintf("%v", idNum))
	}

	// If there is inconsistency b/w id and screen name disregard screen name
	if !c.CheckUserValueIntegrity(v){
		id := v.Get("user_id")
		idNum, err := strconv.ParseInt(id, 10, 64)
		screenName, err = c.GetScreenNameFromID(idNum)
		if err != nil{
			return err
		}
		v.Set("screen_name", screenName)
	}
	return nil
}


func (u User) GetEmail()string{
	return u.Email
}

func (u User) GetProfileImageURL() string{
	return u.ProfileImageURL
}

func (u User) GetBackgroundImageURL() string{
	return u.ProfileBackgroundImageURL
}

func (u User) IsVerified() bool{
	return u.Verified
}

func (u User ) GetDescription() string{
	return u.Description
}

func (u User) GetScreenName() string{
	return u.ScreenName
}

func (u User) GetID() string{
	return u.IdStr
}

func (u User) IsProtected() bool{
	return u.Protected
}

func (u User) GetStatusCount() int64{
	return u.StatusesCount
}

func (u User) GetFavoritesCount() int{
	return u.FavouritesCount
}

func (u User) GetFollowingCount() int{
	return u.FriendsCount
}

func (u User) GetFollowersCount() int{
	return u.FollowersCount
}



