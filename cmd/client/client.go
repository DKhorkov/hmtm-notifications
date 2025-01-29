package main

import (
	"context"
	"fmt"

	"github.com/DKhorkov/hmtm-notifications/api/protobuf/generated/go/notifications"

	"github.com/DKhorkov/libs/requestid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	notifications.EmailsServiceClient
}

func main() {
	clientConnection, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", "0.0.0.0", 8040),
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		panic(err)
	}

	client := &Client{
		EmailsServiceClient: notifications.NewEmailsServiceClient(clientConnection),
	}

	ctx := metadata.AppendToOutgoingContext(context.Background(), requestid.Key, requestid.New())

	emails, err := client.GetUserEmailCommunications(ctx, &notifications.GetUserEmailCommunicationsIn{
		UserID: 1,
	})
	fmt.Printf("Emails: %+v\nErr: %v\n", emails, err)
}
