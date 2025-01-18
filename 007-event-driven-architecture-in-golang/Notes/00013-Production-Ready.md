# Part 3: Production Ready

In this last part, we will cover the topics of testing, deployment, and observability. We will begin by discussing testing strategies and going over the different kinds of tests we can use to ensure our application works as intended. Next, we will refactor the application from a modular monolith into microservices that can be deployed into a cloud environment. Then, we will update the application so that it can be monitored using logging, metrics, and distributed traces.

This part consists of the following chapters:

- **Chapter 10**: Testing
- **Chapter 11**: Deploying Applications to The Cloud
- **Chapter 12**: Monitoring and Observability

# Chapter 10: Testing

In Part 2 of this book, we took an entirely synchronous application and transformed it into an asynchronous application using events and messaging. Our application is more resilient and agile but has gained some new libraries and dependencies as a result.

Testing an asynchronous application can pose some unique challenges but remains within reach by following testing best practices. In this chapter, we will look at testing the MallBots application from the unit test level and writing executable specifications using the Gherkin language.

## In this chapter, we will cover the following topics:

- Coming up with a testing strategy
- Testing the application and domain with unit tests
- Testing dependencies with integration testing
- Testing component interactions with contract tests
- Testing the application with end-to-end tests

## Technical Requirements

You will need to install or have installed the following software to run the application or to try this chapter’s examples:

- The Go programming language version 1.18+
- Docker

