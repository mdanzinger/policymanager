package actions

import (
	"github.com/mdanzinger/policymanager"
	"os"
)

const (
	ActionDeleteName = "delete_file"
)

// Delete deletes a given file
type Delete struct{}

func (d Delete) Name() string {
	return ActionDeleteName
}

// Execute implements the policy.Action interface
func (d Delete) Execute(v interface{}) error {
	file, ok := v.(string)
	if !ok {
		return policy.ErrUnexpectedActionInput
	}

	return os.Remove(file)
}

func init() {
	policy.AddAction(ActionDeleteName, func() policy.Action {
		return Delete{}
	})
}
