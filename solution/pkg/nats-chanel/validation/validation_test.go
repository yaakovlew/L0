package validation

import (
	"encoding/json"
	"solution/pkg/model"
	"testing"
)

func TestValidation(t *testing.T) {
	testCount := make(map[string]string)
	testCount["FirstName"] = "Pedro"
	testCount["SecondName"] = "Bonucci"
	b, _ := json.Marshal(testCount)
	if ValidationMessage(b) {
		t.Error("Function need return false, got ", ValidationMessage(b))
	}
	user := model.User{"Andre", "Pirlo", "Apirlo@mail.ru"}
	b, _ = json.Marshal(user)
	if !ValidationMessage(b) {
		t.Error("Function need return true, got ", ValidationMessage(b))
	}
}
