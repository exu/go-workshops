# Dockerizing

docker run --name some-rethink -v "$PWD:/data" -d rethinkdb

docker run --name some-app --link some-rethink:rdb -d application-that-uses-rdb
