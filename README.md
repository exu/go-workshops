# Golang Training Resources

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


## Github style - project structure ([code for: Basics Importing](010-basics-importing))

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


## Package managers ([code for: Package Management](011-package-management))

Currently most advanced in go ecosystem is `dep` https://github.com/golang/dep

You can init your project:

	$ dep init
	$ ls
	Gopkg.toml Gopkg.lock vendor/

after that dep will add vendor dir where all depndencies will be loaded (In go after 1.5 `vendor` dir have preference over "github style `$GOPATH` based directories - when compiler will not find library in vendor dir it'll try to take it from `$GOPATH`).

For more details please refer to `dep` documentation at https://golang.github.io/dep/docs/daily-dep.html

## Variables ([code for: Basics Variables](020-basics-variables))

In go variables can be declared and assigned in one simple way `a := 1`, compiler will detect type of variable based on its value.


## Constants ([code for: Basics Constants](030-basics-constants))

Use `const` to define new contant, mechanics looks like in other languages. What is worth to mention that we have `iota` keyword which could be used as some kind of iterator in constant definition.


# ([code for: Basics Overriding Internal Types](035-basics-overriding-internal-types))



## Funkcje ([code for: Basics Functions](040-basics-functions))

Functions in Go are "First class citizen".

- function definition
- multiple returns
- named multiple returns
- defered calls
- variadic functions

## Loops ([code for: Basics Loops](040-basics-loops))

In go there is only one loop keyword: `for`. It's often used with `range` keyword to iterate over array like elements.

## Packages initialisation ([code for: Basics Init](041-basics-init))

`func init() {}` is responsible for package initialisation, it's called only once when import statement for given package is called.


## Closures aka anonymous functions ([code for: Basics Closures](042-basics-closures))

Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.


## Tablice ([code for: Basics Arrays](050-basics-arrays))

In Go, an `array` is a **numbered sequence of elements** of a specific length. Arrays are "low level" data structures with slices over them which simplifies creating and managing.


## Slices ([code for: Basics Slices](051-basics-slices))


Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

Unlike arrays, slices are typed only by the elements they contain (not the number of elements). To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued).

sources:

- https://blog.golang.org/slices
- https://github.com/golang/go/wiki/SliceTricks
- https://blog.golang.org/go-slices-usage-and-internals
- http://research.swtch.com/godata
- http://blog.golang.org/go-slices-usage-and-internals
- http://www.dotnetperls.com/slice-go


## Mapy ([code for: Basics Maps](055-basics-maps))

One of the most useful data structures in computer science is the hash table. Many hash table implementations exist with varying properties, but in general they offer fast lookups, adds, and deletes. Go provides a built-in map type that implements a hash table.


## `make` and `new` keywords ([code for: Basics New And Make](059-basics-new-and-make))

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

## Struktury ([code for: Basics Structs Defining](060-basics-structs-defining))

A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a struct, non-blank field names must be unique.


## Struktury - Kompoozycja ([code for: Basics Struct Composition](062-basics-struct-composition))

Kompozycja taki pattern chyba znacie ?


## Struct tags (annotations like) ([code for: Basics Struct Tags](064-basics-struct-tags))

A tag for a field allows you to attach meta-information to the field which can be acquired using reflection. Usually it is used to provide transformation info on how a struct field is encoded to or decoded from another format (or stored/retrieved from a database), but you can use it to store whatever meta-info you want to, either intended for another package or for your own use.


# ([code for: Basics Anonymous Structs](065-basics-anonymous-structs))



## Interface'y ([code for: Basics Interfaces](065-basics-interfaces))

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

## Error handling ([code for: Basics Errors](067-basics-errors))

There is no exceptions in Go, errors are returned by value, or aggregated in intermediate objects. In go error is simply value which should be handled programatically as quick as possible.

Sources:
- https://blog.golang.org/errors-are-values
- http://blog.golang.org/error-handling-and-go
- http://davidnix.io/post/error-handling-in-go/


## Panics ([code for: Basics Panics](068-basics-panics))

- Used when we want to stop the program.
- We can check if there was a `panic` occurence in function defer chain


## Strings ([code for: Basics Strings](070-basics-strings))

In Go, a string is in effect a **read-only slice of bytes**.

