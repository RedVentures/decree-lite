package decree

import (
	"github.com/hopkinsth/go-ruler"
	"encoding/json"
)

// Decree main object
type Decree struct {
	rules *Rules
}

// Result is what is returned in the decree format
type Result struct {
	OverallCollectionResult bool
	OverallCollectionResultEvaluatedString string
	RuleCollectionDefinition string
	ResultSetResults string
}

// NewDecree returns a new instance of decree
func NewDecree(rules *Rules) *Decree {
	if rules != nil {
		return &Decree{
			rules,
		}
	}

	return &Decree{}
}

// CreateRulesFromJSON creates a ruleset based on a given decreerule
func CreateRulesFromJSON(jsonBytes []byte) (*Rules, error) {

	var rules *Rules

	err := json.Unmarshal(jsonBytes, &rules)

	if err != nil {
		return nil, err
	}
    
    return rules, nil

}

// Runner runs the decree rule sets on a given context
func (d *Decree) Runner(context map[string]interface{}) (resultCollection Result, err error) {

	// Currently only supporting 1 rule set. No support for AND, OR, or nested rules
	jsonString, err := json.Marshal(d.rules.RuleSets[0].Rule)
    if err != nil {
		return resultCollection, err
	}
	
	engine,_ := ruler.NewRulerWithJson(jsonString)

	result := engine.Test(context)

	resultCollection.OverallCollectionResult = result

	return resultCollection, nil
}