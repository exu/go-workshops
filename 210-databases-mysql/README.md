## Instalacja

Uruchom

```
cd docker/mysql
docker-compose start
mysql -uroot -proot -P7701 -h127.0.0.1 < init.sql
mysql -uroot -proot -P7701 -h127.0.0.1
```


### GORM całkiem nieźle rozbudowany ORM

detale na https://github.com/jinzhu/gorm


### GORP - Go Object Relational Persistence

gdy potrzebujesz czegoś lżejszego

https://github.com/go-gorp/gorp
