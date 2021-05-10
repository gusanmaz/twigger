package twigger

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const(
	MediaText    = "text" // No use for now
	MediaPicture = "picture"
	MediaCollage = "collage"
	MediaVideo   = "video"
)

var tweetElems = map[string]string{
	MediaText: "text",
	MediaPicture:"picture",
	MediaCollage: "pictures",
	MediaVideo: "video",
}

func (c *Connection) PublishTextTweet(text string) (int64, error){
	tw, err := c.Client.PostTweet(text, nil)
	if err == nil{
		c.InfoLog.Printf("Text tweet (ID:%v) is successfully published.\n", tw.Id)
		c.InfoLog.Printf("Content of text tweet: %v\n", text)
	}else{
		c.ErrLog.Printf("Text tweet with following text: [%v] couldn't be published!\n")
		c.ErrLog.Printf("Error message: %v\n", err)
	}
	return tw.Id, err
}

func (c *Connection) PublishTextTweetAsReply(text string, replyTweetID int64) (int64, error){
	v := url.Values{}
	v.Set("in_reply_to_status_id", fmt.Sprintf("%v", replyTweetID))
	tw, err := c.Client.PostTweet(text, nil)
	if err == nil{
		c.InfoLog.Printf("Text tweet (ID:%v) is successfully published as a reply to tweet (ID:%v).\n", tw.Id, replyTweetID)
		c.InfoLog.Printf("Content of text tweet: %v\n", text)
	}else{
		c.ErrLog.Printf("Text tweet with following text: [%v] couldn't be published!\n")
		c.ErrLog.Printf("Tweet's reply tweet ID: %v\n", replyTweetID)
		c.ErrLog.Printf("Error message: %v\n", err)
	}
	return tw.Id, err
}

func (c *Connection) publishMediaTweetGeneric(mediaType string, filePaths []string, text string, replyToTweetID int)(int64, error){
	c.InfoLog.Printf("*** Publishing of %v tweet with text: (%v) has initiated. ***\n", mediaType, text)

	selectionMsg := ""
	selectionMsgTemp := "A %v tweet may not contain more than %v %v.\n"
	selectionMsgTemp += "Path(s) of all media that is requested for inclusion in this tweet: %v\n"
	selectionMsgTemp += "Paths of the media selected for inclusion in this tweet: %v\n"
	selectedFilePaths := filePaths
	
	if mediaType == MediaCollage && len(filePaths) > 4{
		selectedFilePaths = filePaths[:4]
		selectionMsg = fmt.Sprintf(selectionMsgTemp, mediaType, 4, tweetElems[mediaType], filePaths, selectedFilePaths)
	}
	if (mediaType == MediaPicture || mediaType == MediaVideo)  && len(filePaths) > 1{
		selectedFilePaths = filePaths[:1]
		selectionMsg = fmt.Sprintf(selectionMsgTemp, mediaType, 1, tweetElems[mediaType], filePaths, selectedFilePaths)
	}
	if selectionMsg == ""{
		selectionMsg = fmt.Sprintf("Path(s) of the media that will be included in this tweet: %v\n", filePaths)
	}
	
	c.InfoLog.Printf(selectionMsg)

	client := c.Client
	uploadFunc := c.UploadImage
	if mediaType == MediaVideo{
		uploadFunc = c.UploadVideo
	}
	
	mediaIDs := make([]string, len(selectedFilePaths))
	for i, filepath := range selectedFilePaths{
		mediaID, _ := uploadFunc(filepath)
		mediaIDs[i] =  strconv.FormatInt(mediaID, 10)
	}

	v := url.Values{}
	v.Set("media_ids", strings.Join(mediaIDs, ","))
	// in reply to tweet id value equal to or less than 0 indicates this tweet is not a reply to a some specific tweet
	if replyToTweetID > 0{
		v.Set("in_reply_to_status_id", fmt.Sprintf("%v", replyToTweetID))
	}
	
	result, err := client.PostTweet(text, v)
	if err != nil {
		c.ErrLog.Printf("%v tweet publishing task has failed!\nError message: %v\n", mediaType, err)
		return result.Id, err
	} else {
		c.InfoLog.Printf("%v tweet publishing task has successfully completed.\n", mediaType)
		return result.Id, nil
	}
}

func (c *Connection) PublishPictureTweet(filepath string, text string) (int64, error){
	return c.publishMediaTweetGeneric(MediaPicture, []string{filepath}, text, -1)
}

func (c *Connection) PublishCollageTweet(filePaths []string, text string) (int64, error){
	return c.publishMediaTweetGeneric(MediaCollage, filePaths, text, -1)
}

func (c *Connection) PublishPictureTweetAsReply(filepath string, text string, replyTweetID int) (int64, error){
	return c.publishMediaTweetGeneric(MediaPicture, []string{filepath}, text, replyTweetID)
}

func (c *Connection) PublishCollageTweetAsReply(filePaths []string, text string, replyTweetID int) (int64, error){
	return c.publishMediaTweetGeneric(MediaCollage, filePaths, text, replyTweetID)
}

func (c *Connection) PublishVideoTweet(filepath string, text string) (int64,error){
	return c.publishMediaTweetGeneric(MediaVideo, []string{filepath}, text, -1)
}

func (c *Connection) PublishVideoTweetAsReply(filepath string, text string, replyTweetID int) (int64,error){
	return c.publishMediaTweetGeneric(MediaVideo, []string{filepath}, text, replyTweetID)
}



