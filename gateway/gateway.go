package main

import (
	"context"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"echo-grpc/proto/echopb"

	"google.golang.org/grpc"
)

const (
	portNumber           = "9000"
	grpcServerPortNumber = "9001"
)

func main() {

	ctx := context.Background()
	mux := runtime.NewServeMux()

	// mux.HandlePath("/openapiv2/", "", openAPIServer("http"))
	http.HandleFunc("/openapiv2/", openAPIServer("http"))

	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	if err := echopb.RegisterEchoServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:"+grpcServerPortNumber,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	log.Printf("start HTTP on %s port", portNumber)
	if err := http.ListenAndServe(":"+portNumber, mux); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}

func openAPIServer(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
			glog.Errorf("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		glog.Infof("Serving %s", r.URL.Path)
		p := strings.TrimPrefix(r.URL.Path, "/openapiv2/")
		p = path.Join(dir, p)
		http.ServeFile(w, r, p)
	}
}
