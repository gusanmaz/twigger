package main

import (
	"flag"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/gusanmaz/twigger"
	"github.com/gusanmaz/twigger/genjson"
	"net/url"
	"os"
)

func main(){
	InitCommand()

	flag.StringVar(&queryTypeFlag, "output", queryOutputDefFilePath , queryOutputFilePathUsage)
	flag.StringVar(&outputFilePathFlag, "o", queryOutputDefFilePath, queryOutputFilePathUsage + shorthand)

	flag.StringVar(&keywordFlag, "keyword", keywordDef, keywordUsage)
	flag.StringVar(&queryTypeFlag, "k", keywordDef, keywordUsage + shorthand)

	flag.StringVar(&queryTypeFlag, "type", queryTypeDef, queryTypeUsage)
	flag.StringVar(&queryTypeFlag, "t", queryTypeDef, queryTypeUsage + shorthand)

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

	v := url.Values{}
	v.Set("result_type", queryTypeFlag)
	tweets, err := conn.GetXTweetsByQuery(keywordFlag, v)
	if err != nil{
		fmt.Fprintln(os.Stderr, "Terminating program...")
		os.Exit(1)
	}

	err = genjson.SaveDataTo(tweets, outputFilePathFlag)
	if err != nil{
		fmt.Fprintf(os.Stderr, "Cannot save query tweets into file %v", outputFilePathFlag)
		fmt.Fprintln(os.Stderr, "Terminating program...")
		os.Exit(1)
	}
}

