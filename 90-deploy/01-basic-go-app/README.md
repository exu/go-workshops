# Używanie dockera


## Budowanie kontenera

Dla go istnieją oficjalne obrazy dodajemy do naszego pliku
`Dockerfile` następującą linijkę:

    FROM golang:onbuild


następnie odpalamy:

    docker build .

Obraz ma ustawioną binarkę `app` jako główny `CMD` Dockerfile'a



## Trochę lepsze budowanie kontenera

Użyjemy "docker comopse" aby skonfigurować zależności pomiędzy
naszymi aplikacjami:

    app:
      build: .
      ports:
      - "8080:8080"

umieść powyższy kod w pliku  `docker-compose.yml` i uruchom:

    docker-compose up

docker-compose pokaże informacje o budowaniu aplikacji:

    ...
    ...
    Successfully built db28da3b5081
    Attaching to dockerize_app_1
    app_1 | + exec app

## Co się stało?

- Nasz kontener specjalnie zbudowanu dla Go zbudował sobie środowisko
  w oparciu o LXC (taka wirtualizacja na wspólnym kernelu z systemem
  z bardzo małym narzutem dodatkowych zasobów)

- sieć kontenera została połączona z naszym localhostem na porcie
  8080 - został stworzony mapping portu hosta (ip dynamiczne)
  kontenera do localhost:8080

- Został uruchomiony build aplikacji wewnątrz kontenera dla entrypointa
 `app.go`

- został uruchomiony plik binarny `app` (domyślny plik wyjściowy
  po zbudowaniu pliku `app.go` przez `go build app.go`

Od teraz możemy używać naszej aplikacji tak jakby była uruchomiona na
naszym lokalnym hoście

    ❯ curl localhost:8080/something
    Hi there,something!%

Uruchomiony kontener dostaje również swój adres IP możemy bez problemu
do niego się dostać:

    ❯ curl 172.17.0.11:8080/something
    Hi there,something!%
