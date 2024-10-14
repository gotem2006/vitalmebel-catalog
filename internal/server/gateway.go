package server

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/gotem2006/vitalmebel-catalog/pkg/catalog"
)

func createGatewayServer(grpcAddr, gatewayAddr string) *http.Server {
	conn, err := grpc.NewClient(
		grpcAddr,
		// grpc.WithUnaryInterceptor(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to dial server")
	}

	mux := runtime.NewServeMux()
	if err := pb.RegisterApartmentServiceHandler(context.Background(), mux, conn); err != nil {
		log.Fatal().Err(err).Msg("Failed registration handler")
	}

	gatewayServer := &http.Server{
		Addr:    gatewayAddr,
		Handler: mux,
	}

	return gatewayServer
}
