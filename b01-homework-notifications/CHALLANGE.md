# Challange to complete for Those who want MOAR!

Write apllication which push notfications to clients.

## Architecture

Use framework, libs, patterns, architecture of your choice.

### Notification receivers

- Client 1 - browser with EventSource client
- Client 2 - document database (feel free to chose whatever you want)
             take for consideration that I'll need to run it
- Client 3 - Log file
- Client 4 - REST API accepting JSON data powered POST request
             with "message" key and "app_name" key (will be your name)
             (you should write it Yourself) URL should be easily configurable

### Api backend

JSON powered REST API with http transport


### How it should work (business requirements)

We should be able to push JSON POST notification request to Our API
let it be simple JSON with key "message" e.g.  `{"message": "Our super message"}`

Our API shold pass this message to all from our clients
(use known techniqes *or better* techniques out of scope of this training)



### How to complete?

Send link to repo (on github bitbucket whatever with public access) and README file how to
run it (should run as simply as possible).

Write requirements for running your app (take for consideration that there is
not only your system architecture). How to build it? From where get binary if it's
prebuild? etc. (youre developer you know it already)



## For the best of the best (A+ task)

Should run on docker containers / kubernetes cluster / Helm (minimum 2 containers with docker-compose usage)

## Rewards

Several best solutions receive certificates, fame, prestige.
(Dziwki koks i tajski boks będzie tylko następstwem tych powyżej)

## Important notice

If I'll be not able to run your solution It'll be DISQUALIFICATED
and all your fame will go away, so be carefull and do it good.
(pro tip: use containers :)

## Deadline

9 days from receiving task on friday (2 weekends)
