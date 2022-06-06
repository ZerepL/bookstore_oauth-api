# Bookstore OAuth API
 
A Golang microservice that will be used as an OAuth API to centralize authentication.

## Architecture

![Arch](./miscs/oauth.png)

## Requirements

### Standalone

* [Cassandra](https://cassandra.apache.org/_/index.html)

### Running a container

* [Docker](https://docs.docker.com/get-docker/)
* [Cassandra image](https://hub.docker.com/_/cassandra)

## Before running

This app collect some data from env, bellow you can find a list of all vars and their values:

* TBD

|       Variable       |   Description   |
|:--------------------:|:---------------:|
|                      |  Database User  |
|                      |  Database Pass  |
|                      |  Database URL   |
|                      | Database Schema |

### Creating the container

* TBD

While the same folder of Dockerfile, run:

``` shell
docker build --tag bookstore_oauth-api:latest .
```

## Running


``` shell
go run *.go
```

### Running as container

* TBD

## API

* TBD

## TODO

* Create deployment for K8s
* Create Docker file
* Implement API
* Implement Cassandra integration

## Credits

This microservice is based in this [course](https://www.udemy.com/course/golang-how-to-design-and-build-rest-microservices-in-go/)