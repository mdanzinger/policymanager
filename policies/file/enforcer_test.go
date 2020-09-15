package file

import (
	"os"
	"github.com/mdanzinger/policymanager"
	"github.com/mdanzinger/policymanager/pkg/duration"
	"github.com/mdanzinger/policymanager/policies/file/actions"
	"github.com/mdanzinger/policymanager/policies/file/conditions"
	"github.com/mdanzinger/policymanager/policies/file/testutil"
	"testing"
	"time"
)

func TestEnforcer_EnforceDeleteAction(t *testing.T) {
	// create test data
	testFile := testutil.NewTestFile(t,
		testutil.WithFileAge(time.Now().Add(time.Hour*5*-1)),
		testutil.Closed(),
	)

	testPolicy := policy.DefaultPolicy{
		PolType:       "file",
		PolActions:    []policy.Action{actions.Delete{}},
		PolConditions: []policy.Condition{conditions.FileAge{OlderThan: duration.Duration{time.Hour}}},
		PolSpec: &Spec{
			Target: []string{testFile.Name()},
		},
	}

	// create enforcer
	enforcer := Enforcer{}

	// enforce test policy
	err := enforcer.Enforce(testPolicy)
	if err != nil {
		t.Fatalf("unexpected error enforcing policy: %s", err.Error())
	}

	// assert test file was deleted
	if _, err := os.Stat(testFile.Name()); !os.IsNotExist(err) {
		t.Fatalf("expected enforcer to delete file, but file exists")
	}
}
