package emails

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/DKhorkov/hmtm-notifications/api/protobuf/generated/go/notifications"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
	customgrpc "github.com/DKhorkov/libs/grpc"
	"github.com/DKhorkov/libs/logging"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// RegisterServer handler (serverAPI) connects EmailsServer to gRPC server:.
func RegisterServer(gRPCServer *grpc.Server, useCases interfaces.UseCases, logger *slog.Logger) {
	notifications.RegisterEmailsServiceServer(gRPCServer, &ServerAPI{useCases: useCases, logger: logger})
}

type ServerAPI struct {
	// Helps to test single endpoints, if others is not implemented yet
	notifications.UnimplementedEmailsServiceServer
	useCases interfaces.UseCases
	logger   *slog.Logger
}

func (api ServerAPI) GetUserEmailCommunications(
	ctx context.Context,
	in *notifications.GetUserEmailCommunicationsIn,
) (*notifications.GetUserEmailCommunicationsOut, error) {
	emailCommunications, err := api.useCases.GetUserEmailCommunications(ctx, in.GetUserID())
	if err != nil {
		logging.LogErrorContext(
			ctx,
			api.logger,
			fmt.Sprintf("Error occurred while trying to get Email Communications for User with ID=%d", in.GetUserID()),
			err,
		)

		return nil, &customgrpc.BaseError{Status: codes.Internal, Message: err.Error()}
	}

	processedEmailCommunications := make([]*notifications.Email, len(emailCommunications))
	for i, communication := range emailCommunications {
		processedEmailCommunications[i] = &notifications.Email{
			ID:      communication.ID,
			UserID:  communication.UserID,
			Email:   communication.Email,
			Content: communication.Content,
			SentAt:  timestamppb.New(communication.SentAt),
		}
	}

	return &notifications.GetUserEmailCommunicationsOut{Emails: processedEmailCommunications}, nil
}
