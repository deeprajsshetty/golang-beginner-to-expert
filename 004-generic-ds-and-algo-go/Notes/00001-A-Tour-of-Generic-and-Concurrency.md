## Chapter Overview

This chapter introduces the syntax and semantics of generics in Go. Many coding examples are presented that illustrate this new and powerful feature of Go. This sets the stage for the continued use of generics throughout the book.

Concurrency in Go is also reviewed in this chapter. Many coding examples are presented along with benchmarks that contrast the performance of algorithms with and without concurrency. This also sets the stage for the continued use of concurrency throughout the book.

# 1.1 Brief History and Description of Go

Go is a relatively new programming language released in late 2009 and developed at Google by Robert Griesemer (a Swiss computer scientist who helped create Google’s V8 JavaScript engine), Rob Pike (a Canadian computer scientist and part of the Unix team at Bell Labs and creator of the Limbo programming language), and Ken Thompson (creator of the Unix operating system and the B programming language).

The Go programming language is sometimes called Golang. Why? The domain “go.org” wasn’t available at the time the language was released, so golang.org (a mix of Go and language) was born. The official name of the language is Go, but the Twitter tag is #golang. Go figure!

One of the major goals in creating Go was to produce an easily readable, strong, and statically typed language with garbage collection and fast compilation and execution speed particularly suited for concurrent applications.

The **goroutine** is a lightweight process that requires less memory overhead than a normal thread seen in other languages such as Java and C#. A concurrent Go program may spawn thousands of goroutines running on a much smaller number of threads. The **channel** construct (to be explained later in this chapter) allows information to be passed into and taken out of goroutines and is used to synchronize these concurrent lightweight processes. Although parallel processing is not the primary objective of goroutines, they can be used to approximate this on a shared memory multiprocessor computer.

Go is a platform-independent language that runs on various Unix platforms including MacOS and also runs on MS Windows. Go applications compile to a binary executable so they can be distributed to a customer without having to package an interpreter and runtime libraries as is the case with Python and other interpreted languages.

Go, like many recent languages, is a public open source project. There are a bevy of free tools that are downloadable. New packages are constantly being released, so much of the power of the language resides outside the language in the plethora of high-quality packages available to the programmer. In this sense, Go is like Python.

Among the tools that are available are high-quality editors, debuggers, and IDEs. Go requires a prescribed format, so the `gofmt` tool is often integrated into various code editors. Having a standard code format provides a huge advantage to Go programming teams as well as solo programmers inspecting the code written by others.

## What’s Missing in Go?

So what is missing in Go? What is its downside? Up until the most recent and perhaps most important new release, Version 1.18, Go lacked genericity. With this new release of Go, this major shortcoming is gone.

Now one can build an algorithm or data structure that does not have to be modified every time the underlying information to be stored changes. Data structure and algorithm implementations can focus on the core logic needed to manipulate the information. A new syntax associated with generics allows a programmer to precisely describe the requirements that data must satisfy to be stored in a particular data structure. This furthers a programmer’s ability to have a program specify its intent in the code itself.

The use of constrained and unconstrained generic parameters is introduced and illustrated in the next section.

# 1.2 Introducing Generic Parameters

In this section, we present a series of examples that introduce and illustrate the use of generic-type parameters, both unconstrained and constrained.

In the first several code listings, we present a set of related problems of adding a new student to an existing slice of students. First, we add just the name of the student to our existing slice. Next, we add the student’s ID number to a slice containing ID numbers. Next, we add a struct containing name, ID, and age to an existing slice of structs. Then finally, we bring generics on stage and show the simplification that is achievable using a generic-type parameter.

## Adding a New Student by Name

Consider the simple Go application given in Listing 1-1.

```
package main
import(
    "fmt"
)
func addStudent(students []string, student string) []string {
    return append(students, student)
}
func main() {
    students := []string{} // Empty slice
    result := addStudent(students, "Michael")
    result = addStudent(result, "Jennifer")
    result = addStudent(result, "Elaine")
    fmt.Println(result)
}
/*
Output:
[Michael Jennifer Elaine]
*/

Listing 1-1A slice of students

```

The function `addStudent` takes a slice of string representing the current collection of students as the first parameter and a string representing a new student to be added to the collection as the second parameter. The `append` function is used to add the new student to the existing slice, and that result is returned.

## Adding a New Student by ID Number

Suppose we wish to specify the slice of students using their ID number, an `int`, instead of their name, a `string`.

    We would need to modify Listing 1-1 as shown in Listing 1-2.

```
package main
import(
    "fmt"
)
func addStudent(students []string, student string) []string {
    return append(students, student)
}
func addStudentID(students []int, student int) []int {
    return append(students, student)
}
func main() {
    students := []string{} // Empty slice
    result := addStudent(students, "Michael")
    result = addStudent(result, "Jennifer")
    result = addStudent(result, "Elaine")
    fmt.Println(result)
    students1 := []int{} // Empty slice
    result1 := addStudentID(students1, 155)
    result1 = addStudentID(result1, 112)
    result1 = addStudentID(result1, 120)
    fmt.Println(result1)
}
/* Output
[Michael Jennifer Elaine]
[155 112 120]
*/

Listing 1-2Adding student IDs

```

The logic in function `addStudentID` is essentially the same as in function `addStudent`. Only the base type of the slice is changed from `string` to `int`.

## Adding a New Student by Student Struct

And to take this one step further, suppose we define a `Student` type as:

```
type Student struct {
    Name string
    ID int
    age float64
}

```

    and we modify Listing 1-2 as shown in Listing 1-3.

```

package main
import(
    "fmt"
)
type Student struct {
    Name string
    ID int
    age float64
}
func addStudent(students []string, student string) []string {
    return append(students, student)
}
func addStudentID(students []int, student int)  []int {
    return append(students, student)
}
func addStudentStruct(students []Student, student Student) []Student {
    return append(students, student)
}
func main() {
    students := []string{} // Empty slice
    result := addStudent(students, "Michael")
    result = addStudent(result, "Jennifer")
    result = addStudent(result, "Elaine")
    fmt.Println(result)
    students1 := []int{} // Empty slice
    result1 := addStudentID(students1, 155)
    result1 = addStudentID(result1, 112)
    result1 = addStudentID(result1, 120)
    fmt.Println(result1)
    students2 := []Student{} // Empty slice
    result2 := addStudentStruct(students2, Student{"John", 213, 17.5} )
    result2 = addStudentStruct(result2,  Student{"James", 111, 18.75} )
    result2 = addStudentStruct(result2,  Student{"Marsha", 110, 16.25} )
    fmt.Println(result2)
}
/* Output
[Michael Jennifer Elaine]
[155 112 120]
[{John 213 17.5} {James 111 18.75} {Marsha 110 16.25}]
*/

Listing 1-3Adding Student type to the mix

```

Having to add a new function each time we wish to add a new underlying data type to our various student collections is tedious and a major downside to earlier versions of Go.
