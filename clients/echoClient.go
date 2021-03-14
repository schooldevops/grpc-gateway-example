package main

import (
	"context"
	"fmt"
	"log"

	"echo-grpc/proto/echopb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Echo Client Luanched.")

	cc, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	
	defer cc.Close()

	c := echopb.NewEchoServiceClient(cc)

	echo := &echopb.EchoRequest{
		Message: "kido",
	}

	result, err := c.Echo(context.Background(), echo)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	fmt.Printf("Result is: %v", result.GetResult())


}