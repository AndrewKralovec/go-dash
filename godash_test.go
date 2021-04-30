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
	{1, 1},
}
var point = Point{0, 0}

/************************************
#region
First Tests
************************************/

func TestFirstString(t *testing.T) {
	target := stringArray[0]
	want := regexp.MustCompile(target)
	f := First(stringArray)

	if f == nil || !want.MatchString(f.(string)) {
		t.Fatalf(`value :%v does not match :%v`, want, f)
	}
}

func TestFirstInt(t *testing.T) {
	target := intArray[0]
	f := First(intArray)

	if f == nil || target != f {
		t.Fatalf(`value :%v does not match :%v`, target, f)
	}
}

func TestFirstPoint(t *testing.T) {
	target := structArray[0]
	f := First(structArray)

	if f == nil || target != f {
		t.Fatalf(`value :%v does not match :%v`, target, f)
	}
}

/************************************
#endregion
************************************/

/************************************
#region
Last Tests
************************************/
func TestLastString(t *testing.T) {
	target := stringArray[len(stringArray)-1]
	want := regexp.MustCompile(target)
	l := Last(stringArray)

	if l == nil || !want.MatchString(l.(string)) {
		t.Fatalf(`value :%v does not match :%v`, target, l)
	}
}

func TestLastInt(t *testing.T) {
	target := intArray[len(intArray)-1]
	l := Last(intArray)

	if l == nil || target != l {
		t.Fatalf(`value :%v does not match :%v`, target, l)
	}
}

func TestLastPoint(t *testing.T) {
	target := structArray[len(structArray)-1]
	l := Last(structArray)

	if l == nil || target != l {
		t.Fatalf(`value :%v does not match :%v`, target, l)
	}
}

/************************************
#endregion
************************************/

/************************************
#region
Each Tests
************************************/
func TestEachString(t *testing.T) {
	Each(stringArray, func(i interface{}, v interface{}) {
		target := stringArray[i.(int)]
		if v != target {
			t.Fatalf(`value :%v does not match :%v`, target, v)
		}
	})
}

func TestEachInt(t *testing.T) {
	Each(intArray, func(i interface{}, v interface{}) {
		target := intArray[i.(int)]
		if v != target {
			t.Fatalf(`value :%v does not match :%v`, target, v)
		}
	})
}

func TestEachStruct(t *testing.T) {
	Each(structArray, func(i interface{}, v interface{}) {
		target := structArray[i.(int)]
		if v != target {
			t.Fatalf(`value :%v does not match :%v`, target, v)
		}
	})
}

func TestEachWithStruct(t *testing.T) {
	Each(point, func(k interface{}, v interface{}) {
		if k != "X" && k != "Y" {
			t.Fatalf(`key :%v not found in :%v`, k, point)
		}
		if v != point.X && v != point.Y {
			t.Fatalf(`value :%v not found in :%v`, v, point)
		}
	})
}

func TestEachMatchingElements(t *testing.T) {
	var visted []interface{}
	Each(structArray, func(i interface{}, v interface{}) {
		target := structArray[i.(int)]
		if v != target {
			t.Fatalf(`value :%v does not match :%v`, target, v)
		}

		visted = append(visted, v)
	})
	vSize := len(visted)
	stSize := len(structArray)
	if vSize != stSize {
		t.Fatalf(`size :%v does not match :%v`, vSize, stSize)
	}

	for i, v := range structArray {
		target := visted[i]
		if v != target {
			t.Fatalf(`value :%v does not match :%v`, target, v)
		}
	}

}

/************************************
#endregion
************************************/

/************************************
#region
Includes Tests
************************************/
func TestIncludesString(t *testing.T) {
	target := stringArray[0]
	has := Includes(stringArray, target)

	if !has {
		t.Fatalf(`could not find value :%v in :%v`, target, stringArray)
	}
}

func TestIncludesInt(t *testing.T) {
	target := intArray[0]
	has := Includes(intArray, target)

	if !has {
		t.Fatalf(`could not find value :%v in :%v`, target, intArray)
	}
}

func TestIncludesStruct(t *testing.T) {
	target := structArray[0]
	has := Includes(structArray, target)

	if !has {
		t.Fatalf(`could not find value :%v in :%v`, target, structArray)
	}
}

/************************************
#endregion
************************************/

/************************************
#region
Find Tests
************************************/
func TestFindString(t *testing.T) {
	target := stringArray[0]

	v := Find(stringArray, func(val interface{}) bool {
		return val.(string) == target
	})

	if v != target {
		t.Fatalf(`value :%v does not match :%v`, target, v)
	}
}

func TestFindInt(t *testing.T) {
	target := intArray[0]

	v := Find(intArray, func(val interface{}) bool {
		return val.(int) == target
	})

	if v != target {
		t.Fatalf(`value :%v does not match :%v`, target, v)
	}
}

func TestFindStruct(t *testing.T) {
	target := structArray[0]

	v := Find(structArray, func(val interface{}) bool {
		return val.(Point).X == target.X
	})

	if v != target {
		t.Fatalf(`value :%v does not match :%v`, target, v)
	}
}

/************************************
#endregion
************************************/
