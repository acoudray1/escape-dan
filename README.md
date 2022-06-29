# Escape Dan

*back-end part of the app*

code in golang with postgreSQL database
using docker to develop and deploy

## Deployment

to first build the docker:`docker-compose build`

to init the golang modules: `docker compose run --rm app go mod init github.com/aicypxxx/escape-dan-back`

### other commands

run the following command to start: `docker-compose up`
run the following command to stop: `docker-compose down`
to check services running: `docker ps`
to get the db ip: `docker inspect <container-id> | Select-String IPAddress`