FROM golang:latest

LABEL maintainer="Tony Duchars <tony.duchars@gmail.com>"

WORKDIR /goserv

COPY go.mod ./

RUN go mod download

COPY . .

EXPOSE 9090

CMD [ "go run ./product-api/main.go" ]