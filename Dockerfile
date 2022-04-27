FROM golang:1.18-alpine
RUN apk --no-cache add build-base git gcc

WORKDIR /app
COPY . /app

ENV GIN_MODE=release
ENV PORT=8010

CMD ls /app
RUN go build /app

EXPOSE 8010

ENTRYPOINT ["app"]
