package twigger

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

type SimpleUser struct{
	ID int64
	IDStr string
	Name string
	ScreenName string
}

func (t Tweet) GetTextURLs()[]string{
	urls := []string{}
	if t.Entities.Urls != nil{
		for _, v := range t.Entities.Urls{
			urls = append(urls, v.Display_url)
		}
	}
	return urls
}

func (t Tweet) GetUserMentions()[]SimpleUser{
	users := []SimpleUser{}
	if t.Entities.Urls != nil{
		for _, m := range t.Entities.User_mentions{
			s := SimpleUser{
				ID:         m.Id,
				IDStr:      m.Id_str,
				Name:       m.Name,
				ScreenName: m.Screen_name,
			}
			users = append(users, s)
		}
	}
	return users
}

func (t Tweet) GetMedias()[]string{
	urls := []string{}
	if t.ExtendedEntities.Media != nil{
		for _, m := range t.ExtendedEntities.Media{
			urls = append(urls, m.Media_url_https)
		}
	}
	return urls
}

func (t Tweet) GetHashtags()[]string{
	hashtags := []string{}
	if t.Entities.Hashtags != nil{
		for _, tags := range t.Entities.Hashtags{
			hashtags = append(hashtags, tags.Text)
		}
	}
	return hashtags
}

func (t Tweet) GetText()string{
	return t.Text
}

func (t Tweet) GetFullText()string{
	return t.FullText
}

func (t Tweet) DownloadMediaTo(dirPath string) error{
	finfo, err := os.Stat(dirPath)
	if err != nil || finfo.IsDir() == false{
		err := os.Mkdir(dirPath, 0700)
		if err != nil{
			return errors.New("cannot create the directory")
		}
	}

	mediaURLs := t.GetMedias()

	for _, url := range mediaURLs{
		response, err := http.Get(url)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			return errors.New("Received non 200 response code")
		}

		parts := strings.Split(url, "/")
		fileName := parts[len(parts) - 1]
		//Create a empty file
		file, err := os.Create(path.Join(dirPath, fileName))
		if err != nil {
			return err
		}
		defer file.Close()

		//Write the bytes to the fiel
		_, err = io.Copy(file, response.Body)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t Tweet)DownloadMedia(){
	dirPath := t.IdStr
	t.DownloadMediaTo(dirPath)
}

func (t Tweet)ContainsPhoto()bool{
	if len(t.ExtendedEntities.Media) == 0{
		return false
	}
	for _, v := range t.ExtendedEntities.Media{
		if v.Type == "photo"{
			return true
		}
	}
	return false
}

func (t Tweet)ContainsVideo()bool{
	if len(t.ExtendedEntities.Media) == 0{
		return false
	}
	for _, v := range t.ExtendedEntities.Media{
		if v.Type == "video"{
			return true
		}
	}
	return false
}

func (t Tweet)ContainsGIF()bool{
	if len(t.ExtendedEntities.Media) == 0{
		return false
	}
	for _, v := range t.ExtendedEntities.Media{
		if v.Type == "animated_gif"{
			return true
		}
	}
	return false
}

func (t Tweet)ContainsOnlyText()bool{
	return !(t.ContainsGIF() || t.ContainsPhoto() || t.ContainsVideo())
}



