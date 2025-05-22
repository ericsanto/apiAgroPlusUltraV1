package jsonutil

import "encoding/json"

func ConvertStringToJson(stringResponse string, target interface{}) error {

	return json.Unmarshal([]byte(stringResponse), &target)
}
