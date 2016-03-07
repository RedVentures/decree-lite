package decree

import (
	"encoding/json"
	"errors"
)

// Rules is the collection of rules to run
type Rules struct {
	RuleName string
	RuleSets []RuleSets
}

// RuleSets contains what to do on true and false and also a slice of rules to run
type RuleSets struct {
	OnFalse Decision
	OnTrue Decision
	Rule []Rule
}

// Decision is what to do on result
type Decision struct {
	DecicionType string `json:"type"`
	Value string
}

// Rule individual to run and where to compare
type Rule struct {
	Value interface{} `json:"value"`
	Path string `json:"path"`
	Comparator string `json:"comparator"`
}

// DecreeRule is the original rule that comes from decree
type DecreeRule struct {
	EvaluateTo interface{} 
	PropertyToEvaluate string
	Rule string
}

// UnmarshalJSON transforms decree rules into rules to be used by go-ruler
func (r *Rule) UnmarshalJSON(data []byte) error {
	
	var aux DecreeRule
	
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	r.Path = aux.PropertyToEvaluate
	r.Value = aux.EvaluateTo

	// add more rules here if need be right now only need matches
	switch aux.Rule {
	case "matches_regex":
		r.Comparator = "matches"
	case "equals":
		r.Comparator = "eq"
	default:
		return errors.New("Unknown rule: " + aux.Rule)
	}

    return nil
}