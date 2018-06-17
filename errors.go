package refutil

import "errors"

// ErrArgumentNotSlice is telling you that argument supplied to function
// is not a slice or array
var ErrArgumentNotSlice = errors.New("argument is not slice")
