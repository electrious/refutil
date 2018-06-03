# Electrious Reflect Util

`reflect` functionality we often use in our projects.
Who knows maybe you will find it useful as well.

## Package `refutil`

### `IsEqual(a, b interface{})`

this method is intended to compare 2 interfaces if they are same or not.
Stolen from [Testify Assert](https://github.com/stretchr/testify/tree/master/assert).
Difference between `reflect.DeepEqual` and `refutil.IsEqual` is that `reflect.IsEqual`
dont care about underlying type as long as types are convertible.

```go
refutil.IsEqual(uint(1), 1) // true
```

of course this is possible with any type, not just numeric ones.

### `Len(a interface{})`

Len tell length of object and return if possible to get len on object

```go
refutil.Len([]uint{1}) // returns 1, true
refutil.Len("hello") // returns 2, true
refutil.Len(struct{}{}) // returns 0, false - you cant really get len of struct
```

### `Nil(a interface {})`

returns if object is nil or not. Look at example why you want to use it.

```go
// Your custom error
type MyOwnError struct {}
func(err *MyOwnError) Error() string {return ""}

// Create new error of error interface
err := errors.New("ok")
err = (*MyOwnError)(nil)
// note err is still error interface
err == nil // false
refutil.IsNil(err) // true
```

### `PathTo(a interface {})`

With `reflect.PkgPath(...)` it not works with when value is pointer.

```go
type K struct{}
refutil.PathTo(&K{}) // returns "github.com/electrious/refutil"
```

### `Dereference(a interface {})`

Will get value of `interface{}` if `interface{}` is pointer to something.

```go
type K struct{}
k := K{}
refutil.Dereference(&k) == refutil.Dereference(k) // true
refutil.Dereference(nil) == refutil.Dereference((*K)(nil)) // true
```

### `Index(a, b interface{})`

Index will return index of element in array, map, string or any other searchable.
It uses Equal. First argument return index if not found return `-1` second argument
return if source is searchable.

```go
refutil.Index([]uint{1,2,3}, 2) // returns 2, true
refutil.Index(struct{}{}, 2) // returns -1, false
```

### `IndexSame(a, b interface{})`

Just like Index except types needs to be equal to make this work.

```go
refutil.Index([]uint{1,2,3}, uint(2)) // returns 2, true
refutil.Index([]uint{1,2,3}, 2) // returns -1, true
```

### `Contains(a, b interface{})`

Just like Index but if result not found (eg. result of index `-1`) then return false instead.

```go
refutil.Contains([]uint{1,2,3}, 2) // returns found: true, ok: true
refutil.Contains(struct{}{}, 2) // returns false, false
```

### `ContainsSame(a, b interface{})`

Just like IndexSame and same as contains.

```go
refutil.ContainsSame([]uint{1,2,3}, uint(2)) // returns 2, true
refutil.ContainsSame([]uint{1,2,3}, 2) // returns false, true
```


### `IsZero(a interface{})`

Will check if underlying value is Zero. For any kind of type. Check zero_test.go to see
behavior in action.

```go
var header []string
refutil.IsZero(uint(2)) // returns false
refutil.IsZero(uint(0)) // returns true
refutil.IsZero(nil) // true
refutil.IsZero((*time.Time)(nil)) // true
refutil.IsZero(header) // true
refutil.IsZero([]string{}) // true
refutil.IsZero([]string{"1"}) // false
// and os on
```

## Package `structs`

Package structs is copied from great [https://github.com/fatih/structs](https://github.com/fatih/structs).
It is using `refutil` package to determine few things and to not repeat self too much.
Otherwise almost same.

```go
type Server struct {
	Name        string `json:"name,omitempty"`
	ID          int
	Enabled     bool
	users       []string // not exported
	http.Server          // embedded
}

server := &Server{
	Name:    "gopher",
	ID:      123456,
	Enabled: true,
}
// Convert a struct to a map[string]interface{}
// => {"Name":"gopher", "ID":123456, "Enabled":true}
m := structs.Map(server)

// Convert the values of a struct to a []interface{}
// => ["gopher", 123456, true]
v := structs.Values(server)

// Convert the names of a struct to a []string
// (see "Names methods" for more info about fields)
n := structs.Names(server)

// Convert the values of a struct to a []*Field
// (see "Field methods" for more info about fields)
f := structs.Fields(server)

// Return the struct name => "Server"
n := structs.Name(server)

// Check if any field of a struct is initialized or not.
h := structs.HasZero(server)

// Check if all fields of a struct is initialized or not.
z := structs.IsZero(server)

// Check if server is a struct or a pointer to struct
i := structs.IsStruct(server)
```


## Package `unsafe`

Package unsafe provide simple functionality to bypass runtime check for private fields.
It will get raw pointer so you can access anything you want.
Very useful for debugging. But with this great power comes great responsibility.
It is called `unsafe` for purpose.

Originally invented in [go-spew](https://github.com/davecgh/go-spew).

### `unsafe.ReflectValue`

```go
type K struct {
    a int
}
v := unsafe.ReflectValue(reflect.ValueOf(&K))
// v is now reflect unsafe value
```

## Credit

- Richard Hutta
- Mat Ryer
- Tyler Bunnell
- Fatih Arslan

Thank you guys for your great work.
