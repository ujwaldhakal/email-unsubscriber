FROM golang:latest as base

WORKDIR /app

COPY .. /app
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD go