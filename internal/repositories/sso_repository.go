package repositories

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/DKhorkov/hmtm-sso/api/protobuf/generated/go/sso"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
)

func NewGrpcSsoRepository(client interfaces.SsoGrpcClient) *GrpcSsoRepository {
	return &GrpcSsoRepository{client: client}
}

type GrpcSsoRepository struct {
	client interfaces.SsoGrpcClient
}

func (repo *GrpcSsoRepository) GetUserByID(ctx context.Context, id uint64) (*entities.User, error) {
	response, err := repo.client.GetUser(
		ctx,
		&sso.GetUserIn{
			ID: id,
		},
	)

	if err != nil {
		return nil, err
	}

	return repo.processUserResponse(response), nil
}

func (repo *GrpcSsoRepository) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	response, err := repo.client.GetUsers(
		ctx,
		&emptypb.Empty{},
	)

	if err != nil {
		return nil, err
	}

	users := make([]entities.User, len(response.GetUsers()))
	for i, userResponse := range response.GetUsers() {
		users[i] = *repo.processUserResponse(userResponse)
	}

	return users, nil
}

func (repo *GrpcSsoRepository) processUserResponse(userResponse *sso.GetUserOut) *entities.User {
	return &entities.User{
		ID:                userResponse.GetID(),
		DisplayName:       userResponse.GetDisplayName(),
		Email:             userResponse.GetEmail(),
		EmailConfirmed:    userResponse.GetEmailConfirmed(),
		Phone:             userResponse.Phone,
		PhoneConfirmed:    userResponse.GetPhoneConfirmed(),
		Telegram:          userResponse.Telegram,
		TelegramConfirmed: userResponse.GetTelegramConfirmed(),
		Avatar:            userResponse.Avatar,
		CreatedAt:         userResponse.GetCreatedAt().AsTime(),
		UpdatedAt:         userResponse.GetUpdatedAt().AsTime(),
	}
}
