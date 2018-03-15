## Using 3rd parties

In go we can get packages to our $GOPATH with use of `go get` command.


## DEP

```sh
go dep init
```

`dep` is fully-fledged dependency manager. It downloads all dependencies source code to `vendor` directory and then compilator includes those code during compilation process.
