package policy

import (
	"encoding/json"
	"fmt"
)

type Condition interface {
	Evaluate(v interface{}) (bool, error)
	Name() string
}

type Conditions []Condition

func (c Conditions) MarshalJSON() ([]byte, error) {
	var jsonConditions []conditionJSON
	for _, cond := range c {
		data, err := json.Marshal(cond)
		if err != nil {
			return data, err
		}

		jsonCond := conditionJSON{
			Type: cond.Name(),
			Data: data,
		}

		jsonConditions = append(jsonConditions, jsonCond)
	}

	return json.Marshal(jsonConditions)
}

// UnmarshalJSON implments the unmarshaller interface
func (c *Conditions) UnmarshalJSON(data []byte) error {
	// TODO: Clean this up a bit, expose proper errors
	var jsonConditions []conditionJSON

	if err := json.Unmarshal(data, &jsonConditions); err != nil {
		return err
	}

	for _, jsCon := range jsonConditions {
		factory, ok := ConditionRegistry[jsCon.Type]
		if !ok {
			return fmt.Errorf("unable to find condition %s in registry", jsCon.Type)
		}

		condition := factory()
		if err := json.Unmarshal(jsCon.Data, &condition); err != nil {
			return err
		}

		*c = append(*c, condition)
	}
	return nil
}

type ConditionFactory func() Condition

type conditionJSON struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

var ConditionRegistry = map[string]ConditionFactory{}

// AddCondition adds a condition to the registry.
func AddCondition(name string, factory ConditionFactory) {
	ConditionRegistry[name] = factory
}
