# Build BINARY stage
FROM golang:1.13-alpine AS build-go

RUN apk add --no-cache --update git

ENV PROJ=/go/src/example-service

COPY ./go.mod $PROJ/
COPY ./go.sum $PROJ/
COPY ./cmd $PROJ/cmd
COPY ./pkg $PROJ/pkg

RUN cd $PROJ/cmd && GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o example-service && mv example-service /tmp/


# Build IMAGE stage
FROM alpine:3.10

COPY --from=build-go /tmp/example-service /data/server/

WORKDIR /data/server/

EXPOSE 8081

CMD ["./example-service"]
