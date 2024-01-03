package utils

import "encoding/json"

func MarshalToString(data interface{}) string {
	if data == nil {
		return ""
	}
	jsonString, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonString)
}
