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

/************************************
#region
Each Benchmarks
************************************/

func BenchmarkNativeEachInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		intSlice := GenerateIntSlice(10000)

		b.StartTimer()
		for i, v := range intSlice {
			DoNothing(i, v)
		}
	}
}

func BenchmarkEachInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		intSlice := GenerateIntSlice(10000)

		b.StartTimer()
		Each(intSlice, func(i interface{}, v interface{}) {
			DoNothing(i, v)
		})
	}
}

func BenchmarkNativeEachStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		structSlice := GenerateStructSlice(10000)

		b.StartTimer()
		for i, v := range structSlice {
			DoNothing(i, v)
		}
	}
}

func BenchmarkEachStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		structSlice := GenerateStructSlice(10000)

		b.StartTimer()
		Each(structSlice, func(i interface{}, v interface{}) {
			DoNothing(i, v)
		})
	}
}

/************************************
#endregion
************************************/
