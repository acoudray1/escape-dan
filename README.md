# Escape Dan

*back-end part of the app*

code in golang with postgreSQL database
using docker to develop and deploy

## Project architecture
the base architecture is based on this [github repository](https://github.com/golang-standards/project-layout)

## Deployment

to first build the docker:`docker-compose build`
to init the golang modules: `docker-compose run --rm app go mod init github.com/aicyp/escape-dan-back`
run the following command to start: `docker-compose up`
run the following command to stop: `docker-compose down`

### Air setup

run Air: `docker-compose run --rm app air init`

### Other commands
tidy up the modules: `docker compose run --rm app go mod tidy`
to check services running: `docker ps`

go in the container shell: `docker run --rm -it <image hash/name> /bin/sh`

-----

https://firehydrant.com/blog/develop-a-go-app-with-docker-compose/
https://medium.com/easyread/today-i-learned-golang-live-reload-for-development-using-docker-compose-air-ecc688ee076
https://golangdocs.com/golang-docker