It's important to state right up front that a string holds arbitrary bytes. It is not required to hold Unicode text, UTF-8 text, or any other predefined format. As far as the content of a string is concerned, it is exactly equivalent to a slice of bytes.

More on https://blog.golang.org/strings


## Go routines ([code for: Basics Goroutines](080-basics-goroutines))

A goroutine is a lightweight thread of execution.

Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives (channels).

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.


## Using 3rd parties ([code for: Basics 3rd Party Packages](090-basics-3rd-party-packages))

In go we can get packages to our $GOPATH with use of `go get` command.


## DEP

```sh
go dep init
```

`dep` is fully-fledged dependency manager. It downloads all dependencies source code to `vendor` directory and then compilator includes those code during compilation process.


# ([code for: Basics Pointers](090-basics-pointers))



## Channels ([code for: Concurrency Channels](100-concurrency-channels))

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

## Other concurrency patterns ([code for: Concurrency Other](101-concurrency-other))

You've lerned about channels primitive on previous chapter of this workshops, now it's time to learn about some other ways of creating concurrent programs in go.

We'll be doing some code with:
- data races errors
- atomic counters
- mutexes
- wait groups


## Date time ([code for: Stdlib Date Time](110-stdlib-date-time))

Go has very powerful standard library, first and one of awesome library are is Date time.

# ([code for: Stdlib Os Processes](110-stdlib-os-processes))



# ([code for: Stdlib Cli](114-stdlib-cli))



# Streams - Przykłady ([code for: Stdlib Streams](115-stdlib-streams))

## Readers

implementują `io.Reader`

## Writers

implementują `io.Writer`.

koniec.


## Podstawowe operacje na IO, Bufory ([code for: Stdlib Io](116-stdlib-io))

### Przykład z `bufio`

### Directory traversal

### Łączenie plików za pomocą buforów


# ([code for: Stdlib Logging](120-stdlib-logging))



# ([code for: Stdlib Http](140-stdlib-http))



# ([code for: Stdlib Http Middlewares](141-stdlib-http-middlewares))



# ([code for: Stdlib Encoding Json](150-stdlib-encoding-json))



# ([code for: Stdlib Encoding Xml](151-stdlib-encoding-xml))



# ([code for: Stdlib Templates](170-stdlib-templates))



# ([code for: Stdlib Rand](180-stdlib-rand))



# ([code for: Stdlib Regexp](180-stdlib-regexp))



# ([code for: Stdlib Context](181-stdlib-context))



# ([code for: Stdlib Sort](181-stdlib-sort))



# ([code for: Stdlib Signal](182-stdlib-signal))



# Task ([code for: Task Http Response Reader](199-task-http-response-reader))

MVP:

- Write HTTP response reader from given sources
- You can implement it with channels.
- Push results in JSON format to os.Stdout

Homework for ambitious ones:

- Do it with given interval (e.g. 10 seconds)
- Pass last results back through WWW REST JSON API (e.g. localhost:8080/statuses)


## Instalacja ([code for: Databases Mysql](210-databases-mysql))

Uruchom

```
cd docker/mysql
docker-compose start
mysql -uroot -proot -P7701 -h127.0.0.1 < init.sql
mysql -uroot -proot -P7701 -h127.0.0.1
```


## ORMs in Go ([code for: Databases Orm](215-databases-orm))

### GORM całkiem nieźle rozbudowany ORM - na GH ~ tyle gwiazdek co doctrine

detale na https://github.com/jinzhu/gorm


### GORP - Go Object Relational Persistence

gdy potrzebujesz czegoś lżejszego

https://github.com/go-gorp/gorp


# Przykłady MongoDB ([code for: Databases Mongodb](220-databases-mongodb))


# RethinkDB ([code for: Databases Rethinkdb](230-databases-rethinkdb))

Jedną z ciekawych funkcjonalności RethinkDB jest możliwość
monitorowania zmian na kolekcji

# Dockerizing

## Uruchom rethinka jako kontener

```
docker run --name some-rethink -v "$PWD:/data" -d rethinkdb
```

mozemy również zmapować porty do lokalnej maszynki

```
docker run --name rethink -p 28015:28015 -v "$PWD/data:/data" -d rethinkdb
```

wtedy instacja kontenera będzie widoczna pod adresem `localhost:28015`

## Zlinkuj go w swojej aplikacji

