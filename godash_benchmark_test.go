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

// avoid compiler optimisations
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
