# Lesson 1 - How to dockerize your application


## Build container

There are many of docker images which can help you with most of your problems :)
Put line below in your `Dockerfile`

    FROM golang:1.4-onbuild


Next run

    go build .

to build your docker container, docker will build and deploy app inside container and automatically set
`app` binary as main `CMD`



## Do it little better

We an use docker comopse to configure dependencies between our little architecture
(now it's only one service but there will be more in next lessons)

    app:
      build: .
      ports:
      - "8080:8080"

put code above in `docker-compose.yml` file and run

    docker-compose up

docker-compose will output that we've built our app

    ...
    ...
    Successfully built db28da3b5081
    Attaching to dockerize_app_1
    app_1 | + exec app

## What can I do now?

You've created http powered web application

it's mapped from container port 8080 to our local machine port 8080

    ❯ curl localhost:8080/something
    Hi there,something!%

We can also go to docker container ip address (ip will be probably different for each machine)

    ❯ curl 172.17.0.11:8080/something
    Hi there,something!%
