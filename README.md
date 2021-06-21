# literalutil
A package containing direct methods for structure literals.

For example:
```go
s := lu.String("Hello world!")
s = s.Replace("world", "golang")
```

## Current supported structure types:
 * string
 * Array

## Additional features:
 * Conditional ternary

### Contional ternary example:
```go
// returning all types...
conditional := true
fmt.Printf("Output: %v\n", lu.Ternary(conditional, Array{1, 2, 3}, Array{"a", "b", "c"}))
// Output: [1, 2, 3]

// returning a string
conditional = false
fmt.Printf("Output: %v\n", lu.STernary(conditional, "Hello world!", "Hello golang!"))
// Output: Hello golang!

// returning an integer
conditional = true
fmt.Printf("Output: %v\n", lu.ITernary(conditional, 1, 2))
// Output: 1
```
