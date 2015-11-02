## Instalacja

Uruchom

```
cd docker/mysql
docker-compose start
mysql -uroot -proot -P7701 -h127.0.0.1 < init.sql
mysql -uroot -proot -P7701 -h127.0.0.1
```
