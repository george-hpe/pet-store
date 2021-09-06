package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gk-hpe/pet-store/cmd"
	"github.com/gk-hpe/pet-store/petstorepb"
	"github.com/gk-hpe/pet-store/server"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

const addr = "0.0.0.0:50051"

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	petstorepb.RegisterStoreServiceServer(s, server.New())

	// Serve gRPC Server
	fmt.Println("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	log.Fatal(cmd.Run(addr))
}
