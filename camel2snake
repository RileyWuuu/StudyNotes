package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	j := `
	{	"version": "1.0",
		"rules": [{"resource": {"path": "/api/data/documents"},
			"allowOrigins": [ "http://this.example.com", "http://that.example.com" ],
			"allowMethods": [ "GET" ],
			"allowCredentials": true
		}]
	}`
	jsonMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(j), &jsonMap)
	if err != nil {
		log.Fatal(err)
	}
	convertJson(jsonMap)
}
func convertJson(data map[string]interface{}) (result map[string]interface{}) {
	result = map[string]interface{}{}
	converter := regexp.MustCompile("([a-z0-0])([A-Z])")
	for i, x := range data {
		convertToSnake := strings.ToLower(converter.ReplaceAllString(i, "${1}_${2}"))
		switch x.(type) {
		case []interface{}:
			tmp := []interface{}{}

			for _, obj := range x.([]interface{}) {
				switch obj.(type) {
				case map[string]interface{}:
					tmp = append(tmp, convertJson(obj.(map[string]interface{})))
				default:
					tmp = append(tmp, obj)
				}
			}
			result[convertToSnake] = tmp
		case map[string]interface{}:
			temp := convertJson(x.(map[string]interface{}))
			result[convertToSnake] = temp
		default:
			result[convertToSnake] = x
		}
	}
	fmt.Println(result)
	return result
}
