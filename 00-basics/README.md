
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


## Go Tools

### Included

- go test - wbudowane testowanie
- go fmt - code formater - tylko jeden słuszny coding standard (https://golang.org/pkg/fmt/)
- gocode - autocomplete service (https://github.com/nsf/gocode)
- go vet - znajduje błędy (http://godoc.org/golang.org/x/tools/cmd/vet)
- go oracle - wyszukiwanie zależności (http://golang.org/s/oracle-user-manual)
- godoc - generator dokumentacji

### IDEs

- TODO IDE - sprawdźić IntelliJ Go plugin
- LiteIDE
- TODO SublimeText
- TODO Atom

### Auto reload

- BRA
