# Event-Driven Architecture in Golang

## Preface

Companies are adopting event-driven architecture (EDA) as their web applications grow in size and complexity. Applications that communicate using events are easier to develop and scale. Adding or developing your application around real-time interactions becomes easier with EDA.

Direct point-to-point communication between microservices inevitably leads to the development of a distributed monolith, which is just a monolith with extra and unnecessary complexity. EDA is an architecture that helps organizations decouple microservices and avoid developing another distributed monolith.

Choosing a new architecture for your next application or deciding to refactor an existing one can be fraught with known and unknown challenges. This book's goal is to provide you with enough examples and knowledge to give you a great head start should you decide to take the development of an EDA.

In this book, we will discuss and cover EDA concepts and related topics with the help of a small modular monolith demonstration application. We will use this application to take a journey through the concepts and topics to convert the synchronous mechanisms used by the application into asynchronous communication mechanisms.

## Who this book is for

This architecture book is for developers working with microservices, or those architecting and designing new applications built with microservices. Intermediate-level knowledge of Go is required to make the most of the examples and concepts in this book. Developers with a background in any programming language and experience working with microservices will still find the concepts and explanations useful.

## What this book covers

### Chapter 1: Introduction to Event-Driven Architectures

- Introduces EDA.

### Chapter 2: Supporting Patterns in Brief

- Covers helpful patterns such as domain-driven design, domain-centric architectures, and application architectures.

### Chapter 3: Design and Planning

- Explores ways to discover the capabilities and features of an application using EventStorming and other methods.

### Chapter 4: Event Foundations

- Introduces the Mallbots modular monolith application and domain events.

### Chapter 5: Tracking Changes with Event Sourcing

- Introduces event sourcing and guides you through the development of event-sourced aggregates.

### Chapter 6: Asynchronous Connections

- Covers adding asynchronous communication using event messages.

### Chapter 7: Event-Carried State Transfer

- Expands on the use of message-based communication between components.

### Chapter 8: Message Workflows

- Covers the concept of distributed transactions and introduces orchestrated sagas.

### Chapter 9: Transactional Messaging

- Explores the use of message inboxes and outboxes to reduce data loss.

### Chapter 10: Testing

- Discusses the concept of a testing strategy and guides you through testing an event-driven application.

### Chapter 11: Deploying Applications to the Cloud

- Covers the use of infrastructure as code and deploying an application as microservices.

### Chapter 12: Monitoring and Observability

- Discusses how to monitor a distributed application and make it observable with logging, metrics, and distributed tracing.

## To get the most out of this book

This book is written with the expectation that you can execute the demonstration application to understand and view the code changes that have been made in each chapter as the application is refactored. To get the most out of the book, it is recommended you read the chapters in order, as the chapters will reference code that has been modified in the previous chapter.

## Software/Hardware Covered in the Book

### Operating System Requirements

- **Go 1.18+**
- **Windows**, **macOS**, or **Linux**

### Tools

- **Docker 20.10.x**
  - Windows, macOS, or Linux
- **NATS 2.4**
  - Windows, macOS, or Linux

Most of the development for this book was done in Windows 10, but the code was tested to run in Windows Subsystem for Linux 2 (WSL 2) on Ubuntu 20.04 and tested to run on a Mac. You are expected to run the application and its dependencies within a Docker Compose environment. Instructions for using Docker are given wherever possible to minimize installing new software on your machine.

## Code Repository

If you are using the digital version of this book, we advise you to type the code yourself or access the code from the book’s GitHub repository (a link is available in the next section). Doing so will help you avoid any potential errors related to copying and pasting code.

You can follow the author on GitHub (https://github.com/stackus) or make a connection with them on LinkedIn (https://www.linkedin.com/in/stackmichael).

### Download the Example Code Files

You can download the example code files for this book from GitHub at [https://github.com/PacktPublishing/Event-Driven-Architecture-in-Golang](https://github.com/PacktPublishing/Event-Driven-Architecture-in-Golang). If there’s an update to the code, it will be updated in the GitHub repository.

We also have other code bundles from our rich catalog of books and videos available at [https://github.com/PacktPublishing/](https://github.com/PacktPublishing/). Check them out!

### Download the Color Images

We also provide a PDF file that has color images of the screenshots and diagrams used in this book. You can download it here: [https://packt.link/qgf1O](https://packt.link/qgf1O).

## Conventions Used

There are a number of text conventions used throughout this book.

### Code in Text

Indicates code words in text, database table names, folder names, filenames, file extensions, pathnames, dummy URLs, user input, and Twitter handles. Here is an example:

> “If all the participants have responded positively, then the coordinator will send a COMMIT message to all of the participants and the distributed transaction will be complete.”

### A Block of Code

A block of code is set as follows:

```sql
BEGIN;
-- execute queries, updates, inserts, deletes …
PREPARE TRANSACTION 'bfa1c57a-d99d-4d74-87a9-3aaabcc754ee';
```
