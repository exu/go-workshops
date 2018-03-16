default: build

build:
	cd docker/mysql && docker-compose build
	cd docker/rethinkdb && docker-compose build
	cd docker/mongo && docker-compose build
	cd docker/redis && docker-compose build

doc:
	go run generate.go && git add README.md && git commit -m "generated and updated main README.md file"
