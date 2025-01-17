# Preface

In the preface of the first edition, I wrote:

> My first choice for a book title was _Boring Go_ because, properly written, Go is boring…

Boring does not mean trivial. Using Go correctly requires an understanding of how its features are intended to fit together. While you can write Go code that looks like Java or Python, you’re going to be unhappy with the result and wonder what all the fuss is about. That’s where this book comes in. It walks through the features of Go, explaining how to best use them to write idiomatic code that can grow.

Go remains a small language with a small feature set. It still lacks inheritance, aspect-oriented programming, function overloading, operator overloading, pattern matching, named parameters, exceptions, and many additional features that complicate other languages. So why does a book on a boring language need an update?

There are a few reasons for this edition. First, just as boring doesn’t mean trivial, it also does not mean unchanging. Over the past three years, new features, tools, and libraries have arrived. Improvements like structured logging, fuzzing, workspaces, and vulnerability checking help Go developers create reliable, lasting, maintainable code. Now that Go developers have several years of experience with generics, the standard library has started to include type constraints and generic functions to reduce repetitive code. Even the `unsafe` package has been updated to make it a little safer. Go developers need a resource that explains how to best use these new features.

Secondly, some aspects of Go weren’t done justice by the first edition. The introductory chapter didn’t flow as well as I’d like. The rich Go tool ecosystem wasn’t explored. And first-edition readers asked for exercises and additional sample code. This edition attempts to address those limitations.

Finally, the Go team has introduced something new, and, dare I say, exciting. There’s now a strategy that allows Go to keep the backward compatibility required for long-term software engineering projects while providing the ability to introduce backward-breaking changes to address long-standing design flaws. The new for loop variable scoping rules introduced in Go 1.22 are the first feature to take advantage of this approach.

Go is still boring, it’s still fantastic, and it’s better than ever. I hope you enjoy this second edition.

---

## Who Should Read This Book

This book is targeted at developers who are looking to pick up a second (or fifth) language. The focus is on people who are new to Go. This ranges from those who don’t know anything about Go other than it has a cute mascot, to those who have already worked through a Go tutorial or even written some Go code. The focus for _Learning Go_ isn’t just how to write programs in Go; it’s how to write Go idiomatically. More experienced Go developers can find advice on how to best use the newer features of the language. The most important thing is that the reader wants to learn how to write Go code that looks like Go.

Experience is assumed with the tools of the developer trade, such as version control (preferably Git) and IDEs. Readers should be familiar with basic computer science concepts like concurrency and abstraction, as the book explains how they work in Go. Some of the code examples are downloadable from GitHub, and dozens more can be tried out online on The Go Playground. While an internet connection isn’t required, it is helpful when reviewing executable examples. Since Go is often used to build and call HTTP servers, some examples assume familiarity with basic HTTP concepts.

While most of Go’s features are found in other languages, Go makes different tradeoffs, so programs written in it have a different structure. _Learning Go_ starts by looking at how to set up a Go development environment, and then covers variables, types, control structures, and functions. If you are tempted to skip over this material, resist the urge and take a look. It is often the details that make your Go code idiomatic. Some of what seems obvious at first glance might actually be subtly surprising when you think about it in depth.

---

## Conventions Used in This Book

The following typographical conventions are used in this book:

- **Italic**  
  Indicates new terms, URLs, email addresses, filenames, and file extensions.

- **Constant width**  
  Used for program listings, as well as within paragraphs to refer to program elements such as variable or function names, databases, data types, environment variables, statements, and keywords.

- **Constant width bold**  
  Shows commands or other text that should be typed literally by the user.

- **Constant width italic**  
  Shows text that should be replaced with user-supplied values or by values determined by context.

- **Tip**  
  This element signifies a tip or suggestion.

- **Note**  
  This element signifies a general note.

- **Warning**  
  This element indicates a warning or caution.

---

## Using Code Examples

Supplemental material (code examples, exercises, etc.) is available for download at [https://github.com/learning-go-book-2e](https://github.com/learning-go-book-2e).
