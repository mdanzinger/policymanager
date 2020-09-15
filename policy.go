package policy

import (
	"encoding/json"
	"fmt"
)

type Policy interface {
	ID() string
	Name() string
	Description() string
	Tags() []string
	Type() string
	Actions() Actions
	Conditions() Conditions
	Spec() Spec
}

// Default implementation of a Policy
type DefaultPolicy struct {
	PolID          string     `json:"id"`
	PolName        string     `json:"name"`
	PolDescription string     `json:"description"`
	PolType        string     `json:"type"`
	PolTags        []string   `json:"tags"`
	PolActions     Actions    `json:"actions"`
	PolConditions  Conditions `json:"conditions"`
	PolSpec        Spec       `json:"spec"`
}

func (d *DefaultPolicy) UnmarshalJSON(data []byte) error {
	var jsonPolicy = struct {
		PolID          string          `json:"id"`
		PolName        string          `json:"name"`
		PolDescription string          `json:"description"`
		PolType        string          `json:"type"`
		PolTags        []string        `json:"tags"`
		PolActions     Actions         `json:"actions"`
		PolConditions  Conditions      `json:"conditions"`
		Spec           json.RawMessage `json:"spec"`
	}{
		PolActions:    Actions{},
		PolConditions: Conditions{},
	}

	if err := json.Unmarshal(data, &jsonPolicy); err != nil {
		return err
	}

	factory, ok := SpecRegistry[jsonPolicy.PolType]
	if !ok {
		return fmt.Errorf("spec not found for policy type %s", jsonPolicy.PolType)
	}

	spec := factory()
	if err := json.Unmarshal(jsonPolicy.Spec, spec); err != nil {
		return err
	}

	*d = DefaultPolicy{
		PolID:          jsonPolicy.PolID,
		PolName:        jsonPolicy.PolName,
		PolDescription: jsonPolicy.PolDescription,
		PolType:        jsonPolicy.PolType,
		PolTags:        jsonPolicy.PolTags,
		PolActions:     jsonPolicy.PolActions,
		PolConditions:  jsonPolicy.PolConditions,
		PolSpec:        spec,
	}
	return nil
}

// TODO: Implement custom Marshal/Unmarshal functions to correctly map conditions and policy specs

func (d DefaultPolicy) Tags() []string {
	return d.PolTags
}

func (d DefaultPolicy) ID() string {
	return d.PolID
}

func (d DefaultPolicy) Name() string {
	return d.PolName
}

func (d DefaultPolicy) Description() string {
	return d.PolDescription
}

func (d DefaultPolicy) Type() string {
	return d.PolType
}

func (d DefaultPolicy) Actions() Actions {
	return d.PolActions
}

func (d DefaultPolicy) Conditions() Conditions {
	return d.PolConditions
}

func (d DefaultPolicy) Spec() Spec {
	return d.PolSpec
}
