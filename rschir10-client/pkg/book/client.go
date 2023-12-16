package book

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	desc "ubulllka.com/rschir10-client/pkg/book/proto"
)

const bookUrl = "book:50052"

type ServiceClient struct {
	Client desc.BookV1Client
}

func InitServiceClient() desc.BookV1Client {
	cc, err := grpc.Dial(bookUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return desc.NewBookV1Client(cc)
}