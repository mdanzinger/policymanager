package file

import (
	"fmt"
	"path/filepath"
	"github.com/mdanzinger/policymanager"
)

const (
	// SpecName contains the file spec name
	SpecName = "file"

	// SpecResolvedFiles is the key to the policy spec resolved files
	SpecResolvedFiles = "resolved_files"
)

type Spec struct {
	// Target directory to enforce policy. This can be a a hard path to a file, a directory, or a regex pattern (/path/*.log), etc.
	Target []string `json:"target"`
}

func (s *Spec) Name() string {
	return SpecName
}

// Parse returns a parsed policy spec
func (s *Spec) Parse() (map[string]interface{}, error) {
	spec := make(map[string]interface{})

	files, err := s.resolveFiles()
	if err != nil {
		return nil, fmt.Errorf("resolving files: %w", err)
	}
	spec[SpecResolvedFiles] = files

	return spec, nil
}

func (s *Spec) resolveFiles() ([]string, error) {
	var files []string

	for _, target := range s.Target {
		tmpFiles, err := filepath.Glob(target)
		if err != nil {
			return nil, err
		}
		files = append(files, tmpFiles...)
	}

	return files, nil
}

func init() {
	policy.AddSpec(SpecName, func() policy.Spec {
		return &Spec{}
	})
}
