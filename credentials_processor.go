package twigger

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"strings"
)

type Credentials struct{
	APIKey string `json:"APIKey" yaml:"APIKey"`
	APISecret string `json:"APISecret" yaml:"APISecret"`
	AccessToken string `json:"accessToken" yaml:"accessToken"`
	AccessSecret string `json:"accessSecret" yaml:"accessSecret"`
	BearerToken  string `json:"bearerToken" yaml:"bearerToken"`
}

func LoadJSONCredentials(filepath string) (Credentials, error){
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil{
		return Credentials{},err
	}
	bytes, err := io.ReadAll(f)
	if err != nil{
		return Credentials{},err
	}
	t := Credentials{}
	err = json.Unmarshal(bytes,&t)
	if err != nil{
		return t,err
	}
	return t,nil
}

func LoadYAMLCredentials(filepath string) (Credentials, error){
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil{
		return Credentials{},err
	}
	bytes, err := io.ReadAll(f)
	if err != nil{
		return Credentials{},err
	}
	t := Credentials{}
	err = yaml.Unmarshal(bytes,&t)
	if err != nil{
		return t,err
	}
	return t,nil
}

func LoadCredentials(filepath string)(Credentials, error){
	if strings.HasSuffix(filepath, "json"){
		return LoadJSONCredentials(filepath)
	}else if strings.HasSuffix(filepath, "yaml"){
		return LoadYAMLCredentials(filepath)
	}else{
		err := errors.New("Credentials file format is not recognized.")
		return Credentials{}, err
	}
}









