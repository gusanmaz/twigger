package cmd
// This file is common file for each command and has no use in cmd package for now.
import (
	"flag"
	"fmt"
	"github.com/gusanmaz/twigger"
	"os"
)

const(
	credsDefValue = "creds.json"
	credsUsage    = "File path for Twitter API credentials. File should be in JSON or YAML format"

	screenNameDefValue = ""
	screenNameUsage    = "Screen name of the twitter user"

	logFileDefPath = "/dev/null"
	logFilePathUsage = "Filepath for log file."

	favOutputDefFilePath   = "user_favs.json"
	favOutputFilePathUsage = "Filepath for output file that consists of data about user's favorite tweets"

	tweetsOutputDefFilePath   = "user_tweets.json"
	tweetsOutputFilePathUsage = "Filepath for output file that consists of data about user's tweets"

	timelineOutputDefFilePath   = "user_timeline.json"
	timelineOutputFilePathUsage = "Filepath for output file that consists of data about user's timeline tweets"

	friendsOutputDefFilePath   = "user_friends.json"
	friendsOutputFilePathUsage = "Filepath for output file that consists of data about user's friends"

	followersOutputDefFilePath   = "user_followers.json"
	followersOutputFilePathUsage = "Filepath for output file that consists of data about user's followers"

	queryOutputDefFilePath   = "query.json"
	queryOutputFilePathUsage = "Filepath for output file that consists of tweets with keyword"

	nFavsDef = twigger.LimitFavs
	nFavsUsage = "Number of recent favorites asked for."

	nTweetsDef = twigger.LimitTweets
	nTweetsUsage = "Number of recent tweets asked for."

	nTimelineDef = twigger.LimitTweets
	nTimelineUsage = "Number of recent timeline tweets asked for."

	nFollowersDef = -1 // All followers
	nFollowersUsage = "Number of followers asked for."

	nFriendsDef = -1 // All friends
	nFriendsUsage = "Number of friends asked for."

	keywordDef = "apple"
	keywordUsage = "keyword to search for among tweets"

	queryTypeDef = "top"
	queryTypeUsage = "type of the search for tweet search based on keyword. Valid values: [recent, top, mixed]"

	shorthand = " (shorthand)"
)

var(
	credsFlag string
	screenNameFlag string
	outputFilePathFlag string
	logFilePathFlag string
	nFavsFlag int
	nTweetsFlag int
	nFollowersFlag int
	nFriendsFlag int
	nTimelineFlag int
	keywordFlag string
	queryTypeFlag string
)

var logFile *os.File

func InitCommand(){
	flag.StringVar(&credsFlag, "credentials", credsDefValue, credsUsage)
	flag.StringVar(&credsFlag, "c", credsDefValue, credsUsage + shorthand)

	flag.StringVar(&logFilePathFlag, "log", logFileDefPath, logFilePathUsage)
	flag.StringVar(&logFilePathFlag, "l", logFileDefPath, logFilePathUsage + shorthand)

	logFile, err := os.OpenFile(logFilePathFlag, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		fmt.Printf("Cannot open or create log file (%v)\n", logFilePathFlag)
		fmt.Println("Program output will only be written to standard output and standard error.")
	}
	defer logFile.Close()
}
