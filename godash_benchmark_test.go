package godash

import (
	"math/rand"
	"testing"
)

func GenerateIntSlice(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(1e9))
	}
	return s
}

func GenerateStructSlice(n int) Points {
	s := make(Points, 0, n)
	for i := 0; i < n; i++ {
		point := Point{rand.Intn(1e9), rand.Intn(1e9)}
		s = append(s, point)
	}
	return s
}

// Do nothing
func DoNothing(data ...interface{}) {}

// avoid compiler optimizations
var boolResult bool
var intResult int
var interfaceResult interface{}
var structResult Point
var intSliceResult []int
var structSliceResult Points
var interfaceSliceResult []interface{}

/************************************
#region
Each Benchmarks
************************************/

func BenchmarkNativeEachInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		intSlice := GenerateIntSlice(1000)

		b.StartTimer()
		for i, v := range intSlice {
			DoNothing(i, v)
		}
	}
}

func BenchmarkEachInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		intSlice := GenerateIntSlice(1000)

		b.StartTimer()
		Each(intSlice, func(i interface{}, v interface{}) {
			DoNothing(i, v)
		})
	}
}

func BenchmarkNativeEachStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		structSlice := GenerateStructSlice(1000)

		b.StartTimer()
		for i, v := range structSlice {
			DoNothing(i, v)
		}
	}
}

func BenchmarkEachStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		structSlice := GenerateStructSlice(1000)

		b.StartTimer()
		Each(structSlice, func(i interface{}, v interface{}) {
			DoNothing(i, v)
		})
	}
}

/************************************
#endregion
************************************/

/************************************
#region
Filter Benchmarks
************************************/

func BenchmarkNativeFilterInt(b *testing.B) {
	b.StopTimer()

	var r []int
	intSlice := GenerateIntSlice(1000)
	target := intSlice[0]

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		r = make([]int, 0)
		for _, val := range intSlice {
			if val == target {
				r = append(r, target)
			}
		}
	}

	intSliceResult = r
}

func BenchmarkFilterInt(b *testing.B) {
	b.StopTimer()

	var r []interface{}
	intSlice := GenerateIntSlice(1000)
	target := intSlice[0]

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		r = Filter(intSlice, func(val interface{}) bool {
			return val.(int) == target
		})
	}

	interfaceSliceResult = r
}

func BenchmarkNativeFilterStruct(b *testing.B) {
	b.StopTimer()

	var r Points
	structSlice := GenerateStructSlice(1000)
	target := structSlice[0]

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		r = make(Points, 0)
		for _, val := range structSlice {
			if val.X == target.X {
				r = append(r, target)
			}
		}
	}

	structSliceResult = r
}

func BenchmarkFilterStruct(b *testing.B) {
	b.StopTimer()

	var r []interface{}
	structSlice := GenerateStructSlice(1000)
	target := structSlice[0]

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		r = Filter(structSlice, func(val interface{}) bool {
			return val.(Point).X == target.X
		})
	}

	interfaceSliceResult = r
}

/************************************
#endregion
************************************/

/************************************
#region
Find Benchmarks
************************************/

func BenchmarkNativeFindInt(b *testing.B) {
	b.StopTimer()

	var r int
	intSlice := GenerateIntSlice(1000)
	middle := ((len(intSlice) - 1) / 2)
	target := intSlice[middle]

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		for _, val := range intSlice {
			if val == target {
				r = target
				break
			}
		}
	}

	intResult = r
}

func BenchmarkFindInt(b *testing.B) {
	b.StopTimer()

	var r interface{}
	intSlice := GenerateIntSlice(1000)
	middle := ((len(intSlice) - 1) / 2)
	target := intSlice[middle]

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		r = Find(intSlice, func(val interface{}) bool {
			return val.(int) == target
		})
	}

	interfaceResult = r
}

func BenchmarkNativeFindStruct(b *testing.B) {
	b.StopTimer()

	var r Point
	structSlice := GenerateStructSlice(1000)
	middle := ((len(structSlice) - 1) / 2)
	target := structSlice[middle]

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		for _, val := range structSlice {
			if val.X == target.X {
				r = target
				break
			}
		}
	}

	structResult = r
}

func BenchmarkFindStruct(b *testing.B) {
	b.StopTimer()

	var r interface{}
	structSlice := GenerateStructSlice(1000)
	middle := ((len(structSlice) - 1) / 2)
	target := structSlice[middle]

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		r = Find(structSlice, func(val interface{}) bool {
			return val.(Point).X == target.X
		})
	}

	interfaceResult = r
}

/************************************
#endregion
************************************/

/************************************
#region
Includes Benchmarks
************************************/

func BenchmarkNativeIncludesInt(b *testing.B) {
	b.StopTimer()

	var r bool
	intSlice := GenerateIntSlice(1000)
	middle := ((len(intSlice) - 1) / 2)
	target := intSlice[middle]

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		for _, val := range intSlice {
			if val == target {
				r = true
				break
			}
		}
	}

	boolResult = r
}

func BenchmarkIncludesInt(b *testing.B) {
	b.StopTimer()

	var r bool
	intSlice := GenerateIntSlice(1000)
	middle := ((len(intSlice) - 1) / 2)
	target := intSlice[middle]

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		r = Includes(intSlice, target)
	}

	interfaceResult = r
}

func BenchmarkNativeIncludesStruct(b *testing.B) {
	b.StopTimer()

	var r bool
	structSlice := GenerateStructSlice(1000)
	middle := ((len(structSlice) - 1) / 2)
	target := structSlice[middle]

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		for _, val := range structSlice {
			if val.X == target.X {
				r = true
				break
			}
		}
	}

	boolResult = r
}

func BenchmarkIncludesStruct(b *testing.B) {
	b.StopTimer()

	var r bool
	structSlice := GenerateStructSlice(1000)
	middle := ((len(structSlice) - 1) / 2)
	target := structSlice[middle]

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		r = Includes(structSlice, target)
	}

	interfaceResult = r
}

/************************************
#endregion
************************************/
