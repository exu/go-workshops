# Goland Training

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
- zintegrowane sprawdzanie poprawności kodu `go vet`
- `gocode` serwer do intellisense - nie musisz miec IDE aby ci podpowiadało
  możesz pisać w swoim ulubionym edytorze (Sublime, Atom, Vim)
- szybkie budowanie binarek
- ciekawa standardowa biblioteka template/html, performant www servers, json, xml
  streams, io, buffers, first class citizen concurrency
- kompilacja na wiele maszyn (cross-compilation)
- łatwy i przyjemny setup środowiska (edytory, ide, code completition server)
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


## Importowanie i packages [BASICS IMPORTING CODE](010-basics-importing)

W go package jest zbiorem plików z dyrektywą `package nazwa`


## Zmienne [BASICS VARIABLES CODE](020-basics-variables)

W go zmienne nie muszą mieć z góry określonego typu, możemy przypisać
zmienną w postaci `a := 1` kompilator będzie wiedział że ma doczynienia
z typem int.


## Stałe [BASICS CONSTANTS CODE](030-basics-constants)

Obsługa standardowa, fajna rzecz `iota` (taki autoincrement)


# [OVERRIDING INTERNAL TYPES CODE](035-overriding-internal-types)



## Funkcje [BASICS FUNCTIONS CODE](040-basics-functions)

Funkcje w go to "First class citizen".


## Pętle [BASICS LOOPS CODE](040-basics-loops)

W go istnieje tylko jedna pętla: `for`. Wykorzystywana jest jednak w różnych
wariantach, często używana ze słowem kluczowym `range`


## Inicjalzacja package'u [BASICS INIT CODE](041-basics-init)

`func init() {}` odpowiada za zainicjowanie paczki, jest wykonywana tylko
przy pierwszym imporcie paczki.


## Tablice [BASICS ARRAYS CODE](050-basics-arrays)

Tablice w Go to niskopoziomowa struktura danych, najczęściej z nich nie korzystasz
zamiast tego wykorzystujemy slice'y nakładkę na typy tablicowe znacznie ułatwiającą
pracę z nimi


## Slices [BASICS SLICES CODE](051-basics-slices)

Slice to nakładka na tablicę często reprezentuje część tablicy, przy dynamicznym
tworzeniu nie musisz się zajmować pojemnością slice'a,

źródła:
- http://blog.golang.org/go-slices-usage-and-internals
- https://blog.golang.org/slices
- http://www.dotnetperls.com/slice-go


## Mapy [BASICS MAPS CODE](055-basics-maps)

Mapy są statycznie typowane jak inne struktury w go, jeżeli potrzebujesz
przechowywać różne typy w jednej mapie możesz uzyć  `interface{}` oznaczjącego
dowolny typ


## Struktury [BASICS STRUCTS DEFINING CODE](060-basics-structs-defining)

Struktury to podstawowy typ danych w go, większość driverów do storage'u
pozwala kodować i dekodować do struktur. są bardzo użyteczne, dzięki nim
możemy zamodelować pseudo-obiektowość, często używane do kompzycji
oprogramowania.


## Struktury - Kompoozycja [BASICS STRUCT COMPOSITION CODE](062-basics-struct-composition)

Kompozycja taki pattern chyba znacie ?


## Tagi - annotacje [BASICS STRUCT TAGS CODE](064-basics-struct-tags)

Struktury w go posiadają możliwość opisywania pól dodatkowymi
"tagami" które następnie możemy wykorzystać programistycznie
wykorzystywane często przy enkodowaniu i dekodowaniu z formatów danych (JSON, yml)
jak i z różnych silników baz danych.


## Interface'y [BASICS INTERFACES CODE](065-basics-interfaces)

Go posiada tzw "implicit interfaces", oznacza to że jeżeli dana struktura
implementuje metody z interfejsu automatycznie staje się obikektem o typie
równym typie intefejsu.


## Zarządzanie błędami [BASICS ERRORS CODE](067-basics-errors)

W go nie ma exceptionów, błędy są zwracane poprzez wielokrotne wartości
lub agregowane w obiektach jeżeli zachodzi taka potrzeba. Preferuje
się podejście jak najszybszej obsługi błędów. W Go błąd jest wartością
na którą masz reagować.

Źródła:
- https://blog.golang.org/errors-are-values
- http://blog.golang.org/error-handling-and-go


## Panics [BASICS PANICS CODE](068-basics-panics)

