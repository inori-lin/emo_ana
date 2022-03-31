package cred

import (
	"log"
	"user/server/conf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func TLSInterceptor() grpc.ServerOption {

	creds, err := credentials.NewServerTLSFromFile(conf.PemPath, conf.KeyPath)
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	return grpc.Creds(creds)
}
