package testutil

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

type option func(t *testing.T, f *os.File)

func WithFileAge(age time.Time) option {
	return func(t *testing.T, f *os.File) {
		if err := os.Chtimes(f.Name(), age, age); err != nil {
			t.Fatalf("setting file age: %s", err)
		}
	}
}

func Closed() option {
	return func(t *testing.T, f *os.File) {
		if err := f.Close(); err != nil {
			t.Fatalf("closing file: %s", err)
		}
	}
}

// NewTestFile creates a temporary file for use in tests
func NewTestFile(t *testing.T, opts ...option) *os.File {
	testFile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal("unable to create test file: %w", err)
	}

	for _, opt := range opts {
		opt(t, testFile)
	}

	t.Cleanup(cleanup(testFile))
	return testFile
}

func cleanup(file *os.File) func() {
	return func() {
		file.Close()
		os.Remove(file.Name())
	}
}
