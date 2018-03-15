## Package managers

Currently most advanced in go ecosystem is `dep` https://github.com/golang/dep

You can init your project:

	$ dep init
	$ ls
	Gopkg.toml Gopkg.lock vendor/

after that dep will add vendor dir where all depndencies will be loaded (In go after 1.5 `vendor` dir have preference over "github style `$GOPATH` based directories - when compiler will not find library in vendor dir it'll try to take it from `$GOPATH`).

For more details please refer to `dep` documentation at https://golang.github.io/dep/docs/daily-dep.html