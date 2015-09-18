# RethinkDB

Jedną z ciekawych funkcjonalności RethinkDB jest możliwość
monitorowania zmian na kolekcji

# Dockerizing

## Uruchom rethinka jako kontener
docker run --name some-rethink -v "$PWD:/data" -d rethinkdb

## Zlinkuj go w swojej aplikacji
docker run --name some-app --link some-rethink:rdb -d application-that-uses-rdb
