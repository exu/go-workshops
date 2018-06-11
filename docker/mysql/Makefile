default: up

up:
	docker-compose rm --force
	docker-compose up

db:
	mysql -uroot -proot -P7701 -h127.0.0.1 < init.sql

cli:
	mysql -uroot -proot -P7701 -h127.0.0.1 gotraining
