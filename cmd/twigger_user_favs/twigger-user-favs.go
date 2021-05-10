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
	flag.StringVar(&screenNameFlag, "screenname", screenNameDefValue, screenNameUsage)
	flag.StringVar(&screenNameFlag, "s", screenNameDefValue, screenNameUsage + shorthand)

	flag.StringVar(&outputFilePathFlag, "output", favOutputDefFilePath, favOutputFilePathUsage)
	flag.StringVar(&outputFilePathFlag, "o", favOutputDefFilePath, favOutputFilePathUsage + shorthand)

	flag.IntVar(&nFavsFlag, "n", nFavsDef, nFavsUsage)

	flag.Parse()

	creds, err := twigger.LoadJSONCredentials(credsFlag)
	if err != nil{
		fmt.Fprintln(os.Stderr, "Terminating program...")
		os.Exit(1)
	}

	conn,  err := twigger.NewConnection(creds, logFile, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Terminating program...")
		os.Exit(1)
	}
	conn.Client.SetLogger(anaconda.BasicLogger)

	if screenNameFlag == "" {
		screenNameFlag = conn.User.ScreenName
	}


	tweets, err := conn.GetRecentNFavoritesFromScreenName(screenNameFlag, nFavsFlag)
	if err != nil{
		fmt.Fprintln(os.Stderr, "Terminating program...")
		os.Exit(1)
	}

	err = genjson.SaveDataTo(tweets, outputFilePathFlag)
	if err != nil{
		fmt.Fprintf(os.Stderr, "Cannot save favorites of user %v into file %v", screenNameFlag, outputFilePathFlag)
		fmt.Fprintln(os.Stderr, "Terminating program...")
		os.Exit(1)
	}
}
