package policy

import (
	"encoding/json"
	"fmt"
)

// Action are performed by an enforcer if conditions are met
type Action interface {
	Execute(v interface{}) error
	Name() string
}

type Actions []Action

func (a *Actions) UnmarshalJSON(data []byte) error {
	var actionsRaw []string
	if err := json.Unmarshal(data, &actionsRaw); err != nil {
		return err
	}

	for _, actionName := range actionsRaw {
		factory, ok := ActionRegistry[actionName]
		if !ok {
			return fmt.Errorf("unable to find action type %s", actionName)
		}
		*a = append(*a, factory())
	}

	return nil
}

type ActionFactory func() Action

var ActionRegistry = map[string]ActionFactory{}

func AddAction(name string, factory ActionFactory) {
	ActionRegistry[name] = factory
}