- Używamy gdy chcemy zatrzymać program.
- Możemy podpiąć sprawdzenie do "defer chain" czy panic wystąpił


## Stringi [BASICS STRINGS CODE](070-basics-strings)

Podstawowa struktura danych w większości problemów programistycznych
tu też jest i ma wszystko czego potrzebujesz.


## Go routines [BASICS GOROUTINES CODE](080-basics-goroutines)

Concurrency. Goroutine to lekki "wątek" uruchamiany wewnątrz programu
Goroutines są bardzo lekkie więc uruchomienie równolegle wielu tysięcy nie
kosztuje nas zbyt wiele zasobów.


## Używanie tzw 3rd parties [BASICS 3RD PARTY PACKAGES CODE](090-basics-3rd-party-packages)

Go posiada zintegrowany package manager (bez wersjonowania jeszcze niestety)
`go get` pozwala nam w łatwy sposób ściągnąć paczki oraz je uruchamiać.
dzięki powyższej komendzie zostaną ściągnięte wszystkie zależności związane
z packagem `main`


# [BASICS POINTERS CODE](090-basics-pointers)



# Example project structure [BASICS PROJECT STRUCTURE CODE](090-basics-project-structure)


# [STDLIB CLI CODE](100-stdlib-cli)



# [STDLIB CONCURRENCY CODE](100-stdlib-concurrency)



# [STDLIB DATE TIME CODE](110-stdlib-date-time)



# [STDLIB OS PROCESSES CODE](110-stdlib-os-processes)



# Streams - Przykłady [STDLIB STREAMS CODE](115-stdlib-streams)

## Readers

implementują `io.Reader`

## Writers

implementują `io.Writer`.

koniec.


# [STDLIB LOGGING CODE](120-stdlib-logging)



# [STDLIB HTTP CODE](140-stdlib-http)



# [STDLIB ENCODING JSON CODE](150-stdlib-encoding-json)



# [STDLIB ENCODING XML CODE](151-stdlib-encoding-xml)



## Podstawowe operacje na IO, Bufory [STDLIB IO CODE](160-stdlib-io)

### Przykład z `bufio`

### Directory traversal

### Łączenie plików za pomocą buforów


# [STDLIB TEMPLATES CODE](170-stdlib-templates)



# [STDLIB REGEXP CODE](180-stdlib-regexp)



# [DATABASES MYSQL CODE](210-databases-mysql)



# Przykłady MongoDB [DATABASES MONGODB CODE](220-databases-mongodb)


# RethinkDB [DATABASES RETHINKDB CODE](230-databases-rethinkdb)

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


# [DATABASES REDIS CODE](240-databases-redis)



# [TESTING UNIT CODE](300-testing-unit)



# [TESTING HTTP CODE](310-testing-http)



# [TESTING BENCHMARKING CODE](320-testing-benchmarking)



# [TESTING BLACKBOX CODE](340-testing-blackbox)



# [TESTING CHROMEDRIVER CODE](350-testing-chromedriver)



# [TESTING PARALLEL BENCHMARK CODE](380-testing-parallel-benchmark)



# [FULLSTACK HTML AND ANGULAR CODE](510-fullstack-html-and-angular)



# [FULLSTACK REST ANGULAR RESOURCE CODE](520-fullstack-rest-angular-resource)



# [FULLSTACK JSON EVENT STREAM CODE](530-fullstack-json-event-stream)



# [FULLSTACK WEBSOCKETS CODE](540-fullstack-websockets)



# [FULLSTACK WIKI CODE](560-fullstack-wiki)



# [LIBS BRA CODE](610-libs-bra)



# [LIBS VEGETA CODE](620-libs-vegeta)



# go readline implementation [LIBS READLINE CODE](630-libs-readline)

https://github.com/chzyer/readline


## Validator package [LIBS VALIDATOR CODE](630-libs-validator)

github.com/go-playground/validator


# [LIBS TERMBOX CODE](640-libs-termbox)



## Caddy webserver [LIBS CADDY CODE](650-libs-caddy)


# [HOW TO RUN ON PRODUCTION CODE](800-how-to-run-on-production)



# [DEBUGGING EXPVAR CODE](910-debugging-expvar)



# [DEBUGGING DELVE CODE](950-debugging-delve)



# [SHOOTING YOURSELF IN THE FOOT CODE](999-shooting-yourself-in-the-foot)

