# Podstawy

## Instalacja

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
