package main

import (
	"context"
	"log"
	"user/client/auth"
	"user/client/conf"
	pb "user/client/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const Address string = "10.0.16.7:4396"

var grpcClient pb.UserServiceClient

func main() {

	creds, err := credentials.NewClientTLSFromFile(conf.PemPath, "")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	token := auth.Token{
		Value: "grpc.auth.token",
	}
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&token))

	if err != nil {
		log.Fatalf("net.Connect err:%v", err)
	}
	defer conn.Close()
	grpcClient = pb.NewUserServiceClient(conn)
	//调用方法
	userLogin()
}

func userLogin() {
	req := pb.UserRequest{
		UserName: "inori",
		Password: "lin1264.",
	}
	res, err := grpcClient.UserLogin(context.Background(), &req)
	if err != nil {
		log.Fatalf("call route err:%v", err)
	}
	log.Println(res)
}
