# Podstawy

## Instalacja Go

Ważne aby mieć nową wersję, nie gwarantuję że na starszych będą działać te przykłady
(choć powinny)


Pobierz paczkę dla twojego systemu:

- Linux: https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz
- MacOS: https://storage.googleapis.com/golang/go1.5.1.darwin-amd64.pkg

Przykład dla Linux'a:
```
wget https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.5.1.linux-amd64.tar.gz
```


- dodaj `export PATH=$PATH:/usr/local/go/bin` do czegoś wykonywalnego
przy uruchomieniu systemu (zazwyczaj ~/.bashrc)

Szczegóły instalacji na https://golang.org/doc/install


Ustawmy "workspace" dla naszych projektów go,
możesz to zrobić w pliku twojego shella z reguły `.bashrc`

```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
```

otwórz terminal ponownie aby bashrc się wczytał.

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
