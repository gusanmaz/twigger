package twigger

import (
	"github.com/ChimeraCoder/anaconda"
	"io"
	"log"
	"net/url"
	"os"
	"time"
)

type Connection struct {
	Client *anaconda.TwitterApi

	logFile       *os.File
	logStdoutFile *os.File
	logStdErrFile *os.File
	InfoLog       *log.Logger
	ErrLog        *log.Logger

	User         *anaconda.User
	CreationTime int64
}

func NewConnection(c Credentials, logFile, logOutFile, logErrFile *os.File) (*Connection, error) {
	anacondaClient := anaconda.NewTwitterApiWithCredentials(c.AccessToken, c.AccessSecret, c.APIKey, c.APISecret)

	newConn := Connection{Client: anacondaClient, CreationTime: time.Now().Unix()}
	newConn.logFile = logFile
	newConn.logStdoutFile = logOutFile
	newConn.logStdErrFile = logErrFile
	//newConn.Client.

	infoMW := io.MultiWriter(logFile, logOutFile)
	errMW := io.MultiWriter(logFile, logErrFile)

	newConn.InfoLog = log.New(infoMW, "Twigger INFO:", log.LstdFlags)
	newConn.ErrLog = log.New(errMW, "Twigger ERROR:", log.LstdFlags)

	_, err := anacondaClient.VerifyCredentials()
	if err != nil {
		newConn.User = nil
		newConn.ErrLog.Printf("Connection error. Error message: %v \n", err)
	} else {
		user, _ := anacondaClient.GetSelf(url.Values{})
		screenName := user.ScreenName
		newConn.User = &user
		newConn.InfoLog.Printf("Connection is successfully established for Twitter user: %v\n", screenName)
	}
	return &newConn, err
}
