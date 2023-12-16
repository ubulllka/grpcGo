FROM golang:1.21

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go build -o /go/bin/app ./cmd/rschir10-book-server

EXPOSE 5005

CMD ["/go/bin/app"]