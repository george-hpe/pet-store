package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gk-hpe/pet-store/petstorepb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// Run runs the gRPC-Gateway.
func Run(addr string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint.
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := petstorepb.RegisterStoreServiceHandlerFromEndpoint(ctx, mux, addr, opts)
	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}

	fmt.Println("Serving gRPC-Gateway on https://localhost:8081")

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}
