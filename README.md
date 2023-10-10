# simple-rest-service


## Running the service
`go run .`


## checkout mockgen for mocking.
go install github.com/golang/mock/mockgen@v1.6.0
mockgen -source=services/ping_service.go -destination=mocks/ping_service.go

## Docker

### Building & Running Docker images
https://docs.docker.com/language/golang/build-images/

```
docker-compose up --build
```
