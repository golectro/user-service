package grpc

import (
	"fmt"
	"golectro-user/internal/delivery/grpc/handler"
	"golectro-user/internal/delivery/grpc/interceptor"
	proto "golectro-user/internal/delivery/grpc/proto/address"
	"golectro-user/internal/usecase"
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func StartGRPCServer(addressUC *usecase.AddressUseCase, port int, viper *viper.Viper) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryRequestIDInterceptor(),
			interceptor.UnaryLoggingInterceptor(addressUC.Log),
			interceptor.UnaryAuthInterceptor(viper),
		),
	)

	userHandler := &handler.AddressHandler{UseCase: addressUC}
	proto.RegisterAddressServiceServer(grpcServer, userHandler)

	log.Printf("gRPC server listening at :%d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
