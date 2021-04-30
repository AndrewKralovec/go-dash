package godash

import (
	"regexp"
	"testing"
)

var testingData = []string{"peach", "apple", "pear", "plum"}

func TestFirst(t *testing.T) {
	target := testingData[0]
	want := regexp.MustCompile(target)
	f := First(testingData)

	if f == nil || !want.MatchString(f.(string)) {
		t.Fatalf(`does not match`)
	}
}
