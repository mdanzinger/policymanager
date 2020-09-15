package policy

// Enforcer enforces a given policy
type Enforcer interface {
	Enforce(p Policy) error
}

var EnforcerRegistry = map[string]Enforcer{}

func AddEnforcer(policyType string, enforcer Enforcer) {
	EnforcerRegistry[policyType] = enforcer
}

// ResolveEnforcer returns the enforcer for a given policy type
func ResolveEnforcer(p Policy) (Enforcer, error) {
	enforcer, ok := EnforcerRegistry[p.Type()]
	if !ok {
		return nil, NewErrEnforcerNotFound(p.Type())
	}
	return enforcer, nil
}
