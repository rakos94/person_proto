package helper

import (
	"encoding/json"
	"fmt"
)

// ConvertRequest convert request to other struct (m)
func ConvertRequest(i interface{}, m interface{}) {
	byteArray, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(string(byteArray)), &m)
	if err != nil {
		fmt.Println(err)
	}
}

// ConvertStruct convert struct to other struct (m)
func ConvertStruct(i interface{}, m interface{}) {
	byteArray, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(string(byteArray)), &m)
	if err != nil {
		fmt.Println(err)
	}
}

// ConvertMap convert struct to map
func ConvertMap(i interface{}) map[string]interface{} {
	var result map[string]interface{}
	inrec, _ := json.Marshal(i)
	json.Unmarshal(inrec, &result)

	return result
}
