# Podstawy

## Instalacja Go

TODO !!!!! version bump !!!!

Ważne aby mieć nową wersję, nie gwarantuję że na starszych będą działać te przykłady
(choć powinny)

Postępuj zgodnie z dla twojego systemu:
https://golang.org/doc/install

Jest kilka sposobów na zainstalowanie go, my aby wszyscy mieli najnowszą wersję
po prostu podamy naszemu systemowi golang prosto w pysk

- `wget https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz`

- `tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`

- dodaj `export PATH=$PATH:/usr/local/go/bin` do czegoś wykonywalnego
przy uruchomieniu systemu (zazwyczaj ~/.bashrc)

- spróbuj uruchomić `go version`


## Instalacja Docker

Aby uruchomić bazy danych wykorzystywane w szkoleniu (jeżeli nie posiadasz
zainstalowanych natywnie) należy zainstalować dockera oraz docker-compose
(tool do zarządzania environmentem kontenerów)


TODO !!!!

## Instalacja Docker Compose

TODO !!!!


## Inicjalizacja zawarości repozytoriów dockera przed szkoleniem

Aby nie przeciążać sieci przed szkoleniem należałoby odpalić `docker-compose build` dla każdego z kontenerów (każdy ściąga ~ kilkaset MB)

aby to zrobić uruchom make w głownym katalogu repozytorium

```
make
```
