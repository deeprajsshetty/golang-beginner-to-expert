# 9. Transactional Messaging

In this book, we have transformed an application into an asynchronous one, which removes a lot of issues that arose from the tightly coupled and temporally bound nature of synchronous communication. Nearly all communication in the application is now made with a message brokered through NATS JetStream, providing loose coupling for the application components. However, despite all the advances we have made, we still face issues that all distributed systems suffer from.

In this chapter, we are going to discuss the following main topics:

- Identifying problems faced by distributed applications
- Exploring transactional boundaries
- Using an Inbox and Outbox for messages
- Technical requirements

You will need to install or have installed the following software to run the application or to try the examples:

- The Go programming language, version 1.18+
- Docker

The code for this chapter can be found at [GitHub Repository](https://github.com/PacktPublishing/Event-Driven-Architecture-in-Golang/tree/main/Chapter09).

## Identifying problems faced by distributed applications

In every distributed application, there are going to be places where interactions take place between components that reside in different bounded contexts or domains. These interactions come in many forms and can be synchronous or asynchronous. A distributed application could not function without a method of communication existing between the components.

In the previous chapter, we looked at some different ways, such as using sagas, to improve the overall reliability of complex operations that involve multiple components. The reliability we added is at the operation level and spans multiple components, but it does not address reliability issues that happen at the component level.

Let’s take a look at what affects reliability in synchronous and asynchronous distributed applications and what can be done to address the problem.

### Identifying problems in synchronous applications

In the first version of the MallBots application from much earlier in the book, we only had synchronous communication between the components, and that looked something like this:

![alt text](image-129.png)

Figure 9.1 – Synchronous interaction between the Payments and Ordering modules

In the preceding example, the Payments module receives a request for an invoice to be marked as paid. When an invoice is to be paid, a synchronous call is made to the Ordering module to also update the order status to Completed. Before returning to the user, the updated invoice is saved back into the database.

### Identifying problems in asynchronous applications

After fully updating the application to use asynchronous communication, we see that the same action now looks like this:

![alt text](image-130.png)

Figure 9.2 – Asynchronous interactions of the Payment module

In the updated version of the Pay Invoice action, we publish the `InvoicePaid` event instead of calling the Ordering module directly. The Ordering module is listening for this event, and when it receives the event, it will take the same action as before.

In both implementations, we can run into problems if the updated invoice data is not saved in the database. Also, for both implementations, you might think that a solution such as making a second call to the Ordering module or publishing an event to revert the change might help. However, there are a number of things that would make that solution improbable:

- The Ordering module has made or triggered irreversible modifications
- The second call or published event could also fail
- The failure is exceptional and there is no opportunity to take any corrective action

Adopting an asynchronous event-driven approach has not necessarily made the situation any better. In fact, things might be worse off now. When we publish the `InvoicePaid` event, we no longer directly know who has consumed the event and which other changes have happened as a result.

This is called a **dual write problem**, and it will exist everywhere we make a change in the application state that requires the change to be split in two or more places. When a write fails while multiple writes are being made, the application can be left in an inconsistent state that can be difficult to detect.

### Examining potential ways to address the problem

Sagas provide distributed transaction-like support for applications, and when an error occurs, they can be compensated. Using a small two-step saga might help here, but sagas are an operation consistency solution and not a component consistency solution. They cannot help with recovering a missing message or missing write in the database. That would be like using a sledgehammer to swat a fly.

Reordering the writes so that the more likely-to-fail write happens first does leave the more reliable ones to follow, but swapping the order of the writes will not help because the second write to whichever destination it is can always fail. Unless the action that is reordered to happen first is of little to no consequence, reordering the actions is not going to really be of any help.

Writing everything into the database while using a transaction could help if everything we were writing was in the database. However, in both of our examples, we are dealing with either a gRPC call or the publishing of a message. If we could convert all of our writes into ones that could be directed at our database, this approach would hold some promise.

### The singular write solution

If writing into multiple destinations is the problem, then it is reasonable to think that writing to a single destination could be a solution. Even more important than having a single write is that we need to have a single transactional boundary around what we will be writing.

We need a mechanism to make sure that the messages we publish must be added to the database in addition to or instead of being sent to NATS JetStream. We also need a way to create a single transaction boundary for this write to take place in. This means that for our application, we must create a transaction in PostgreSQL, into which we will put our writes in order to combine them into a single write operation.

This is called the **Transactional Outbox pattern**, or sometimes just the **Outbox pattern**. With it, we will be writing the messages we publish into a database alongside all of our other changes. We will be using a transaction so that the existing database changes that we would normally be making, and the messages, are written atomically. To do this, we will need to make the following changes to the modules:

- Setting up a single database connection and transaction that is used for the lifetime of an entire request
- Intercepting and redirecting all outgoing messages into the database

With these two items, we have our plan. First up, we will look at the implementation of a transactional boundary around each request, followed by the implementation of a transactional outbox.

### Exploring transactional boundaries

Starting with the more important part first, we will tackle how to create a new transaction for each request into our modules, whether these come in as messages, a gRPC call, or the handling of a domain event side effect. As we are using grpc-gateway, all of the HTTP requests are proxied to our gRPC server and will not need any special attention.

Creating a transaction is not the difficult part. The challenge will be ensuring the same transaction is used for every database interaction for the entire life of the request. With Go, our best option is going to involve using the **context** to propagate the transaction through the request. Before going into what that option might look like, we should also have a look at some of the other possible solutions:

1. We can toss out the option of using a global variable right away. Beyond being a nightmare to test, they will also become a problem to maintain or refactor as the application evolves.
2. A new parameter could be added to every method and function to pass the transaction along so that it can eventually be passed into the repositories and stores. This option, in my opinion, is completely out of the question because the application would become coupled to the database. It also would go against the clean architecture design we are using and make future updates more difficult.
3. A less ideal way to use a context with a transaction value within it would be to modify each of our database implementations to look for a transaction in the provided context. This would require us to update every repository or store and require all new implementations to look for a transaction. Another potential problem with this is we cannot drop in any third-party database code because it will not know to look for our transaction value.
4. A more ideal way to use the context, in my opinion, is to create a repository and store instances when we need them and to pass in the transaction using a new interface in place of the `*sql.DB` struct we are using now. Using a new interface will be easy and will result in very minimal changes to the affected implementations. Getting the repository instances created with the transactions will be handled by a new **dependency injection (DI)** package that we will be adding. The approach I am going to take will require a couple of minor type changes in the application code with the rest of the changes all made to the composition roots and entry points.

### How the implementation will work

A new DI package will be created so that we can create either singleton instances for the lifetime of the application or scoped instances that will exist only for the lifetime of a request. We will be using a number of scoped instances for each request so that the transactions we use can be isolated from other requests.

Before jumping further into how this implementation works, I should mention that this approach is not without a couple of downsides. The first is from the jump in complexity that using a DI package brings to the table. We will also be making use of a lot of type assertions because the container will be unaware of the actual types it will be asked to deliver.

Most of the updates will be made to the composition roots, which I believe softens the impact of those downsides somewhat.

Our implementation is going to require a new DI package. We will want the package to provide some basic features such as the management of scoped resources so that we can create our transactional boundaries.
