## Bounds Check Elimination
Go is a memory safe language. In array/slice element indexing and subslice operations, Go runtime will check whether or not the involved indexes are out of range. If an index is out of range, a panic will be produced to prevent the invalid index from doing harm. This is called bounds check. Bounds checks make our code run safely, on the other hand, they also make our code run a little slower.

Since Go Toolchain 1.7, the standard Go compiler has used a new compiler backend, which based on SSA (static single-assignment form). SSA helps Go compilers effectively use optimizations like BCE (bounds check elimination) and CSE (common subexpression elimination). BCE can avoid some unnecessary bounds checks, and CSE can avoid some duplicate calculations, so that the standard Go compiler can generate more efficient programs. Sometimes the improvement effects of these optimizations are obvious.

## run sample file 00.go
``` 
go run -gcflags="-d=ssa/check_bce/debug=1" 00.go
```

## output 
``` 
./00.go:4:7: Found IsInBounds
./00.go:5:7: Found IsInBounds
./00.go:6:7: Found IsInBounds
./00.go:10:7: Found IsInBounds
./00.go:16:7: Found IsInBounds
./00.go:21:7: Found IsInBounds
```

## Common patterns
* Loop upper bound is known
* Contiguous slice access
* Slice accesses offset from a base index
* Iterate multiple slices with known upper bound
* Slice range can be derived