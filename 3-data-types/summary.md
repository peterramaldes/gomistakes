* When reading existing code, bear in mind that integer literals starting with 0 are octal numbers. ALso, to improve readability, make octal integers explicit by prefixing them with 0o.
* Because integer overflows and underflows are handled silently in Go, you can implement your own functions to catch them.
* Making floating-point comparisons within a given delta can ensure that you code is portable.
* When performing addition or subtraction, group the operations with a similar order of magnitude to favor accuracy. Also, perform multiplication and division before addition and subtraction.
* Understanding the difference between slice length and capacity should be part a Go developer's core knowledge. The slice length is the number of available elements in a slice, whereas the slice capacity is the number of elements in the backing array.
* When creating a slice, initialize it with a given length or capacity if its length is already known. This reduces the number of allocations and improves performance. The same logic goes for maps, and you need to initialize their size.
* Using copy or the full slice expression is a way to prevent `append` from creating conflicts if two different functions use slices backed by the same array. However, only a slice copy prevents memory leaks if you want to shrink a large slice.
* To copy one slice to another using the `copy` built-in function, remember that the number of copied elements corresponds to the minimum between the two slice's length.
* Working with a slice of pointers or structs with pointers fields, you can avoid memory leaks by marking as nil the elements excluded by a slicing operation.
* To prevent common confusions such as when use the `encoding/json` or the `reflect` package, you need to understand the difference between nil and empty slices. Both are zero-length, zero-capacity slices, but only a nil slices doesn't require allocation.
* To check if a slice doesn't contain any element, check its length. This check works regardless of whether the slices is nil or empty. The same goes for maps.
* To design unambiguous APIs, you shouldn't distinguish between nil and empty slices.
* A map can always grow in memory, but it never shrinks. Hence, if it leads to some memory issues, you can try different options, such as forcing Go to re-create the map or using pointers.
* To compare types in Go, you can use the `==` and `!=` operators if two types are comparable: Booleans, numerals, strings, pointers, channels, and structs are composed entirely of comparable types. Otherwise, you can either use `reflect.DeepEqual` and pay the price of reflection or use custom implementations and libraries.

Reference: 
Harsanyi, Teiva. 100 Go Mistakes and How to Avoid Them (p. 205). Manning. Kindle Edition.