package decree

import (
	"io/ioutil"
	"encoding/json"
	"testing"
)



func TestNewDecree(t *testing.T) {

	var session map[string]interface{}

	sessionBytes, err := ioutil.ReadFile("samples/session.json")
    if err != nil {
        t.Error(err.Error())
    }

    err = json.Unmarshal(sessionBytes, &session)
	if err != nil {
		t.Error(err.Error())
	}

    jsonBytes, err := ioutil.ReadFile("samples/rules.json")
    if err != nil {
        t.Error(err.Error())
    }

    rules,err := CreateRulesFromJSON(jsonBytes)
    if err != nil {
        t.Error(err.Error())
    }

    decree := NewDecree(rules)

    result, err := decree.Runner(session)

    if err != nil {
    	t.Error(err.Error())
    }

    if result.OverallCollectionResult == false {
		t.Error("Expected result of true, received false")
	}
    
}

