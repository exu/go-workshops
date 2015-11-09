default: up

up:
	docker-compose rm --force
	docker-compose up

db:
	PGPASSWORD=root psql -U root -p 7705 -h localhost < initdb.sql
	PGPASSWORD=root psql -U root -p 7705 -h localhost gotraining < init.sql

cli:
	PGPASSWORD=root psql -U root -p 7705 -h localhost gotraining
