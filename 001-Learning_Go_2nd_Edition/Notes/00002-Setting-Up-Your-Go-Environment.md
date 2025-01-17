# Chapter 1. Setting Up Your Go Environment

Every programming language needs a development environment, and Go is no exception. If you’ve already built a Go program or two, then you have a working environment, but you might have missed out on some of the newer techniques and tools. If this is your first time setting up Go on your computer, don’t worry; it’s easy to install Go and its supporting tools. After setting up the environment and verifying it, you will build a simple program, learn about the different ways to build and run Go code, and then explore some tools and techniques that make Go development easier.

---

## Installing the Go Tools

To build Go code, you need to download and install the Go development tools. You can find the latest version of the tools at the [downloads page](https://golang.org/dl/) on the Go website. Choose the download for your platform and install it. The `.pkg` installer for Mac and the `.msi` installer for Windows automatically install Go in the correct location, remove any old installations, and put the Go binary in the default executable path.

### Tip

If you are a Mac developer, you can install Go using Homebrew with the command:

```bash
brew install go
```

Windows developers who use Chocolatey can install Go with the command:

```bash
choco install golang
```

The various Linux and BSD installers are gzipped TAR files and expand to a directory named `go`. Copy this directory to `/usr/local` and add `/usr/local/go/bin` to your `$PATH` so that the `go` command is accessible:

```bash
$ tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
$ echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.bash_profile
$ source $HOME/.bash_profile

```

You might need root permissions to write to `/usr/local`. If the `tar` command fails, rerun it with:

```bash
sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
```

Go programs compile to a single native binary and do not require any additional software to be installed in order to run them. This is in contrast to languages like Java, Python, and JavaScript, which require you to install a virtual machine to run your program. Using a single native binary makes it a lot easier to distribute programs written in Go.

This book doesn’t cover containers, but developers who use Docker or Kubernetes can often package a Go app inside a scratch or distroless image. You can find details in Geert Baeke’s blog post [downloads page](https://blog.baeke.info/2021/03/28/distroless-or-scratch-for-go-apps/)

You can validate that your environment is set up correctly by opening up a terminal or command prompt and typing:

```bash
$ go version
```

If everything is set up correctly, you should see something like this printed:

```bash
go version go1.20.5 darwin/arm64
```

This tells you that this is Go version 1.20.5 on macOS. (Darwin is the operating system at the heart of macOS, and arm64 is the name for the 64-bit chips based on ARM’s designs.) On x64 Linux, you would see:

```bash
go version go1.20.5 linux/amd64
```

# Troubleshooting Your Go Installation

If you get an error instead of the version message, it’s likely that you don’t have Go in your executable path, or you have another program named `go` in your path.

On macOS and other Unix-like systems, use `which go` to see the `go` command being executed, if any. If nothing is returned, you need to fix your executable path.

If you’re on Linux or BSD, it’s possible you installed the 64-bit Go development tools on a 32-bit system or the development tools for the wrong chip architecture.

## Go Tooling

All of the Go development tools are accessed via the `go` command. In addition to `go version`, there’s a compiler (`go build`), code formatter (`go fmt`), dependency manager (`go mod`), test runner (`go test`), a tool that scans for common coding mistakes (`go vet`), and more. These tools are covered in detail in Chapters 10, 11, and 15.

For now, let’s take a quick look at the most commonly used tools by writing everyone’s favorite first application: **“Hello, World!”**

### Note

Since the introduction of Go in 2009, several changes have occurred in the way Go developers organize their code and their dependencies. Because of this churn, there’s lots of conflicting advice, and most of it is obsolete (for example, you can safely ignore discussions about `GOROOT` and `GOPATH`).

For modern Go development, the rule is simple: you are free to organize your projects as you see fit and store them anywhere you want.
