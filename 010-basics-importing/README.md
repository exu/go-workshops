## Github style struktura projektu

w Go projekty są uporządkowane zgodnie z "github style" - częścią ścieżki
jest adres serwera na którym hostowane jest projekt / biblioteka

```
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

Zmienna środowiskowa `$GOPATH` decyduje gdzie się znajdują te katalogi w
twoim systemie.


## Importowanie i packages

W go package jest zbiorem plików z dyrektywą `package nazwa`
Package default'owo jest importowany po pełnej ścieżce ścieżce

    import "github.com/exu/go-workshops/010-basics-importing/sub"

## Pobieranie i instalowanie

```
go install external.package.com/uri/to/package
```