The code for this chapter can be found at:  
[GitHub - Chapter 10 Code](https://github.com/PacktPublishing/Event-Driven-Architecture-in-Golang/tree/main/Chapter10)

## Coming Up with a Testing Strategy

For applications such as MallBots, we should develop a testing strategy that tests whether the application code does what it is supposed to be doing. It should also check whether various components communicate and interact with each other correctly and that the application works as expected:

![alt text](image-140.png)

Figure 10.1 – Our testing strategy as a pyramid or ziggurat

## Our Testing Strategy

Our testing strategy will have four parts:

1. **Unit tests**
2. **Integration tests**
3. **Contract tests**
4. **End-to-end tests**

Unit tests are a no-brainer addition to our strategy; we want to ensure the code we write does what we intend it to do. We want to test the input and output from the module core and include an integration test to test the dependencies that it uses. We will use contract tests to detect any breaking changes to the application’s many APIs and messages that tie the modules together. Finally, we want to run tests that check that the application is functioning as per stakeholder expectations and will use end-to-end (E2E) testing.

There are additional levels and forms of testing that we could include, such as component testing. This would be used to test each module in isolation – that is, like an E2E test but limited to just that module. We may also see some manual tests take place or have the testing or development teams work through scenarios to perform exploratory testing. We could also stress or load test the application, which could be added to the strategy later as the application matures.

### Unit Tests

Unit tests should make an appearance in any testing strategy. They are used to test code for correctness and to locate problems with application and business logic implementations. In a testing strategy, they should take up the bulk of the testing efforts. These tests should be free of any dependencies, especially any I/O, and make use of test doubles such as mocks, stubs, and fakes. The system under test for a unit test should be very small; for example, individual functions and methods.

#### System Under Test

At each level of testing, we use the term **system under test** (SUT) to describe the component or components being tested. For unit tests, the SUT may be a function, whereas for E2E testing, it would encompass the application and any external APIs involved. Generally, the SUT expands in scope or application coverage the higher up you go in the testing pyramid.

Any application can benefit from having unit tests in its testing strategy. If used sparingly, extremely fast-running tests can focus on logic and algorithms that are complex or critical to the success of the business.

### Integration Tests

Next up is integration testing, where, instead of focusing on the logic, you will focus on testing the interactions between two components. Typically, you must test the interactions between a component with one of its dependencies. Testing that your ORM or repository implementations work with a real database would be an example of an integration test. Another example would be testing that your web interface works with application or business logic components. For an integration test, the SUT will be the two components with any additional dependencies replaced with mocks.

Applications with complex interactions in their infrastructure can benefit from the inclusion of integration tests in the testing strategy. Testing against real infrastructure can be difficult or too time-consuming, so teams may decide to not do so or only develop a few critical path tests.

### Contract Tests

A distributed application or a modular monolith like ours is going to have many connection points between the microservices or modules. We can use contract tests built by consumers’ expectations of an API or message to verify whether a producer has implemented those expectations. Despite being rather high on the testing pyramid, these contract tests are expected to run just as fast as unit tests since they do not deal with any real dependencies or test any logic. The SUT for a contract will be either the consumer and its expectation, or the producer and its API or message verification.

Distributed applications will benefit the most from adding contract tests to the testing strategy. These tests are not just for testing between microservices – they can also be used to test your UI with its backend API.

### End-to-End Tests

E2E tests are used to test the expected functionality of the whole application. These tests will include the entire application as the SUT. E2E tests are often extensive and slow. If your application includes a UI, then that too will become part of the tests because they intend to test the behaviors and correctness of the application from the point of view of the end user. The correctness being tested for is how the application performs and not like a unit test’s correctness of how the application does it.

Teams that take on the effort of maintaining fragile and costly tests are rewarded with confidence that the application is working as expected during large operations that can span the whole application.

### In the upcoming sections, we will explore each of the testing methods present in our testing strategy.

---

## Testing the Application and Domain with Unit Tests

The system under test for a unit test is the smallest unit we can find in our application. In applications that are written in Go, this unit will be a function or method on a struct:

![alt text](image-141.png)

Figure 10.2 – The scope of a unit test

As shown in Figure 10.2, only the function code is being tested. Any dependencies that the code under test requires must be provided as a test double such as a mock, a stub, or a fake dependency. Test doubles will be explained a little later in the **Creating and Using Test Doubles in Our Tests** section.

Each test should focus on testing only one path through the function. Even for moderately complex functions, this can result in a lot of duplication in your testing functions. To help with this duplication, the Go community has adopted **table-driven tests** to organize multiple tests of a single piece of code under test into a single test function.

### Table-Driven Testing

This method of testing was introduced to the Go community by Dave Cheney in his similarly named blog post, ["Prefer table driven tests"](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests). Table-driven tests are made up of two parts – a table of test cases and the testing of those test cases.

#### The Table of Test Cases

A slice of structs that contains the test inputs and outputs is called a test case. The following listing shows a table with two test cases using a map to build the table:

```go
tests := map[string]struct {
    input string
    want  int
}{
    "word":  {input: "ABC", want: 3},
    "words": {input: "ABC ABC", want: 6},
}
```

If we used a slice instead of the map, then we would want to include an additional field in the struct to hold a string that is used as the subtest’s name.

### Testing Each Test Case

The actual testing will depend on how the unit needs to be tested. However, there is some simple boilerplate code that we should use so that we can make sense of the test failures, should they pop up.

In the following code block, the highlighted code is the simple boilerplate that is used to run through each test case:

```go
for name, tc := range tests {
    t.Run(name, func(t *testing.T) {
        // arrange, act, and assert
    })
}
```

In the loop, we use the subtesting feature to run each test case under the heading of the original test function name. The following output is an example of running the `AddItem` application tests for the Shopping Baskets module:

--- PASS: TestApplication_AddItem (0.00s)
--- PASS: TestApplication_AddItem/NoBasket (0.00s)
--- PASS: TestApplication_AddItem/NoProduct (0.00s)
--- PASS: TestApplication_AddItem/NoStore (0.00s)
--- PASS: TestApplication_AddItem/SaveFailed (0.00s)
--- PASS: TestApplication_AddItem/Success (0.00s)
PASS

### The AddItem Test and Test Doubles

The `AddItem` test has five test cases that test how the input to the function might be handled under different conditions. This test can be found in the `/baskets/internal/application/application_test.go` file.

The application that `AddItem` is defined on has several dependencies, and each of those is replaced with test doubles so that we can avoid dealing with any real I/O. We also want to intercept calls into the dependencies to control which path through the `AddItem` method we are testing.

We will want to use a test double that is not only able to intercept the calls but also able to send back programmed responses. There are several kinds of test doubles, so let’s look at them and see which works best for us here.

### Creating and Using Test Doubles in Our Tests

Test doubles are tools we can use to isolate the system or code under test from the rest of the system around it.

These tools come in different forms, each useful for different testing scenarios:

- **Fakes** implement the same functionality as the real dependency. An in-memory implementation of a repository could stand in and take the place of a PostgreSQL implementation so that the test does not rely on any real I/O.
- **Stubs** are like fakes, but the stub implementation responds with static or predictable responses.
- **Spies** work like an observable proxy of the real implementation. A spy can be used to report back the input, return values, and the number of calls that it received. Spies may also help with recording the inputs and outputs that were seen for later use.
- **Mocks** mimic the real implementation, similar to a fake, but do not provide an alternative implementation. Instead, a mock is configured to respond in certain ways to specific inputs. Then, like a spy, it can be used to assert whether the proper inputs were received, the right number of calls were made, and no unexpected calls or inputs were encountered.

Fakes and stubs can be used when the interaction with the dependency is not important to the test, whereas spies and mocks should be used when the input and responses matter.

### Working with Mocks

For our unit test, we will use mocks. To create the mocks that we will use, we will use the **Testify mocks package** ([https://github.com/stretchr/testify](https://github.com/stretchr/testify)). This will provide the mocking functionality, along with the **mockery tool** ([https://github.com/vektra/mockery](https://github.com/vektra/mockery)) to make generating them a breeze. The mockery tool can be installed with the following command:

```bash
go install github.com/vektra/mockery/v2@v2.14.0

```

Each module that will be tested using mocks will have the following line added to its `generate.go` file; for example, `/baskets/generate.go`:

```go
//go:generate mockery --all --inpackage --case underscore
```

This `go:generate` directive will look for the interfaces defined within the directory and subdirectories and create mocks of them. `--inpackage` and `--case underscore` will configure the tool to create the mocks in the current package using underscores in the filename. The `--all` flag will make the tool generate a mock for each interface that is found. When mockery creates mocks next to the interface, it will add a `Mock` prefix to each interface that it found in a file with a prefix of `mock_`. For example, the `Application` interface is mocked as `MockApplication`, and that mock will be found in `/baskets/internal/application/mock_application.go`.

Organizing and naming test doubles comes down to preferences in most cases. My preference is to place them next to the interfaces and implementations that they double. Another preference is to keep the naming simple and use either a prefix or suffix to identify the type of test double that you are dealing with.

With the mocks created, we need to use them in our tests. To do that, we will include a new field in our test case structs so that they can be configured for each test case:

```go
type mocks struct {
    baskets   *domain.MockBasketRepository
    stores    *domain.MockStoreRepository
    products  *domain.MockProductRepository
    publisher *ddd.MockEventPublisher[ddd.Event]
}

tests := map[string]struct {
    ...
    on      func(f mocks)
    wantErr bool
}{...}
```

In the previous listing, a named struct has been created with field types of the actual mocks; using the interfaces here will not help us since we want the concrete mock implementations. Then, in the anonymous struct that defines our test cases, we have added a function that accepts the mocks struct. With these additions, any test case that is expected to make calls into a mock can do so during the Arrange portion of the test function.

### Arrange, Act, and Assert

The **Arrange-Act-Assert (AAA)** pattern is a simple yet powerful way to build your tests. It breaks up a test function into three parts for better maintainability and readability. The Arrange portion is where the test is set up, the Act portion is where the target is called or interacted with, and the Assert portion is where the final test or verification happens. By following this pattern, it is easy to spot test functions that are doing more than testing one thing at a time. This pattern is also known as **Given-When-Then**.

In the test function, we must set up the mocks, execute the method, and perform our assertions using the following code:

```go
// Arrange
m := mocks{
    baskets:   domain.NewMockBasketRepository(t),
    stores:    domain.NewMockStoreRepository(t),
    products:  domain.NewMockProductRepository(t),
    publisher: ddd.NewMockEventPublisher[ddd.Event](t),
}
if tc.on != nil {
    tc.on(m)
}
a := New(m.baskets, m.stores, m.products, m.publisher)

// Act
err := a.AddItem(tc.args.ctx, tc.args.add)

// Assert
if (err != nil) != tc.wantErr {
    t.Errorf("AddItem() unexpected error = %v", err)
}
```

During each test case run, we will create new mocks but leave them alone if no function has been defined to configure them. A mock will fail the test if it is called and we have not configured any calls. This will be helpful because we do not need to remember which calls or mocks have been set up, nor which to change when we make changes to the code under test.

The **Mockery** package has generated constructors for our mocks that accept the test variable. Using the constructors, we do not need to include any additional assertions for the Assert portion of the test function. When the test completes, each mock will be automatically checked to ensure that the exact number of calls were made into it and that the calls included the correct inputs:

![alt text](image-142.png)

Figure 10.3 – The SUT for the AddItem method

To test the `AddItem` method on `Application`, we must provide an application instance with all the dependencies that it needs and then pass additional parameters to the `AddItem` method. The method only returns an error, so using a mock double instead of any of the others makes the most sense. Without mocks, we would not be able to see into the method.

### Testing dependencies with integration testing

An application is made up of many different components; some of those components are external to the application. These external components can be other services, or they can be infrastructure that needs to be in place for the application to function properly.

It is hard to find any application built for the web that does not interact with infrastructure. Actually, it’s impossible – the web server or API gateway that the application would use to receive requests would fall into the definition of infrastructure.

We cannot test these dependencies using a unit test because if we replaced the dependency with any kind of test double, then it would not be a true test of that dependency.

In an **integration test**, both components are the **SUT**, which means we need to test with real infrastructure when possible:

![alt text](image-143.png)

Figure 10.4 – The scope of an integration test

Unlike the unit tests, which were expected to be very simple, at least in terms of what the environment must provide so that they can run, **integration tests** ramp up the complexity a great deal.

A wrong way to develop these tests is to assume each developer that might run them has the same environment as you do. That would mean the same OS, the same installed tools, and the same locally accessible services with the same configurations and permissions. You would then write the test leaving all of those environmental expectations out and leaving out how to run the tests in the documentation or some other form of knowledge share.

A better way to write **integration tests** would be to bring what is necessary to run the test into that test without requiring any test-specific environment and setup.

### Incorporating the dependencies into your tests

A lot of the services and infrastructure in use today are available as a Docker container; this could be a real production-ready container such as the ones for many databases. Some containers are purpose-built to aid in development or testing efforts. An example of a container that can help with development would be **LocalStack** ([localstack.cloud](https://localstack.cloud)), a container that provides a local development and testing environment for many of the offerings from AWS.

Using **Docker** and containers is a great way to bring these dependencies into your local environment. However, the challenge in using them in tests is that we want to be able to control their state and may also want to set them up in different ways to support different tests. However, we need to know how to incorporate these containers into our tests.

#### Option 1 – Manually by using Docker Compose files

We can create a **Docker Compose** file for our tests, such as `test.docker-compose.yml`, that will stand up everything we will need to connect to for the integration tests that we’ll write. This should make it easy for every developer to have the dependencies available, and so long as everyone remembers to start up the environment, they should also have no issues running the tests. Volumes can be destroyed during the environment teardown so that previous runs do not affect others.

The downsides of this option begin with the **Compose file** itself. If a problem exists when standing up the entire environment, then someone will need to make changes to it before they can test. There may also be issues running the tests multiple times, so tearing down the environment to stand it back up again might be necessary, which could take a considerable amount of time. To tackle this, we can take a different approach.

#### Option 2 – Internalizing the Docker integration

There is a solution we can use that will not only allow us to use different containers or compose environments for different tests but also remove the step of having to run a Docker command before executing any integration tests.

**Testcontainers-Go** ([golang.testcontainers.org](https://golang.testcontainers.org)) is a library that makes it possible to start up a container or compose an environment that is controlled by code that we can include in our tests.

The benefits of this option are that we will always have a pristine environment to run our tests in and subsequent runs will not need to wait for any containers or volumes to be reset. The other is the containers will always be started and removed when the test is run. This means that there is no longer any need to maintain documentation on how to prepare a local environment to run tests. This is the better option in my opinion, but it will require some initial setup, as well as some resetting or cleanup between each test.

### Running tests with more complex setups

Our **integration tests** will likely end up being a little more complex than the unit tests we have previously worked with. We may require certain actions to occur at the start of the run and the end; likewise, we need actions to run before and after each test. This is not a difficult task by any stretch. We can write the test harness ourselves, but whatever we write should also contain tests. Instead, we can use an existing test harness that handles all of this for us.

This harness is the **Testify suite** package. When we are using this new harness, we can continue to use table-driven tests, but we need to manage the state setup and reset ourselves.

To start using **Testify suites**, create a new struct and include `suite.Suite` as an anonymous field. Then, create a simple test function to run the suite:

```go
type productCacheSuite struct {
    suite.Suite
    // [optional] any additional fields
}

func TestProductCacheRepository(t *testing.T) {
    suite.Run(t, &productCacheSuite{})
}
```

We can include additional fields in the struct that can be accessed by the test methods.

### Testing ProductCacheRepository

We will use all of the aforementioned methods to test the interaction between the **PostgreSQL** implementation of **ProductCacheRepository** and PostgreSQL:

![alt text](image-144.png)

Figure 10.5 – The integration test for ProductCacheRepository

# gRPC Client Integration Test Setup

This implementation uses a connection to the database and has a dependency on the `ProductRepository` interface. In the application, this is implemented as a gRPC client, which will fetch a `Product` instance when it cannot be found in the database. For this integration test, that dependency will be mocked. Before writing our first test, we need to configure the suite so that our tests can use a real database while remaining isolated from each other.

## Suite Composition

Inside the `productCacheSuite` struct, we will add the following additional fields:

- **container**: Holds the reference to the PostgreSQL container we have started.
- **db**: A real database connection to PostgreSQL. It will be used to reset the database between tests.
- **mock**: An instance of `MockProductRepository`. If there are other dependencies to mock or fake, a more specific name could be used.
- **repo**: A real instance of the PostgreSQL implementation that we intend to test.

These fields will be accessible to our tests and the methods we will use to set up the suite and each individual test.

## Suite Setup

We must begin by setting up the suite with some fields that will be available to all tests. The first of these is the database connection. To make that connection, we need to have a database we can connect to. The following is how the PostgreSQL container is started up:

```go
const dbUrl = "postgres://***:***@localhost:%s/mallbots"
s.container, err = testcontainers.GenericContainer(ctx,
    testcontainers.GenericContainerRequest{
        ContainerRequest: testcontainers.ContainerRequest{
            Image:        "postgres:12-alpine",
            ExposedPorts: []string{"5432/tcp"},
            Env: map[string]string{
                "POSTGRES_PASSWORD": "***",
            },
            Mounts: []testcontainers.ContainerMount{
                testcontainers.BindMount(
                    initDir,
                    "/docker-entrypoint-initdb.d",
                ),
            },
            WaitingFor: wait.ForSQL(
                "5432/tcp",
                "pgx",
                func(port nat.Port) string {
                    return fmt.Sprintf(
                        dbUrl,
                        port.Port(),
                    )
                },
            ).Timeout(5 * time.Second),
        },
        Started: true,
    },
)
```

This listing will start up a new container from the `postgres:12-alpine` image. Like the service entry in the `docker-compose.yml` file, we must provide it a hostname and initialize it with some files that will be mounted in the container.

The `WaitingFor` configuration is used to block the startup process until the database is truly ready for requests. In the `docker-compose.yml` file, we achieve a similar effect using a small `wait-for` script.

The `testcontainers-go` package can also stand up services defined in a Docker Compose file. While we won't be making use of this feature in our case, you can learn more about it in the [Testcontainers Go documentation](https://golang.testcontainers.org/features/docker_compose/).

Once the container is running and we are waiting for it to become available, we can proceed to establish a connection to the database.

## Test Setup

Before each test, a new mock is created, which is then injected, along with the database connection, into the `ProductCacheRepository`:

```go
func (s *productCacheSuite) SetupTest() {
    s.mock = domain.NewMockProductRepository(s.T())
    s.repo = NewProductCacheRepository(
        "baskets.products_cache",
        s.db,
        s.mock,
    )
}
```

We are keeping a reference to the mock because we will need to configure it during tests to expect specific calls. If we did not maintain this reference, there would be no way to configure the mock repository from within the `ProductCacheRepository` during testing.

## Test Teardown

Every test should run in isolation, with the same initial state. For tests involving the database, we will be creating new rows, updating rows, or deleting them. Without resetting the database between tests, we might encounter situations where the order of tests affects subsequent test outcomes (i.e., one test might inadvertently affect another).

To ensure a clean slate for each test, we perform a database reset by truncating the relevant tables:

```go
func (s *productCacheSuite) TearDownTest() {
    _, err := s.db.ExecContext(
        context.Background(),
        "TRUNCATE baskets.products_cache",
    )
    if err != nil {
        s.T().Fatal(err)
    }
}
```

To keep things simple, we will `TRUNCATE` any tables we work with in the tests. This approach is safe as long as the test suite always uses a dedicated PostgreSQL container that exists only for the duration of the test suite.

## Suite Teardown

Once all tests have finished running, we no longer need the database connection, and the container should be cleaned up and removed. The following code accomplishes this:

```go
func (s *productCacheSuite) TearDownSuite() {
    err := s.db.Close()
    if err != nil {
        s.T().Fatal(err)
    }

    err = s.container.Terminate(context.Background())
    if err != nil {
        s.T().Fatal(err)
    }
}
```

In reverse order from what happened in the `SetupSuite()` method, we close the database connection and then terminate the container, which removes it and any volumes we might have created.

## The Tests

With all the setup and teardown taken care of, our tests are going to be simple and to the point, much like the unit tests were. The following listing shows the test for the rebranding functionality:

```go
func (s *productCacheSuite) TestPCR_Rebrand() {
    // Arrange
    _, err := s.db.Exec("INSERT ...")
    s.NoError(err)

    // Act
    s.NoError(s.repo.Rebrand(
        context.Background(),
        "product-id",
        "new-product-name",
    ))

    // Assert
    row := s.db.QueryRow("SELECT ...", "product-id")
    if s.NoError(row.Err()) {
        var name string
        s.NoError(row.Scan(&name))
        s.Equal("new-product-name", name)
    }
}
```

We can access any of the fields defined in the suite and can even organize the tests in AAA (Arrange, Act, Assert) fashion. During the Arrange phase of this test, we use the database connection to insert a new product cache record that is then acted upon in the next phase. The suite also has access to all the usual Testify assert functions, and we can skip importing that package in favor of using the assertion methods directly from the suite itself.

## Breaking Tests into Groups

Integration tests do not need to run quickly, and for good reason. Integration tests typically need to deal with I/O, which is not exactly fast or predictable. Skipping or excluding the longer-running tests will be necessary if you want to keep the wait for test feedback as low as possible when developing new logic or features.

### Ways to Break Long-Running Tests into Groups

There are three ways to break long-running tests into groups or exclude them when running fast-running unit tests.

#### 1. Running Specific Directories, Files, or Tests

You can specify specific files, directories, and even individual tests when using the `go test` command. This option will not permanently break your tests up into different groups that can be run separately, but outside of using your IDE, it presents the easiest way to target individual tests.

For example, to run all the application tests for the Shopping Baskets module, you would use the following command:

```bash
go test ./baskets/internal/application

```

To run only the `RemoveItem` test, you would add `-run "RemoveItem$"` to the command:

```bash
go test ./baskets/internal/application -run "RemoveItem$"
```

We can target specific table-driven subtests as well. To run only the `NoProduct` subtest for the `RemoveItem` test, we can use `RemoveItem/NoProduct$`. For the following command, I have moved into the `internal` directory:

```bash
go test ./application -run "RemoveItem/NoProduct$"

```

### Using Regex to Target Specific Tests

In the previous two command examples, I used a Regex to search for the test to run. You can target a group of tests with a well-written Regex. The test tool makes it very easy to target specific tests when we need to be very focused on a test or a collection of tests.

## Go Build Constraints

We can use the conditional compilation build constraints to create groups of our tests. These constraints are normally used to build our programs for different OSs or CPU architectures, but we can also use them for our tests because the tests and the application are both compiled when we run the tests.

Because this is accomplished by adding a special comment to the top of our files, we can only group tests together by files; we cannot create any subgroups of the tests within the files.

To group tests into an integration grouping, we can add the following with a second blank line to the top of the test file:

```go
//go:build integration

```

### Rules for Using Go Build Constraints

The following are a few rules that need to be followed for the compiler to recognize the comment as a build constraint:

- There must not be any spaces between the single-line comment syntax and `go:build`. Multiline comment syntax will not work.
- The constraint must be followed by a blank line.
- The constraint must be the first line in the file.

Once this is done, the file will now be ignored when we run the test command. The examples from the previous section would all ignore the tests in this file, even if we were to target the file and tests specifically.

To run the tests now, we will need to pass the `-tags` option into the test command, like so:

```bash
go test ./internal/postgres -tags integration

```

### Combining Multiple Tags with Build Constraints

You can combine multiple tags to create subgroups using the build constraints by taking advantage of the Boolean operators that it supports. We can modify the constraint so that the database tests are run with all integration tests or can be run by themselves:

```go
//go:build integration || database

```

A file with this constraint could be run using any of the following commands:

```bash
go test ./internal/postgres -tags integration

go test ./internal/postgres -tags database

go test ./internal/postgres -tags integration,database

```

Using build constraints is a powerful and easy way to create groups of tests. Without the `-tags` option, any file that uses a build constraint will be ignored. This also has the downside of potentially skipping tests that are broken and not knowing it. The constraints at the top of the file can include typos or logical errors caused by incorrect operator usage.

When using build constraints, it is best to keep it simple.

## Using the Short Test Option

The final method we will look at is the short test mode, which is built into the test tool. To enable short mode, you can simply pass in the `-short` option to any test command that you run. By itself, nothing happens, but if you include a check in your tests, you can exclude the longer-running tests from running. The test tool itself is not able to determine which tests are long-running tests; that determination is up to you.

We can skip long-running tests by using a block of code such as this:

```go
func TestProductCacheRepository(t *testing.T) {
    if testing.Short() {
        t.Skip("short mode: skipping")
    }
    suite.Run(t, &productCacheSuite{})
}

```

The entire suite of tests will be skipped when the following command is run:

```bash
go test ./internal/postgres -tags integration -short

```

Checking for short mode can be added to any individual tests and subtests as well; we do not need to limit ourselves to tests run via a suite. Skipping tests with short mode allows us to be more selective about which test or tests are ultimately skipped.

The downside to using short mode is that the long-running tests are included by default, and we need to enable short mode to skip them. Another downside is that the option can be either on or off; there is no way to split your tests into more than two groups.

All three options I’ve mentioned can be used together. You could treat short mode as a way to skip tests that are just a little longer when running unit tests, and likewise for the other kinds of tests when used with the `-tags` option.

By using Docker containers, we can test more of our application by including real infrastructure in our tests, and by grouping the tests, we can exclude them when we want to run very fast unit tests. This form of testing will be too fragile to test integrations much larger than infrastructure interactions. For that testing, we can turn to contract tests.

## Testing Component Interactions with Contract Tests

We have chosen to create our application using the modular monolith pattern. Here, we have each module communicate with the others using either gRPC or by publishing messages into a message broker. These would be very common to see on any distributed application but would be rare or not used at all on a monolith or small application.

What’s more common to see across all applications is the REST API we use. This demonstration application does not have any true UI, but we have the API to support one. This API represents a third form of communication in our application, which is between an API provider and the API consumer.

We could test these interactions using integration tests since the definition of what an integration test covers is testing the interactions or integration between two components. However, the integration tests we wrote before tested smaller components, and the scope for the system under test was not very large. They are larger than the unit tests before them but are still small.

![alt text](image-145.png)

Figure 10.6 – System under test for an integration test of two modules

Testing even two modules together using an integration test would be a very large jump regarding how much of the system under test would now be forced into the scope of the test. There may be several real dependencies that are too difficult to replace with a test double, and a real dependency would need to be stood up and used for the test.

Another possible but very likely issue with using integration tests in this manner is that we could be testing two components that have entirely different development teams and release schedules.

We want to minimize the extraneous components that get included in the test scope. This means we should only target the REST API if that is what we are interested in testing. The same goes for messaging; we should test whether we are receiving the messages that we expect and leave the rest of the module out of the equation.

Contract testing allows us to focus on the APIs between providers and consumers, just like the integration tests do, but it allows us to run the tests in isolation, similar to a unit test.

## Contract Testing

Contract testing comes in two forms:

- **Consumer-driven contract testing (CDCT)**, which is when the contract is developed using the expectations of the consumers
- **Provider-driven contract testing (PDCT)**, which is when the contract is developed using the provided API of the provider

We will be using **CDCT** for our testing.

![alt text](image-146.png)

Figure 10.7 – System under test for consumer-driven contract testing

Contract testing is broken down into two parts: **consumer expectations** and **provider verifications**. Between the two sits the **contract** that is produced by the consumer when using consumer-driven testing.

- On the **consumer side**, real requests are made to the mock provider, which will respond with simulated responses.
- On the **provider side**, the real requests will now be used as simulated requests, and the provider will respond with real responses that are verified against the expected responses recorded in the contract.

Because the consumer is creating expectations, there would be no value in only running the consumer side without the provider verifying those expectations. Each side, both consumer and provider, has different contract testing goals.

## Consumer Expectations

The consumers of an API will uniquely use that API. This could mean that they use a fraction of the provided API endpoints or messages, and it could also mean that they are using only a portion of the data they are provided.

Consumers should write their expectations based on what they use. This allows providers that are tested with contract testing to know what endpoints and data are being used by the consumers.

Consumers' expectations will change over time, as will the contracts. Processes can be set up in your CI/CD pipeline so that these changed contracts can be automatically verified with the provider to ensure that there are no issues in deploying the updated consumer into production.

## Provider Verifications

Providers will be given one or more contracts to verify their API support. Each contract that they receive will expect different things from their API, different collections of endpoints, or different simulated requests.

The providers will be expected to implement the tests to verify the simulated requests against their real API. However, they may use whatever test doubles they need so that they don’t have to stand up their entire module or microservice.

When a consumer’s contract is verified, this can be shared with the consumer so that they know it will be OK to deploy with its API usage. Likewise, a provider, having passed all of the contract verifications it was presented with, will have confidence in knowing it too can be deployed without any issues.

## Not Building Any Silos

Contract testing does not eliminate any necessary communication regarding the integrations between teams; it helps them know about and get to the issues quickly. With contract testing, we achieve a high level of confidence on both sides that the integration is working.

When issues are discovered during verification, it is expected that the teams will have some dialog. Consumers can make mistakes and have incorrect expectations, which could mean there is room to improve or add API documentation. Providers may make a breaking change and will need to cooperate with the affected consumers to coordinate updates and releases.

## Contract Testing with Pact

Just like using the Testify suite package for our more complex test setups, we will use a tool called **Pact** ([https://pact.io](https://pact.io)) to handle a lot of the concerns outside of our tests. Pact provides libraries for many languages, which is handy for testing a JavaScript UI with your Go backend. Several tools can be used locally by the developers, as well as in the CI/CD process, to provide the promised confidence that deployments can happen without issues.

### Pact Broker

**Pact Broker** ([https://docs.pact.io/pact_broker](https://docs.pact.io/pact_broker)) is an application we can start up in our environment to share contracts, as well as provide feedback for consumers stating that their contracts have been verified by the provider.

![alt text](image-147.png)

Figure 10.8 – Pact Broker showing our example integrations

Pact Broker can also be integrated with your CI/CD process to automate the testing of providers when a consumer has created or updated a contract. Likewise, consumers can be automatically tested when a provider has made changes to their API.

![alt text](image-148.png)

Figure 10.9 – Contract creation and verification flow using Pact Broker

Pact Broker may be installed locally using a Docker image, though you may use the hosted version with a free account at [https://pactflow.io/](https://pactflow.io/).

## CLI Tools

Pact will take care of creating and running the mock provider and consumer, but this functionality will require the necessary Pact CLI tools to be installed and available. For more information, refer to the official documentation: [Pact CLI Documentation](https://docs.pact.io/implementation_guides/cli). You may choose either a Docker image or a Ruby-based standalone version.

## Additional Go Tools

The provider example for the asynchronous tests uses an updated version of the Go libraries. If you would like to follow along and run these tests, you will need to install the `pact-go` installer and use it to download some additional dependencies:

```bash
go install github.com/pact-foundation/pact-go/v2@2.x.x

pact-go -l DEBUG install
```

The two preceding commands will download some files that will allow the updated provider verifications to run.

At the time of writing this book, the version used was tagged as `v2.x.x`. The minor and patch version values are `x`.

## REST Consumer and Provider Example

First, we will test a simple JavaScript client against the REST API provided by the Shopping Baskets module. We do not have a real UI to add tests to, but we can create a small JavaScript client library. For contract testing, we would only want to work with the client library anyhow, so this is not a big problem.

We will focus on a couple of endpoints for this demonstration:

```javascript
const axios = require("axios");

class Client {
  constructor(host = "http://localhost:8080") {
    this.host = host;
  }

  startBasket(customerId) {
    return axios.post(`${this.host}/api/baskets`, { customerId });
  }

  addItem(basketId, productId, quantity = 1) {
    return axios.put(`${this.host}/api/baskets/${basketId}/addItem`, {
      productId,
      quantity,
    });
  }
}

module.exports = { Client };
```

This JavaScript client is ready to be used in the latest single-page application (SPA) frontend and deployed to production. Before we deploy this client, it needs to be tested against the REST API.

Now, instead of starting up the real REST API server and running tests, we want to create individual interactions and test those against a mock provider. These interactions will then be used to produce a contract that is shared with the provider, so it may verify every interaction from its point of view. This allows us to test these interactions just as swiftly as our unit tests.

To better explain these interactions, we will look at one from `/baskets/ui/client.spec.js` for the UI consumer tests in the Shopping Baskets module:

```javascript
provider
  .given("a store exists")
  .given("a product exists", { id: productId })
  .given("a basket exists", { id: basketId })
  .uponReceiving("a request to add a product with a negative quantity")
  .withRequest({
    method: "PUT",
    path: `/api/baskets/${basketId}/addItem`,
    body: {
      productId: productId,
      quantity: -1,
    },
    headers: { Accept: "application/json" },
  })
  .willRespondWith({
    body: MatchersV3.like({
      message: "the item quantity cannot be negative",
    }),
    headers: { "Content-Type": "application/json" },
    status: 400,
  });
```

In the previous listing, we are building an interaction for a call to the `AddItem` endpoint. We expect to receive an error when we include a negative quantity in our request.

Here is what each method is doing when building the interaction:

- **`given()`** is used to signal to the provider that a certain state should be configured or used to respond to the simulated request when it is verifying the contract. Of the four methods shown, only `given()` is optional. It is used in the code example three times, with two of the calls including static data that should be used in place of the state the provider would generate.
- **`uponReceiving()`** sets up a unique name for this expectation.

- **`withRequest()`** defines the exact request that will be used by both the consumer tests and provider verification tests. In the consumer tests, it is compared with the real request that will be made to the mock provider. Then, in the provider tests, it will be used as a simulated request from the mock consumer against the provider.

- **`willRespondWith()`** is the expected response. We build it using matchers, creating an expectation based on what is important to the consumer. In the consumer tests, this response will be returned by the mock provider, and in the provider tests, the real response is verified against it. The real error response from the `AddItem` endpoint includes more than the `message` property, but we match only the one value that we care about.

The interaction is then tested using your preferred testing library. We will only be able to truly test one side of the interaction right now, which involves verifying that the request we send to the mock provider is exactly as we said it would be:

```javascript
it("should return an error message", () => {
  return provider.executeTest((mockServer) => {
    const client = new Client(mockServer.url);
    return client.addItem(basketId, productId, -1).catch(({ response }) => {
      expect(response.status).to.eq(400);
    });
  });
});
```

To test the interaction with the consumer, we use the real client code to create and send a request to the mock provider. The response can be checked as well, and in this case, we catch the expected error response. If we don’t, then an uncaught exception could occur, which would throw off our test.

When all of our consumer tests are passing, a contract will be generated using the consumer and provider names, such as `baskets-ui-baskets-api.json`. This contract will need to be shared with the provider somehow so that the other half of the tests can take place. Contracts can be shared via the filesystem, by hosting them, or they can be published to Pact Broker.

### Verifying the Contract with a Provider

To verify a contract with a provider, we need to receive simulated requests. However, we need to return real responses from a real provider. This means that we need to stand up just enough of the provider so that real responses can be built and returned to the mock consumer. The provider tests are located in the `/baskets/internal/rest/gateway_contract_test.go` file.

For the Shopping Baskets module, we can start up the gRPC and HTTP servers, use test doubles for all of the application dependencies, and still be able to generate real responses. This provider will need to be running in the background so that the mock consumer can send the interactions that each consumer contract has defined.

### Example: Verifying with a Provider

When performing the verifications for simple APIs, we could start up the provider, configure the verifier, feed in contracts, and be done with our test:

```go
verifier.VerifyProvider(t, provider.VerifyRequest{
    Provider:                   "baskets-api",
    ProviderBaseURL:            "http://127.0.0.1:9090",
    ProviderVersion:            "1.0.0",
    BrokerURL:                  "http://127.0.0.1:9292",
    BrokerUsername:             "pactuser",
    BrokerPassword:             "***",
    PublishVerificationResults: true,
})

```

### Verifying Interactions with Provider States

The configured verifier in the prior listing will connect the mock consumer to the provider running on port 9090, then look for contracts published to our Pact Broker that belong to the `baskets-api` provider. If every interaction is verified for a contract, then we publish that success back to Pact Broker.

However, if any consumers have made interactions that make use of the provider state, as we did in our `baskets-ui` consumer using `given()`, then those states need to be supported; otherwise, the interactions cannot be verified.

### Handling Provider States

For example, to verify the `AddItem` endpoint, we will need to populate the test doubles with a basket, product, and store records. Using provider states will require communication and collaboration between teams. Documentation could be written that lists the state options that the provider supports. Failing these verification tests could block a provider from deploying, so the use of new provider states should be communicated and documented in all cases.

Provider states may optionally accept parameters that allow consumers to customize the interactions that they send and expect to receive back. The following state is used by the consumer:

```javascript
given("a basket exists", { id: basketId });
```

This is supported by the provider with the following:

```go
// ... inside provider.VerifyRequest{}
StateHandlers: map[string]models.StateHandler{
    "a basket exists": func(_ bool, s models.ProviderState)
        (models.ProviderStateResponse, error) {
        b := domain.NewBasket("basket-id")
        if v, exists := s.Parameters["id"]; exists {
            b = domain.NewBasket(v.(string))
        }
        b.Items = map[string]domain.Item{}
        b.CustomerID = "customer-id"
        if v, exists := s.Parameters["custId"]; exists {
            b.CustomerID = v.(string)
        }
        b.Status = domain.BasketIsOpen
        if v, exists := s.Parameters["status"]; exists {
            b.Status = domain.BasketStatus(v.(string))
        }
        baskets.Reset(b)
        return nil, nil
    },
},

```

## Expected State

Supporting the expected state for products and stores should be enough to verify the provider for the current UI consumer.

## Negative Quantity Test

When the `AddItem` endpoint is verified against the interaction with a negative quantity value, it will produce the following result:

### Scenario

**Given:**

- A store exists
- A product exists
- A basket exists

**When:**

- A request is made to add a product with a negative quantity

**Then:**

- The response has:
  - Status code: `400 (Bad Request)`
  - Headers:
    - `Content-Type: application/json`
  - Body: A matching body indicating the error

This result comes from the simulated request being sent to our real provider, which responded exactly how it would under normal conditions. The real response was then compared with the expected response, and it all passed.

With that, we have tested both a real request and a real response and have confirmed that they will work both as intended and expected. The REST API will work for every consumer that has created a contract, giving the provider confidence that it can be deployed without it breaking any consumers.

## Message Consumer and Provider Example

Contracts can also be developed by the consumers of asynchronous messages. We will want to expect messages from the consumers and verify that the providers will send the right messages. With asynchronous messaging, there will be no request portion to the test, but only an incoming message to process. Likewise, for the provider, we will not receive any request for a message, so the testing pattern changes slightly.

We will create tests for the messages that the **Store Management** module publishes and test message consumption in both the **Shopping Baskets** and **Depot** modules.

### Consumer Test Locations

The consumer tests are located in the following files:

- `/baskets/internal/handlers/integration_event_contract_test.go`
- `/depot/internal/handlers/integration_event_contract_test.go`

These two modules receive messages from the Store Management module, which we will discuss later.

### Expected Message Contract

For each message that a consumer expects to receive, we must create an expected message entry in our contract with the following code:

```go
message := pact.AddAsynchronousMessage()
for _, given := range tc.given {
    message = message.GivenWithParameter(given)
}
assert.NoError(t, message.
    ExpectsToReceive(name).
    WithMetadata(tc.metadata).
    WithJSONContent(tc.content).
    AsType(&rawEvent{}).
    ConsumedBy(msgConsumerFn).
    Verify(t),
)
```

The `GivenWithParameter()` and `ExpectsToReceive()` methods should be familiar to you if you read through the REST example.

`WithJSONContent()` is one of several methods we can use in Go to provide the expected message to the test. The content that we provide as our expected content is built using matchers. We can also use `WithMetadata()` to provide expectations for the headers or extra information that is published along with the content. This can be seen in the following example for the test of the `StoreCreated` event:

```go
metadata: map[string]string{
    "subject": storespb.StoreAggregateChannel,
},
content: Map{
    "Name": String(storespb.StoreCreatedEvent),
    "Payload": Like(Map{
        "id":       String("store-id"),
        "name":     String("NewStore"),
        "location": String("NewLocation"),
    }),
}

```

The `AsType()` method is a convenient way to convert the JSON that results from the matchers into something we can more easily work with and is optional.

Contract testing messaging will not use a mock provider or consumer, which is what we did in the REST example. The consumers will only be receiving messages and are not expected to send anything back. We will not be using a mock provider this time; instead, we will use a function that we provide to `ConsumedBy()` to test that our expected message will work.

The idea remains the same as in the REST example: we want to test that the message can be consumed. If it cannot, then we need to fix the message, application, or test.

To test that the events we receive work, we will need to turn `rawEvent` into an actual `ddd.Event` event, which means also converting the JSON payload into a `proto.Message` protocol. First, we need to register the `storespb.*` messages using a JSON Serde instead of the Protobuf Serde we typically use:

```go
reg := registry.New()
err := storespb.RegistrationsWithSerde(
    serdes.NewJsonSerde(reg),
)

```

Then, in the function that we provide to the `ConsumedBy()` method, we will deserialize the JSON into the correct `proto.Message`:

```go
msgConsumerFn := func(contents v4.MessageContents) error {
    event := contents.Content.(*rawEvent)
    data, err := json.Marshal(event.Payload)
    if err != nil { return err }
    payload, err := reg.Deserialize(event.Name, data)
    if err != nil { return err }
    return handlers.HandleEvent(
        context.Background(),
        ddd.NewEvent(event.Name, payload),
    )
}

```

The test will fail if the built event is not handled as expected. For extra measure, we use mocks that are passed into the handlers to test whether the right calls are being made when we call down into the handlers.

The contracts that we produce from message testing will not contain interactions and cannot be verified using a provider test, which is what we used in the REST example. The providers will use the description and any provider states to construct the message that is expected by consumers. There will not be any requests coming in.

Like the Shopping Basket REST provider, we want to avoid manually generating the message and should stand up enough of the module to create messages for us. We should verify that the processes that produce messages will continue to produce the right messages into the right streams as the application changes.

Just as we did in the REST provider test, we will create a verifier that will connect to Pact Broker, fetch the contracts that belong to the provider, verify the messages, then publish the results of the verifications back to Pact Broker.

The Store Management module provider verification tests can be found in the `/stores/internal/handlers/domain_events_contract_test.go` file. The key differences between this test file and the one for the REST contracts are that we do not start any mock consumer or start the provider listening on any ports. Message verification will also require that we implement each description string that the consumers have used in their contracts, such as "a StoreCreated message," as a message handler:

![alt text](image-149.png)

Figure 10.10 – Verifying the StoreCreated message

To verify the `StoreCreated` message, we can make a call into `CreateStore()` that will fire off the domain event, which, in turn, publishes the expected message. Using a `FakeMessagePublisher` test double, we can retrieve the last published message to complete the verification process.

The message payload, `proto.Message`, is serialized using a JSON Serde, similar to what we used in the consumer tests. We need to use the same methods for encoding when we create these message tests, and JSON is currently the best option for the content that we want to verify. Other formats could be used, but the Pact tools support JSON the best and the matchers only work with JSON.

Our entire message handler for the `StoreCreated` event message looks like this:

```go
"a StoreCreated message": func(
        states []models.ProviderState,
    ) (message.Body, message.Metadata, error) {
    // Arrange
    dispatcher := ddd.NewEventDispatcher[ddd.Event]()
    app := application.New(
        stores, products, catalog, mall, dispatcher,
    )
    publisher := am.NewFakeMessagePublisher[ddd.Event]()
    handler := NewDomainEventHandlers(publisher)
    RegisterDomainEventHandlers(dispatcher, handler)
    cmd := commands.CreateStore{
        ID:       "store-id",
        Name:     "NewStore",
        Location: "NewLocation",
    }
    // Act
    err := app.CreateStore(context.Background(), cmd)
    if err != nil { return nil, nil, err }
    // Assert
    subject, event, err := publisher.Last()
    if err != nil { return nil, nil, err }
    return rawEvent{
            Name:    event.EventName(),
            Payload: reg.MustSerialize(
                event.EventName(), event.Payload(),
            ),
        }, map[string]any{
            "subject": subject,
        }, nil
},
```

The real, albeit `rawEvent` event, is returned, along with a map for the metadata containing the subject that the message, if it had been published, would have been published into.

With that, we have completed the message verification process. We have taken a contract containing the expected messages for a pair of consumers and verified them with the provider. The results are automatically published to Pact Broker. If configured, Pact Broker could then inform the CI/CD processes to allow deployments to proceed.

Contract testing allows us to test integrations between components very quickly and with a lot less effort than if we had used a more traditional integration test approach. We can test the integration between two components, but we still need to test the operations that span multiple interactions.

## Testing the application with end-to-end tests

The final form of testing we will cover is end-to-end (E2E) testing. E2E testing will encompass the entire application, including third-party services, and have nothing replaced with any test doubles. The tests should cover all of the processes in the application, which could result in very large tests that take a long time to complete:

![alt text](image-150.png)

Figure 10.11 – The scope of an end-to-end test

E2E testing takes many forms, and the one we will be using is a features-based approach. We will use Gherkin, introduced in Chapter 3, Design and Planning, to write plain text scenarios that should cover all essential flows throughout the application.

## Relationship with behavior-driven development

You can do behavior-driven development (BDD) without also doing E2E testing, and vice versa. These two are sometimes confused with each other or it’s thought that they are the same. BDD, as a practice, can be used at all levels of the testing pyramid and not just for the final E2E tests or the acceptance tests. Whether or not to also employ BDD, and perhaps TDD, is a tangential decision for any particular level of testing in your testing strategy:

![alt text](image-151.png)

Figure 10.12 – The double-testing loop with BDD and TDD

BDD is also associated with the Gherkin language, and it has become dominant because of how the user stories BDD uses are created. We will be using Gherkin to write our features and their related scenarios, but again, this does not mean that we will be doing BDD. Gherkin can also be used for our unit or integration tests. Instead of using table-driven tests or a library to run tests as a suite, they could be written as plain text tests.

## E2E test organization

Our first step in E2E testing is to create feature specifications and then record them in our feature test files using Gherkin. There is no standard for organizing these feature files, but if we consider that an application uses multiple repositories because it is a distributed application that also uses microservices, then organizing all of the features into a repository might make sense. We only have one repository, so we will organize all of the features and other E2E-related test files under `/testing/e2e`.

## Making executable specifications out of our features

To make a feature file an executable specification, we will use the `godog` library, which is the official Cucumber ([https://cucumber.io](https://cucumber.io)) library for Go. With this library, we can write a `TestEndToEnd` function that will be executed using the `go test` command.

We will also need clients for each of the REST APIs. Normally, E2E tests would involve interacting with some end user UI, but our little application has none to work with at the moment. The REST clients can be generated using the `go-swagger` ([https://github.com/go-swagger/go-swagger](https://github.com/go-swagger/go-swagger)) tool, which can be installed along with the other tools we have used in this book by running the following command from the root of the code for this chapter:

```bash
make install-tools
```

The actual command to generate the clients is then added to the `generate.go` file for each module. The added command looks something like the following, with added line breaks to make it easier to read:

```go
//go:generate swagger generate client -q
  -f ./internal/rest/api.swagger.json
  -c storesclient
  -m storesclient/models
  --with-flatten=remove-unused
```

The generate command in the previous listing will create an entire REST client that is ready to be pointed at the Store Management REST API.

The final step of turning features into executable specifications is to implement each step and then register the implementation with the library.

## Example step implementation

Let’s say we have the following feature:

```gherkin
Feature: Register Customer
  Scenario: Registering a new customer
    Given no customer named "John Smith" exists
    When I register a new customer as "John Smith"
    Then I expect the request to succeed
    And expect a customer named "John Smith" to exist
```

We have four steps that we need to implement and register. To implement the registration of a new customer, we can start with a function signature, like this:

```go
func iRegisterANewCustomerAs(name string)
```

The string that is enclosed within the double quotes would be passed as the `name` parameter. Steps can have several parameters, and those parameters can be of several different Go types. Gherkin Docstrings and Tables are supported and can be passed in as well. The name of the function does not matter to the library and can be anything.

The function can be standalone or be part of a struct if you want to capture and use some test state, for example. We can also have an error return value if the step should fail:

```go
func iRegisterANewCustomerAs(name string) error
```

After we have implemented our step, we will need to register it so that when `godog` runs across the step statement, it knows what function will be expected to handle it:

```go
// ctx is a *godog.ScenarioContext
ctx.Step(
    `^I register a new customer as "([^"]*)"$`,
    iRegisterANewCustomerAs,
)
```

The step statements may be provided as strings and are interpreted as regular expressions, or directly as a compiled `*regexp.Regexp`. This is so that the parameters can be parsed out and passed into the step function.

## What to test or not test

E2E testing sits very high on the testing pyramid, and we should not try to write features covering everything that the application does or can do. Start with the critical flows to the business and then go from there. The identified flows will have several tests associated with them, not just one. You will want to consider what conditions can affect it and write tests to cover those conditions.

Some flows may not automate very well and should be left for the testers to run through manually.

## Summary

Testing an event-driven application is no harder than testing a monolithic application when you have a good testing strategy. In this chapter, we covered the application, domain, and business logic using unit tests. These tests make up the bulk of our testing force. We follow up our unit tests with integration tests, which help uncover issues with how our components interact. Using tools such as Testcontainers-Go can help reduce the effort required to run the tests, and using libraries such as the Testify suite can help reduce the test setup and teardown complexities.

A distributed application, whether it is event-driven like ours or synchronous, gains a lot from including contract testing in the testing strategy. Having confidence in how you are using or have made expectations of a provider without the mess and complexities of standing the provider up is a time saver many times over. Finally, including E2E testing in any form will give the team and stakeholders confidence that the application is working as intended.

In the next chapter, we will cover deploying the application into a Kubernetes environment. We will be using Terraform so that our application can be deployed to any cloud provider that provides Kubernetes services. We will also break a module out of the monolith into a microservice so that we can deploy it.

# 11 Deploying Applications to the Cloud

In this book, we have worked with the MallBots application as a modular monolith and have only experienced running it locally using Docker Compose. In this chapter, we will be breaking the application into microservices. We will update the Docker Compose file so that we can run either the monolith or the microservices. Then, we will use Terraform, an Infrastructure as Code (IaC) tool, to stand up an environment in AWS and deploy the application there.

In this chapter, we are going to cover the following topics:

- Turning the modular monolith into microservices
- Installing the necessary DevOps tools
- Using Terraform to configure an AWS environment
- Deploying the application to AWS with Terraform

## Technical requirements

You will need to install or have installed the following software to run the application or to try the examples:

- The Go programming language version 1.18+
- Docker
- The Kubernetes CLI tools
- Terraform
- The AWS CLI
- The PostgreSQL CLI tools

We have a lot more tool requirements for this chapter and will be covering download locations and installation within the chapter for each new tool. The code for this chapter can be found at [GitHub - Chapter 11](https://github.com/PacktPublishing/Event-Driven-Architecture-in-Golang/tree/main/Chapter11).

## Turning the modular monolith into microservices

Our application, while it is modular, is a monolith. It is built as a single executable and can be deployed as a single application. There is nothing wrong with that, but faced with scaling issues, we have only one knob we can adjust. If we broke the application up by turning each module into its own microservice, then when faced with scaling issues, we would have finer control over how the application can be deployed to support the load.

Turning our application into microservices will have many steps to it but will not be difficult:

1. We will need to refactor the monolith construct used to initialize each module.
2. We will make some small updates to the composition root of each module.
3. We will then update each module so it can run as standalone.
4. After we are done with these steps, we will update the Docker Compose file and make other small changes so that the two experiences, running the monolith or running the microservices, are the same.

## Refactoring the monolith construct

Our motivation for updating this part of the application is so that we can continue to run the monolith after we have turned each module into a microservice:

![alt text](image-152.png)

Figure 11.1 – Docker Compose with either a monolith or microservices

The monolith is built using the `/cmd/mallbots` main package. Up to this point, we have used a local `app` struct in that package to provide each module the resources that they required. The unexported `app` struct implements the `Monolith` interface, and this interface was used in each module’s `Startup()` method.

## Docker Compose version

The Docker Compose command, `docker compose`, that I am using is available from the Compose V2 release. If this command is not available, you can use the older version by putting a hyphen between the words as follows: `docker-compose`. The arguments used in the examples will not change when using the older version of the command.

Using the `app` struct as a template, we will create a new shared `System` struct in a new `/internal/system` directory and package:

![alt text](image-153.png)

Figure 11.2 – The types and interfaces of the system package

This new package also duplicates the interfaces that were found in the monolith package. From the monolith package, the old `Monolith` interface is renamed `Service` because it will serve a new, more general service need.

We can see from Figure 11.2 that the `System` struct has a lot of the same functionality, with some new exported methods, as the original `app` struct that it will be replacing. We did not bring over anything to do with managing the modules. Working with the modules will remain a monolith-only concern, and we use the following to reimplement the existing functionality for `/cmd/mallbots`:

```go
type monolith struct {
    *system.System
    modules []system.Module
}

func (m *monolith) startupModules() error {
    for _, module := range m.modules {
        ctx := m.Waiter().Context()
        err := module.Startup(ctx, m)
        if err != nil { return err }
    }
    return nil
}
```

The original `app` was initialized using some small functions. This initialization can be moved into functions alongside or as methods on `System`. These can all be called from a new constructor for `System` such as the following:

```go
func NewSystem(cfg config.AppConfig) (*System, error) {
    s := &System{cfg: cfg}
    if err := s.initDB(); err != nil { return nil, err }
    if err := s.initJS(); err != nil { return nil, err }
    s.initMux()
    s.initRpc()
    s.initWaiter()
    s.initLogger()
    return s, nil
}
```

The original `monolith main.go` file can now be switched over to use `System` instead of `app`, and the unused functions and the `monolith.go` file are removed. The `internal/monolith` directory can also be removed. Any lingering references to its package will be addressed in the upcoming section when we turn our attention to the modules.

## Updating the composition root of each module

Every module uses a `Startup()` method to initialize itself to run with the resources that the monolith has provided. Our update will be a small one. We will be moving the code within `Startup()` to a new `Root()` function. Then, we create a call to it from `Startup()`, and it will be as though nothing has changed:

```go
func (m *Module) Startup(
    ctx context.Context, mono system.Service,
) (err error) {
    return Root(ctx, mono)
}

func Root(
    ctx context.Context, svc system.Service,
) (err error) {
    // …
}
```

This simple change will allow us to reuse the composition root code for the other method of running the module, running it as a standalone microservice. We do not need to maintain the composition root in this way if we do not want to continue to run the monolith. If a real-world application were to be organized this way and the opportunity presented itself, why not keep the option to run a single process? Being able to continue to run the application as a monolith would allow us to avoid some of the trade-offs with a microservices architecture. For example, local development becomes more resource-intensive because more resources will be required to have each service running. Also, attaching a debugger to debug a single process is much easier than attaching multiple debuggers to multiple processes.

## Making each module run as a service

Each module will be made to run standalone by adding `/module/cmd/service` and a new `main` package to them. These additions are little more than copies of the monolith version. We remove anything to do with the management of modules and are left with the following:

```go
func main() {
    if err := run(); err != nil {
        fmt.Printf(
            "baskets exited abnormally: %s\n", err,
        )
        os.Exit(1)
    }
}

func run() (err error) {
    var cfg config.AppConfig
    cfg, err = config.InitConfig()
    if err != nil { return err }
    s, err := system.NewSystem(cfg)
    if err != nil { return err }
    defer func(db *sql.DB) {
        if err = db.Close(); err != nil { return }
    }(s.DB())
    err = s.MigrateDB(migrations.FS)
    if err != nil { return err }
    s.Mux().Mount("/",
        http.FileServer(http.FS(web.WebUI)),
    )
    err = baskets.Root(s.Waiter().Context(), s)
    if err != nil { return err }
    fmt.Println("started baskets service")
    defer fmt.Println("stopped baskets service")
    s.Waiter().Add(
        s.WaitForWeb,
        s.WaitForRPC,
        s.WaitForStream,
    )
    return s.Waiter().Wait()
}
```

We replaced the setup of the modules with a single call to this module’s `Root()` function.

Thanks to moving the bulk of the initialization of the system to the constructor, starting up the monolith or each service does not take much. Again, we must consider what trade-offs were made by refactoring things this way. If the microservices begin to diverge in the resources that they need, then we may end up initializing resources for dependencies that we do not have. `System` is a simple construct that starts up everything the same way – when the need arrives, it can be updated to be smarter about what should be initialized and what should not.

Every module could be run standalone at this point, but we would run into a few issues if we tried to copy the monolith service for each new service into the `docker-compose.yml` file.

Running our services and having the same experience as running the monolith will require a few more updates to be made.

## Updates to the Dockerfile build processes

We have only a single `Dockerfile` that builds the monolith. Going forward, we also need a way to compile the individual services. To accomplish this, I will use an additional `Dockerfile` that will make use of build arguments to target the right service to build.

The new Dockerfile will be named `Dockerfile.microservices` and live alongside the current one in `/docker`:

```Dockerfile
ARG svc
FROM golang:1.18-alpine AS builder
ARG svc
WORKDIR /mallbots
COPY go.* ./
RUN go mod download
COPY .. ./
RUN go build -ldflags="-s -w" -v -o service \
    ./${svc}/cmd/service

FROM alpine:3 AS runtime
COPY --from=builder /mallbots/docker/wait-for .
RUN chmod +x /wait-for
COPY --from=builder /mallbots/service /mallbots/service
CMD ["/mallbots/service"]
```

This is a multi-stage Dockerfile. In our first stage called `builder`, we compile the service into a binary. In the second stage, we copy the `wait-for` utility, which is used to wait for the database to be available, and the newly compiled binary. By using this Dockerfile, we keep the containers we produce very small, which helps with transferring them and loading them, among other things.

To build the specific service we want, we use the `--build-arg=svc=<service>` command-line argument with the `docker build` command as follows:

```bash
docker build -t baskets --file docker/Dockerfile.microservices --build-arg=svc=baskets .
```

This command would build the Shopping Baskets microservice and make it available as `baskets` in our Docker repository.

## Updates to the Docker Compose file

We will need to update the `docker-compose.yml` file so that each microservice can be started much like the monolith was. First, we need to add in each service using a block of YAML such as the following:

```yaml
baskets:
  container_name: baskets
  hostname: baskets
  image: baskets
  build:
    context: .
    dockerfile: docker/Dockerfile.microservices
    args:
      service: baskets
  ports:
    - "8080:8080"
  expose:
    - "9000"
  environment:
    ENVIRONMENT: development
    PG_CONN: <DB CONNECTION PARAMS>
    NATS_URL: nats:4222
  depends_on:
    - nats
    - postgres
  command:
    - "./wait-for"
    - "postgres:5432"
    - "--"
    - "/mallbots/service"
```

Similar blocks are added for the other modules-turned-microservices that we want to start up (in total, nine new blocks of YAML are added). Secondly, we want to be able to start either the monolith or the microservices version of our application. To do that, we can use the **profiles** feature of Docker Compose to selectively start services.

At the end of the monolith services block, we append the following YAML:

```yaml
services:
  monolith:
    # ... existing YAML
    profiles:
      - monolith
```

We can do the same for each new service, except using `microservices` instead:

```yaml
services:
  baskets:
    # ... existing YAML
    profiles:
      - microservices
```

With those last edits made to the `docker-compose.yml` file, we can start the monolith or start the microservices version of our application.

### Starting the monolith

Running the following command will run NATS, PostgreSQL, the Pact Broker, and then only the monolith service:

```bash
docker compose --profile monolith up
```

It is the same experience we are used to running, only that now we also need to include the `--profile monolith` part to get it.

### Starting the microservices

Running the following command will appear the same at first, with a lot more containers to build:

```bash
docker compose --profile microservices up
```

However, it will fail to run, ending with the following error message:

```bash
Bind for 0.0.0.0:8080 failed: port is already allocated
```

![alt text](image-154.png)

Figure 11.3 – Host and container ports for docker compose services

To give each microservice a unique but memorable new host port, we will use a sequence starting with the baskets entry down to the stores entry. For the baskets entry, we will use port 8081 and for stores, we will be using 8089 for its host port. The container port should remain as it is for all microservices.

Running the compose command again starts up the environment. Requests need to be sent to the correct port for the service now. If we attempt to open the Swagger UI, we will run into our second problem. We cannot load the OpenAPI specifications as we could before when we were running the monolith. The local specification for each service can be loaded, but we will not be able to view them all as we could before:

![alt text](image-155.png)

Figure 11.4 – The Swagger UI experience is broken

# Adding a Reverse Proxy to Docker Compose

Our fix has allowed our microservices to run, but the overall experience is far from the same as it was with the monolith. To fix the current problem with loading the OpenAPI specifications and return the experience to what it was before, we need to add a reverse proxy service to the `docker-compose.yml` file.

A reverse proxy will take the requests we send in and direct them to one of our microservices. The client will only interact with the reverse proxy and will not be aware of the microservices behind it.

## Adding a Reverse Proxy to the Compose Environment

We can quickly set up a reverse proxy using **Nginx**. Nginx is a popular web server, reverse proxy, and load balancer application. We only need to set up a reverse proxy today, and thankfully, it is going to be quite easy to do.

### Step 1: Define the Configuration File

First, we define a configuration file for the application called `/docker/nginx.conf`:

```nginx
worker_processes 1;
events { worker_connections 1024; }
http {
  sendfile on;

  upstream docker-baskets {
    server baskets:8080;
  }

  # ... plus upstream blocks for each other microservice

  server {
    listen 8080;

    # Reverse proxy for baskets service
    location /api/baskets {
      proxy_pass     http://docker-baskets;
      proxy_redirect off;
    }

    location /baskets-spec/ {
      proxy_pass     http://docker-baskets;
      proxy_redirect off;
    }

    # ... plus location block pairs for each other microservice

    # Reverse proxy for the Swagger UI files
    location / {
      proxy_pass     http://docker-baskets;
      proxy_redirect off;
    }
  }
}
```

I have only included the baskets microservice as an example in the configuration file example, but each microservice would need to have an upstream configuration and a pair of location configurations so that the reverse proxy could properly redirect the requests to where they need to go. A final location block is used to serve the Swagger UI from any microservice.

Secondly, we need to remove the port configurations for each microservice and add the reverse proxy as a new service to the docker-compose.yml file:

```bash
reverse-proxy:
container_name: proxy
hostname: proxy
image: nginx:alpine
ports: - '8080:8080'
volumes: - './docker/nginx.conf:/etc/nginx/nginx.conf'
profiles: - microservices
```

## Fixing the Reverse Proxy and gRPC Connections

We only want the reverse proxy to start up with the other microservices, so it is also given the **microservices profile**. The configuration file we created for Nginx is mounted at the appropriate place for the application to find it.

### Initial Setup

At this point, we've fixed the initial experience, and the Swagger UI is usable again. However, if we run the **E2E tests**, we'll encounter a final issue. The services that use gRPC fallbacks are still dialing into gRPC connections that point back to the local gRPC connection.

When the **Shopping Baskets** service tries to make a fallback call to the **Store Management** service to locate a product, it ends up calling itself. We need to provide the correct addresses for these services to ensure that the connections are properly established.

### Fixing the gRPC Connections

To fix this, we will provide the addresses of the gRPC servers used by other services through a new environment variable called `RPC_SERVICES`. This value will contain a map of service names and their respective addresses. The format will look like this:

```bash
RPC_SERVICES="STORES=stores:9000,CUSTOMERS=customers:9000"
```

## Updating the RPC Configuration for gRPC Service Resolution

To ensure proper gRPC connections are made, we need to update the `/internal/rpc/config.go` file with the necessary logic for resolving service addresses. Below is the code we will add to handle the `RPC_SERVICES` environment variable:

### Modifying `config.go`

In the `/internal/rpc/config.go` file, add the following code:

```go
type RpcConfig struct {
    // ... snipped existing fields
    Services
}

type Services map[string]string

// Service returns the address of the specified service.
func (c RpcConfig) Service(service string) string {
    if address, ok := c.Services[service]; ok {
        return address
    }
    return c.Address() // Fallback to default address if service is not found
}

// Decode parses a string containing service name/address pairs and populates the Services map.
func (s *Services) Decode(v string) error {
    services := map[string]string{}
    pairs := strings.Split(v, ",")
    for _, pair := range pairs {
        p := strings.TrimSpace(pair)
        if len(p) == 0 {
            continue
        }
        kv := strings.Split(p, "=")
        if len(kv) != 2 {
            return fmt.Errorf("invalid pair: %q", p)
        }
        services[strings.ToUpper(kv[0])] = kv[1]
    }
    *s = services
    return nil
}
```

The Services type will use the custom decoder, `Decode()`, to turn the service pairs into a usable map. A `Service()` method is also added to the `RpcConfig` struct for convenience so it will be easier to fetch the correct service address when we need to.

Now, we need to update the Shopping Baskets and Notifications composition roots to use the correct address for the gRPC connection that they are dialing into. Here is the updated connection from Shopping Baskets to the Store Management service:

```go
container.AddSingleton("storesConn",
     func(c di.Container) (any, error) {
          return grpc.Dial(
               ctx,
               svc.Config().Rpc.Service("STORES"),
          )
     },
)
```

Now, all that is left to do is to add the new `RPC_SERVICES` environment variable to each service in the `docker-compose.yml` file:

```yaml
# ... snipped other configuration
environment:
  # ... snipped other variables
  RPC_SERVICES: "STORES=stores:9000
    ,CUSTOMERS=customers:9000"
```

Rebuild the microservice containers and restart the compose environment, and now our E2E tests all pass again. Likewise, trying to add an item to a basket with an invalid product identifier in the Swagger UI also behaves as expected, if you care to verify things are working that way.

Our application can now run as a monolith or as a suite of microservices. To recap, these are the steps we took to get here:

1. We refactored the monolith startup code into a shared service startup library.
2. Each module got a new service command, with an updated composition root.
3. To build the new services, a new Dockerfile was created that used build arguments so that a single Dockerfile could be used for all services.
4. The `docker-compose.yml` file was updated to include each service, and we used Docker Compose profiles to start either the monolith or the microservices.
5. A reverse proxy was added so we could reach all services with a single address.
6. We updated the gRPC configuration so we could provide the right gRPC server addresses to the gRPC clients.

We will also want to run our application in the cloud, and we have many providers to choose from. Amazon Web Services (AWS), at https://aws.amazon.com, is the oldest, largest, and most well-known cloud provider. There are other big names to choose from, such as Google Cloud Platform (GCP), at https://cloud.google.com, and Azure Cloud, at https://azure.microsoft.com. Smaller or regional cloud providers are also available, such as DigitalOcean (https://www.digitalocean.com), OVHcloud (https://ovhcloud.com), and Hetzner (https://www.hetzner.com).

From all of these options, we will be using AWS, partly because of its status as the top cloud provider and partly because it is also the one I know best. However, before we do that, we will need to install and get a little familiar with some new tools.

### Installing the necessary DevOps tools

The plan is to deploy the application in its microservices form to AWS. For most developers, learning about every service offering in AWS is not something they focus on – taking off their software developer hat and putting on their system administrator hat, so to speak. To make things easier, we will be relying on an application called Terraform, which is an IaC (Infrastructure as Code) tool. We will be able to define what our application needs with code and then let it do the heavy lifting of pulling all the right levers and pushing all the right buttons for us.

We will also need a few more tools to help us:

- **The AWS CLI**, `aws`, is how we will authorize ourselves with AWS.
- **Helm** is a tool that will let us use packages called Charts to deploy some complex machinery into Kubernetes.
- We will be using a **PostgreSQL** database in the cloud and will want the PostgreSQL client `psql` installed to help set it up.
- To view our Kubernetes cluster, we will use an application called **K9s**, which is a Terminal UI (TUI) that makes it very easy to navigate around the cluster.
- We will also need a tool called **Make**, which is a small application runner that helps us turn large or multistep commands into ones that are easy to remember and run.

If you do not already have these applications installed, I have two options for you to install them. The first option is to keep your local system clean of additional applications by using a Docker container with all of the applications already installed or to find and install them yourself.

If you are going to be following along and you are on Windows, I recommend the first option.

Regardless of which option you choose, you will also need an AWS account. Visit https://portal.aws.amazon.com/billing/signup to create a free account with AWS. Let us check both options.

### Installing every tool into a Docker container

This is the easier route to take and it also keeps your local system clean of any applications you are not likely to be using again. This option will compile a Docker container called `deploytools`, which will then be made available with a shell command alias.

To start, you need to either be using macOS or Linux or be able to open a PowerShell in Windows. A non-PowerShell Command Prompt in Windows will not work.

To start, go into `deployment/setup-tools` in your Terminal or PowerShell window.

You will now need to execute the right script for your OS. macOS and Linux users should run the following command:

```bash
source set-tool-alias.sh
```

PowerShell users should run this command:

```powershell
.\win-set-tool-alias.ps1
```

Both do the same things. During the first run, the `deploytools` container will be built; subsequent runs will rebuild the container only if it is missing or the Dockerfile has changed. It will then set up the `deploytools` command. This is a temporary command that will stop working when you close the Terminal or window. To get it back, you just need to run the correct script command again from the `deployment/setup-tools` directory.

Once you have your alias, you can verify it works by running the following:

```bash
deploytools terraform -version
```

You should see the Terraform version printed out, looking something like this:

```bash
Terraform v1.2.9
on linux_amd64
```

If you see that, then the container and command are ready for use.

When you are using this option, you need to prefix the commands in the following sections with the `deploytools` command. Let’s take this command as an example:

```bash
deploytools aws configure

```

Turn it into this command:

```bash
deploytools aws configure

```

Speaking of which, you will still need to configure your AWS credentials. You will find instructions to do so in the **Creating and configuring your AWS credentials** section that comes a little later.

### Next, let’s look at the other option.

### Installing the tools into your local system

We will need a few tools to support our plans to deploy the application as microservices in AWS. All of these tools are available for Linux, macOS, and Windows OSs; only the download location or installer will be different. Using them will be the same.

#### Installing and configuring the AWS CLI

The first tool we will want to install is the **AWS CLI**. You can find instructions for your OS at [AWS CLI Getting Started](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html).

Once you have downloaded the tool, we need to set up a user and configure it in your shell.

### Creating and configuring your AWS credentials

We will be using this user to access the AWS services from Terraform and the CLI:

1. Sign into your AWS account.
2. Access the **Identity and Access Management (IAM)** service.
3. Click on **Users** on the sidebar.
4. Click on the **Add User** button to begin creating your user.
5. Give the user a name, such as `mallbots_user`, choose **Access Key** as the credential type, and then click on the **Next** button.
6. Choose to attach existing policies directly, select the **AdministratorAccess** permission, and then click on the **Next** button.
7. You may add any tags you wish – then, click on the **Next** button.
8. Confirm that the user has **Programmatic access** and is using the **AdministratorAccess** policy and then click on the **Create user** button.
9. On this next screen, you should download the credentials as a `.csv` file. **Do not close this page before getting the credentials**, as this is the only time you will be given the opportunity to retrieve them.

When you are done testing the deployment of the application and have properly removed all of the resources, you can then remove this user. There are no charges incurred by keeping this user account present.

Next, we will use `aws-cli` to configure your shell with the credentials you just downloaded. Locate and open the `.csv` file with your credentials to use them with the following command:

```bash
aws configure
```

You will be prompted to enter the user’s access key ID, secret access key, default region, and default output format. The keys are found in the `.csv` file you downloaded. You may leave the default values blank if you wish. For the default region, there are many options to select from and you should select the region that is nearest to you.

To verify that you have entered the credentials correctly, use the following command to fetch a list of users from IAM:

```bash
aws iam list-users
```

If you see a list of users, including the user you just made, then `aws-cli` is ready to use.

### Installing Terraform

The installers for Terraform can be found at [Terraform Installation Guide](https://learn.hashicorp.com/tutorials/terraform/install-cli). Version 1.2.8 was used for the examples in this book.

Once Terraform has been installed for your operating system, it requires no configuration and is ready to use.

### Installing Helm

Some of the configurations we will be using will be in the form of **Helm charts**, which are collections of files that describe Kubernetes resources. Instead of creating new custom Terraform code, we can rely on the battle-tested community versions.

Helm install instructions can be found at [Helm Installation Guide](https://helm.sh/docs/intro/install/). Version 3.9.4 was used for the examples in this book.

### Installing the tools to access Kubernetes clusters

We will be deploying the application into a Kubernetes cluster and instead of navigating the AWS console to keep an eye on things, we will install some tools to make things easier.

The first tool we will install is **K9s**, a TUI application that makes it very easy to browse the various resources, such as Pods, Ingresses, and Services, that will be part of the Kubernetes cluster. The install instructions can be found at [K9s Installation Guide](https://k9scli.io/topics/install/). Version 0.26.3 was used for the examples in this book.

The second optional tool to install is **kubectl** and the installers can be found at [kubectl Installation Guide](https://kubernetes.io/docs/tasks/tools/). Version 1.25.0 was used for the examples in this book.

Other tools will need a Kubernetes configuration before they can be used and we will be able to fetch one after we deploy the infrastructure to AWS.

### Installing the tools to initialize PostgreSQL

We will be using the **PostgreSQL CLI tool** `psql` to initialize the databases and set up the schemas and users after we deploy the application infrastructure. The `psql` tool comes with the PostgreSQL server installation. We do not need to install the PostgreSQL server, so if you are given the option, choose to only install the command-line tools. The PostgreSQL installers can be found at [PostgreSQL Download Page](https://www.postgresql.org/download/). Not every installer will put the `psql` tool in your path; you will have to either move the file or add the install location to your path. Version 14.5 was used for the examples in this book.

We now have our environment ready to execute the deployment scripts and configurations to deploy the infrastructure and application up to AWS.

### Using Terraform to configure an AWS environment

The **MallBots** application is going to be run from **AWS Elastic Kubernetes Service (EKS)**, a managed Kubernetes environment. The IaC (Infrastructure as Code) to create the infrastructure is going to be found in the `/deployment/infrastructure` directory.

We will be configuring a small typical AWS environment across two Availability Zones (AZs):

![alt text](image-156.png)

Figure 11.5 – Our AWS infrastructure

In the **infrastructure** directory, there are several Terraform files. Altogether, they are going to be used to set up the following in AWS:

- **Docker repositories** with Elastic Container Service (ECS). We will be uploading the built microservice images here.
- A **Kubernetes cluster** in EKS. We will be deploying our application here from images stored in ECS.
- A **PostgreSQL database** using Relational Database Service (RDS). A single instance will serve all of the microservice databases and schemas.
- Additional components such as a **Virtual Private Cloud (VPC)** and its subnets, security groups, roles, and policies to both permit and lock down access.

When running the next commands, you need to be in the `/deployment/infrastructure` directory.

### Preparing for the deployment

Terraform is capable of deploying thousands of different kinds of resources, but it cannot do it by itself. We will need to install the libraries that our specific project needs, and to do that, we need to run the following command:

```bash
make ready
```

This will run both the `terraform init` and `terraform validate` commands. The `init` command will download the libraries and executables needed by the scripts that have been written to build our environment. The `validate` command will also validate our scripts are correct.

The next Terraform command that we run is going to ask for some input from us. Instead of providing the input each time we run it, we can provide the values automatically with a variables file. Create a file named `terraform.tfvars` and put the following lines into it:

```hcl
allowed_cidr_block = "<Your Public IP Address>/32"
db_username = "<Preferred DB username>"
region = "<Your Nearest Or Preferred AWS Region>"
lb_image_repository = "<AWS Regional Image Repository>"
```

The first variable is used to limit access to the resources that are created to your IP or a block of IP addresses. If you only want to allow your public IP, then keep `/32` at the end – for example, `192.168.13.13/32`. The DB username will be used along with a generated password to connect to the database to initialize it in a subsequent step. The final two variables should be set to the AWS Region that works best for you. You can find which repository to enter at [AWS EKS Add-ons Images](https://docs.aws.amazon.com/eks/latest/userguide/add-ons-images.html).

It is not critical you create this file, but if you do not, then you will be prompted for the values each time you create a new Terraform deployment plan.

### A look at the AWS resources we are deploying

The AWS resources that we will be deploying are broken up into different files, so let’s run through each file and cover the major resources that will be installed and configured by the Terraform code within them:

- **Application Load Balancer (ALB)**: The `alb.tf` file sets up a service account on the Kubernetes cluster that will be used by the ALB. The file also contains a Helm resource that will install the ALB using a Chart.
- **Elastic Container Registry (ECR)**: The `ecr.tf` file sets up private image registries for each of the nine services we will be deploying. It will also build and push each service up into the newly created registries.
- **EKS**: The `eks.tf` file is responsible for creating our Kubernetes cluster. It makes use of a Terraform module, which is a collection of other Terraform scripts, to build the necessary resources from one resource definition. Some AWS IAM policies and roles are configured in this file for the cluster to support the installation of the ALB.
- **RDS**: `rds.tf` will set up a serverless PostgreSQL database and make it available to the Kubernetes cluster. The database will also be accessible by us or anyone else who has an IP address allowed by the `allowed_cidr_block` value.
- **Security groups**: The `security_groups.tf` file will set up our security group that will limit access to our resources from the internet. Whatever `allowed_cidr_block` we provide will be the only set of IP addresses that will be able to reach our database, cluster, and any other resources we have set up.
- **VPC**: The `vpc.tf` file will create a set of networks, connect them with routing, and also use our security group to limit access to them. These networks will be used by the Kubernetes cluster to deploy Pods, by the database, and by the application. The VPC will be installed across two AZs to improve our deployed resource resiliency by being installed in different data centers.

I have included the URL as a comment above each resource and module that is being used so you can visit and learn more about the resources being installed or learn about what other configuration options are available.

### Next up is to deploy all of this infrastructure into AWS.

### Deploying the infrastructure

To create our deployment plan using the variable provided in the `terraform.tfvars` file, or when prompted, and to deploy it into AWS, we run the following command:

```bash
make deploy
```

This command will execute the plan and apply Terraform commands. These will be followed up with a command to fetch the cluster configuration so that we can connect to it with K9s. The plan that Terraform creates will contain approximately 87 resources. During the apply stage, Terraform will make use of the plan and will immediately begin the process of creating, configuring, connecting, and verifying each resource. Terraform will do its best to create resources concurrently when it can, but this process will take some time to complete – around **15 to 20 minutes**.

### Usage costs warning

Running these Terraform commands will create AWS resources that are not covered by any free tiers. You will begin incurring usage costs from the moment you execute the `make deploy` command. You will continue to be charged until you destroy the infrastructure with `make destroy`. Running this demo for a few hours will cost roughly **$2 to $5** depending on the region that it is run in.

As it creates resources, Terraform will output logs of what is happening so that you are not left in the dark. You can also see some progress if you go into the AWS console and view the various services, such as EKS, RDS, and ECS.

If the process is interrupted or something times out, Terraform will end with an error. If it does, you can rerun the `make deploy` command to get things back on track in most cases.

When it is done, it will output any outputs we have defined to the screen as long as they are not marked sensitive. Some of these outputs will be used in our second phase of deploying the application.

### Viewing the Kubernetes environment

At this point, the infrastructure is completely set up. We can go into the AWS console to view various things, but if we try to view the Kubernetes cluster in EKS, it may say our user does not have permission to view the components. This is expected because we only gave the user we created permissions and not our main AWS account user. To view the cluster components, we will need to run the following command to bring up the K9s UI:

```bash
k9s
```

It might take a moment to load up completely, but after it is done loading, we should see something like this:

![alt text](image-157.png)

Figure 11.6 – The K9s terminal application showing the running Pods

To navigate around the components, you start the command with a colon and then the type of components you would want to view – for example, typing `:deployments` and hitting Enter will show the list of deployments in the cluster, and `:services` will show the running services.

To exit K9s, type `:quit` and then hit Enter.

If you are familiar with `kubectl` or would prefer to work from the CLI instead, then to view the list of deployments, you can use the following command:

```bash
kubectl get deployment -n kube-system
```

This will display a short list of deployments in the `kube-system` namespace. Likewise, we can view the list of services using this command:

```bash
kubectl get services -n kube-system
```

This would display a short list of services.

Using either K9s or kubectl, we should see some load balancer resources installed with `load-balancer-aws-load-balancer-controller` in the list of deployments and `aws-load-balancer-webhook-service` in the list of services. Seeing these means that we know our infrastructure is ready.

### Deploying the application to AWS with Terraform

To deploy the application, we will need to switch to the `/deployment/application` directory.

Similar to what we did for the infrastructure, we will prepare Terraform by installing the libraries that deploying the application will require by running the following command:

```bash
make ready
```

### Getting to know the application resources to be deployed

As we did for the infrastructure, we have broken up the resources we will be deploying into multiple files.

#### Database setup

For the database, we will initialize the shared triggers and that action can be found in the `database.tf` file.

#### Kubernetes setup

In Kubernetes, components can be organized into namespaces. This can help when you have multiple applications, when you have multiple users and want to restrict access, or when you are using the cluster for multiple purposes. Our application will be deployed into the `mallbots` namespace. In K9s, we can filter what we see by namespace to make it easier to locate just our application components.

As with our local development environment, the services will be using environment variables. Most of those variables are the same for each service, and in Kubernetes, we can create ConfigMaps for data that we want to share. A config map is created with the common environmental variables, such as `ENVIRONMENT`, `WEB_PORT`, and `NATS_URL`. We will pass this config map into each microservices deployment resource.

Lastly, in the `kubernetes.tf` file, we define an ingress on the ALB, for the Swagger UI. Just as with our local experience, we will be able to visit a single URL to access all of the microservices and Swagger.

#### NATS setup

In the `nats.tf` file, we create a deployment for NATS using the same container we used in the Docker Compose environment. A persistent volume claim, a little bit like the Docker volumes, is also set up for NATS to record its data. This way, if the deployment was restarted, we would not lose any messages. The NATS deployment is made available using a service component. A service defines how a deployment may be accessed.

#### Microservices setup

Each microservice is kept in its own file using a filename pattern such as `svc_<service>.tf`.

Instead of using a static database password as we do in our local environment, each service uses a randomly generated password. These passwords are generated each time we plan and deploy the application freshly. Updating the application and redeploying will reuse the password from the Terraform state data. The random passwords are used within the initialized service database resource.

Kubernetes config maps are not good places to put secrets such as database passwords. They are not stored with any encryption, so it is possible the data could be seen. For things such as passwords, we have secrets that do use encryption and are less likely to be seen or understood if they are leaked. For the `PG_CONN` environment variable, we create a secret and store each microservice separately.

As with NATS, each microservice has a deployment and a service component. Unlike NATS, most services also have an ingress setup so that they are also available at the exposed address provided by the ALB. Services such as `cosec` and `notifications` do not have any ingress defined because they do not expose any APIs.

### Deploying the application

To deploy the application in the waiting infrastructure, we run the following command:

```bash
make deploy
```

The application deployment consists of approximately 57 resources. This deployment will not take as long as the infrastructure deployment but will still clock in at around 5 to 10 minutes.

If you have K9s open, you can watch as the deployments come online and you can see the ingresses being added, using the `:deployments` and `:ingress` commands, respectively.

To view the list of deployments using `kubectl`, you would use the following command:

```bash
kubectl get deployment -n mallbots
```

The list of deployments should look like this, with different values in the AGE column after all the deployments are done:

```bash
NAME            READY   UP-TO-DATE   AVAILABLE   AGE
baskets         1/1     1            1           18m
cosec           1/1     1            1           18m
customers       1/1     1            1           18m
depot           1/1     1            1           18m
nats            1/1     1            1           19m
notifications   1/1     1            1           18m
ordering        1/1     1            1           18m
payments        1/1     1            1           18m
search          1/1     1            1           18m
stores          1/1     1            1           18m

```

Please note that we are viewing the deployments in the `mallbots` namespace and not the `kube-system` namespace this time.

When the deployment has been completed, Terraform will output the address you can find in the Swagger UI. We are not deploying the application to any particular domain, so this address will be generated. If you missed the address, you could retrieve it using this command:

```bash
terraform output swagger_url
```

Opening this Swagger UI will be exactly the same as the experience we have locally. That is why IaC and repeatable deployments are so popular.

The application and infrastructure will only be accessible to your IP address, but leaving it running will continue to cost you money. Thankfully, solving this is also made easy using Terraform.

## Tearing down the application and infrastructure

Running this application in Kubernetes and using the infrastructure resources will continue to rack up costs by the hour, so when you are done with the MallBots application, you should tear it all down. When Terraform makes changes to an environment, it keeps a state file – in our case, kept locally – so that it can minimize the changes it needs to make when the Terraform files are changed, and new `plan` and `apply` commands are run. The state is also used to locate the resources that need to be destroyed when we are done with them.

Start with the application deployment first. Go into the `/deployment/application` directory and run the following command:

```bash
make destroy
```

As with the deployment process, this can take some time to complete. When it does complete, we can run the same command from the `/deployment/infrastructure` directory.

After the second command completes, your AWS account should be back to how it was before we started this journey. You can verify by signing into your AWS account on the AWS console and by visiting RDS to make sure there are no database instances, ECR to verify that there are no repositories, and EKS to see that the cluster has been completely removed. Anything you find across AWS you can view tags for; if you see the `Application: MallBots` tag, then it was something left behind. I ran and reran the deployment and tear-down steps over a dozen times, and Terraform always did an excellent job restoring my account to how it was.

## Summary

In this chapter, we converted the modular monolith application into a microservices application. We modified the modules in such a way that we could continue to run the application as a monolith or with microservices. This is not exactly a goal most teams have, but we could do it, so we did. A real application would likely begin to diverge and maybe pick up new microservices that are written in different languages, which would make keeping the monolith around an unlikely outcome.

We also set up our environment to deploy our application into the cloud using either a containerized approach or installing the necessary tools directly onto our system. We used these tools to stand up the infrastructure that our application needed to be run on top of first. Then, as a second step, we deployed the application itself to AWS. The experiences between the locally running application as a monolith, as microservices, and as a cloud deployment remained exactly the same.

In the next chapter, we will be learning how to monitor the performance of our application and to track requests as they flow through it using causation and correlation identifiers.

## 12 Monitoring and Observability

In this final chapter, we will cover how to monitor the performance and health of the services and the application as a whole. The most common approach to monitoring is to use logging to record information and errors. Following logging, the second most common approach is to record metrics such as CPU usage, request latency, and more. We will be looking into these forms of monitoring and will also take a look at an additional form of monitoring, known as distributed tracing.

In this chapter, we will also introduce **OpenTelemetry** and learn about its goals, and how it works. We will then add it to our application to record each request as it works its way through the application.

We will end by looking at the tools that are used to consume the data produced by our monitoring additions – that is, **Jaeger**, **Prometheus**, and **Grafana**.

In this chapter, we are going to cover the following main topics:

- What are monitoring and observability?
- Instrumenting the application with OpenTelemetry and Prometheus
- Viewing the monitoring data

### Technical requirements

You will need to install or have installed the following software to run the application or try the examples provided:

- The Go programming language version 1.18+
- Docker version 20+
- Docker Compose version 2+
- The code for this chapter can be found at [GitHub - Chapter12](https://github.com/PacktPublishing/Event-Driven-Architecture-in-Golang/tree/main/Chapter12).

---

### What are monitoring and observability?

Most deployments perform monitoring using logging and metrics. This allows an organization to track the application’s performance, usage, and health. It is also used to detect failure states in an application. Monitoring is about reacting to the analysis of the data that is being collected:

![alt text](image-158.png)

Figure 12.1 – Basic monitoring of a service

Some examples of monitoring include the following:

- Kubernetes checking whether a container is still running or responding by performing a health check
- Tracking the query performance when swapping out one database for another
- Automatically scaling services based on the CPU and memory usage
- Sending alerts when the error rate of an endpoint exceeds a certain threshold

The data that is produced from your monitoring efforts is fed into dashboards so that basic questions can be answered. The data is also used to configure alerts so that when a problem is developing, staff can be notified to take the appropriate action.

Monitoring works with predetermined logs and metrics. Its weakness is dealing with the unexpected. If we know ahead of time that a process has the potential to consume large amounts of CPU, then we can include that in our monitoring; otherwise, it will be a blind spot. Putting this another way, if we can predict that we will have problems in a certain part of the application, then we can instrument it so that it can be monitored.

The data that is collected from the various monitoring efforts is often used in isolation. As a result, it can sometimes be difficult to correlate an event across different sets of data. Keeping the data isolated is not done by choice; the tools themselves are typically not designed to interact with other tools or other forms of data.

The purpose of monitoring is to answer "What happened?" and "Why did it happen?". However, when we need to correlate the data ourselves across different tools, it is not always easy. For example, let’s say that a team receives an alert about a service experiencing rising CPU usage. To determine the root cause, the team could look at related dashboards to determine a timeframe, and then search the logs to locate any errors during the timeframe reported in the dashboards. If the team fails to locate any errors, a new search to look over all the logs to spot a trend would be necessary. This is a typical approach many teams take, and it can be a workable solution for most applications. Distributed applications only make the problem of locating the root cause of an issue more difficult. With a distributed system, requests travel through multiple services and use a variety of communication methods.

As the application grows in complexity, so does the need to monitor for more things. Making accurate predictions about all of the places that will be problematic can be extremely difficult. You need instrumentation that will be able to answer the questions about the unknown unknowns or to provide answers without asking explicit questions. This is where **observability** and **distributed tracing** enter the picture.

---

### The Three Pillars of Observability

Observability is made up of three pillars. We discussed the first two – **logs** and **metrics** – in the previous section; in this section, we will be covering the third: **traces**. Traces are recordings of a request as it moves through the application.

Together with logs and metrics, traces give you a complete picture of the state of the application:

- **Logs** tell you why your application is in a given state.
- **Metrics** tell you how long your application has been in a given state.
- **Traces** tell you what is impacted by being in a given state.

A trace may begin with the client being at the first entry point in the application or even somewhere in between. The trace will be given some kind of identifier that is passed along as the request that it tracks makes its way through the application.

---

### How Tracing Works

We will work with an example where the trace starts as it enters the application’s backend. When a brand-new request comes in, an identifier is generated for it; for example, `abcd`. At the same time, **correlation** and **causation** identifiers are also assigned the same value:

![alt text](image-159.png)

Figure 12.2 – Tracing with request, correlation, and causation identifiers

The purpose of a **correlation identifier** is to correlate all requests back to a single originating request. A **causation identifier** is used to point back from a follow-up request to the request that came before it.

As the request makes its way through the application, the **correlation identifier** never changes. The **causation identifier** will always point back to the call that preceded it. Requests into the application can fork. Here, we follow the same rules that have already been laid out; no new rules are required to handle branches that can occur during a request.

These identifiers can then be logged with other log messages. If you are building your tracing implementation manually, then this is how you might record how a request flows through the application. You will not be able to construct a **span**, a representation of a call, or another unit of work with log messages alone.

Tools such as **Jaeger** can visualize a trace, giving you an entirely new view of your application that you can’t see from metrics or the logs themselves.

In a visualization of a trace, you can see the different spans along the **y-axis**, which could be different processes that were run in different components. Along the **x-axis**, you can see the element of **time** so that you can get a sense of how long it took to log those processes:

![alt text](image-160.png)

Figure 12.3 – Visualization of a request traced through the application broken into multiple spans

You could develop your own tracing implementation, but I would suggest otherwise. There is a lot more than simply being able to visualize an **icicle diagram**, or **upside-down flame graph**, of the different spans that make up the trace. Traces cannot be created from log messages, so you would be developing an entirely new instrumentation method for your application. Traces are also very information-rich and can be annotated with bits of information and even record errors that have occurred at specific points.

Thankfully, you do not need to start from scratch to instrument your application. The **OpenTelemetry project** ([https://opentelemetry.io](https://opentelemetry.io)) exists for this purpose, with the goal of merging the instrumentation for logging, metrics, and tracing into a single unified API.

## Instrumenting the application with OpenTelemetry and Prometheus

Our application has already been set up with a logger, but we need traces and metrics to achieve observability. The **OpenTelemetry project** aims to support all three (logging, traces, and metrics) in the Go SDK, but at the time of writing this book and version **v1.10**, only tracing is stable. So, we will leave our logger in place and interact directly with **Prometheus** for our metrics. We will begin with **OpenTelemetry** and distributed tracing.

### Adding distributed tracing to the application

Getting started with OpenTelemetry is very easy; first, we will need to create a connection to the **collector**. In our application, we will have one running and available at the default port. The monolith or microservices will use the following environment variables to configure themselves:

```env
OTEL_SERVICE_NAME: mallbots
OTEL_EXPORTER_OTLP_ENDPOINT: http://collector:4317
```

The OpenTelemetry SDK we will use will look for specific variables in the environment that all begin with the **OTEL** prefix, short for OpenTelemetry. The two variables shown in the preceding snippet are the minimum we will need to run our demonstration.

- `OTEL_SERVICE_NAME` should be set to a unique name for the application or component. Here, we are setting it to `mallbots` for the monolith application. For the services, we will use their package names.
- `OTEL_EXPORTER_OTLP_ENDPOINT` is set to the OpenTelemetry Protocol (OLTP) endpoint, which defaults to looking for the collector on **localhost**. In our case, we have set it to the hostname in the Docker Compose environment. We will be communicating with the collector via the OLTP protocol on the specified port (4317).

An **OpenTelemetry collector** is a vendor-agnostic service that provides instrumentation data collection, processing, and exporting functionality. A single collector can replace the need to run, configure, and connect to multiple agents to instrument your application.

### The demonstration is local only

There is no advantage to running the demonstration of the application instrumentation in AWS, so this demonstration is expected to be run **locally** in your **Docker Compose** environment.

### Initializing OpenTelemetry in the code

Now, we can update the `/internal/system` code to initialize the connection to the collector. This involves configuring the **OpenTelemetry SDK** to collect and export tracing data. Here is how you can update your system's initialization function:

```go
func (s *System) initOpenTelemetry() error {
    // Create an exporter to send tracing data to the OpenTelemetry collector.
    exporter, err := otlptracegrpc.New(
        context.Background(),
    )
    if err != nil {
        return err
    }

    // Create a Tracer Provider with the batcher exporter.
    s.tp = sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
    )

    // Set the global Tracer Provider for OpenTelemetry.
    otel.SetTracerProvider(s.tp)

    // Set the text map propagator for context propagation (trace context & baggage).
    otel.SetTextMapPropagator(
        propagation.NewCompositeTextMapPropagator(
            propagation.TraceContext{},
            propagation.Baggage{},
        ),
    )
    return nil
}
```

### Setting Up OpenTelemetry with gRPC in Go

The `initOpenTelemetry()` method sets up a gRPC connection to the collector. No host or address information is needed because it is already configured in the environment. The method also sets up the tracer so that trace data is sent to the collector in batches, which improves performance and should be used in most cases.

### Tracer Provider Setup

The tracer provider, `s.tp`, is then set as the default. This allows us to interact with the tracer anywhere in the application without needing to pass a reference into structures or include a value in the context. This makes it easier to adopt the library in your application.

### Trace Propagation

The function finishes by setting the default for how trace data should be propagated. Both the normal trace context data and any optional baggage (e.g., additional metadata) will be read and saved as maps of strings.

### Updating gRPC Server with Middleware

To automatically create new spans, the gRPC server and message handlers must be updated with new middleware.

For the gRPC server, the OpenTelemetry Go library provides client and server interceptors that can be added quickly. To set up the server, include the following in the gRPC server initializer:

```go
s.rpc = grpc.NewServer(
    grpc.UnaryInterceptor(
        otelgrpc.UnaryServerInterceptor(),
    ),
    // If there are streaming endpoints, also add:
    // grpc.StreamInterceptor(
    //    otelgrpc.StreamServerInterceptor(),
    // ),
)
```

### Adding the Interceptor for gRPC Clients

Adding the interceptor for clients is straightforward. The interceptors for clients are added as a Dial option.

Here's an example of how to add the interceptor for unary and streaming RPCs:

```go
func Dial(ctx context.Context, endpoint string) (
    conn *grpc.ClientConn, err error,
) {
    return grpc.DialContext(ctx, endpoint,
        grpc.WithTransportCredentials(
            insecure.NewCredentials(),
        ),
        grpc.WithUnaryInterceptor(
            otelgrpc.UnaryClientInterceptor(),
        ),
        // If there are streaming endpoints, also add:
        // grpc.WithStreamInterceptor(
        //     otelgrpc.StreamClientInterceptor(),
        // ),
    )
}
```

### Instrumenting Message Publishers and Subscribers with OpenTelemetry

The OpenTelemetry library does not provide ready-made middleware for custom message publishers and subscribers. However, creating custom middleware for this purpose is straightforward.

In a new package, `/internal/amotel`, which is named to indicate an instrumentation relationship with the `/internal/am` package, we define the `OtelMessageContextInjector()` and `OtelMessageContextExtractor()` functions. These functions are used to inject and extract trace context for messages.

We use the `OtelMessageContextInjector()` for all outgoing messages. As a result, every `MessagePublisher` constructor call will be updated to receive this new middleware. Here's an example of how to integrate it:

```go
am.NewMessagePublisher(
    stream,
    amotel.OtelMessageContextInjector(),
    tm.OutboxPublisher(outboxStore),
)
```

### Ordering Middleware for Message Publishers and Subscribers

When applying middleware to message publishers and subscribers, it's important to consider the order in which they are applied. If the middleware for the outbox and the new instrumentation is applied in the wrong order, the messages stored in the outbox may not be modified with the correct metadata.

For message subscribers, we use the `OtelMessageContextExtractor()` to extract trace context from incoming messages. This ensures that the trace context is correctly propagated when processing the message.

Here's how we integrate the extractor into the `MessageSubscriber` constructor calls:

```go
am.NewMessageSubscriber(
    stream,
    amotel.OtelMessageContextExtractor(),
)
```

### No Ordering Concerns for Subscriber Middleware

In contrast to message publishers, there are no ordering concerns with existing middleware for message subscribers. As long as we cover every constructor for the publishers and subscribers, our application will correctly output span data for each traced request.

### Adding Additional Trace Data

Earlier, I mentioned that traces can contain more data. Now, we can explore how additional data is added to the trace. For example, in the `/baskets/internal/grpc/server.go` file, we see that server calls are updated with new instrumentation. Specifically, in the `CheckoutBasket()` function, trace attributes are set as follows:

```go
span := trace.SpanFromContext(ctx)
span.SetAttributes(
    attribute.String("BasketID", request.GetId()),
    attribute.String("PaymentID", request.GetPaymentId()),
)
```

### Working with Spans and Adding Attributes

In the first line of the code, we retrieve the current span. If no span exists, the library returns a no-op span, which is a span that performs no operation. This ensures that our code doesn’t break even if tracing is not set up. The next lines then annotate this span with values that are important to the gRPC request.

These attributes are not meant to be recalled later like context values, but are instead sent to the trace collector. They provide valuable information that can help diagnose requests when visualized in tools such as Jaeger.

To interact with the existing span, we use `trace.SpanFromContext(ctx)`. However, there may be scenarios where you need to create a new span, for example, for processes that require their own trace. To create a new span, you can use the following code:

```go
ctx, span := otel.GetTracerProvider().
    Tracer("pkg_name").
    Start(ctx, "span_name")
```

This will grab the default tracer provider, then create a new tracer with whatever name you want to give it. But the best practice is to use the fully qualified package name. Then, a new span will be started, using any span from the context as the parent, and any name you wish to use.

Unless you know you need to create a new span, it is best to work with the existing span from the context.

Traces can also be annotated with events. Here, events are annotations that also have time components. This is very similar to logging, but the data is encapsulated entirely within the trace data. These too can be visualized in the graphs and diagrams produced by the trace tools. The event is visualized as either a line or dot on the span it was recorded to. Using events adds another dimension to the data that makes the flow of time more apparent:

![alt text](image-161.png)

Figure 12.4 – Spans annotated with events

The domain event handlers in each module will record additional information about the events they handled and the amount of time that it took. The following excerpt is from the `/baskets/internal/handlers/domain_events.go` file:

```go
span := trace.SpanFromContext(ctx)
defer func(started time.Time) {
    if err != nil {
        span.AddEvent(
            "Error encountered handling domain event",
            trace.WithAttributes(
                errorsotel.ErrAttrs(err)...,
            ),
        )
    }
    span.AddEvent("Handled domain event",
        trace.WithAttributes(
            attribute.Int64(
                "TookMS",
                time.Since(started).Milliseconds(),
            ),
    ))
}(time.Now())
span.AddEvent("Handling domain event",
    trace.WithAttributes(
        attribute.String("Event", event.EventName()),
    ),
)
```

In the preceding snippet, we are adding events before and after the event is handled. If handling the domain event results in an error, then a third event is going to be added with information about the error itself. When these two or three events are displayed in the graphs, they are positioned proportionately to the entire trace when they occurred.

I could have also recorded the error directly to the span using `RecordError()`. Doing this would change the status of the span to reflect that an error was encountered. Likewise, I could also directly set the status of the span when an error existed with `SetStatus()`. I do not want to use either here because I only want to record the fact an error occurred. The middleware that is used for the gRPC server and on `MessageSubscriber` will take care of calling both of those functions if the error hasn’t been handled already. Once you record an error to a span or set the status to the error level, you cannot undo it. So, it is best to let the code that created the span take care of doing both.

This is all the distributed tracing we will be adding in this chapter, but do experiment with updating a module or two to play around with creating new spans, adding attributes, and recording events.

### To instrument the application with OpenTelemetry, we made the following updates:

- Created a default `TracerProvider` struct in the `internal/system` package, which is configured using environment variables using a new method named `initOpenTelemetry()`
- Added gRPC interceptors to the server and client dialers to propagate the trace context for gRPC requests
- Added middleware to `MessagePublisher` and `MessageSubscriber` to propagate the trace context for messages
- In each gRPC server, we annotated the spans with relevant request data
- The domain event handlers were updated to bookend the handling of the domain events by recording events in the span

Next, we will learn how to report metrics about the application to Prometheus.

## Adding metrics to the application

We will be using [Prometheus](https://prometheus.io/) to instrument our application to report metrics. Prometheus is quick to set up and just as quick to use.

To begin, we need to set up an endpoint on the HTTP server so that Prometheus can fetch the metrics we will be publishing. Unlike OpenTelemetry, which uses a push model to send data to the collector, Prometheus uses a pull model and will need to be told where to look for metrics data.

To provide Prometheus an endpoint to fetch the data, we need to import the `promhttp` package and then add the handler it provides to the HTTP server. We must modify the `/internal/system/system.go` file to add the endpoint:

```go
import (
  "github.com/prometheus/client_golang/prometheus/promhttp"
)
// ... much further down
func (s *System) initMux() {
    s.mux = chi.NewMux()
    s.mux.Use(middleware.Heartbeat("/liveness"))
    s.mux.Method("GET", "/metrics", promhttp.Handler())
}
```

Prometheus expects to find metrics at the `/metrics` path by default, but that can be changed when you configure Prometheus to fetch the data.

The Go client for Prometheus automatically sets up a bunch of metrics for our application. Hitting that endpoint will display a dizzying list of metrics that were set up for us for free. We can also set up custom metrics for our application; to demonstrate, we will start with the messaging system.

When publishing a message, we must use a counter to record a total count and a count for that specific message:

```go
counter := promauto.NewCounterVec(
    prometheus.CounterOpts{
        Namespace: serviceName,
        Name:      "sent_messages_count",
        Help:      fmt.Sprintf(
            "The number of messages sent by %s",
            serviceName,
        ),
    },
    []string{"message"},
)
```

This is setting up a monotonically increasing counter that is broken up into partitions using a message value. The message value will be whatever is returned by calling `MessageName()` on the outgoing message. The service name is used as a namespace to avoid collisions when we are reporting metrics from the monolith. The namespace will be prefixed to the counter name, changing its name to something like `baskets_sent_messages_count`.

We are also using the `promauto` package to register these new metrics automatically with the default registry. If we were not using the `promauto` package and were using the `prometheus` package instead, we would need to include the following line to register the counter:

```go
prometheus.MustRegister(counter)
```

To record both the total count and the individual message count, we can use the following two lines:

```go
counter.WithLabelValues("all").Inc()
counter.WithLabelValues(msg.MessageName()).Inc()
```

Each time a message is published, we increment two partitions – the `all` partition and the message-specific partition.

The values kept in the counter will be lost when the service is restarted, and that is fine in most cases. Counter metrics are typically going to be watched for trends such as increasing too quickly, staying level over time, and so forth. The actual value of the counter rarely comes into play.

On the receiving side, we will use a similar counter to record how many messages have come in:

```go
counter := promauto.NewCounterVec(
    prometheus.CounterOpts{
        Namespace: serviceName,
        Name:      "received_messages_count",
        Help:      fmt.Sprintf(
            "The number of messages received by %s",
            serviceName,
        ),
    },
    []string{"message", "handled"},
)
```

This time, the counter has a second label called `handled`, which will be used to further split the count into successfully handled messages and the ones that produced an error. We are also interested in how long it takes to handle a message, so we will use another type of metric: a histogram.

Histograms are used to track length-like values such as request duration or message size. They are configured with buckets that will store the counts. We will use one to record the time it takes to handle each incoming message:

```go
histogram := promauto.NewHistogramVec(
    prometheus.HistogramOpts{
        Namespace: serviceName,
        Name:      "received_messages_latency_seconds",
        Buckets:   []float64{
            0.01, 0.025, 0.05, 0.1,
            0.25, 0.5, 1, 2.5, 5,
        },
    },
    []string{"message", "handled"},
)
```

Like the counter, we will use two labels to partition the histogram. The `Buckets` field is optional, and Prometheus provides a default bucket setup very similar to what’s shown in the preceding code example.

To record all of the metrics for the incoming messages, we will use the following code. This will record four metrics when handling a message:

```go
handled := strconv.FormatBool(err == nil)
counter.WithLabelValues("all", handled).Inc()
counter.WithLabelValues(
    msg.MessageName(), handled,
).Inc()
histogram.WithLabelValues(
    "all", handled,
).Observe(time.Since(started).Seconds())
histogram.WithLabelValues(
    msg.MessageName(), handled,
).Observe(time.Since(started).Seconds())
```

For each metric, we record the `all` partition and the specific message partition. To determine whether the message was handled properly, we check whether the `err` value is `nil`. This will record the metrics on a lot of partitions, which can be useful in setting up detailed dashboards.

The metrics are recorded using middleware that lives in the `/internal/amprom` package. Using this middleware is going to be the same as using the OpenTelemetry middleware we created. For the publisher, we can add it before the outbox middleware:

```go
am.NewMessagePublisher(
    stream,
    amotel.OtelMessageContextInjector(),
    amprom.SentMessagesCounter("baskets"),
    tm.OutboxPublisher(outboxStore),
)
```

Then, we can use the same ID we used for the NewMessageSubscriber constructor by adding it either before or after the OpenTelemetry middleware:

```go
am.NewMessageSubscriber(
    stream,
    amotel.OtelMessageContextExtractor(),
    amprom.ReceivedMessagesCounter("baskets"),
)
```

We will be able to create detailed dashboards showing the number of messages being used in our application and how long it takes our application to process each one.

Speaking of dashboards, they are not only used by the engineers working on the application but also by people from other departments. It is common to expose metrics about how much product is being produced, or how many customers have registered. We can add those kinds of metrics as well.

In the composition root for the Customers module, we can add a counter for customers_registered_count:

```go
customersRegistered := promauto.NewCounter(
prometheus.CounterOpts{
Name: "customers_registered_count",
},
)
```

There’s no need for a namespace or partitions this time; we can use a simple counter. We want to use this counter to count every successful registration that is made. We could pass the counter into the application, then increment the counter if there was no error being returned by the RegisterCustomer() method by checking the results with a deferred function. This would not be my first choice on how to go about this. The MallBots application is a relatively simple application and the Application struct in the real application may already be dealing with a lot of dependencies. My preference is to create a wrapper for the Application struct that will be used for this counter and any other metric we want to add. This keeps the concerns separated and keeps the existing Application tests unchanged. It also means we can test the wrapper in isolation.

The wrapper will only intercept the RegisterCustomer() method, letting all of the other methods pass through unaffected:

```go
type instrumentedApp struct {
App
customersRegistered prometheus.Counter
}
func NewInstrumentedApp(
app, customersRegistered prometheus.Counter,
) App {
return instrumentedApp{
App: app,
customersRegistered: customersRegistered,
}
}
func (a instrumentedApp) RegisterCustomer(
ctx context.Context, register RegisterCustomer,
) error {
err := a.App.RegisterCustomer(ctx, register)
if err != nil { return err }
a.customersRegistered.Inc()
return nil
}
```

To use this instrumented application, we need to wrap the application instance in the composition root:

```go
application.NewInstrumentedApp(
application.New(
customersRepo,
domainDispatcher,
),
customersRegistered,
)
```

Other modules can be updated to record metrics such as counting the number of baskets started by the users or counting the number of new products made available by the stores.

Let’s recap what we did to add Prometheus metrics to the application:

    1. An endpoint was added to the HTTP service so that Prometheus can retrieve our metrics
    2. Middleware was added to add metrics for the published and received messages
    3. Middleware was included in the constructors for the MessagePublisher and MessageSubscriber interfaces
    4. Additional application counters were created, such as the registered customer counter
    5. An application wrapper was used to instrument the application without modifying it

In this section, we added distributed tracing and metrics to our application. This covers all three pillars of observability since the application already had logging. Everything that was added should have no measurable impact on the application; if it does, we will now be able to monitor it.

In the next section, we will learn about the tools we can use to view the data that’s now being reported about the application.

### Viewing the monitoring data

The application will now be producing a lot of data; to view this data, we need to collect it or, in the case of Prometheus, retrieve it.

The Docker Compose environment was updated with four new services, as follows:

    1. The OpenTelemetry collector, which will collect trace and span data
    2. Jaeger to render the traces
    3. Prometheus to collect and display metrics data
    4. Grafana to render dashboards based on the metrics data

The OpenTelemetry collector will also provide Prometheus metrics about the traces and spans it collects:

![alt text](image-162.png)

Figure 12.5 – The additional monitoring services

We have already configured the modules to connect with the collector so that is ready to go. For Prometheus, we still need to configure it to retrieve the metrics from each microservice. The configuration file, /docker/prometheus/prometheus-config.yml, will need to be updated so that it contains a job for each microservice we want to scrape. For the Shopping Baskets microservice, we must add the following under the scrape_configs heading:

```bash
- job_name: baskets
  scrape_interval: 10s
  static_configs:
    - targets:
      - 'baskets:8080'
```

There are a lot more options we could set here, but these are all we will need for now.

At this point, we can start up the Docker Compose environment, then use the Swagger UI to make some requests. However, making individual requests with the Swagger UI could take some time; we need to build up enough data to give us some idea of what collecting data from an active application might look like.

Instead, we can use a small application that can be found under /cmd/busywork to simulate several users making requests to perform several different activities. The application is nothing fancy and you are encouraged to modify it to simulate whatever interactions you like.

With the MallBots application already running locally with Docker Compose, start the busywork application by running the following:

```
cd cmd/busywork

go run .
```

Five clients will be started up and will begin making requests. You can increase the number of clients by passing in the -clients=n flag, with an upper limit of 25. To end the busywork application, use Ctrl + C or Cmd + C; this will kill the process.

Now, we can look at some of the data that is being produced, starting with Jaeger. Open http://localhost:8081 in your browser to open Jaeger. You should see a UI like this:

![alt text](image-164.png)

Figure 12.6 – The Jaeger UI

Toward the left, under Service, select the customers service and click the Find Traces button. Doing this will show several traces in a timeline view and as a list. In the timeline, the size of the circle signifies the size of the trace. The larger the circle, the more spans that were involved. Also, the height of the circles signifies the duration of the trace. This is an example of a search for traces that involve the customers service; your graph will be different because the busywork clients will be randomly interacting with the MallBots application:

![alt text](image-165.png)

Figure 12.7 – Traces that involved the customers service

If you do not have any of the larger trace circles, as shown in the preceding figure, wait for a moment and perform a new search; eventually, one will appear. These larger circles are from the create order saga execution coordinator. If you click on one, it will open up the trace details screen for that trace. From the details screen, we can see how the services all worked together to accomplish the task of creating a new order:

![alt text](image-166.png)

Figure 12.8 – A portion of the create order process shown in Jaeger

Clicking on one of the rows in the graph will provide you with additional details. If we click on the first row for baskets basketspb.BasketService/CheckoutBasket, we will be able to see the additional data we recorded to the span using the gRPC service’s CheckoutBasket() method. Under Tags, we will find the BasketID and PaymentID properties, which were used for this request. Under Logs, we will find the events that were recorded to the span ordered by time, with all times relative to the start of the trace.

Remember when we added the bookend events to the handling of the domain events? If you compare the log timestamps for the two events with the timestamps of the next child span, you will see that the second log correctly shows it occurred after the child span had been completed.

A lot of data recorded is with each trace and that is its major downside. Recording a trace can be very demanding on the disk to store them, the CPU to process them, and the network to collect them. To lessen this resource demand, traces are sampled and only some are saved. Deciding to save a trace is either head-based, during the initial parent span creation, or tail-based, where the decision can be made by a child span at any time. OpenTelemetry only supports head-based decision-making. The upside is that it is easier to implement and work with, but the downside is that it drops traces that include errors that might be worth checking out. One day, OpenTelemetry might offer tail-based decision-making, but until it does, you should continue to use logging to capture important errors.

We also have the metrics to check out in Prometheus. Opening http://localhost:9090 in your browser will present you with the following UI:

![alt text](image-167.png)

Figure 12.9 – The Prometheus UI

Performing a search for cosec_received_messages_count will return results similar to this:

![alt text](image-168.png)

Figure 12.10 – Searching for the received messages counts for the cosec service

You could also try searching for go_gc_duration_seconds to see the garbage collector metrics for each microservice or any other metric you can think of. Like the traces, we are dealing with a very large amount of data – not as much as with the traces, but certainly a large number of metrics.

Searching for metrics in Prometheus and viewing the raw data is not very compelling. That is why we also have Grafana running. Opening https://localhost:3000/ and then browsing for dashboards will show the two dashboards that are installed under the MallBots folder. The Application dashboard will display some panels that will give you insights into how active the application is, and will display several panels showing the rates of incoming and outgoing messages for a few services:

![alt text](image-169.png)

Figure 12.11 – The MallBots Application dashboard

How much activity you see in the dashboard will depend on how many clients you have running in the busywork application and the random interactions that the clients are performing.

The other dashboard that is available is the OpenTelemetry Collector dashboard, which will provide some details about how much work the collector is doing.

With a small to moderate amount of work, we added a massive amount of instrumented data to our application that gives incredible insight into the inner workings of the application.

### Summary

In this final chapter, we learned about monitoring and observability. We were introduced to the OpenTelemetry library and learned about its goals of making applications observable easier. We also learned about distributed tracing and how it is one of the three pillars of observability.

Later, we added both distributed tracing and metrics to the application using OpenTelemetry and Prometheus. With a little work, both forms of instrumentation were added to the application. To demonstrate this new instrumentation, we made use of a small application to simulate users making requests while we were free to view the recorded data in either Jaeger or Prometheus.

This chapter concludes the adventure we started, which involved taking a synchronous application and refactoring it to turn it into a fully asynchronous application that could be deployed to AWS and be completely observable.
