# Go training go go go ...

## Agenda

## [Podstawy](00-basics)

## @TODO [Zarządzanie paczkami / zależnościami](01-package-management)

## [HTTP](04-http)

## [JSON](05-json)

## [Strumienie]((15-streams))

## [Jak strukturyzować projekt?](09-project-structure)

## [Kanały](10-channels)

## [Testowanie i dokumentacja](20-testing)

## [Przykłady Full-stack Angular+Go](30-full-stack) czyli jak to ograć z frontendem

## [Cli apps](40-cli)

## [Bazy Danych](70-databases)

## [Debugowanie](80-debugging)

## [Deployment aplikacji](90-deploy)


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
    w go 1.5 została  dodana flaga która pozwoli na ładowanie w podobny
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
