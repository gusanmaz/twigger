package main

import (
	"flag"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/gusanmaz/twigger"
	"github.com/gusanmaz/twigger/genjson"
	"os"
)

func main(){
	InitCommand()

	flag.StringVar(&outputFilePathFlag, "output", timelineOutputDefFilePath, timelineOutputFilePathUsage)
	flag.StringVar(&outputFilePathFlag, "o", timelineOutputDefFilePath, timelineOutputFilePathUsage + shorthand)

	flag.IntVar(&nTimelineFlag, "n", nTimelineDef, nTimelineUsage)

	flag.Parse()

	creds, err := twigger.LoadJSONCredentials(credsFlag)
	if err != nil{
		fmt.Fprintln(os.Stderr, "Terminating program...")
		os.Exit(1)
	}

	conn, err := twigger.NewConnection(creds, logFile, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Terminating program...")
		os.Exit(1)
	}
	conn.Client.SetLogger(anaconda.BasicLogger)

	tweets, err := conn.GetRecentNHomeTimelineTweets(nTimelineFlag)
	if err != nil{
		fmt.Fprintln(os.Stderr, "Terminating program...")
		os.Exit(1)
	}

	err = genjson.SaveDataTo(tweets, outputFilePathFlag)
	if err != nil{
		fmt.Fprintf(os.Stderr, "Cannot save timeline tweets of user %v into file %v", conn.User.ScreenName, outputFilePathFlag)
		fmt.Fprintln(os.Stderr, "Terminating program...")
		os.Exit(1)
	}
}
