package twigger

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"io/ioutil"
	"regexp"
)

func (c *Connection) uploadImage(filepath string) (int64, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return -1, err
	}

	mimeData, err := mimetype.DetectFile(filepath)
	if err != nil {
		return -1, err
	}
	acceptReg := regexp.MustCompile("(?i)jpeg|jpg|png|webp|gif")
	if !acceptReg.MatchString(mimeData.String()) {
		return -1, errors.New("Twitter only accepts jpeg,png,webp,gif mime type image uploads.")
	}

	client := c.Client
	mediaResponse, err := client.UploadMedia(base64.StdEncoding.EncodeToString(data))
	if err != nil {
		return -1, err
	}
	return mediaResponse.MediaID, nil
}

func (c *Connection) UploadImage(filepath string) (int64, error) {
	c.InfoLog.Printf("* Upload of %v is initiated by %v. *\n", filepath, c.User.ScreenName)
	mediaId, err := c.uploadImage(filepath)
	if err != nil {
		c.ErrLog.Printf("Upload of %v has failed! Error: %v\n", filepath, err)
		return mediaId, err
	}
	c.InfoLog.Printf("Upload of %v has been completed successfully.\n", filepath)
	return mediaId, err
}

func (c *Connection) uploadVideo(filepath string) (int64, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return -1, err
	}

	const VideoSizeLimit = (500 * 1000 * 1000) - 200 // Max allowed video size is 500MB as of May 2021
	if len(bytes) > VideoSizeLimit {
		return -1, errors.New(fmt.Sprintf("%v exceeds 500MB video upload limit constraint of Twitter!\n", filepath))
	}

	mimeData, err := mimetype.DetectFile(filepath)
	if err != nil {
		return -1, err
	}
	acceptReg := regexp.MustCompile("quicktime|mp4")
	if !acceptReg.MatchString(mimeData.String()) {
		return -1, errors.New("Twitter only accepts mp4 and quicktime mime type videos uploads.")
	}

	totalBytes := len(bytes)
	media, err := c.Client.UploadVideoInit(totalBytes, mimeData.String())
	if err != nil {
		return -1, err
	}

	mediaMaxLen := 1 * 1024 * 1024
	segment := 0
	for i := 0; i < totalBytes; i += mediaMaxLen {
		var mediaData string
		if i+mediaMaxLen < totalBytes {
			mediaData = base64.StdEncoding.EncodeToString(bytes[i : i+mediaMaxLen])
		} else {
			mediaData = base64.StdEncoding.EncodeToString(bytes[i:])
		}
		if err = c.Client.UploadVideoAppend(media.MediaIDString, segment, mediaData); err != nil {
			break
		}
		segment += 1
	}
	if err != nil {
		return -1, err
	}

	video, err := c.Client.UploadVideoFinalize(media.MediaIDString)
	if err != nil {
		return -1, err
	}
	return video.MediaID, nil
}

func (c *Connection) UploadVideo(filepath string) (int64, error) {
	c.InfoLog.Printf("* Upload of %v is initiated by %v. *\n", filepath, c.User.ScreenName)
	mediaId, err := c.uploadVideo(filepath)
	if err != nil {
		c.ErrLog.Printf("Upload of %v has failed! Error: %v\n", filepath, err)
		return mediaId, err
	}
	c.InfoLog.Printf("Upload of %v has been completed successfully.\n", filepath)
	return mediaId, err
}
