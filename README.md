# golang-mod
golang-mod repository is simple golang rest-api application which is used for go rest-api sample backend service. <br /> <br />
### serving endpoints
the following list consists of predefined endpoints
| endpoints  | 
| -------------
| /home  |
| /feed  |
<br /><br />
### required environment variables
set the following environment variables for service runtime
| env var name  | value type |
| ------------- | ------------- |
| PORT_NUMBER  | string  |
<br /><br />

### how to build the application
**1. using makefile**<br /> <br />
we need to build the image first:
```shell
$ make build-builder && make build-service
```
set the port number for container and host to be bound and exposed.
```shell
$ export PORT_NUMBER=8080
```
run the application container.
```shell
$ make run-service
```
<br /><br />**2. docker compose**
```shell
$ docker-compose up --build
```