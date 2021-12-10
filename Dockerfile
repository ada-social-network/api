# Start by building the application.
FROM golang:1.17-bullseye as build

ARG VERSION=dev

WORKDIR /go/src/app

ADD go.mod .
ADD go.sum .

RUN go mod download

ADD . /go/src/app

RUN go build -o /go/bin/ada-api -ldflags="-X 'main.version=${VERSION}'" .

# Now copy it into our base image.
FROM debian

RUN mkdir -p /usr/local/ada/data

COPY --from=build /go/bin/ada-api /usr/local/ada/api

RUN ln -s /usr/local/ada/api /usr/local/bin/ada-api

ENTRYPOINT ["ada-api"]

CMD ["--sqlite-dsn", "/usr/local/ada/data/gorm.db"]