# build from source
FROM golang:1.21-alpine3.18 AS build-stage
RUN apk --no-cache add build-base git gcc

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY .. ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /simple-rest-service

FROM golang:1.21-alpine3.18 AS build-release-stage
RUN apk add --no-cache bash
WORKDIR /
COPY --from=build-stage /simple-rest-service /simple-rest-service
EXPOSE 8010
ENTRYPOINT ["/simple-rest-service"]
