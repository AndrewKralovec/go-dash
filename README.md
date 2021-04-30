# Go Dash
Experimental go lang lodash implementation. Reflection is heavily used.

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

