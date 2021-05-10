## Twigger

* Twigger is helper library written in Go that aims to make common Twitter API calls easier. 
* Twigger is based on Twitter API library [anaconda](https://github.com/ChimeraCoder/anaconda).
* Twigger also provides easy to use CLI commands for some most common Twitter API calls.
* Neither Twigger API nor Twigger CLI commands are exhaustive for now.

## Example Usage of Twigger API

```go
creds, err := twigger.LoadJSONCredentials("test_resources/login.json")

homeDir, _ := os.UserHomeDir()
homeDir = homeDir + "/"

logFile, _ := os.OpenFile(homeDir + "out.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
defer logFile.Close()

conn,  err := twigger.NewConnection(creds, logFile, os.Stdout, os.Stderr)
if err != nil{
    log.Fatalln("Cannot create a new connection!")
}

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
```

## Installing Twigger CLI Commands

1. `go get github.com/gusanmaz/twigger`

2. `cd $GOPATH/src/github.com/gusanmaz/twigger`

3. `find ./cmd/twigger* -type d | xargs go install`

## Twigger CLI Usage Examples

* Getting all recent favorites of a user

`twigger-user-favs -credentials creds.json  -screenname nixcraft -output nixcraft_favs.json `

* Getting all recent tweets of a user

`twigger-user-tweets -credentials creds.json  -screenname nixcraft -output nixcraft_tweets.json `

* Getting all friends of a user

`twigger-user-friends -credentials creds.json  -screenname nixcraft -output nixcraft_friends.json `

* Getting 2000 followers of a user

`twigger-user-tweets -credentials creds.json  -screenname nixcraft -n 2000 -output nixcraft_followers.json `

## Author

Güvenç Usanmaz

## License

MIT License


