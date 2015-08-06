+TITLE: Go resources

#  Go lang workshops

## Intro

Slajd 1
Dlaczego Go?

-   łatwy deploy aplikacji (kod nie ma zależności - jenda binarka)
-   brak wojny o code style  `gofmt`
-   zintegrowany package manager `go get`
-   szybkie budowanie binarek
-   ciekawa standardowa biblioteka template/html, performant www servers, json, xml
    streams, io, buffers, first class citizen concurrency

Slajd 2 Dodatkowo ...
-   Kompilacja na wiele maszyn (cross-compilation)
-   łatwy i przyjemny setup środowiska (edytory, ide, code completition server)
-   bardzo niski próg wejścia aby zacząć pisać
-   idiomatyczny
-   wielu dev killerów pisze w go, więc community trochę bardziej inteligentne niż PHP
-   hype w internetach

Slajd 3 :)
Dlaczego nie Go ?
-   brak zarządzania wersjami w package managerze (go get only honors URLs?)
    3rd party - `godep`
    w go 1.5 ma zostać dodana flaga która pozwoli na ładowanie w podobny
    sposób jak godep to robi
    (istnieje co prawda zarządzanie na poziomie pkg server
    przykład mongo "gopkg.in/mgo.v2/bson")
    zmieniamy API podbijamy wersję, API kompatybilne
-   często jeszcze młode biblioteki przykład gin i skopany cache w contrib repo

## Golang basics

Każdy na osobnym slajdzie (1-2 slajdy per item):
-   how to get help `godoc`
-   static typing (types)
-   for if switch and other keywords
-   function return values
-   variadic functions
-   custom types
-   structs
-   interfaces
-   channels
-   timers
-   sql

## Porównanie z PHP

```
|                      | PHP                                           | Go                                                                     |
|----------------------+-----------------------------------------------+------------------------------------------------------------------------|
| Głowna architektura  | Reqest - response app                         | App servers                                                            |
| Typ                  | Skryptowy (z OPCache)                         | Kompilowany                                                            |
| Typowanie            | Dynamiczne (static z dupy w PHP7)             | Statyczne                                                              |
|                      |                                               |                                                                        |
| Zaprojektowany jako: | HTML generator                                | łatwo zarządzalny następca C dla większych projektów                   |
|                      |                                               |                                                                        |
| Szybkość pisania     | Niektóre rzeczy da się zrobić dużo szybciej   | Które tu zrozumiesz jak działają :) i jest szansa                      |
|                      | niż w Go niektóre nie                         | na zwiększenie wydajności (jak nie spierdzielisz)                      |
|                      |                                               |                                                                        |
| Główny paradygmat    | Pseudo Java OO                                | Pseudo-funkcyjny / Pseudo-Obiektowość da się zamodelować na struct'ach |
|                      |                                               |                                                                        |
| RPSy                 | zamodelowanie Event-loop'a kopnie cie w tylek | "Concurrency as CORE feature" WWW serwer może być wystarczająco szybki |
|                      |                                               |                                                                        |
| Modelowanie danych   | Co kraj to obyczaj                            | Tu zauważyłem tendencje do mapowania obiektow/encji na strukturki      |
|                      |                                               | wystarczy wymienić "annotacje"                                         |
|                      |                                               |                                                                        |
```


## Github style project structure

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

i've put them into $HOME :)

## Env Setup (on native machine)

jest kilka sposobów na zainstalowanie go, my aby wszyscy mieli najnowszą wersję
po prostu podamy naszemu systemowi golang prosto w pysk

wget binary

add paths to bashrc/zshrc

voila!

spróbuj `go`




## Standard library

### HTTP

"Helloł Łerld"

### JSON

{"message": "Helloł łerld"}

### Template with angular example

<html><input ng-model="message"></a></html>
<js>$scope.message = "Helloł łerld"</js>

### Channels

simple queued logging system

-   REST post message
-   put message to queue
-   append message to file

one file

### Streams

writer?

### Shell binares - flags

simple shell app (maybe tree implementation)





## Useful libs:

### framework

- echo
- gin
- beego





## Tooling

### IDEs

- TODO IDE - sprawdźić IntelliJ Go plugin
- LiteIDE
- TODO SublimeText
- TODO Atom

### Auto reload

BRA





## Testing

-   Unit testing
-   Mocking HTTP server
-   Chrome driver integration example






## Deploy

### push and run

-   build binary localy (you can change architecture)
    Cross compiling even to ARM or Android
-   push to server
-   run on server

no deps needed!

### docker

-   Przykład postawienia maszynki przez docker-compose
    (base system only)
-   Wspomnienie CoreOS (bez przykładów - osobny temat)

### Spróbuj postawić na OpenShift

### use server on DO?

buy one :) get money from TPM :P

Tu musi być już struktura ustawiona i działający kompilator
każdy kompiluje przykłady,
żeby nie klepać z łapki checkout z taga
git co l001 l002 itp







## Debugging

-   GDB jak w C (słabe)
-   Delve? (brakepoint ustawiany z łapy zależności trzeba przekazywać) pewnie

ktoś to ogarnie w jakimś IDE :P






## Profiling

-   ? function call graph
-   cośtam flame ?


## Workshops wyspawać pomysł kursantów

jak nikt nie da nic ciekawego kilka pomysłów związanych z robotą:

-   Może pobierzmy z kolejki mongo na nmelu dane z jakiejs kolejki i zrobmy cos z nimi
-   Migracja części danych do mongo/cassandry
-   próba przepisania wolnej miracji w PHPie na go
