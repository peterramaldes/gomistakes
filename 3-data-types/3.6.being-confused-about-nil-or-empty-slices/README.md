* `var s []string` if we aren't sure about the final length and the slice can be empty.
* `[]string(nil)` as syntactic sugar to create a nil and empty slice.
* `make([]string, length)` if the future length is known.