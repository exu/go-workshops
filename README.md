# Golang Training

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


## Github style struktura projektu [BASICS IMPORTING CODE](010-basics-importing)

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


## Package managers [PACKAGE MANAGEMENT CODE](011-package-management)

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


## Zmienne [BASICS VARIABLES CODE](020-basics-variables)

W go zmienne nie muszą mieć z góry określonego typu, możemy przypisać
zmienną w postaci `a := 1` kompilator będzie wiedział że ma doczynienia
z typem int.


## Stałe [BASICS CONSTANTS CODE](030-basics-constants)

Obsługa standardowa, fajna rzecz `iota` (taki autoincrement)


# [BASICS OVERRIDING INTERNAL TYPES CODE](035-basics-overriding-internal-types)



## Funkcje [BASICS FUNCTIONS CODE](040-basics-functions)

Funkcje w go to "First class citizen".


## Pętle [BASICS LOOPS CODE](040-basics-loops)

W go istnieje tylko jedna pętla: `for`. Wykorzystywana jest jednak w różnych
wariantach, często używana ze słowem kluczowym `range`


## Inicjalzacja package'u [BASICS INIT CODE](041-basics-init)

`func init() {}` odpowiada za zainicjowanie paczki, jest wykonywana tylko
przy pierwszym imporcie paczki.


# [BASICS CLOSURES CODE](042-basics-closures)



## Tablice [BASICS ARRAYS CODE](050-basics-arrays)

Tablice w Go to niskopoziomowa struktura danych, najczęściej z nich nie korzystasz
zamiast tego wykorzystujemy slice'y nakładkę na typy tablicowe znacznie ułatwiającą
pracę z nimi


## Slices [BASICS SLICES CODE](051-basics-slices)

Slice to "nakładka" na tablicę, trzyma do niej referencję
jak przypiszesz slice do slice'a to będą wskazywać na tą
samą tablicę.

źródła:

- https://blog.golang.org/slices
- https://github.com/golang/go/wiki/SliceTricks
- https://blog.golang.org/go-slices-usage-and-internals
- http://research.swtch.com/godata
- http://blog.golang.org/go-slices-usage-and-internals
- http://www.dotnetperls.com/slice-go


## Mapy [BASICS MAPS CODE](055-basics-maps)

Mapy są statycznie typowane jak inne struktury w go, jeżeli potrzebujesz
przechowywać różne typy w jednej mapie możesz uzyć  `interface{}` oznaczjącego
dowolny typ


# [BASICS NEW AND MAKE CODE](059-basics-new-and-make)



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


# [BASICS ANONYMOUS STRUCTS CODE](065-basics-anonymous-structs)



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
- http://davidnix.io/post/error-handling-in-go/


## Panics [BASICS PANICS CODE](068-basics-panics)

- Używamy gdy chcemy zatrzymać program.
- Możemy podpiąć sprawdzenie do "defer chain" czy panic wystąpił


## Stringi [BASICS STRINGS CODE](070-basics-strings)

Podstawowa struktura danych w większości problemów programistycznych
tu też jest i ma wszystko czego potrzebujesz.

Więcej na https://blog.golang.org/strings


## Go routines [BASICS GOROUTINES CODE](080-basics-goroutines)

Concurrency. Goroutine to lekki "wątek" uruchamiany wewnątrz programu
Goroutines są bardzo lekkie więc uruchomienie równolegle wielu tysięcy nie
kosztuje nas zbyt wiele zasobów.


## Używanie tzw 3rd parties [BASICS 3RD PARTY PACKAGES CODE](090-basics-3rd-party-packages)

Go posiada zintegrowany package manager (bez wersjonowania jeszcze niestety)
`go get` pozwala nam w łatwy sposób ściągnąć paczki oraz je uruchamiać.
dzięki powyższej komendzie zostaną ściągnięte wszystkie zależności związane
z packagem `main` (wrzucone niestety do globalnej sciezki z GOPATH)


## DEP

```
go dep init 
```

ściągnie automatycznie zależności  i wrzuci wszystko do katalogu `vendor`


# [BASICS POINTERS CODE](090-basics-pointers)



# Example project structure [BASICS PROJECT STRUCTURE CODE](090-basics-project-structure)


# [CONCURRENCY CHANNELS CODE](100-concurrency-channels)



# [CONCURRENCY OTHER CODE](101-concurrency-other)



# [STDLIB DATE TIME CODE](110-stdlib-date-time)



# [STDLIB OS PROCESSES CODE](110-stdlib-os-processes)



# [STDLIB CLI CODE](114-stdlib-cli)



# Streams - Przykłady [STDLIB STREAMS CODE](115-stdlib-streams)

## Readers

implementują `io.Reader`

## Writers

implementują `io.Writer`.

koniec.


## Podstawowe operacje na IO, Bufory [STDLIB IO CODE](116-stdlib-io)

### Przykład z `bufio`

### Directory traversal

### Łączenie plików za pomocą buforów


# [STDLIB LOGGING CODE](120-stdlib-logging)



# [STDLIB HTTP CODE](140-stdlib-http)



# [STDLIB HTTP MIDDLEWARES CODE](141-stdlib-http-middlewares)



# [STDLIB ENCODING JSON CODE](150-stdlib-encoding-json)



# [STDLIB ENCODING XML CODE](151-stdlib-encoding-xml)



# [STDLIB TEMPLATES CODE](170-stdlib-templates)



# [STDLIB RAND CODE](180-stdlib-rand)



# [STDLIB REGEXP CODE](180-stdlib-regexp)



# [STDLIB CONTEXT CODE](181-stdlib-context)



