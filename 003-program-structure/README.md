# Program Structure

In this chapter, we’ll go into more detail about the basic structural elements of a Go program. The example programs are intentionally simple, so we can focus on the language without getting sidetracked by complicated algorithms or data structures.

## Names

In Go, names are used for functions, variables, constants, types, statement labels, and packages. Following a consistent naming convention is crucial for code clarity and maintainability.

### Basic Rules for Names

- Names must start with a letter (Unicode-defined) or an underscore (_).
- Subsequent characters can include letters, digits, and underscores.
- Go is case-sensitive; "heapSort" and "Heapsort" are considered different names.

### Reserved Keywords

Go has 25 keywords that are reserved and cannot be used as names. Some of these keywords include:

- break
- default
- func
- interface
- select
- case
- defer
- go
- map
- struct
- chan
- else
- goto
- package
- switch
- const
- fallthrough
- if
- range
- type
- continue
- for
- import
- return
- var

### Predeclared Names

Go includes approximately three dozen predeclared names for built-in constants, types, and functions. While these names are not reserved, it is essential to use them wisely to prevent confusion. Some of these predeclared names include:

#### Constants:
- true
- false
- iota
- nil

#### Types:
- int
- int8
- int16
- int32
- int64
- uint
- uint8
- uint16
- uint32 
- uint64
- uintptr
- float32
- float64
- complex128
- complex64
- bool
- byte
- rune
- string
- error

#### Functions:
- make
- len
- cap
- new
- append
- copy
- close
- delete
- complex
- real
- imag
- panic
- recover

### Scope and Visibility

- Entities declared within a function are local to that function.
- Declarations made outside of a function are visible in all files of the package to which they belong.
- The case of the first letter in a name determines its visibility across package boundaries.
- Names starting with an uppercase letter are exported and can be accessed outside their package.

### Naming Style

- In Go, it is a convention to use "camel case" when forming names by combining words. Interior capital letters are preferred over interior underscores.
- Acronyms and initialisms like ASCII and HTML should always be rendered in the same case for consistency.
- Examples: htmlEscape, HTMLEscape, or escapeHTML (not escapeHtml).

### Summary

In summary, adhering to Go's naming conventions and standards is essential for writing clean and maintainable code. Key takeaways include:

- Names must follow specific rules regarding characters and case sensitivity.
- Keywords and predeclared names have special significance and should be used with care.
- Scope and visibility of names depend on their declaration context.
- Consistent use of "camel case" and casing for acronyms enhances code readability.

By following these guidelines, you can ensure code consistency and make your Go programs more accessible to others.

## Declarations

A declaration names a program entity and specifies some or all of its properties. There are four major kinds of declarations: 
- var 
- const 
- type 
- func. 

For more detailed information and examples, please consult the sample programs located in the "declaration" folder.