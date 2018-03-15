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


## Init

    go get github.com/exu/go-workshops
    cd $GOPATH/exu/go-workshops


## Github style - projects structure

In Go ecosystem projects are organised in "github-style" - part of module path is path to server addess e.g.:

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
