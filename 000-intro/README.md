## Dlaczego Go?

Kto jest głównym pomysłodawcą języka
- Rob Pike (Unix, UTF-8)
- Ken Thompson (Unix author, UTF-8, B lang)
- Robert Griesemer (V8, Java Hotspot (Oracle Java), GFS)

lista kontrybutorów https://golang.org/AUTHORS

Dlaczego go zostało stworzone przez Google?
Chcieli rozzwiązać problemy z DUŻYMI aplikacjami które mieli napisane w Google:
- przyspieszyć development
- multicore systems

źródła:
- https://golang.org/doc/faq#What_is_the_purpose_of_the_project
- https://talks.golang.org/2012/splash.article


## Charakterystyka Go
- Statycznie kompilowany (jedna binarka ze wszystkimi zależnościami)
- Garbage Collected
- Silnie typowany
- Funkcyjno - Pseudo obiektowa hybryda

## Dlaczego jest warte uwagi?
- łatwy deploy aplikacji (kod nie ma zależności - jenda binarka) + statics
- brak wojny o code style  `gofmt`
- zintegrowany package manager `go get`
- zintegrowane sprawdzanie poprawności kodu `go vet` także `golint` (github.com/golang/lint)
- `gocode` serwer do intellisense - nie musisz miec IDE aby ci podpowiadało
  możesz pisać w swoim ulubionym edytorze (Sublime, Atom, Vim)
- szybkie budowanie binarek
- ciekawa standardowa biblioteka template/html, performant www servers, json, xml
  streams, io, buffers, first class citizen concurrency
- kompilacja na wiele maszyn (cross-compilation)
- łatwy i przyjemny setup środowiska (edytory, ide, code completition server)
- wbudowane testy
- bardzo niski próg wejścia aby zacząć pisać
- hype w internetach, jeden z większych wzrostów w trendach na githubie oraz google trends
  (Tiobe kłamie!)

## Dlaczego jest niefajne?
- brak zarządzania wersjami w package managerze (go get only honors URLs?)
    3rd party - `godep`
- w go 1.5 została dodana flaga która pozwoli na ładowanie w podobny
    sposób jak godep to robi
    (istnieje co prawda zarządzanie na poziomie pkg server
    przykład mongo "gopkg.in/mgo.v2/bson")
    zmieniamy API podbijamy wersję, API kompatybilne
- często jeszcze młode biblioteki przykład gin i skopany cache w contrib repo


## Init

    go get github.com/exu/go-workshops
    cd $GOPATH/exu/go-workshops

## Porównanie z PHP

Architektura web aplikacji
- PHP: Reqest - response app
- Go:  Serwery aplikacyjne (a'la NodeJs, Ruby, Python)

Typ
- PHP: Skryptowy (z OPCache)
- Go: Kompilowany

Typowanie
- PHP: Dynamiczne (static z dupy w PHP7)
- Go:  Statyczne

Zaprojektowany jako:
- PHP: HTML generator
- Go:  Łatwo zarządzalny następca C dla większych projektów

Entry level
- PHP: Trochę trzeba się naumieć aby przykrywać niekompetencje frameworków innymi design patternami
- Go: niski dla podstawowych, trzeba zrozumieć concurrency i kilka dziwnych rozwiązań (np: err handling)

Główny paradygmat
- PHP: Pseudo Java OO
- Go: Pseudo-funkcyjny / Pseudo-Obiektowość da się zamodelować na struct'ach

RPSy
- PHP: Req-Response boli, zamodelowanie Event-loop'a kopnie cie w tylek
- Go: "Concurrency as CORE feature" WWW serwer może być wystarczająco szybki

Użycie w świecie:
- PHP: Tu wygrywa bardzo dużo softu, mało ciekawych projektów
- Go: Jak już coś wychodzi o czym słychać to z reguły jest fajne :)


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

- IntelliJ Go plugin
  - https://github.com/go-lang-plugin-org/go-lang-idea-plugin/wiki/Documentation
  - http://blog.pivotal.io/labs/labs/setting-google-go-plugin-intellij-idea-13-os-x-10-8-5
    testowane na IntelliJ 14.1.4

- LiteIDE
- TODO SublimeText
- TODO Atom

### Auto reload

- BRA



## Testing

- Unit
- Http
- Bdd tools
- Blackbox testing
- Benchmarking
- Chromedriver example

### Assertion libs

Go nie posiada wbudowanej biblioteki do asercji, istnieje za to
wiele projektów open source:

- http://onsi.github.io/gomega/
- https://github.com/assertgo/assert
- https://github.com/stretchr/testify


## Database drivers

Storages:
- Mongo
- RethinkDB
- Redis

- MySQL?
- Cassandra?


## Biblioteki

### workflow

- BRA https://github.com/Unknwon/bra

### Web frameworks

- echo (MUX)
- gin (MUX)
- beego (większy z ORMem)

### Stress testing

- vegeta

## Go debugging

### Expvar i expvarmon

live app monitoring

### W dokumentacji jest integracja z GDB

ale to jest trochę słabe


##3 Delve

Filmik autora z Gophercona:
https://www.youtube.com/watch?v=InG72scKPd4

### Profiling

https://github.com/bradfitz/talk-yapc-asia-2015/blob/master/talk.md
