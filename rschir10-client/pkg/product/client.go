package product

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	desc "ubulllka.com/rschir10-client/pkg/product/proto"
)

const productUrl = "product:50051"

type ServiceClient struct {
	Client desc.ProductV1Client
}

func InitServiceClient() desc.ProductV1Client {
	cc, err := grpc.Dial(productUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return desc.NewProductV1Client(cc)
}