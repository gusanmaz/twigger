package genjson

import "os"
import "encoding/json"

func SaveDataTo(data interface{}, filepath string)error{
	bytes, err := json.Marshal(data)
	if err != nil{
		return err
	}
	err =  os.WriteFile(filepath, bytes, 0666)
	if err != nil{
		return err
	}
	return nil
}

func LoadDataFrom(filepath string)(interface{},error){
	//emptyInterface := interface{}
	bytes, err := os.ReadFile(filepath)
	if err != nil{
		return nil, err
	}
	var ret interface{}
	err = json.Unmarshal(bytes, &ret)
	if err != nil{
		return nil,err
	}
	return ret, nil
}

