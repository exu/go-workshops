## MySQL

Install instructions to run this section:

```sh
cd docker/mysql
docker-compose up
mysql -uroot -proot -P7701 -h127.0.0.1 < init.sql
mysql -uroot -proot -P7701 -h127.0.0.1
```
