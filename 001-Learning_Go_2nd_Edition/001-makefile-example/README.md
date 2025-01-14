# Automated Build Process with Make for Go Developers

## Introduction

This README discusses the importance of automating the software build process and the use of the "make" tool as a solution for Go developers.

## IDEs vs. Automation

While Integrated Development Environments (IDEs) offer a user-friendly experience, they can be challenging to automate for creating repeatable builds. Modern software development practices prioritize the ability to run builds reliably by anyone, anywhere, at any time.

## Good Software Engineering Practice

Requiring automation tooling for builds is considered a fundamental software engineering practice. It helps prevent the common scenario where developers absolve themselves of build problems with the statement, "It works on my machine!"

## Scripted Build Steps

To achieve automation and repeatability, developers should use scripts to define and specify the sequence of steps required to build a software program.

## "Make" for Go Developers

In the world of Go programming, developers often turn to the "make" tool as their solution for specifying and organizing build steps. "Make" has a long history and has been used to automate build processes on Unix systems since 1976.

## Learning More

If you want to learn more about writing Makefiles, you can check out an informative tutorial by Chase Lambert. Please note that this tutorial includes a small amount of C code to explain the concepts.

[Makefile Tutorial by Chase Lambert](https://makefiletutorial.com/)

## Drawback

One drawback to Makefiles is that they are exceedingly picky. You must indent the steps in a target with a tab. They are also not supported out-of-the-box on Windows. If you are doing your Go development on a Windows computer, you need to install make first. The easiest way to do so is to first install a package manager like Chocolatey and then use it to install make (for Chocolatey, the command is choco install make.)

## Conclusion

Automating the build process is crucial in modern software development to ensure reliability and repeatability. Go developers can benefit from using "make" to define and execute build steps consistently.

Feel free to explore this repository for examples and resources on using "make" in your Go projects.
