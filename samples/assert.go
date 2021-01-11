package samples

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func AssertFilesEqual(t *testing.T, expected, actual string) bool {
	expectedFile, err := ioutil.ReadFile(expected)
	if err != nil {
		t.Errorf("Error reading expected file %v %v", expected, err)
		return false
	}

	actualFile, err := ioutil.ReadFile(actual)
	if err != nil {
		t.Errorf("Error reading actual file %v %v", actual, err)
		return false
	}

	if !bytes.Equal(expectedFile, actualFile) {
		t.Errorf("Files '%v' and '%v' are different", expected, actual)
		return false
	}

	return true
}
