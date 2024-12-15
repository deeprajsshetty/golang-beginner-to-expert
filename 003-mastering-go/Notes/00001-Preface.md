# Mastering Go - Fourth Edition

Welcome to the fourth edition of _Mastering Go_! As the Go programming language continues to evolve and gain popularity, I am delighted to present this updated edition of the book.

In the rapidly changing landscape of technology, Go has emerged as a language of choice for building scalable, performant, and maintainable software. Whether you are a seasoned Go developer looking to deepen your expertise or a newcomer eager to master the intricacies of the language, this book is your comprehensive guide.

## Preface

If you have an older edition of _Mastering Go_, apart from the first edition (which is quite outdated), don’t throw it away just because the fourth edition is out. Go has not changed significantly, and both the second and third editions remain useful and relevant. However, the fourth edition is better in many ways, including up-to-date information on the latest Go versions, which you will not find in the previous editions.

This edition covers exciting topics such as writing RESTful services, building a statistics application, and working with eBPF. It also includes a brand new chapter on fuzz testing and observability.

I have aimed to strike the right balance between theory and hands-on content, but only you, the reader, can tell if I have succeeded. I encourage you to complete the exercises at the end of each chapter and share feedback to help improve future editions.

Thank you for choosing _Mastering Go_, Fourth Edition. Let's dive in and unlock the full potential of Go together. Happy coding!

## Who This Book Is For

This book is aimed at:

- Amateur and intermediate Go programmers who want to take their Go knowledge to the next level.
- Developers from other programming languages who want to learn Go.

If this is your first programming book, you may face some challenges in following the content. A second reading may be required to fully absorb the material.

Learning by doing is a key principle in mastering any programming language. The book includes practical examples and exercises to help you apply key concepts in real-world scenarios.

Expect to work, learn, and fail — then work, learn, and fail some more. After all, life and programming share many similarities.

## What This Book Covers

- **Chapter 1: A Quick Introduction to Go**

  - History and advantages of Go
  - Basic Go programming concepts and creating a simple statistics application

- **Chapter 2: Basic Go Data Types**

  - Numeric and non-numeric data types
  - Arrays, slices, pointers, constants, and random data generation

- **Chapter 3: Composite Data Types**

  - Working with maps, structures, and CSV files
  - Data persistency in the statistics application

- **Chapter 4: Go Generics**

  - Writing generic functions and defining generic data types

- **Chapter 5: Reflection and Interfaces**

  - Reflection, interfaces, and type methods
  - Error handling, type assertions, and type switches

- **Chapter 6: Go Packages and Functions**

  - Creating Go packages, modules, and using `defer`
  - Creating a Go package for SQLite3 and exploring Go Workspaces

- **Chapter 7: Telling a UNIX System What to Do**

  - Working with command-line arguments, signals, and file I/O
  - Converting the statistics application into a command-line utility

- **Chapter 8: Go Concurrency**

  - Goroutines, channels, pipelines, and the sync package
  - Detecting race conditions and managing goroutine timeouts

- **Chapter 9: Building Web Services**

  - Web servers, clients, and working with HTTP timeouts
  - Creating a web service version of the statistics application

- **Chapter 10: Working with TCP/IP and WebSocket**

  - TCP/IP protocols, WebSocket, and RabbitMQ
  - Developing practical servers and clients

- **Chapter 11: Working with REST APIs**

  - Defining and developing REST APIs
  - Building concurrent RESTful servers and command-line utilities

- **Chapter 12: Code Testing and Profiling**

  - Code testing, profiling, and optimization techniques
  - Using `go:generate` and identifying unreachable code

- **Chapter 13: Fuzz Testing and Observability**

  - Introduction to fuzz testing and observability tools

- **Chapter 14: Efficiency and Performance**

  - Benchmarking code and understanding the Go memory model
  - Working with eBPF for observability and performance

- **Chapter 15: Changes in Recent Go Versions**

  - Language changes, new features, and improvements

- **Appendix: The Go Garbage Collector**
  - Understanding the Go Garbage Collector and its impact on performance

## To Get the Most Out of This Book

This book requires:

- A modern computer with a recent version of Go installed.
- Familiarity with your operating system, filesystem, and `git`.

Most of the presented code runs on Mac OS X, Linux, and Microsoft Windows without modifications.

As you learn, I encourage you to experiment, ask questions, and engage actively with the material. Go offers a unique blend of simplicity and power, and I am confident this book will help you become a proficient Go developer.

## Download the Example Code Files

The code bundle for this book is hosted on GitHub at:  
[https://github.com/mactsouk/mGo4th](https://github.com/mactsouk/mGo4th)

Other code bundles from Packt Publishing are available here:  
[https://github.com/PacktPublishing/](https://github.com/PacktPublishing/)

## Download the Color Images

For a PDF containing color images of the screenshots and diagrams used in this book, you can download it here:  
[https://packt.link/gbp/9781805127147](https://packt.link/gbp/9781805127147)

## Conventions Used

Throughout the book, the following conventions are used:

- **CodeInText**: Indicates code words in text, filenames, paths, user input, etc.
  - Example: `Mount the downloaded WebStorm-10*.dmg disk image file...`
- **Code Blocks**: Code blocks are displayed as follows:

  ```go
  package main
  import "fmt"

  func doubleSquare(x int) (int, int) {
      return x * 2, x * x
  }
  ```
