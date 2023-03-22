package converse

import (
	"encoding/json"
	"fmt"
)

// return Key=Value
func ParseToString_KeyValue(key string, value interface{}) string {
	return fmt.Sprintf(" %v=%+v", key, value)
}

// return Value
func ParseToString(value interface{}) string {
	return fmt.Sprintf("%+v", value)
}

func JsonMarshalIndent(data interface{}) string {
	jsonMarshalIndent, _ := json.MarshalIndent(data, "\t", "\t")
	return fmt.Sprintf("%v", string(jsonMarshalIndent))
}

func JsonMarshal(data interface{}) string {
	jsonMarshal, _ := json.Marshal(data)
	return fmt.Sprintf("%v", string(jsonMarshal))
}

func JsonUnmarshal(data string) interface{} {
	var response interface{}
	json.Unmarshal([]byte(data), &response)
	return response
}
