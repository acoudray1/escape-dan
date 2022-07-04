# Escape Dan

*back-end part of the app*

code in golang with postgreSQL database
using docker to develop and deploy

## Project architecture
the base architecture is based on this [github repository](https://github.com/golang-standards/project-layout)

## Docker management
### Deployment

to first build the docker:`docker-compose build`
to init the golang modules: `docker-compose run --rm app go mod init github.com/aicyp/escape-dan-back`
run the following command to start: `docker-compose up`
run the following command to stop: `docker-compose down`

### Development
tidy up the modules: `docker compose run --rm app go mod tidy`
to rebuild only changed containers: `docker-compose up -d --build`

### Other commands
to check services running: `docker ps`

start container and access its shell: `docker run --rm -it <image hash/name> /bin/sh`
go in the running container shell: `docker exec -it <image hash/name> /bin/sh`

copy file from docker container to local machine: `docker cp <id>:/go/src/github.com/aicyp/escape-dan-back/<file_name> Y:\Documents\Code\escape-dan-back\`

-----

https://blog.logrocket.com/how-to-build-a-restful-api-with-docker-postgresql-and-go-chi/