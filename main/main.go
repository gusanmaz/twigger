package main

import (
	"github.com/gusanmaz/twigger"
	"log"
	"os"
)

func main(){
	creds, err := twigger.LoadJSONCredentials("test_resources/real_login.json")

	homeDir, _ := os.UserHomeDir()
	homeDir = homeDir + "/"

	logFile, _ := os.OpenFile(homeDir + "out.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer logFile.Close()

	conn,  err := twigger.NewConnection(creds, logFile, os.Stdout, os.Stderr)
	if err != nil{
		log.Fatalln("Cannot create a new connection!")
	}

	conn.PublishVideoTweet("test_resources/video.mp4", "this is a video test tweet and will be deleted in minutes.")
	return

	tweets, _ := conn.GetAllRecentTweetsFromScreenName("BBCWorld")
	tweets.Save(homeDir + "bbc_tweets.json")

	favs, _   := conn.GetAllRecentFavoritesFromScreenName("philosophybites")
	favs.Save(homeDir + "philosophybites_favs.json")

	friends, _ := conn.GetAllFriendsFromScreenName("twitterVideo")
	friends.Save(homeDir + "twitter_video_friends.json")
	followers, _ := conn.GetAllFollowersFromScreenName("randomartdaily")
	followers.Save(homeDir + "randomartdaily_followers.json")

	followers, _ = conn.GetNFollowersFromScreenName(2000, "BBCEarth")
	followers.Save(homeDir + "bbcearth_2000_followers.json")

	popTweets, _ := conn.GetPopularTweetsByQuery("Apple M1")
	popTweets.Save(homeDir + "apple_m1_pop_tweets.json")

	picturePaths := []string{
		"test_resources/antalya.jpg",
		"test_resources/eskisehir.jpg",
		"test_resources/stars.jpg",
		"test_resources/ny.jpg",
		"test_resources/providence.jpg",
	}

	conn.PublishCollageTweet(picturePaths,"This is a test tweet and will be deleted soon.")
	conn.PublishVideoTweet("test_resources/video.mp4", "this is a video test tweet and will be deleted in minutes.")

}
