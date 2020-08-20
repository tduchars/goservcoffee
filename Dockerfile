FROM golang:1.14

LABEL maintainer="Tony Duchars <tony.duchars@gmail.com>"

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

WORKDIR /goservcoffee

COPY go.mod ./

RUN go mod download

COPY . .

EXPOSE 9090

CMD ["go", "run", "/goservcoffee/product-api/main.go"]