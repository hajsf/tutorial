package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	pb "server/proto"
	"sync"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type server struct {
	pb.UnimplementedStreamServiceServer
}

func (s server) FetchResponse(in *pb.Request, srv pb.StreamService_FetchResponseServer) error {
	if p, ok := peer.FromContext(srv.Context()); ok {
		fmt.Println("Client ip is:", p.Addr.String())
	}
	md := metadata.New(map[string]string{"Content-Type": "text/event-stream", "Connection": "keep-alive"})
	srv.SetHeader(md)
	//use wait group to allow process to be concurrent
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(count int64) {
			defer wg.Done()
			//w.Header().Set("Content-Type", "text/event-stream")
			//time sleep to simulate server process time
			time.Sleep(time.Duration(count) * time.Second)
			resp := pb.Response{Result: fmt.Sprintf("Request #%d For Id:%d", count, in.Id)}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
			log.Printf("finishing request number : %d", count)
		}(int64(i))
	}

	wg.Wait()
	return nil
}

func main() {
	gRPcPort := ":50005"
	// Create a gRPC listener on TCP port
	lis, err := net.Listen("tcp", gRPcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterStreamServiceServer(s, server{})

	log.Println("Serving gRPC server on 0.0.0.0:50005")
	grpcTerminated := make(chan struct{})
	// Serve gRPC server
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			close(grpcTerminated) // In case server is terminated without us requesting this
		}
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	//	gateWayTarget := fmt.Sprintf("0.0.0.0%s", gRPcPort)
	conn, err := grpc.DialContext(
		context.Background(),
		gRPcPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()

	// Register custom route for  GET /hello/{name}
	err = gwmux.HandlePath("GET", "/hello/{name}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Write([]byte("hello " + pathParams["name"]))
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Register Greeter
	err = pb.RegisterStreamServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: allowCORS(gwmux),
	}

	log.Println("Serving gRPC-Gateway on http://localhost:8090")

	fmt.Println("run POST request of: http://localhost:8090/v1/example/stream?id=1 or just http://../stream (for id=0)")
	fmt.Println("or run curl -X GET -k http://localhost:8090//v1/example/stream?id=1")

	log.Fatal(gwServer.ListenAndServe()) // <- This line alone could be enough ang no need for all the lines after,
	// depending on the application complexity

	// The above line `log.Fatal(gwServer.ListenAndServe())` is enough, but as
	// the application is probably doing other things and you will want to be
	// able to shutdown cleanly; passing in a context is a good method..
	/*	ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // Ensure cancel function is called eventually

		grpcWebTerminated := make(chan struct{})
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := gwServer.ListenAndServe(); err != nil {
				fmt.Printf("Web server (GRPC) shutdown: %s", err)
			}
			close(grpcWebTerminated) // In case server is terminated without us requesting this
		}()

		// Wait for the web server to shutdown OR the context to be cancelled...
		select {
		case <-ctx.Done():
			// Shutdown the servers (there are shutdown commands to request this)
		case <-grpcTerminated:
			// You may want to exit if this happens (will be due to unexpected error)
		case <-grpcWebTerminated:
			// You may want to exit if this happens (will be due to unexpected error)
		}
		// Wait for the goRoutines to complete
		<-grpcTerminated
		<-grpcWebTerminated
	*/
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		h.ServeHTTP(w, r)
	})
}
