package validation

import (
	"encoding/json"
)

func ValidationMessage(m []byte) (count bool) {
	var f interface{}
	err := json.Unmarshal(m, &f)
	if err != nil {
		return false
	}
	msg := f.(map[string]interface{})
	count = true
	countFN := false
	countSN := false
	countE := false
	for k, v := range msg {
		switch v.(type) {
		case string:
			if k == "FirstName" {
				countFN = true
			} else if k == "SecondName" {
				countSN = true
			} else if k == "Email" {
				countE = true
			} else {
				count = false
			}
		default:
			count = false
		}
	}
	if count == true && countSN == true && countE == true && countFN == true {
		return true
	} else {
		return false
	}
}
