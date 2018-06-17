package refutil

import "errors"

// ErrArgumentNotSlice is telling you that argument supplied to function
// is not a slice or array
var ErrArgumentNotSlice = errors.New("argument is not slice")

// ErrArgumentNotIndexable is telling you that argument supplied to function
// is not a searchable for specific index
var ErrArgumentNotIndexable = errors.New("argument can't be used to search index")
