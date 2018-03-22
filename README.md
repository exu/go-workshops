# Golang Training Resources

This repository contains files needed for managing Go language workshop - it's some kind of quite complete walk-through of Go language. Feel free to look on the code, there are many comments which could be useful for beginners and semi-intermediate Go developers.
## About Go

Who designed Go language?

- Rob Pike (Unix, UTF-8)
- Ken Thompson (Unix author, UTF-8, B lang)
- Robert Griesemer (V8, Java Hotspot (Oracle Java), GFS)

but those above are only ignitors of whole contributions: https://golang.org/AUTHORS

Why `go` was developed by Google? They have a lot of problems in C/Python/Java codebases:

- speed up development
- speed up compiling
- multicore systems

sources:
- https://golang.org/doc/faq#What_is_the_purpose_of_the_project
- https://talks.golang.org/2012/splash.article


## Go language characteristics

- staticaly compiled (one fat binary with all dependencies )
- Garbage Collected
- Strong types
- Functions as first class citizens
- Object Oriented (but without inheritance)

## Why it's worth of your time?

- easy deployment (no dependencies, single binary statically linked)
- no more code style wars - `gofmt`
- integrated package downloader `go get`
- integrated code validation `go vet` and `golint` (github.com/golang/lint)
- `gocode` intellisense server - you don't need fat IDE to write go code, you can use now editor which you love (Sublime, Atom, Vim, Emacs, VSCode)
- very fast compilation - if you're going from JAVA world you'll be really surprised
- quite complete standard library - template/html, performant www servers, json, xml, streams, io, buffers, first class citizen concurrency
- easy to use cross-compilation (x64, ARM, 386, Mac, Windows)
- easy start, bunch of editors, all things simply work
- testing included
- benchmarking of code included
- very low entry bareer
- hype, one of fastest growing language, many new projects are in Go recently
- concurrency
- great documentation generator
- and many many more ...

## Workshop prerequisities

