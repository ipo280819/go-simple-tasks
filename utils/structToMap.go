package utils

import "encoding/json"

func StructToMap(s interface{}) map[string]interface{} {

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(s)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}
