package file

import (
	"fmt"
	"os"

	"github.com/mdanzinger/policymanager"
)

const (
	EnforcerPolicyType = "file"
)

// Enforcer represents a file policy enforcer and implements the policy.Enforcer interface
type Enforcer struct{}

func (e Enforcer) Enforce(p policy.Policy) error {
	spec, ok := p.Spec().(*Spec)
	if !ok {
		return policy.ErrUnexpectedSpecFieldType
	}

	// get files from policy spec
	params, err := spec.Parse()
	if err != nil {
		return err
	}

	files, ok := params[SpecResolvedFiles].([]string)
	if !ok {
		return policy.ErrUnexpectedSpecFieldType
	}

	// enforce policy on each file
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("opening file: %w", err)
		}

		// ensure file satisfies conditions before enforcing the action
		satisfies, err := satisfiesConditions(f, p.Conditions())
		if err != nil {
			f.Close()
			return fmt.Errorf("asserting conditions: %w", err)
		}

		// we need to close the file prior to executing an action as Windows won't allow a file to be deleted if it
		// is opened else. If an action requires access to the descriptor, it will need to open the file there.
		f.Close()

		// if file satisfies conditions, apply action
		if satisfies {
			err := applyActions(f.Name(), p.Actions())

			if err != nil {
				return fmt.Errorf("applying actions: %w", err)
			}
		}
	}

	return nil
}

func applyActions(file string, actions policy.Actions) error {
	for _, action := range actions {
		err := action.Execute(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func satisfiesConditions(f *os.File, conditions policy.Conditions) (bool, error) {
	for _, condition := range conditions {
		satisfies, err := condition.Evaluate(f)
		if err != nil || !satisfies {
			return satisfies, err
		}
	}
	return true, nil
}

func init() {
	policy.AddEnforcer(EnforcerPolicyType, &Enforcer{})
}
