package twigger

import (
	"encoding/json"
	"github.com/ChimeraCoder/anaconda"
	_ "github.com/dghubble/go-twitter/twitter"
	"os"
	"strconv"
)

type Tweet anaconda.Tweet
type Tweets []Tweet
type TweetMap map[string]Tweet

func (tms1 TweetMap) Join(tms2 TweetMap) TweetMap {
	tweetsMap := map[string]Tweet{}
	for _, tweet := range tms1 {
		tweetsMap[tweet.IdStr] = tweet
	}
	for _, tweet := range tms2 {
		tweetsMap[tweet.IdStr] = tweet
	}
	return tweetsMap
}

func (tm1 Tweets) Join(tm2 Tweets) Tweets {
	tweetsMap := map[string]Tweet{}
	for _, tweet := range tm1 {
		tweetsMap[tweet.IdStr] = tweet
	}
	for _, tweet := range tm2 {
		tweetsMap[tweet.IdStr] = tweet
	}
	tweets := make([]Tweet, 0)

	ind := 0
	for _, v := range tweetsMap {
		tweets[ind] = v
		ind++
	}

	return tweets
}

func (tweets Tweets) Save(filepath string) error {
	f, err := os.Create(filepath)
	defer f.Close()
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(&tweets)
	if err != nil {
		return err
	}
	f.Write(bytes)
	return nil
}

func (tweets Tweets) ToMap() TweetMap {
	tweetMap := TweetMap{}
	for _, t := range tweets {
		IDStr := strconv.FormatInt(t.Id, 10)
		tweetMap[IDStr] = t
	}
	return tweetMap
}

func (t TweetMap) Save(filepath string) error {
	f, err := os.Create(filepath)
	defer f.Close()
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(&t)
	if err != nil {
		return err
	}
	f.Write(bytes)
	return nil
}

func (t TweetMap) Load(filepath string) error {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &t)
	if err != nil {
		return nil
	}
	return nil
}

func (t *Tweets) Load(filepath string) error {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &t)
	if err != nil {
		return nil
	}
	return nil
}

type User anaconda.User
type Users []User
type UserMap map[string]User

func (users Users) Save(filepath string) error {
	f, err := os.Create(filepath)
	defer f.Close()
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(&users)
	if err != nil {
		return err
	}
	f.Write(bytes)
	return nil
}

func (u Users) ToMap() UserMap {
	usersMap := UserMap{}
	for _, user := range u {
		IDStr := strconv.FormatInt(user.Id, 10)
		usersMap[IDStr] = user
	}
	return usersMap
}

func (u UserMap) Save(filepath string) error {
	f, err := os.Create(filepath)
	defer f.Close()
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(&u)
	if err != nil {
		return err
	}
	f.Write(bytes)
	return nil
}

func (u UserMap) Load(filepath string) error {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &u)
	if err != nil {
		return nil
	}
	return nil
}
