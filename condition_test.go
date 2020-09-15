package policy_test

import (
	"encoding/json"
	"github.com/mdanzinger/policymanager"
	_ "github.com/mdanzinger/policymanager/policies"
	"testing"
)

func TestConditions_UnmarshalJSON(t *testing.T) {
	rawCondition := `[
   {
      "type":"file_age",
      "data":{
         "older_than":"90d"
      }
   }
]`

	conditions := policy.Conditions{}
	if err := json.Unmarshal([]byte(rawCondition), &conditions); err != nil {
		t.Errorf("unexpected error unmarshalling condition: %s", err)
	}

	// TODO: Assert condition data was unmarshalled correctly
}
