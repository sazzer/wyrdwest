package http

import "encoding/json"

// ParseJSONToMap will parse a JSON string to a Map[string]interface{} structure that can be used for assertions
func ParseJSONToMap(input string) map[string]interface{} {
	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(input), &parsed); err != nil {
		panic(err)
	}
	return parsed
}
