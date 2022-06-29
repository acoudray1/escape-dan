# Escape Dan

*back-end part of the app*

code in golang with postgreSQL database
using docker to develop and deploy

## Project architecture
the base architecture is based on this [github repository](https://github.com/golang-standards/project-layout)

## Deployment

to first build the docker:`docker-compose build`
to init the golang modules: `docker-compose run --rm app go mod init github.com/aicypxxx/escape-dan-back`
run the following command to start: `docker-compose up`
run the following command to stop: `docker-compose down`

### Air setup

run Air: `docker-compose run --rm app air init`

### Other commands

to check services running: `docker ps`