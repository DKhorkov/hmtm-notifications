package emails

import (
	"context"
	"fmt"

	"github.com/DKhorkov/libs/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	customgrpc "github.com/DKhorkov/libs/grpc"

	"github.com/DKhorkov/hmtm-notifications/api/protobuf/generated/go/notifications"
	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
)

// RegisterServer handler (serverAPI) connects EmailsServer to gRPC server:.
func RegisterServer(gRPCServer *grpc.Server, useCases interfaces.UseCases, logger logging.Logger) {
	notifications.RegisterEmailsServiceServer(
		gRPCServer,
		&ServerAPI{useCases: useCases, logger: logger},
	)
}

type ServerAPI struct {
	// Helps to test single endpoints, if others is not implemented yet
	notifications.UnimplementedEmailsServiceServer
	useCases interfaces.UseCases
	logger   logging.Logger
}

func (api ServerAPI) CountUserEmailCommunications(
	ctx context.Context,
	in *notifications.CountUserEmailCommunicationsIn,
) (*notifications.CountOut, error) {
	count, err := api.useCases.CountUserEmailCommunications(ctx, in.GetUserID())
	if err != nil {
		logging.LogErrorContext(
			ctx,
			api.logger,
			fmt.Sprintf(
				"Error occurred while trying to count Email Communications for User with ID=%d",
				in.GetUserID(),
			),
			err,
		)

		return nil, &customgrpc.BaseError{Status: codes.Internal, Message: err.Error()}
	}

	return &notifications.CountOut{Count: count}, nil
}

func (api ServerAPI) GetUserEmailCommunications(
	ctx context.Context,
	in *notifications.GetUserEmailCommunicationsIn,
) (*notifications.GetUserEmailCommunicationsOut, error) {
	var pagination *entities.Pagination
	if in.GetPagination() != nil {
		pagination = &entities.Pagination{
			Limit:  in.GetPagination().Limit,
			Offset: in.GetPagination().Offset,
		}
	}

	emailCommunications, err := api.useCases.GetUserEmailCommunications(ctx, in.GetUserID(), pagination)
	if err != nil {
		logging.LogErrorContext(
			ctx,
			api.logger,
			fmt.Sprintf(
				"Error occurred while trying to get Email Communications for User with ID=%d",
				in.GetUserID(),
			),
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
