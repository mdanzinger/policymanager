package policy_test

import (
	"encoding/json"
	"github.com/mdanzinger/policymanager"
	_ "github.com/mdanzinger/policymanager/policies"
	"testing"
)

func TestDefaultPolicy_UnmarshalJSON(t *testing.T) {
	policyRaw := `{
  "id": "example policy",
  "name": "gamelog deletor",
  "type": "file",
  "description": "Deletes all game log (.gamelog) files that are older than 30 seconds",
  "actions": ["delete_file"],
  "conditions": [{
   "type": "file_age",
    "data": {
      "older_than": "30s"
    }
  }],
  "spec": {
    "target": ["./*.gamelog"]
  }
}`

	pol := policy.DefaultPolicy{}
	if err := json.Unmarshal([]byte(policyRaw), &pol); err != nil {
		t.Errorf("unexpected error unmarshalling policy: %s", err)
	}
	// TODO: assert spec was unmarshalled correctly
}
