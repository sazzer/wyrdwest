package http

import "encoding/json"

// ParseJsonToMap will parse a JSON string to a Map[string]interface{} structure that can be used for assertions
func ParseJsonToMap(input string) map[string]interface{} {
	var parsed map[string]interface{}
	json.Unmarshal([]byte(input), &parsed)
	return parsed
}
