# Go Dash
Experimental go lang lodash implementation. 
Reflection is heavily used in the first iteration, leading some functions to be vastly out performed by their native counter parts. See examples under [Benchmarks](#benchmarks)

## Methods

* [`Each`](#each)
* [`Filter`](#filer)
* [`Find`](#find)
* [`First`](#first)
* [`Includes`](#includes)
* [`Last`](#last)

### `Each`

Iterates over elements of a collection, executing the callback for each element. The callback is invoked with two arguments; (index, value|key, value). 

```go
_.Each([]string{"peach", "apple", "pear", "plum"}, func(i interface{}, v interface{}) {
	fmt.Println(v)
})
// => logs "peach", "apple", "pear", "plum"
_.Each([]Point{{0, 0}, {1, 1}}, func(k interface{}, v interface{}) {
    fmt.Println(k, v)
})
// => logs 00, 11
```

### `Filter` 

Iterates over elements of a collection, returning an array of all elements the callback returns true for. The callback is invoked with one arguments; (value).  
Returns a new array of elements that passed the callback check. 

```go
f := _.Filter([]string{"peach", "apple", "pear", "plum"}, func(v interface{}) {
    return v.(string) == "peach"
})
// => returns "peach" (<[]interface{}(string)>)
```

### `Find` 

Iterates over elements of a collection, returning the first element that the callback returns truey for. The callback is invoked with one arguments; (value).  
Returns the found element, else nil.

```go
f := _.Find([]string{"peach", "apple", "pear", "plum"}, func(v interface{}) {
    return v.(string) == "peach"
})
// => returns "peach" (<interface{}(string)>)
```

### `First`

Gets the first element of array.
Returns the first element of array or nil if provided an empty array.

```go
f := _.First([]string{"peach", "apple", "pear", "plum"})
// => returns "peach" (<interface{}(string)>)
```
### `Includes`

Checks if value is in collection.
Returns true if value is found, else false.

```go
i := _.Includes([]string{"peach", "apple", "pear", "plum"}, "peach")
// => returns true (bool)
```

### `Last`

Gets the last element of array.
Returns the last element of array or nil if provided an empty array.

```go
l := _.Last([]string{"peach", "apple", "pear", "plum"})
// => returns "plum" (<interface{}(string)>)
```

## Benchmarks

* [`Native loop vs Each`](#native-loop-vs-each)
* [`Native searching vs Filter`](#native-searching-vs-filter)
* [`Native searching vs Find`](#native-searching-vs-find)

### `Native loop vs Each`

Native loops is on average 1.5x-1.6x faster.

```shell
BenchmarkNativeEachInt-16                   4092            299358 ns/op          157953 B/op      19744 allocs/op
BenchmarkEachInt-16                         2482            473850 ns/op          157978 B/op      19745 allocs/op

BenchmarkNativeEachStruct-16                3157            363256 ns/op          237954 B/op      19744 allocs/op
BenchmarkEachStruct-16                      2664            543485 ns/op          237978 B/op      19745 allocs/op
```

### `Native searching vs Filter`

Native search is on average 40x-60x faster.

```shell
BenchmarkNativeFilterInt-16              1980960               606.3 ns/op             8 B/op          1 allocs/op
BenchmarkFilterInt-16                      32916             36629 ns/op            8040 B/op       1002 allocs/op

BenchmarkNativeFilterStruct-16           1358168               882.1 ns/op            16 B/op          1 allocs/op
BenchmarkFilterStruct-16                   29234             41097 ns/op           16040 B/op       1002 allocs/op
```

### `Native searching vs Find`

Native search is on average 70x-120x faster.

```shell
BenchmarkNativeFindInt-16                8080290               148.6 ns/op             0 B/op          0 allocs/op
BenchmarkFindInt-16                        65679             18234 ns/op            4024 B/op        501 allocs/op

BenchmarkNativeFindStruct-16             4122471               290.0 ns/op             0 B/op          0 allocs/op
BenchmarkFindStruct-16                     58962             20312 ns/op            8024 B/op        501 allocs/op
```