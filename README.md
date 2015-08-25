#  Go lang workshops

## Intro

Dlaczego Go?

-   łatwy deploy aplikacji (kod nie ma zależności - jenda binarka)
-   brak wojny o code style  `gofmt`
-   zintegrowany package manager `go get`
-   szybkie budowanie binarek
-   ciekawa standardowa biblioteka template/html, performant www servers, json, xml
    streams, io, buffers, first class citizen concurrency

-   Kompilacja na wiele maszyn (cross-compilation)
-   łatwy i przyjemny setup środowiska (edytory, ide, code completition server)
-   bardzo niski próg wejścia aby zacząć pisać
-   idiomatyczny
-   wielu dev killerów pisze w go, więc community trochę bardziej inteligentne niż PHP
-   hype w internetach

Dlaczego nie Go ?
-   brak zarządzania wersjami w package managerze (go get only honors URLs?)
    3rd party - `godep`
    w go 1.5 ma zostać dodana flaga która pozwoli na ładowanie w podobny
    sposób jak godep to robi
    (istnieje co prawda zarządzanie na poziomie pkg server
    przykład mongo "gopkg.in/mgo.v2/bson")
    zmieniamy API podbijamy wersję, API kompatybilne
-   często jeszcze młode biblioteki przykład gin i skopany cache w contrib repo

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


# Agenda

## [Podstawy](/exu/go-workshops/tree/master/00-basics)

## @TODO [Zarządzanie paczkami / zależnościami](/exu/go-workshops/tree/master/01-package-management)

## [HTTP](/exu/go-workshops/tree/master/04-http)

## [JSON](/exu/go-workshops/tree/master/05-json)

## @TODO [Strumienie]((/exu/go-workshops/tree/master/15-streams))

## [Jak strukturyzować projekt?](/exu/go-workshops/tree/master/09-project-structure)

## [Kanały](/exu/go-workshops/tree/master/10-channels)

## [Testowanie i dokumentacja](/exu/go-workshops/tree/master/20-testing)

## [Przykłady Full-stack Angular+Go](/exu/go-workshops/tree/master/30-full-stack) czyli jak to ograć z frontendem

## [Cli apps](/exu/go-workshops/tree/master/40-cli)

## [Bazy Danych](/exu/go-workshops/tree/master/70-databases)

## [Debugowanie](/exu/go-workshops/tree/master/80-debugging)

## [Deployment aplikacji](/exu/go-workshops/tree/master/90-deploy)