# [STDLIB SORT CODE](181-stdlib-sort)



# [STDLIB SIGNAL CODE](182-stdlib-signal)



# Task [TASK HTTP RESPONSE READER CODE](199-task-http-response-reader)

MVP:

- Write HTTP response reader from given sources
- You can implement it with channels.
- Push results in JSON format to os.Stdout

Homework for ambitious ones:

- Do it with given interval (e.g. 10 seconds)
- Pass last results back through WWW REST JSON API (e.g. localhost:8080/statuses)


## Instalacja [DATABASES MYSQL CODE](210-databases-mysql)

Uruchom

```
cd docker/mysql
docker-compose start
mysql -uroot -proot -P7701 -h127.0.0.1 < init.sql
mysql -uroot -proot -P7701 -h127.0.0.1
```


## ORMs in Go [DATABASES ORM CODE](215-databases-orm)

### GORM całkiem nieźle rozbudowany ORM - na GH ~ tyle gwiazdek co doctrine

detale na https://github.com/jinzhu/gorm


### GORP - Go Object Relational Persistence

gdy potrzebujesz czegoś lżejszego

https://github.com/go-gorp/gorp


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



# [DATABASES BOLT CODE](241-databases-bolt)



# [DATABASES POSTGRESQL CODE](250-databases-postgresql)



# [QUEUES RABBITMQ CODE](260-queues-rabbitmq)



# [DATABASES INFLUXDB CODE](270-databases-influxdb)



# [TESTING UNIT TASK CODE](300-testing-unit-task)



# [TESTING UNIT EXAMPLES CODE](302-testing-unit-examples)



# [TESTING UNIT DEPENDENCIES CODE](305-testing-unit-dependencies)



# [TESTING HTTP HANDLER CODE](310-testing-http-handler)



# [TESTING HTTP SERVER CODE](310-testing-http-server)



# [TESTING BENCHMARKING CODE](320-testing-benchmarking)



# [TESTING BLACKBOX CODE](340-testing-blackbox)



# [TESTING CHROMEDRIVER CODE](350-testing-chromedriver)



# [TESTING PARALLEL BENCHMARK CODE](380-testing-parallel-benchmark)



## For testing we can use mockery [TESTING MOCKERY CODE](390-testing-mockery)

https://github.com/vektra/mockery


# [PATTERNS PIPELINE CODE](400-patterns-pipeline)



# [PATTERNS GLOW MAP REDUCE CODE](401-patterns-glow-map-reduce)



# [FULLSTACK HTML AND ANGULAR CODE](510-fullstack-html-and-angular)



# [FULLSTACK REST ANGULAR RESOURCE CODE](520-fullstack-rest-angular-resource)



# [FULLSTACK JSON EVENT STREAM CODE](530-fullstack-json-event-stream)



# [FULLSTACK WEBSOCKETS CODE](540-fullstack-websockets)



# [FULLSTACK WIKI CODE](560-fullstack-wiki)



## Beego [FULLSTACK BEEGO CODE](570-fullstack-beego)

Bee init script - inicjuje podstawową strukturę katalogów.
hot compile.

```
go get github.com/astaxie/beego
go get github.com/beego/bee
bee new hello
cd hello
bee run hello
```


# [LIBS MANGOS CODE](601-libs-mangos)



# [LIBS QUANTILE PERCENTILES CODE](602-libs-quantile-percentiles)



# [LIBS BEEP CODE](610-libs-beep)



# [LIBS BRA CODE](610-libs-bra)



# [LIBS SLACK CODE](611-libs-slack)



# [LIBS VEGETA CODE](620-libs-vegeta)



# go readline implementation [LIBS READLINE CODE](630-libs-readline)

https://github.com/chzyer/readline


# [LIBS TERMBOX CODE](640-libs-termbox)



## Caddy webserver [LIBS CADDY CODE](650-libs-caddy)


# [LIBS HTTP ECHO CODE](651-libs-http-echo)



# [LIBS HTTP IRIS CODE](651-libs-http-iris)



# [LIBS JOBRUNNER CODE](660-libs-jobrunner)



# [LIBS CRON CODE](665-libs-cron)



## Validator package [LIBS VALIDATOR CODE](670-libs-validator)

- https://github.com/go-playground/validator

- https://github.com/go-validator/validator

- https://github.com/asaskevich/govalidator


# [LIBS GOGRAPHVIZ CODE](677-libs-gographviz)



# [LIBS FASTHTTP CODE](680-libs-fasthttp)



# [LIBS UIPROGRESS CODE](681-libs-uiprogress)



# [LIBS GO RPM CODE](690-libs-go-rpm)



# [LIBS GRPC CODE](690-libs-grpc)



# [LIBS LOGRUS CODE](690-libs-logrus)



# [LIBS GO PLUGIN CODE](691-libs-go-plugin)



# [LIBS CONSUL CODE](692-libs-consul)



# [LIBS LANGUAGE BINDINGS CODE](693-libs-language-bindings)



# [LIBS ASTIELECTRON CODE](694-libs-astielectron)



# AWS Lambda Golang [LAMBDA SIMPLE CODE](700-lambda-simple)

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


# [HOW TO RUN ON PRODUCTION CODE](800-how-to-run-on-production)



# Traefik [LOAD BALANCING TRAEFIK CODE](801-load-balancing-traefik)

Load balancer with hot reloading


# [DEBUGGING DELVE CODE](950-debugging-delve)



# [DEBUGGING EXPVAR CODE](951-debugging-expvar)



## Profilowanie [PROFILING CODE](960-profiling)

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


# [SHOOTING YOURSELF IN THE FOOT CODE](999-shooting-yourself-in-the-foot)

