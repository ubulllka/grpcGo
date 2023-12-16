package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "ubulllka.com/rschir10-product-server/pkg/api"
	rep "ubulllka.com/rschir10-product-server/pkg/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051


type server struct {
	desc.UnimplementedProductV1Server
}


func (s *server) GetAll(ctx context.Context, in *desc.EmptyProduct) (*desc.ProductList, error) {
	r, err := rep.GetAll()
	if err != nil {
		log.Fatalf("GetAll faled, err: #{err}")
	}
	return r, nil
}

func (s *server) Get(ctx context.Context, in *desc.ProductId) (*desc.Product, error) {
	r, err := rep.Get(in)
	if err != nil {
		log.Fatalf("Get faled, err: #{err}")
	}
	return r, nil
}

func (s *server) Insert(ctx context.Context, in *desc.Product) (*desc.Product, error) {
	r, err := rep.Insert(in)
	if err != nil {
		log.Fatalf("Insert faled, err: #{err}")
	}
	return r, nil
}

func (s *server) Update(ctx context.Context, in *desc.Product) (*desc.Product, error) {
	r, err := rep.Update(in)
	if err != nil {
		log.Fatalf("Update faled, err: #{err}")
	}
	return r, nil
}

func (s *server) Remove(ctx context.Context, in *desc.ProductId) (*desc.Product, error) {
	r, err := rep.Remove(in)
	if err != nil {
		log.Fatalf("Remove faled, err: #{err}")
	}
	return r, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: #{err}")
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterProductV1Server(s, &server{})
	log.Printf("server listening at #{lis.Addr()}")

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: #{err}")
	}
}