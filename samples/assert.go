package samples

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func AssertBytesAreEqual(t *testing.T, expectedContentAt string, actual []byte) {
	expected, err := ioutil.ReadFile(expectedContentAt)
	if err != nil {
		t.Fatalf("Error reading expected file %v %v", expectedContentAt, err)
	}
	if !bytes.Equal(expected, actual) {
		actualContentAt := expectedContentAt + "-actual.png"
		if err = ioutil.WriteFile(actualContentAt, actual, 0644); err != nil {
			t.Errorf("Failed to store actual content at %s", actualContentAt)
		}
		t.Errorf("Files '%v' and '%v' are different", expectedContentAt, actualContentAt)
	}
}
