package main

import (
	"context"
	"log"
	"net"
	pb "user/proto"

	"user/server/middlewares/auth"
	"user/server/middlewares/cred"
	"user/server/middlewares/recovery"
	"user/server/middlewares/zap"
	"user/server/service"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

type UserServiceServer struct{}

const (
	// Address 监听地址
	Address string = ":4396"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {

	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	grpcServer := grpc.NewServer(cred.TLSInterceptor(),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
			grpc_auth.StreamServerInterceptor(auth.AuthInterceptor),
			grpc_recovery.StreamServerInterceptor(recovery.RecoveryInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
			grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
			grpc_recovery.UnaryServerInterceptor(recovery.RecoveryInterceptor()),
		)),
	)
	pb.RegisterUserServiceServer(grpcServer, &UserServiceServer{})
	log.Println(Address + " net.Listening...")

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

func (s *UserServiceServer) UserLogin(ctx context.Context, req *pb.UserRequest) (*pb.UserDetailResponse, error) {
	return service.UserLogin(ctx, req)
}

func (s *UserServiceServer) UserRegister(ctx context.Context, req *pb.UserRequest) (*pb.UserDetailResponse, error) {
	return service.UserLogin(ctx, req)
}
