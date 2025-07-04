package jsonutil

import "encoding/json"

type JsonUtilsInterface interface {
	ConvertStringToJson(stringResponse string, target interface{}) error
}

type JsonUtils struct {
}

func NewJsonUtils() JsonUtilsInterface {
	return &JsonUtils{}
}

func (ju *JsonUtils) ConvertStringToJson(stringResponse string, target interface{}) error {

	return json.Unmarshal([]byte(stringResponse), &target)
}
