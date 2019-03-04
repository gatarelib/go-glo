package utils

import "encoding/json"

// ToJSON interface to JSON string
func ToJSON(i interface{}) string {
	data, _ := json.Marshal(i)
	return string(data)
}

// ToRawMessage interface to raw bytes
func ToRawMessage(i interface{}) json.RawMessage {
	data, _ := json.Marshal(i)
	return data
}

// ToJSONArray interfaces JSON string
func ToJSONArray(i ...interface{}) string {
	return ToJSON(i)
}
