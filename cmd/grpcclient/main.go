package main

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	desc "proto-example/pkg/api/example"
)

const host = "localhost:85"

func main() {

	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Msgf("can not connect with server %v", err)
	}

	client := desc.NewExampleAPIClient(conn)
	res, err := client.Health(context.Background(), &desc.HealthRequest{})
	if err != nil {
		log.Fatal().Msgf("can not connect with server %v", err)
	}

	fmt.Println(res)
}
