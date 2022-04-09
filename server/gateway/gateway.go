package gateway

import (
	"context"
	"crypto/tls"
	"emo_ana/conf"
	pb "emo_ana/proto"
	"emo_ana/server/swagger"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func ProvideHTTP(endpoint string, grpcServer *grpc.Server) *http.Server {

	ctx := context.Background()

	creds, err := credentials.NewClientTLSFromFile(conf.PemPath, "emo_ana")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}

	dopts := []grpc.DialOption{grpc.WithTranssportCredentials(creds)}
	gwmux := runtime.NewServerMux()
	err = pb.RegisterEmoAnaServiceHandlerFromEndpoint(ctx, gwmux, endpoint, dopts)
	if err != nil {
		log.Fatalf("Register Endpoint err: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", swagger.ServeSwaggerFile)
	swagger.ServeSwaggerUI(mux)

	log.Println(endpoint + " HTTP.Listing with TLS and token...")
	return &http.Server{
		Addr:      endpoint,
		Handler:   grpcHandlerFunc(grpcServer, mux),
		TLSConfig: getTLSConfig(),
	}
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-type"), "application/grpc") {
			grpcServer.ServerHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func getTLSConfig() *tls.Config {
	cret, _ := ioutil.ReadFile(conf.PemPath)
	key, _ := ioutil.ReadFile(conf.KeyPath)

	var demoKeyPair *tls.Certificate
	pair, err := tls.X509KeyPair(cret, key)
	if err != nil {
		grpclog.Fatalf("TLS KeyPair err: %v\n", err)
	}
	demoKeyPair = &pair
	return &tls.Config{
		Certificates: []tls.Certificate{*demoKeyPair},
		NextProtos:   []string{http2.NextProtoTLS},
	}
}