```
docker run --name some-app --link some-rethink:rdb -d application-that-uses-rdb
```


# ([code for: Databases Redis](240-databases-redis))



# ([code for: Databases Bolt](241-databases-bolt))



# ([code for: Databases Postgresql](250-databases-postgresql))



# ([code for: Queues Rabbitmq](260-queues-rabbitmq))



# ([code for: Databases Influxdb](270-databases-influxdb))



# ([code for: Testing Unit Task](300-testing-unit-task))



# ([code for: Testing Unit Examples](302-testing-unit-examples))



# ([code for: Testing Unit Dependencies](305-testing-unit-dependencies))



# ([code for: Testing Http Handler](310-testing-http-handler))



# ([code for: Testing Http Server](310-testing-http-server))



# ([code for: Testing Benchmarking](320-testing-benchmarking))



# ([code for: Testing Blackbox](340-testing-blackbox))



# ([code for: Testing Chromedriver](350-testing-chromedriver))



# ([code for: Testing Parallel Benchmark](380-testing-parallel-benchmark))



## For testing we can use mockery ([code for: Testing Mockery](390-testing-mockery))

https://github.com/vektra/mockery


# ([code for: Patterns Pipeline](400-patterns-pipeline))



# ([code for: Patterns Glow Map Reduce](401-patterns-glow-map-reduce))



# ([code for: Fullstack Html And Angular](510-fullstack-html-and-angular))



# ([code for: Fullstack Rest Angular Resource](520-fullstack-rest-angular-resource))



# ([code for: Fullstack Json Event Stream](530-fullstack-json-event-stream))



# ([code for: Fullstack Websockets](540-fullstack-websockets))



# ([code for: Fullstack Wiki](560-fullstack-wiki))



## Beego ([code for: Fullstack Beego](570-fullstack-beego))

Bee init script - inicjuje podstawową strukturę katalogów.
hot compile.

```
go get github.com/astaxie/beego
go get github.com/beego/bee
bee new hello
cd hello
bee run hello
```


# ([code for: Libs Mangos](601-libs-mangos))



# ([code for: Libs Quantile Percentiles](602-libs-quantile-percentiles))



# ([code for: Libs Beep](610-libs-beep))



# ([code for: Libs Bra](610-libs-bra))



# ([code for: Libs Slack](611-libs-slack))



# ([code for: Libs Vegeta](620-libs-vegeta))



# go readline implementation ([code for: Libs Readline](630-libs-readline))

https://github.com/chzyer/readline


# ([code for: Libs Termbox](640-libs-termbox))



## Caddy webserver ([code for: Libs Caddy](650-libs-caddy))


# ([code for: Libs Http Echo](651-libs-http-echo))



# ([code for: Libs Http Iris](651-libs-http-iris))



# ([code for: Libs Jobrunner](660-libs-jobrunner))



# ([code for: Libs Cron](665-libs-cron))



## Validator package ([code for: Libs Validator](670-libs-validator))

- https://github.com/go-playground/validator

- https://github.com/go-validator/validator

- https://github.com/asaskevich/govalidator


# ([code for: Libs Gographviz](677-libs-gographviz))



# ([code for: Libs Fasthttp](680-libs-fasthttp))



# ([code for: Libs Uiprogress](681-libs-uiprogress))



# ([code for: Libs Go Rpm](690-libs-go-rpm))



# ([code for: Libs Grpc](690-libs-grpc))



# ([code for: Libs Logrus](690-libs-logrus))



# ([code for: Libs Go Plugin](691-libs-go-plugin))



# ([code for: Libs Consul](692-libs-consul))



# ([code for: Libs Language Bindings](693-libs-language-bindings))



# ([code for: Libs Astielectron](694-libs-astielectron))



# AWS Lambda Golang ([code for: Lambda Simple](700-lambda-simple))

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


# ([code for: How To Run On Production](800-how-to-run-on-production))



# Traefik ([code for: Load Balancing Traefik](801-load-balancing-traefik))

Load balancer with hot reloading


# ([code for: Debugging Delve](950-debugging-delve))



# ([code for: Debugging Expvar](951-debugging-expvar))



## Profilowanie ([code for: Profiling](960-profiling))

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


# ([code for: Shooting Yourself In The Foot](999-shooting-yourself-in-the-foot))

