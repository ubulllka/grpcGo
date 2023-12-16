FROM golang:1.21

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go build -o /go/bin/app ./cmd/rschir10-product-server

EXPOSE 50051

CMD ["/go/bin/app"]