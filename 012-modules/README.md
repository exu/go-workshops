## Modules

Modules are quite new in Go (begins in 1.11 as experimental feature)
in `1.12` they are in auto mode what means that if you are in GOPATH
they will be disabled and when you'll be out of GOPATH they will be enabled

In incoming 1.13 version modules will be enabled by default which means that
no more github style code organisation will be needed


### Commands

0. Check `echo $GO111MODULE` first - if value is `off` or `auto` you need to set it to `on`
1. `go mod init` initialize modules in project directory
2. `go mod tidy` cleans your unused modules
3. `go mod download` downloads modules to cache
4. `go mod vendor` vendoring modules from cache

5. `go build`, `go test`, `go run` downloads modules automatically so you don't need to do it
6. `go run -mod vendor` using vedored modules

### Tools

As in middle 2019 editors have still issues with new modules, you need
to install proper `gocode` or can use `gopls` (language server) for
autocompletition
