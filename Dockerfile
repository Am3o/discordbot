FROM golang:1.14.4 AS build

COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -mod=vendor -a -ldflags '-w' -o /app/main

FROM alpine

COPY --from=build /app/main /usr/local/bin/app
RUN /usr/local/bin/app
