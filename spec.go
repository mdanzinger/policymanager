package policy

type Spec interface {
	Parse() (map[string]interface{}, error)
	Name() string
}

type SpecFactory func() Spec

var SpecRegistry = map[string]SpecFactory{}

// AddSpec adds a spec factory to the SpecRegistry and should be done as part of a package's init
func AddSpec(name string, factory SpecFactory) {
	SpecRegistry[name] = factory
}
