package main

import (
	"context"
	"github.com/bufbuild/connect-go"
	"log"
	"net/http"

	greetv1 "github.com/foundatn-io/kata-connect/proto/gen/greet/v1"
	"github.com/foundatn-io/kata-connect/proto/gen/greet/v1/greetv1connect"
)

func main() {
	client := getConnectClient()
	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg)
}

// using the buf-connect protocol
func getConnectClient() greetv1connect.GreetServiceClient {
	return greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
}

// using the gRPC protocol
func getGRPCClient() greetv1connect.GreetServiceClient {
	return greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		connect.WithGRPC(),
	)
}