You can install `golang` and `docker` using your prefered way i.e. your OS package manager (brew, pacman, apt, snap or other) or you can simp[l follow installation instruction on go and docker sites.

### Golang installation

For recent installation instructions please refer to [Golang installation guide](https://golang.org/doc/install)

You'll need `git` to be installed

### Docker install

> Docker is the company driving the container movement and the only container platform provider to address every application across the hybrid cloud. Today’s businesses are under pressure to digitally transform but are constrained by existing applications and infrastructure while rationalizing an increasingly diverse portfolio of clouds, datacenters and application architectures. Docker enables true independence between applications and infrastructure and developers and IT ops to unlock their potential and creates a model for better collaboration and innovation.



For recent docker installation please follow [Docker installation guide](https://docs.docker.com/install/) for your OS.

Docker is needed for some parts of workshops for running databases or other needed dependencies. Also will be needed to complete homework.

### Docker Compose Installation

> Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application’s services. Then, with a single command, you create and start all the services from your configuration.

To install docker compose please follow [Docker compose installation guide](https://docs.docker.com/compose/install/)

## Init workshops to play on your machine with code

    go get github.com/exu/go-workshops
    cd $GOPATH/exu/go-workshops

## Go Tools

### Included

- go test - included testing framework
- go fmt - code formater - ony one valid coding standard -  [fmt library website](https://golang.org/pkg/fmt/)
- gocode - autocomplete service - [gocode library website](https://github.com/nsf/gocode)
- go vet - code validator and fixer - [vet library website](http://godoc.org/golang.org/x/tools/cmd/vet)
- go oracle - dependencies analyser (can be integrated in editor) [oracle web site](http://golang.org/s/oracle-user-manual)
- godoc - great documentation generator and viewer

### IDEs

- IntelliJ Go plugin
  - https://plugins.jetbrains.com/plugin/9568-go
- GoLand - https://www.jetbrains.com/go/features/
- Emacs - go-mode
- Vim - vim-go
- Visual Studio Code (really great UX)
- LiteIDE
- SublimeText
- Atom

### Auto reload

- BRA (Brilliant Ridiculous Assistant) https://github.com/Unknwon/bra - it's good to setup it when you're working on some webservers to auto reload your app when changes in code are made.


## Github style - project structure ([code for: Basics Importing](https://github.com/exu/go-workshops010-basics-importing))

In go idiomatic way is to organise code in "github style", so part of the path is looking like server address to library. Of course if you want you don't need to do this, but whole ecosystem works that way.

```sh
src/
    github.com
        exu
            mysuperproject
    ioki.com.pl
            mnmel
                 nmelinium
bin/
    superappbinary
pkg/
    compiled packages
```

Environment variable `$GOPATH` is responsible for path to the root dir of `src`, `bin` and `pkg` directories.


## Packages and Importing

`package` in go is in form of files with directive `package package_name` on top of each file. Package by default is imported as full path to package.

    import "github.com/exu/go-workshops/010-basics-importing/sub"

## Getting and installing external packages

To get external package we need to run go install which will get sources and binaries and put them to `src`/`bin`/`pkg` directories

```sh
go install external.package.com/uri/to/package
```

## `sub` example package

As we can see our `sub` package is in directory `sub` (obvious) and have two files; `sub1.go` and `sub2.go` each of them also have `package sub` directive which tells compiler that they are in one package.


## Package managers ([code for: Package Management](https://github.com/exu/go-workshops011-package-management))

Currently most advanced in go ecosystem is `dep` https://github.com/golang/dep

You can init your project:

	$ dep init
	$ ls
	Gopkg.toml Gopkg.lock vendor/

after that dep will add vendor dir where all depndencies will be loaded (In go after 1.5 `vendor` dir have preference over "github style `$GOPATH` based directories - when compiler will not find library in vendor dir it'll try to take it from `$GOPATH`).

For more details please refer to `dep` documentation at https://golang.github.io/dep/docs/daily-dep.html

## Variables ([code for: Basics Variables](https://github.com/exu/go-workshops020-basics-variables))

In go variables can be declared and assigned in one simple way `a := 1`, compiler will detect type of variable based on its value.


## Constants ([code for: Basics Constants](https://github.com/exu/go-workshops030-basics-constants))

Use `const` to define new contant, mechanics looks like in other languages. What is worth to mention that we have `iota` keyword which could be used as some kind of iterator in constant definition.


# ([code for: Basics Overriding Internal Types](https://github.com/exu/go-workshops035-basics-overriding-internal-types))



## Funkcje ([code for: Basics Functions](https://github.com/exu/go-workshops040-basics-functions))

Functions in Go are "First class citizen".

- function definition
- multiple returns
- named multiple returns
- defered calls
- variadic functions

## Loops ([code for: Basics Loops](https://github.com/exu/go-workshops040-basics-loops))

In go there is only one loop keyword: `for`. It's often used with `range` keyword to iterate over array like elements.

## Packages initialisation ([code for: Basics Init](https://github.com/exu/go-workshops041-basics-init))

`func init() {}` is responsible for package initialisation, it's called only once when import statement for given package is called.


## Closures aka anonymous functions ([code for: Basics Closures](https://github.com/exu/go-workshops042-basics-closures))

Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.


## Tablice ([code for: Basics Arrays](https://github.com/exu/go-workshops050-basics-arrays))

In Go, an `array` is a **numbered sequence of elements** of a specific length. Arrays are "low level" data structures with slices over them which simplifies creating and managing.


## Slices ([code for: Basics Slices](https://github.com/exu/go-workshops051-basics-slices))


Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

Unlike arrays, slices are typed only by the elements they contain (not the number of elements). To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued).

sources:

- https://blog.golang.org/slices
- https://github.com/golang/go/wiki/SliceTricks
- https://blog.golang.org/go-slices-usage-and-internals
- http://research.swtch.com/godata
- http://blog.golang.org/go-slices-usage-and-internals
- http://www.dotnetperls.com/slice-go


## Mapy ([code for: Basics Maps](https://github.com/exu/go-workshops055-basics-maps))

One of the most useful data structures in computer science is the hash table. Many hash table implementations exist with varying properties, but in general they offer fast lookups, adds, and deletes. Go provides a built-in map type that implements a hash table.


## `make` and `new` keywords ([code for: Basics New And Make](https://github.com/exu/go-workshops059-basics-new-and-make))

`new(T)` allocates zeroed storage for a new item of type T and returns its address. In Go terminology, it returns a pointer to a newly allocated zero value of type T.

`make(T)` For slices, maps, and channels, make initializes the internal data structure and prepares the value for use.

### `make`

Things you can do with make that you can't do any other way:

- Create a channel
- Create a map with space preallocated
- Create a slice with space preallocated or with len != cap


### `new`

It's a little harder to justify new. The main thing it makes easier is creating pointers to non-composite types. The two functions below are equivalent. One's just a little more concise:

    func newInt1() *int { return new(int) }

    func newInt2() *int {
        var i int
        return &i
    }


Go has multiple ways of memory allocation and value initialization:

    &T{...}, &someLocalVar, new, make

Allocation can also happen when creating composite literals.

new can be used to allocate values such as integers, &int is illegal:

    new(Point)
    &Point{}      // OK
    &Point{2, 3}  // Combines allocation and initialization

    new(int)
    &int          // Illegal

    // Works, but it is less convenient to write than new(int)
    var i int
    &i

In terms of channels there you can use make and new

    p := new(chan int)   // p has type: *chan int
    c := make(chan int)  // c has type: chan int

## Struktury ([code for: Basics Structs Defining](https://github.com/exu/go-workshops060-basics-structs-defining))

A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a struct, non-blank field names must be unique.


## Struktury - Kompoozycja ([code for: Basics Struct Composition](https://github.com/exu/go-workshops062-basics-struct-composition))

Kompozycja taki pattern chyba znacie ?


## Struct tags (annotations like) ([code for: Basics Struct Tags](https://github.com/exu/go-workshops064-basics-struct-tags))

A tag for a field allows you to attach meta-information to the field which can be acquired using reflection. Usually it is used to provide transformation info on how a struct field is encoded to or decoded from another format (or stored/retrieved from a database), but you can use it to store whatever meta-info you want to, either intended for another package or for your own use.


# ([code for: Basics Anonymous Structs](https://github.com/exu/go-workshops065-basics-anonymous-structs))



## Interface'y ([code for: Basics Interfaces](https://github.com/exu/go-workshops065-basics-interfaces))

Go have "implicit interfaces". To implement an interface in Go, we just need to implement all the methods in the interface.

So what is an interface? An interface is two things: it is a set of methods, but it is also a type.

### The `interface{}` type

The `interface{}` type, the empty interface, is the source of much confusion. The interface{} type is the interface that has no methods. Since there is no implements keyword, all types implement at least zero methods, and satisfying an interface is done automatically, all types satisfy the empty interface. That means that if you write a function that takes an `interface{}` value as a parameter, you can supply that function with any value. So this function:

    func DoSomething(v interface{}) {
    // ...
    }

will accept any parameter whatsoever.

### Type assertions

- https://medium.com/golangspec/type-assertions-in-go-e609759c42e1

## Error handling ([code for: Basics Errors](https://github.com/exu/go-workshops067-basics-errors))

There is no exceptions in Go, errors are returned by value, or aggregated in intermediate objects. In go error is simply value which should be handled programatically as quick as possible.

Sources:
- https://blog.golang.org/errors-are-values
- http://blog.golang.org/error-handling-and-go
- http://davidnix.io/post/error-handling-in-go/


## Panics ([code for: Basics Panics](https://github.com/exu/go-workshops068-basics-panics))

- Used when we want to stop the program.
- We can check if there was a `panic` occurence in function defer chain


## Strings ([code for: Basics Strings](https://github.com/exu/go-workshops070-basics-strings))

In Go, a string is in effect a **read-only slice of bytes**.

It's important to state right up front that a string holds arbitrary bytes. It is not required to hold Unicode text, UTF-8 text, or any other predefined format. As far as the content of a string is concerned, it is exactly equivalent to a slice of bytes.

More on https://blog.golang.org/strings


## Go routines ([code for: Basics Goroutines](https://github.com/exu/go-workshops080-basics-goroutines))

A goroutine is a lightweight thread of execution.

Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives (channels).

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.


## Using 3rd parties ([code for: Basics 3rd Party Packages](https://github.com/exu/go-workshops090-basics-3rd-party-packages))

In go we can get packages to our $GOPATH with use of `go get` command.


## DEP

```sh
go dep init
```

`dep` is fully-fledged dependency manager. It downloads all dependencies source code to `vendor` directory and then compilator includes those code during compilation process.


# ([code for: Basics Pointers](https://github.com/exu/go-workshops090-basics-pointers))



## Channels ([code for: Concurrency Channels](https://github.com/exu/go-workshops100-concurrency-channels))

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

    ch <- v    // Send v to channel ch.
    v := <-ch  // Receive from ch, and
            // assign value to v.

(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

    eventsChannel := make(chan int)

By default, **sends and receives block until the other side is ready**. This allows goroutines to synchronize without explicit locks or condition variables.

### Buffered Channels

Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

    ch := make(chan int, 100)

Sends to a buffered channel **block only when the buffer is full**. Receives block when the buffer is empty.

### In this workshops you'll learn about:

- channels mechanics
- receiving data from channel
- channels length
- buffers
- channel closing
- generate data with channels
- timers, tickers with channels
- worker pool
- rate limiting with channels
- selecting proper channel stream
- pipeline and fan-in/fan-out patterns

Feel free to browse [channels chapter source code](100-concurrency-channels/)

## Other concurrency patterns ([code for: Concurrency Other](https://github.com/exu/go-workshops101-concurrency-other))

You've lerned about channels primitive on previous chapter of this workshops, now it's time to learn about some other ways of creating concurrent programs in go.

We'll be doing some code with:
- data races errors
- atomic counters
- mutexes
- wait groups


## Date time ([code for: Stdlib Date Time](https://github.com/exu/go-workshops110-stdlib-date-time))

Go has very powerful standard library, first and one of awesome library are is Date time.

# ([code for: Stdlib Os Processes](https://github.com/exu/go-workshops110-stdlib-os-processes))



## Args and flags ([code for: Stdlib Args](https://github.com/exu/go-workshops114-stdlib-args))

One of important thing when runnein your program is to get some arguments from user.

in Go you can easily get those data using `os.Args` slice or more powerful package `flags`.

Additionally you can get environment variables from `os.Getenv` function in `os` package

# Streams ([code for: Stdlib Streams](https://github.com/exu/go-workshops115-stdlib-streams))

Streams in go are very powerful feature, very large part of standard library is written as some kind of stream reader or stream writer.

Go have two basic interfaces shaping all world of data streams `io.Reader` and `io.Writer`.

In this section of workhops we'll try to cover some of use cases for each of them.


## Basic IO operations ([code for: Stdlib Io](https://github.com/exu/go-workshops116-stdlib-io))

- `bufio` examples
- directory traversal
- merging files with use of `buffers`


## Stdlib Logging package ([code for: Stdlib Logging](https://github.com/exu/go-workshops120-stdlib-logging))

Go has basic logging package to log what's happening in your program.

## HTTP library ([code for: Stdlib Http](https://github.com/exu/go-workshops140-stdlib-http))

Package http provides HTTP client and server implementations.

HTTP is one of fundamental lib in programming. Go has implemented very powerful HTTP primitives which allows creation of full-fledged HTTP powered servers.

Info: there is no routing in stdlib so you need to implement your own or use third party libraries (Gorilla mux, Gin, Echo are ones who can help you)


## HTTP Middlewares ([code for: Stdlib Http Middlewares](https://github.com/exu/go-workshops141-stdlib-http-middlewares))

Go is using the term `middleware`, but each language/framework calls the concept differently. `NodeJS` and `Rails` calls it `middleware`. In the `Java EE` (i.e. Java Servlet), it’s called `filters`. `C#` calls it `delegate handlers`.

Essentially, the middleware **performs some specific function** on the HTTP request or response **at a specific stage in the HTTP pipeline** before or after the user defined controller. Middleware is a design pattern to eloquently add cross cutting concerns like logging, handling authentication, or gzip compression without having many code contact points.


# ([code for: Stdlib Encoding Json](https://github.com/exu/go-workshops150-stdlib-encoding-json))



# ([code for: Stdlib Encoding Xml](https://github.com/exu/go-workshops151-stdlib-encoding-xml))



## Templates ([code for: Stdlib Templates](https://github.com/exu/go-workshops170-stdlib-templates))

In programming we need often some meta templates who help us with interoperability between our code and many output formats. One example could be template engine for generating HTML files for web sites.

In Go there are template engines (yes plural!) implemented in `stdlib`!

We'll go in this chapter by some `html` and `text` template engines.

# ([code for: Stdlib Math](https://github.com/exu/go-workshops180-stdlib-math))



# ([code for: Stdlib Regexp](https://github.com/exu/go-workshops180-stdlib-regexp))



## Context package ([code for: Stdlib Context](https://github.com/exu/go-workshops181-stdlib-context))

Context is very powerful package, in this section i've implemented HTTP client and server which handles cancelling both sides when client e.g. press Ctrl+C during request.

# ([code for: Stdlib Sort](https://github.com/exu/go-workshops181-stdlib-sort))



# ([code for: Stdlib Signal](https://github.com/exu/go-workshops182-stdlib-signal))



# Task ([code for: Task Http Response Reader](https://github.com/exu/go-workshops199-task-http-response-reader))

MVP:

- Write HTTP response reader from given sources
- You can implement it with channels.
- Push results in JSON format to os.Stdout

Homework for ambitious ones:

- Do it with given interval (e.g. 10 seconds)
- Pass last results back through WWW REST JSON API (e.g. localhost:8080/statuses)


## MySQL ([code for: Databases Mysql](https://github.com/exu/go-workshops210-databases-mysql))

Install instructions to run this section:

```sh
cd docker/mysql
docker-compose up
mysql -uroot -proot -P7701 -h127.0.0.1 < init.sql
mysql -uroot -proot -P7701 -h127.0.0.1
```


## ORMs in Go ([code for: Databases Orm](https://github.com/exu/go-workshops215-databases-orm))

### GORM the most popular Go ORM.

gorm have [some nice converntions](http://gorm.io/docs/conventions.html) to start with new project

full documentation is on http://gorm.io/docs/


### GORP - Go Object Relational Persistence

If you need some lighter abstraction on SQL driver (than full-fledged ORM)

https://github.com/go-gorp/gorp


# MongoDB examples ([code for: Databases Mongodb](https://github.com/exu/go-workshops220-databases-mongodb))

to use this examples you'll need to run MongoDB server. you can do this using prepared docker-compose file.

```sh
cd docker/mongo
docker-compose up
mongo localhost:7702
```


# RethinkDB ([code for: Databases Rethinkdb](https://github.com/exu/go-workshops230-databases-rethinkdb))

Why rethink here? I think it'll be worth to point out one of it's nice feature - collection changes.

#@ Dockerizing

### run rethink db as container

```
docker run --name some-rethink -v "$PWD:/data" -d rethinkdb
```

mozemy również zmapować porty do lokalnej maszynki

```
docker run --name rethink -p 28015:28015 -v "$PWD/data:/data" -d rethinkdb
```

wtedy instacja kontenera będzie widoczna pod adresem `localhost:28015`

### Zlinkuj go w swojej aplikacji

```
docker run --name some-app --link some-rethink:rdb -d application-that-uses-rdb
```


# ([code for: Databases Redis](https://github.com/exu/go-workshops240-databases-redis))



# ([code for: Databases Bolt](https://github.com/exu/go-workshops241-databases-bolt))



# ([code for: Databases Postgresql](https://github.com/exu/go-workshops250-databases-postgresql))



# ([code for: Testing Unit Task](https://github.com/exu/go-workshops300-testing-unit-task))



# ([code for: Testing Unit Examples](https://github.com/exu/go-workshops302-testing-unit-examples))



# ([code for: Testing Unit Dependencies](https://github.com/exu/go-workshops305-testing-unit-dependencies))



# ([code for: Testing Http Handler](https://github.com/exu/go-workshops310-testing-http-handler))



# ([code for: Testing Http Server](https://github.com/exu/go-workshops310-testing-http-server))



# ([code for: Testing Benchmarking](https://github.com/exu/go-workshops320-testing-benchmarking))



# ([code for: Testing Parallel Benchmark](https://github.com/exu/go-workshops380-testing-parallel-benchmark))



# ([code for: Patterns Pipeline](https://github.com/exu/go-workshops400-patterns-pipeline))



# ([code for: Patterns Glow Map Reduce](https://github.com/exu/go-workshops401-patterns-glow-map-reduce))



# ([code for: Fullstack Html And Angular](https://github.com/exu/go-workshops510-fullstack-html-and-angular))



# ([code for: Fullstack Rest Angular Resource](https://github.com/exu/go-workshops520-fullstack-rest-angular-resource))



# ([code for: Fullstack Json Event Stream](https://github.com/exu/go-workshops530-fullstack-json-event-stream))



# ([code for: Fullstack Websockets](https://github.com/exu/go-workshops540-fullstack-websockets))



# ([code for: Fullstack Wiki](https://github.com/exu/go-workshops560-fullstack-wiki))



## Beego ([code for: Fullstack Beego](https://github.com/exu/go-workshops570-fullstack-beego))

Bee init script - inicjuje podstawową strukturę katalogów.
hot compile.

```
go get github.com/astaxie/beego
go get github.com/beego/bee
bee new hello
cd hello
bee run hello
```


## Mangos SP message system ([code for: Libs Mangos](https://github.com/exu/go-workshops601-libs-mangos))

> Package mangos is an implementation in pure Go of the SP ("Scalability Protocols") messaging system. This makes heavy use of go channels, internally, but it can operate on systems that lack support for cgo.

> The reference implementation of the SP protocols is available as nanomsg™; there is also an effort to implement an improved and more capable version of nanomsg called NNG™.

> The design is intended to make it easy to add new transports with almost trivial effort, as well as new topologies ("protocols" in SP terminology.)

> At present, all of the Req/Rep, Pub/Sub, Pair, Bus, Push/Pull, and Surveyor/Respondent patterns are supported.

> Additionally, there is an experimental new pattern called STAR available. This pattern is like Bus, except that the messages are delivered not just to immediate peers, but to all members of the topology. Developers must be careful not to create cycles in their network when using this pattern, otherwise infinite loops can occur.

> Supported transports include TCP, inproc, IPC, Websocket, Websocket/TLS and TLS. Use addresses of the form "tls+tcp://:" to access TLS. Note that ipc:// is not supported on Windows (by either this or the reference implementation.) Forcing the local TCP port in Dial is not supported yet (this is rarely useful).

## Perks for Go (golang.org) ([code for: Libs Quantile Percentiles](https://github.com/exu/go-workshops602-libs-quantile-percentiles))

> Perks contains the Go package quantile that computes approximate quantiles over an unbounded data stream within low memory and CPU bounds.


# ([code for: Libs Beep](https://github.com/exu/go-workshops610-libs-beep))



# ([code for: Libs Bra](https://github.com/exu/go-workshops610-libs-bra))



# ([code for: Libs Slack](https://github.com/exu/go-workshops611-libs-slack))



# ([code for: Libs Vegeta](https://github.com/exu/go-workshops620-libs-vegeta))



# go readline implementation ([code for: Libs Readline](https://github.com/exu/go-workshops630-libs-readline))

https://github.com/chzyer/readline


# ([code for: Libs Termbox](https://github.com/exu/go-workshops640-libs-termbox))



## Caddy webserver ([code for: Libs Caddy](https://github.com/exu/go-workshops650-libs-caddy))


# ([code for: Libs Http Echo](https://github.com/exu/go-workshops651-libs-http-echo))



# ([code for: Libs Http Iris](https://github.com/exu/go-workshops651-libs-http-iris))



# ([code for: Libs Jobrunner](https://github.com/exu/go-workshops660-libs-jobrunner))



# ([code for: Libs Cron](https://github.com/exu/go-workshops665-libs-cron))



## Validator package ([code for: Libs Validator](https://github.com/exu/go-workshops670-libs-validator))

- https://github.com/go-playground/validator

- https://github.com/go-validator/validator

- https://github.com/asaskevich/govalidator


# ([code for: Libs Gographviz](https://github.com/exu/go-workshops677-libs-gographviz))



# ([code for: Libs Fasthttp](https://github.com/exu/go-workshops680-libs-fasthttp))



# ([code for: Libs Uiprogress](https://github.com/exu/go-workshops681-libs-uiprogress))



# ([code for: Libs Go Rpm](https://github.com/exu/go-workshops690-libs-go-rpm))



# ([code for: Libs Grpc](https://github.com/exu/go-workshops690-libs-grpc))



# ([code for: Libs Logrus](https://github.com/exu/go-workshops690-libs-logrus))



# ([code for: Libs Go Plugin](https://github.com/exu/go-workshops691-libs-go-plugin))



# ([code for: Libs Consul](https://github.com/exu/go-workshops692-libs-consul))



# ([code for: Libs Language Bindings](https://github.com/exu/go-workshops693-libs-language-bindings))



# ([code for: Libs Astielectron](https://github.com/exu/go-workshops694-libs-astielectron))



# AWS Lambda Golang ([code for: Lambda Simple](https://github.com/exu/go-workshops700-lambda-simple))

https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/

## Remember to build your handler executable for Linux!

    GOOS=linux GOARCH=amd64 go build -o main main.go
    zip main.zip main


## Deploying

https://docs.aws.amazon.com/lambda/latest/dg/deploying-lambda-apps.html

    aws lambda create-function \
    --region us-west-1 \
    --function-name HelloFunction \
    --zip-file fileb://./main.zip \
    --runtime go1.x \
    --tracing-config Mode=Active \
    --role arn:aws:iam::<account-id>:role/<role> \
    --handler main


aws lambda create-function --region us-west-1 --function-name HelloFunction --zip-file fileb://./deployment.zip --runtime go1.x --tracing-config Mode=Active --role arn:aws:iam::270605981035:role/<role> --handler main


# ([code for: How To Run On Production](https://github.com/exu/go-workshops800-how-to-run-on-production))



# Traefik ([code for: Load Balancing Traefik](https://github.com/exu/go-workshops801-load-balancing-traefik))

Load balancer with hot reloading


# Kubernetes deployment on local machine ([code for: How To Run On Kubernetes Cluster](https://github.com/exu/go-workshops810-how-to-run-on-kubernetes-cluster))

We'll try now put our code in working Kubernetes cluster with use of Kubernetes and Minikube on our local machine.

Kubernetes looks great but doing quick developemnt flow could be plain in the ass, if you don't have access to infrastructure of AWS or GCE with prepared pipelines to pass your code to valid infrastructure.

What if we want do develop our containers on local machine? I did'n found the out-of-the box solution but there is quite nice workaround for managing your own registry in [Sharing a local registry with minikube](https://blog.hasura.io/sharing-a-local-registry-for-minikube-37c7240d0615) article.


## Getting started

0. Install kubectl and minikube
1. using docker registry runned on minikube cluster + proxy
     ([source](https://blog.hasura.io/sharing-a-local-registry-for-minikube-37c7240d0615))

            ❯ kubectl create -f kubernetes/kube-registry.yaml

            # forwarding ports is temporary
            ❯ kubectl port-forward --namespace kube-system \
            $(kubectl get po -n kube-system | grep kube-registry-v0 | \
            awk '{print $1;}') 5000:5000


3. Now it's time to build our app. New docker have ability to build multi-stage builds.

    We don't need go in our system, this two step build docker will build binary in one step and put it in second step in small container without dependency.


            ❯ docker build -t localhost:5000/goapp:latest .
            ❯ docker push localhost:5000/goapp

    Our app is now ready to go in our local registry (on our cluster).

4. Now it's time to deploy! We'll use declarative method of managing kubernetes cluster with use of yml files.

    First step: create deployment (i've created file `deployment.yml` in `kubernetes` directory):

    ```yml
        apiVersion: apps/v1 # for versions before 1.9.0 use apps/v1beta2
        kind: Deployment
        metadata:
        name: goapp-deployment
        spec:
        selector:
            matchLabels:
            app: goapp
        replicas: 2 # tells deployment to run 2 pods matching the template
        template: # create pods using pod definition in this template
            metadata:
            labels:
                app: goapp
            spec:
            containers:
            - name: goapp
                image: localhost:5000/goapp:latest
                ports:
                - containerPort: 8080
    ```

    I've used `NodePort` method for exposing

    Now if our deployment is prepared (image: from our local repository), we're ready do sync our definition with kubernetes cluster:

        ❯ kubectl create -f kubernetes/deployment.yml

    our deployment is ready, we can play a little with it.

    - get pods:

            ❯ kubectl get pods -l app=goapp

            NAME                               READY     STATUS    RESTARTS   AGE
            goapp-deployment-684d96ff7-27hct   1/1       Running   0          1h
            goapp-deployment-684d96ff7-ltl7h   1/1       Running   0          1h

    - get deployments

            ❯  kubectl get deployments -l app=goapp
            NAME               DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
            goapp-deployment   2         2         2            2           1h

    - exec something on pod

            ❯ kubectl exec -it goapp-deployment-684d96ff7-27hct sh
            /app # ps aux
            PID   USER     TIME   COMMAND
                1 root       0:00 /bin/sh -c ./goapp
                5 root       0:00 ./goapp
            61 root       0:00 sh
            65 root       0:00 ps aux
            /app #

        Yeah! there is my goapp running, but **its network is exposed only inside Kubernetes cluster**. Now it's time to expose it outside cluster.

    - expose as Service (`NodePort`) - use mapping of some port on cluster node to external cluster ip.

        Now create our service definition file (`kubernetes/servicey.yml`)

        ```yml
            apiVersion: v1
            kind: Service
            metadata:
            name: goapp
            labels:
                app: goapp
            spec:
            type: NodePort
            ports:
                - port: 8080
                nodePort: 30080
            selector:
                app: goapp
        ```

        And synchronise our cluster:

            ❯ kubectl create -f kubernetes/service.yml

        After that we can check if our service is created correctly:

            ❯ kubectl get service goapp
            NAME      TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
            goapp     NodePort   10.106.164.215   <none>        8080:30080/TCP   26m

        `30080` is our port which will be visible outside cluster


        Next we need to get somehow ip address of kubernetes cluster. I'm using minikube so it's quite simple

            ❯ minikube ip
            # we assign to env variable
            ❯ IP=$(minikube ip)


        When we have external cluster IP now we can access our service on given port.

            ❯ IP=$(minikube ip)
            ❯ curl $IP:30080
            Hello World! 2018-03-19 19:15:47.543450202 +0000 UTC from goapp-deployment-684d96ff7-ltl7h

        Yeah it's working.


# ([code for: Debugging Delve](https://github.com/exu/go-workshops950-debugging-delve))



# ([code for: Debugging Expvar](https://github.com/exu/go-workshops951-debugging-expvar))



## Profilowanie ([code for: Profiling](https://github.com/exu/go-workshops960-profiling))

### Command

Profilowanie standardowego programu

```
import (
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func init() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
	}
}

func main() {
	defer pprof.StopCPUProfile()

```


Następnie budujemy binarkę i odpalamy ją z flagą cpuprofile

```
go build command.go && ./command -cpuprofile=data.prof
```

po czym możemy włączyć nasz profiler

```
go tool pprof command data.prof
```

Możemy do naszego programu dodać memory profile

```
var memprofile = flag.String("memprofile", "", "write memory profile to this file")


func memProfile() {
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}
```

informacje o pamięci zbierane są zbierane na końcu więc dorzucamy do defer list
```
func main() {
    defer memProfile()

```

Teraz możemy zbierać informacje o obu profilach

```
go build command.go && ./command -cpuprofile=cpu.prof -memprofile=mem.prof

go tool pprof command cpu.prof

go tool pprof command mem.prof
```


### Web

```
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func handler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 100000000; i++ {
		w.Write([]byte(fmt.Sprintf("%d - %d, ", i, i)))
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

```


run in shell:

```
go tool pprof http://localhost:8080/debug/pprof/profile
```


# ([code for: Shooting Yourself In The Foot](https://github.com/exu/go-workshops999-shooting-yourself-in-the-foot))

