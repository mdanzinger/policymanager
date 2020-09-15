package conditions

import (
	"github.com/mdanzinger/policymanager/pkg/duration"
	"github.com/mdanzinger/policymanager/policies/file/testutil"
	"os"
	"testing"
	"time"
)

func TestFileAge_Evaluate(t *testing.T) {
	var tests = []struct {
		name      string
		condition FileAge
		file      *os.File
		want      bool
		wantErr   bool
	}{
		{
			name:      "file age is older than the condition",
			condition: FileAge{OlderThan: duration.Duration{time.Hour}},
			file:      testutil.NewTestFile(t, testutil.WithFileAge(time.Now().Add(time.Hour*5*-1))),
			want:      true,
			wantErr:   false,
		},
		{
			name:      "file age is not newer than the condition",
			condition: FileAge{OlderThan: duration.Duration{time.Hour * 10}},
			file:      testutil.NewTestFile(t, testutil.WithFileAge(time.Now().Add(time.Hour*5*-1))),
			want:      false,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.condition.Evaluate(tt.file)
			if result != tt.want {
				t.Errorf("Evaluate() got = %v, want %v", result, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
