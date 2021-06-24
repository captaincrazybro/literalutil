# literalutil
A package containing direct methods for structure literals.

For example:
```go
s := lu.String("Hello world!")
s = s.Replace("world", "golang")
```

## Current supported structure types:
 * String
 * Array
 * SArray
 * IArray

## Additional features:
 * Conditional Ternary
 * Single line If and If Else statements

### Contional Ternary example:
```go
// returning all types
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

### Single line If and If Else statements 
```go
// single line if statement
conditional := true
lu.If(conditional, func(){ fmt.Println("Hello world!") })
// Hello world!

// single line if else statement
conditional = false
lu.Ifelse(conditional, func(){ fmt.Println("Hello world!") }, func(){ fmt.Println("Hello golang!") })
// Hello golang!
```