![example workflow](https://github.com/chyngyz-sydykov/go-rating/actions/workflows/ci.yml/badge.svg)
![Go Coverage](https://github.com/chyngyz-sydykov/go-rating/wiki/coverage.svg)

# About the project

This is a one of the microservices for personal pet project as study practice. the whole system consists of 3 microservices:
 - **go-web** works as an api gateway. the endpoints include CRUD actions for book, create endpoint for saving rating. [link](https://github.com/chyngyz-sydykov/go-web)
 - **go-rating** (current project) is a microservice that saves new rating and return list of rating by book id. the communication between go-web and go-rating is done via gRPC 
- **go-recommendation** On Progress third microservice that will hold business logic related with recommendation of books depending on rating and how often the book is edited or created. the communication will be done via RabbitMQ

# Installation

 - clone the repo
 - install docker
 - copy `.env.dist` to `.env`
 - run `docker-compose up --build`
 - if everything is ok, please check `grpcurl -plaintext -d '{"book_id": 123}' localhost:50051 rating.RatingService.GetRatings` in a termninal

# Testing

On initial project setup, please manually create a database and run following command to install uuid extension
`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`


running test `APP_ENV=test go test ./tests/`
run test without cache `go test -count=1 ./tests/`
running test within docker `docker exec -it go_rating_server bash -c "APP_ENV=test go test -count=1 ./tests"`
running the test coverage on local machine `docker exec -it go_rating_server bash "scripts/coverage.sh"`
                `go tool cover -html=coverage/filtered_coverage.out`

# GRPC

the protobuf files are stored in different repo https://github.com/chyngyz-sydykov/book-rating-protos and it is imported via following command.

generate grpc files `docker exec -it go_rating_server bash -c "./generate_protoc.sh"`

check if the service is registered `grpcurl -plaintext localhost:50051 list`. you should see the following in the console `rating.RatingService` 
Call a Specific Method
`grpcurl -plaintext -d '{"book_id": 123}' localhost:50051 rating.RatingService.GetRatings`


in order to communicate with the web microservice on local machine, do following

1. create a local network `docker network inspect grpc-network`
2. after running `docker-compose up` run `docker network inspect grpc-network`
You should see a json with the list of containers ex:
```"Containers": {
            "some_hash": {
                "Name": "go_rating_postgres_db","
            },
            "some_hash": {
                "Name": "go_rest_api",
            },
            "some_hash": {
                "Name": "go_postgres_db",
            },
            "some_hash": {
                "Name": "go_rating_server",
            }
},
```


# Handy commands

To install new package

`go get package_name`

to clean up go.sum run

`go mod tidy`

running project via docker
`docker-compose up --build`
`docker-compose down`

`docker-compose logs -f`