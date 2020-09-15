package conditions

import (
	"fmt"
	"os"
	"github.com/mdanzinger/policymanager"
	"github.com/mdanzinger/policymanager/pkg/duration"
	"time"
)

const (
	ConditionFileAgeName = "file_age"
)

// FileAge returns true if a given file is older than the FileAge.OlderThan
type FileAge struct {
	OlderThan duration.Duration `json:"older_than"`
}

func (f FileAge) Name() string {
	return ConditionFileAgeName
}

// Evaluate checks if the file age is older than the condition's OlderThan duration.
func (f FileAge) Evaluate(v interface{}) (bool, error) {
	file, ok := v.(*os.File)
	if !ok {
		return false, policy.ErrUnexpectedConditionInput
	}

	fileStat, err := file.Stat()
	if err != nil {
		return false, fmt.Errorf("getting file attributes: %w", err)
	}

	return time.Since(fileStat.ModTime()) > f.OlderThan.Duration, nil
}

func init() {
	policy.AddCondition(ConditionFileAgeName, func() policy.Condition {
		return &FileAge{}
	})
}
