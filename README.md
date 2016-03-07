# decree-lite

Takes a formatted rule set and runs them against a map[string]interface{} using go-ruler

## Introduction

This takes a formatted Decree rule set and converts it to be run against go-ruler.

```json
 {
    "RuleName": "Contact",
    "RuleSets": [
        {
            "OnFalse": {
                "Type": "string",
                "Value": "Please update your contact details."
            },
            "OnTrue": {
                "Type": "boolean",
                "Value": "true"
            },
            "Operator": "AND",
            "Rule": [
                {
                    "EvaluateTo": "([0-9])+ (?i)([a-z]){2,}",
                    "PropertyToEvaluate": "session.request.serviceAddress.address1",
                    "Rule": "matches_regex"
                },
                {
                    "EvaluateTo": "[a-z ]+",
                    "PropertyToEvaluate": "session.request.serviceAddress.city",
                    "Rule": "matches_regex"
                },
                {
                    "EvaluateTo": "(?i)[a-z]{2}",
                    "PropertyToEvaluate": "session.request.serviceAddress.state",
                    "Rule": "matches_regex"
                },
                {
                    "EvaluateTo": "[0-9]{5}",
                    "PropertyToEvaluate": "session.request.serviceAddress.zipCode",
                    "Rule": "matches_regex"
                }
            ]
        }
    ]
}
```

This is an example of the decree rule set.