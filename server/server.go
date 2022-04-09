package main

import (
	"context"
	"crypto/tls"
	pb "emo_ana/proto"
	"log"
	"net"

	"emo_ana/server/gateway"
	"emo_ana/server/middlewares/auth"
	"emo_ana/server/middlewares/cred"
	"emo_ana/server/middlewares/recovery"
	"emo_ana/server/middlewares/zap"
	"emo_ana/server/service"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

type EmoAnaServiceServer struct{}

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
	pb.RegisterEmoAnaServiceServer(grpcServer, &EmoAnaServiceServer{})
	log.Println(Address + " net.Listing with TLS and token...")
	httpServer := gateway.ProvideHTTP(Address, grpcServer)

	if err = httpServer.Serve(tls.NewListener(listener, httpServer.TLSConfig)); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

func (s *EmoAnaServiceServer) UserLogin(ctx context.Context, req *pb.UserRequest) (*pb.UserDetailResponse, error) {
	return service.UserLogin(ctx, req)
}

func (s *EmoAnaServiceServer) UserRegister(ctx context.Context, req *pb.UserRequest) (*pb.UserDetailResponse, error) {
	return service.UserLogin(ctx, req)
}

func (s *EmoAnaServiceServer) GetEmoana(ctx context.Context, req *pb.GetEmoanaRequest) (*pb.GetEmoanaResponse, error) {
	return nil, nil
}
