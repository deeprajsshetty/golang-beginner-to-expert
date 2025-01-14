# Installing the Go Tools

To build Go code, you need to download and install the Go development tools. The latest version of the tools can be found at the [downloads page on the Go website](https://go.dev/dl/).

## Installation Instructions

### Choose the download for your platform and install it.
    On Mac, use the .pkg installer.
    On Windows, use the .msi installer.

   These installers automatically install Go in the correct location, remove any old installations, and add the Go binary to the default executable path.

   **Tip:** If you are a Mac developer, you can also install Go using Homebrew with the command `brew install go`. Windows developers using Chocolatey can use `choco install golang`.

### For Linux and BSD users:
    1. Download the gzipped tar file and expand it to a directory named `go`.
    2. Copy this directory to `/usr/local`.
    3. Add `/usr/local/go/bin` to your `$PATH` so that the `go` command is accessible.

   ```shell
   $ tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
   $ echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.bash_profile
   $ source $HOME/.bash_profile
   ```
   **Note:** You might need root permissions to write to /usr/local. If the tar command fails, rerun it with sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz.

### Validate your Go environment setup by opening a terminal or command prompt and typing:

   ```shell
   $ go version
   ```
   If everything is set up correctly, you should see output similar to this:
   
   macOS example:
   - go version go1.20.5 darwin/arm64

   x64 Linux example:
   - go version go1.20.5 linux/amd64

### Troubleshooting Your Go Installation
   - If you encounter an error instead of the version message, it's likely that go is not in your executable path, or you have another program named go in your path. On Unix-like systems, you can use which go to see the go command being executed, if any. If nothing is returned, you need to fix your executable path.
   - If you're on Linux or BSD, ensure that you have installed the correct Go development tools for your system architecture (32-bit or 64-bit).

## Go Tooling

All of the Go development tools are accessed via the go command. In addition to go version, there’s a compiler (go build), code formatter (go fmt), dependency manager (go mod), test runner (go test), a tool that scans for common coding mistakes (go vet) and more.

For now, let’s take a quick look at the most commonly used tools by writing everyone’s favorite first application: Hello World.

### Your First Go Program

Let’s go over the basics of writing a Go program. Along the way, we will see the parts that make up a simple Go program. 

**Making a Go Module**

- The first thing we need to do is create a directory to hold our program. Call it **my-first-program**.
- On the command line, enter the new directory. If your computer’s terminal uses bash or zsh, this looks like:
   ```shell
   $ mkdir my-first-program
   $ cd my-first-program
   ```
- Inside the directory, run the go mod init command to mark this directory as a Go module:
   ```shell
   $ go mod init hello_world
   go: creating new go.mod: module hello_world
   ```
- Go project is called a module. A module is not just source code. It is also an exact specification of the dependencies of the code within the module. Every module has a go.mod file in its root directory.
- Running go mod init creates this file for us. The contents of a basic go.mod file look like this:
   ```
   module hello_world

   go 1.20
   ```
- The go.mod file declares the name of the module, the minimum supported version of Go for the module, and any other modules that your module depends on.
- You shouldn’t edit the go.mod file directly. Instead, use the go get and go mod tidy commands to manage changes to the file.

**go build**

- Now let's write some code! Open up a text editor, enter the following text, and save it inside my-first-program with the file name hello.go:
   ```
   package main

   import "fmt"

   func main() {
   fmt.Println("Hello, world!")
   }
   ```
- Let’s quickly go over the parts of the Go file we created. The first line is a package declaration. Within a Go module, code is organized into one or more packages. The main package in a Go module contains the code that starts a Go program.
- Next there is an import declaration. The import statement lists the packages referenced in this file. We’re using a function in the fmt (usually pronounced “fumpt”) package from the standard library, so we list the package here. Unlike other languages, Go only imports whole packages. You can’t limit the import to specific types, functions, constants, or variables within a package.
- All Go programs start from the main function in the main package. We declare this function with func main() and a left brace. Like Java, JavaScript, and C, Go uses braces to mark the start and end of code blocks.
- The body of the function consists of a single line. It says that we are calling the Println function in the fmt package with the argument "Hello, world!".
- After the file is saved, go back to your terminal or command prompt and type:
   ```shell
   $ go build
   ```
- This creates an executable called hello_world (or hello_world.exe on Windows) in the current directory.
- Run it and you will unsurprisingly see Hello, world! printed on the screen:
   ```shell
   $ ./hello_world
   Hello, world!
   ```
- The name of the binary matches the name in the module declaration. If you want a different name for your application, or if you want to store it in a different location, use the -o flag.
- For example, if we wanted to compile our code to a binary called “hello,” we would use:
   ```shell
   $ go build -o hello
   ```
**go fmt**

- The Go development tools include a command, go fmt, which automatically fixes the whitespace in your code to match the standard format. However, it can’t fix braces on the wrong line. Run it with:
   ```shell
   $ go fmt ./...
   hello.go
   ```
- Using ./... tells a Go tool to apply the command to all the files in the current directory and all subdirectories.
- The go fmt command won’t fix braces on the wrong line because of the semicolon insertion rule. Like C or Java, Go requires a semicolon at the end of every statement. However, Go developers should never put the semicolons in themselves. The Go compiler adds them automatically following a very simple rule described in https://go.dev/doc/effective_go#semicolons.
- The semicolon insertion rule and the resulting restriction on brace placement is one of the things that makes the Go compiler simpler and faster, while at the same time enforcing a coding style. 

**go vet**

- There is a class of bugs where the code is syntactically valid, but quite likely incorrect. The go tool includes a command called go vet to detect these kinds of errors.
- Let’s add one to our program and watch it get detected. Change the fmt.Println line in hello.go to the following:
   ```
   fmt.Printf("Hello, %s!\n")
   ```
- Here, we have a template ("Hello, %s!\n") with a %s placeholder, but no value was specified for the placeholder. This code will compile and run, but it’s not correct.
- One of the things that go vet can detect is whether or not there is a value for every placeholder in a formatting template.
- Run go vet on our modified code and it finds an error:
   ```shell
   $ go vet ./...
   # hello_world
   ./hello.go:6:2: fmt.Printf format %s reads arg #1, but call has 0 args
   ```
- Now that go vet found our bug, it’s easy for us to fix it. Change line 6 in hello.go to:
   ```
   fmt.Printf("Hello, %s!\n", "world")
   ```
- While go vet catches several common programming errors, there are things that it cannot detect. Luckily, there are third-party Go code-quality tools to close the gap.

**go run**

- Compile and run a Go program in a single step.
- Useful for quick testing or running small scripts.
   ```shell
   $ go run main.go
   ```

**go doc**

- Generate documentation for Go packages.
- Helpful for understanding the usage of functions, types, and packages in your codebase.
   ```shell
   $ go doc [package]
   ```
  
*Remember to keep your Go tools and packages up-to-date by periodically running **go get -u** and **go mod tidy** to ensure a smooth development experience.*

### Go Development Environments and IDEs

While it's possible to write Go programs using a simple text editor and the go command, larger projects benefit from advanced development tools and integrated development environments (IDEs). These tools provide features such as automatic code formatting, code completion, type checking, error reporting, and debugging capabilities. Here, we'll discuss two popular Go development environments: Visual Studio Code and GoLand.

#### Visual Studio Code (VS Code)

**Overview:**
- Visual Studio Code, often referred to as VS Code, is a lightweight and highly extensible code editor developed by Microsoft.
- It has a vast and active community, making it one of the most popular choices for Go development.

**Advantages:**

- **Extensions:** VS Code offers a rich ecosystem of extensions. For Go development, the "Go" extension by Microsoft is widely used.
- **Automatic Formatting:** VS Code can be configured to automatically format your Go code on save using the *gofmt* tool.
- **Code Completion:** It provides intelligent code completion, helping you write code more efficiently.
- **Integrated Terminal:** VS Code has a built-in terminal for running *go* commands directly from the editor.
- **Debugging:** Integrated debugging support allows you to set breakpoints and inspect variables.

**Installation:**
- Install VS Code from the official website.
- Install the "Go" extension from the VS Code marketplace.

**Usage:**
- Open your Go project folder in VS Code.
- Configure your Go environment, including the Go SDK and GOPATH.
- Start coding, and the Go extension will provide helpful features like code completion and formatting.

#### GoLand

**Overview:**

- GoLand is a commercial IDE developed by JetBrains, known for its IntelliJ IDEA.
- It's specifically designed for Go development, providing a feature-rich and highly specialized environment.

**Advantages:**

- **Built-In Go Tools:** GoLand includes built-in support for Go tools like go fmt, go vet, and go test.
- **Code Analysis:** Offers extensive code analysis, identifying potential issues and suggesting improvements.
- **Refactoring:** Supports powerful code refactoring tools to help you maintain a clean codebase.
- **Version Control:** Integrated version control support for Git, Mercurial, and more.
- **Code Inspection:** Provides detailed insights into your code's health and performance.

**Installation:**

- Download and install GoLand from the JetBrains website.

**Usage:**

- Open your Go project in GoLand.
- Configure your Go SDK and project settings.
- Leverage the built-in tools and features for efficient Go development.

**Conclusion**

Both Visual Studio Code and GoLand are excellent choices for Go development. Your choice may depend on your preference, project requirements, and budget. VS Code is free and highly customizable, while GoLand offers an extensive and specialized Go development environment.

### Go Compatibility and Development Updates

**Go Development Updates**

Go's development tools have seen regular updates since Go 1.2, with new releases occurring approximately every six months. These releases include patch updates for addressing bugs and security vulnerabilities.

**Incremental Releases**

Go's release strategy is characterized by incremental improvements rather than radical changes. New versions typically introduce small features and enhancements.

**Go Compatibility Promise**

The Go Compatibility Promise is a fundamental commitment by the Go team. It guarantees that there will be no backward-breaking changes to the language or standard library for any Go version starting with 1. Exceptions are made only when changes are necessary for critical bug fixes or security updates.

**Prioritizing Compatibility**

In a keynote talk at GopherCon 2022, Russ Cox emphasized the significance of prioritizing compatibility as a core design decision for Go 1. This commitment ensures that Go programs continue to work even as the language evolves.

**Exception for go Commands**

It's important to note that the Go Compatibility Promise does not extend to the `go` commands. There have been instances of backward-incompatible changes to flags and functionality within these commands, and similar changes may occur in the future.

**Conclusion**

Go's commitment to backward compatibility, along with its incremental development approach, provides stability and confidence to Go developers. While there may be changes in `go` commands, the core language and standard library remain reliable and consistent across versions.

For more details, refer to official Go documentation and release notes.
