package godash

import (
	"regexp"
	"testing"
)

type Point struct {
	X int
	Y int
}
type Points []Point

var stringArray = []string{"peach", "apple", "pear", "plum"}
var intArray = []int{1, 2, 3, 4}
var structArray = Points{
	{0, 0},
	{1, 0},
	{0, 1},
}

/************************************
First Tests
************************************/

func TestFirstString(t *testing.T) {
	target := stringArray[0]
	want := regexp.MustCompile(target)
	f := First(stringArray)

	if f == nil || !want.MatchString(f.(string)) {
		t.Fatalf(`does not match`)
	}
}

func TestFirstInt(t *testing.T) {
	target := intArray[0]
	f := First(intArray)

	if f == nil || target != f {
		t.Fatalf(`does not match`)
	}
}

func TestFirstPoint(t *testing.T) {
	target := structArray[0]
	f := First(structArray)

	if f == nil || target != f {
		t.Fatalf(`does not match`)
	}
}

/************************************
Last Tests
************************************/
func TestLastString(t *testing.T) {
	target := stringArray[len(stringArray)-1]
	want := regexp.MustCompile(target)
	l := Last(stringArray)

	if l == nil || !want.MatchString(l.(string)) {
		t.Fatalf(`does not match`)
	}
}

func TestLastInt(t *testing.T) {
	target := intArray[len(intArray)-1]
	l := Last(intArray)

	if l == nil || target != l {
		t.Fatalf(`does not match`)
	}
}

func TestLastPoint(t *testing.T) {
	target := structArray[len(structArray)-1]
	l := Last(structArray)

	if l == nil || target != l {
		t.Fatalf(`does not match`)
	}
}
