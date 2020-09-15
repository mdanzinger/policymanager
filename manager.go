package policy

// Manager is responsible for manging / persisting policies.
type Manager interface {
	// Create creates and persists a new policy
	Create(p Policy) error

	// Update updates and existing policy
	Update(p Policy) error

	// Delete deletes a policy
	Delete(p Policy) error

	// GetAllWithTag returns policies that contain the tag
	GetAllWithTag(tag string) ([]Policy, error)
}
