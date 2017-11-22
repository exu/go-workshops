## Używanie tzw 3rd parties

Go posiada zintegrowany package manager (bez wersjonowania jeszcze niestety)
`go get` pozwala nam w łatwy sposób ściągnąć paczki oraz je uruchamiać.
dzięki powyższej komendzie zostaną ściągnięte wszystkie zależności związane
z packagem `main` (wrzucone niestety do globalnej sciezki z GOPATH)


## DEP

```
go dep init 
```

ściągnie automatycznie zależności  i wrzuci wszystko do katalogu `vendor`
