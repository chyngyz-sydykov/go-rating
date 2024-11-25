#Installation

 - clone the repo
 - install docker
 - copy `.env.dist` to `.env`
 - run `docker-compose up --build`
 - if everything is ok, please check `grpcurl -plaintext -d '{"book_id": 123}' localhost:50051 rating.RatingService.GetRatings` in a termninal

#Testing

On initial project setup, please manually create a database and run following command to install uuid extension
`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`


running test `APP_ENV=test go test ./tests/`
run test without cache `go test -count=1 ./tests/`
running test within docker `docker exec -it go_rating_server bash -c "APP_ENV=test go test -count=1 ./tests"`

#GRPC

the protobuf files are stored in different repo https://github.com/chyngyz-sydykov/book-rating-protos and it is imported via following command.

generate grpc files `docker exec -it go_rating_server bash -c "./generate_protoc.sh"`

check if the service is registered `grpcurl -plaintext localhost:50051 list`. you should see the following in the console `rating.RatingService` 
Call a Specific Method
`grpcurl -plaintext -d '{"book_id": 123}' localhost:50051 rating.RatingService.GetRatings`


#Handy commands

To install new package

`go get package_name`

to clean up go.sum run

`go mod tidy`

running project via docker
`docker-compose up --build`
`docker-compose down`

`docker-compose logs -f`