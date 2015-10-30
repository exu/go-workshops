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

```
mkdir ~/go
```

Ustawmy "workspace" dla naszych projektów go,
możesz to zrobić w pliku twojego shella z reguły `.bashrc`

```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
```

otwórz terminal ponownie aby bashrc się wczytał.

- spróbuj uruchomić `go version`

Ściągnij pliki szkolenia (wymagany `git`!):

```
go get github.com/exu/go-workshops
```

źródła właśnie zostały pobrane do katalogu:

- `~/go/src/github.com/exu/go-workshops`


## Instalacja Docker

Aby uruchomić bazy danych wykorzystywane w szkoleniu (jeżeli nie posiadasz
zainstalowanych natywnie) należy zainstalować dockera oraz docker-compose
(tool do zarządzania environmentem kontenerów)


Mac OS X - https://docs.docker.com/installation/mac/
Linux - https://docs.docker.com/installation/


## Instalacja Docker Compose

https://docs.docker.com/compose/install/


Przykładowa instalacja na Linuksie:

```
curl -L https://github.com/docker/compose/releases/download/1.5.0rc2/docker-compose-`uname -s`-`uname -m` > docker-compose
sudo mv docker-compose /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

## Inicjalizacja zawartości repozytoriów dockera przed szkoleniem

Uruchom budowanie kontenerów w katalogu szkolenia

```
cd ~/go/src/github.com/exu/go-workshops
make
```

(W zależności od instalacji dockera czasami trzeba użyć `sudo make`)
