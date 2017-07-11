## Package managers

Niestety ta część jest nadal słaba nie ma `virtualenv`'a :(.

### godep

Nie umiem ustawić wersji - wygląda na to że on robi coś
w rodzaju snapshota z aktualnego drzewa zależności
(chyba że coś mnie ominęło w dokumentacji)

### gopm

Wygląda na to że tutaj też nie można ustawić taga ani commita


### gpm + gpv

`gpm` ma buga jeżeli package istnieje w globalnej ścieżce to niestety
się wypie****. Jak bug zostanie poprawiony następna metoda będzie
niepotrzebna ponieważ `gvp` będzie nam zarządzać naszym ENV'em

### gpm + autoenv

Tutaj udało mi się uruchomić wersjonowanie za pomocą prostego tricku
z przepisywaniem `GOPATH` (de facto wszelkie pckage managery działają
na tej zasadzie)

- https://github.com/kennethreitz/autoenv

Ja instalowałem przez `pip`
```
pip install autoenv
```

Następnie należy dodać
```
source /usr/bin/activate.sh
```

do `~/bashrc` lub `~/.zshrc` (lub `.whateveryoureusingrc`)

a w katalogu projektu dodać automatyczne przepisywanie ścieżki na
```
echo 'export GOPATH=$(pwd)/.godeps:$(pwd)' > .env
```

w katalogu wyżej możemy ustawić plik `.env` który będzie wracał do
`root` lub stworzyć alias który będzie


### Przykład

Najpierw ustawiamy lokalnie GOPATH na `GOPATH=$(pwd)/.godeps:$(pwd)`
na potrzeby przykładu możemy to zrobić ręcznie korzystać z autoenv
lub poczekać na usunięcie buga z repo :)

```
package main

import (
	"github.com/exu/supergolib/printer"
)

func main() {
	printer.PrintMe()
}

```


Użyjemy `gpm-example` który ma zależność od biblioteki `github.com/exu/supergolib/printer`
supergolib zostało utworzone wcześniej na githubie, ma 3 tagi `v1`, `v2`, `v23`

definiujemy pliczek z zależnościami

```
github.com/exu/supergolib/printer        v1
```

uruchamiamy
```
gpm
go run echo.go
```

gotowe

Aby zmienić wersję


```
github.com/exu/supergolib/printer        v2
```
uruchamiamy
```
gpm
go run echo.go
```



## Go version manager - zarządzanie wersjami go

- https://github.com/moovweb/gvm



## Glide

- + vendoring (default enabled in 1.6)
