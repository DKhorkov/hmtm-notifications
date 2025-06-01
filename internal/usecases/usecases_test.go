package usecases

import (
	"context"
	"errors"
	"github.com/DKhorkov/libs/pointers"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/DKhorkov/hmtm-notifications/dto"
	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	"github.com/DKhorkov/hmtm-notifications/internal/interfaces"
	mockcontentbuilders "github.com/DKhorkov/hmtm-notifications/mocks/contentbuilders"
	mocksenders "github.com/DKhorkov/hmtm-notifications/mocks/senders"
	mockservices "github.com/DKhorkov/hmtm-notifications/mocks/services"
)

func TestUseCases_GetUserEmailCommunications(t *testing.T) {
	ctrl := gomock.NewController(t)
	emailsService := mockservices.NewMockEmailsService(ctrl)
	ssoService := mockservices.NewMockSsoService(ctrl)
	toysService := mockservices.NewMockToysService(ctrl)
	ticketsService := mockservices.NewMockTicketsService(ctrl)
	verifyEmailBuilder := mockcontentbuilders.NewMockVerifyEmailContentBuilder(ctrl)
	forgetPasswordBuilder := mockcontentbuilders.NewMockForgetPasswordContentBuilder(ctrl)
	ticketUpdatedBuilder := mockcontentbuilders.NewMockTicketUpdatedContentBuilder(ctrl)
	ticketDeletedBuilder := mockcontentbuilders.NewMockTicketDeletedContentBuilder(ctrl)
	emailSender := mocksenders.NewMockEmailSender(ctrl)

	contentBuilders := interfaces.ContentBuilders{
		VerifyEmail:    verifyEmailBuilder,
		ForgetPassword: forgetPasswordBuilder,
		TicketUpdated:  ticketUpdatedBuilder,
		TicketDeleted:  ticketDeletedBuilder,
	}
	senders := interfaces.Senders{
		Email: emailSender,
	}

	useCases := New(
		emailsService,
		ssoService,
		toysService,
		ticketsService,
		contentBuilders,
		senders,
	)

	testCases := []struct {
		name       string
		pagination *entities.Pagination
		userID     uint64
		setupMocks func(
			emailsService *mockservices.MockEmailsService,
			ssoService *mockservices.MockSsoService,
			toysService *mockservices.MockToysService,
			ticketsService *mockservices.MockTicketsService,
			verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
			forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
			ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
			ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
			emailSender *mocksenders.MockEmailSender,
		)
		expected      []entities.Email
		errorExpected bool
	}{
		{
			name:   "success",
			userID: 1,
			pagination: &entities.Pagination{
				Limit:  pointers.New[uint64](1),
				Offset: pointers.New[uint64](1),
			},
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				emailsService.
					EXPECT().
					GetUserCommunications(
						gomock.Any(),
						uint64(1),
						&entities.Pagination{
							Limit:  pointers.New[uint64](1),
							Offset: pointers.New[uint64](1),
						},
					).
					Return([]entities.Email{{ID: 1, UserID: 1}}, nil).
					Times(1)
			},
			expected:      []entities.Email{{ID: 1, UserID: 1}},
			errorExpected: false,
		},
		{
			name:   "error",
			userID: 1,
			pagination: &entities.Pagination{
				Limit:  pointers.New[uint64](1),
				Offset: pointers.New[uint64](1),
			},
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				emailsService.
					EXPECT().
					GetUserCommunications(
						gomock.Any(),
						uint64(1),
						&entities.Pagination{
							Limit:  pointers.New[uint64](1),
							Offset: pointers.New[uint64](1),
						},
					).
					Return(nil, errors.New("not found")).
					Times(1)
			},
			expected:      nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(
					emailsService,
					ssoService,
					toysService,
					ticketsService,
					verifyEmailBuilder,
					forgetPasswordBuilder,
					ticketUpdatedBuilder,
					ticketDeletedBuilder,
					emailSender,
				)
			}

			emails, err := useCases.GetUserEmailCommunications(context.Background(), tc.userID, tc.pagination)
			if tc.errorExpected {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, emails)
			}
		})
	}
}

func TestUseCases_CountUserEmailCommunications(t *testing.T) {
	ctrl := gomock.NewController(t)
	emailsService := mockservices.NewMockEmailsService(ctrl)
	ssoService := mockservices.NewMockSsoService(ctrl)
	toysService := mockservices.NewMockToysService(ctrl)
	ticketsService := mockservices.NewMockTicketsService(ctrl)
	verifyEmailBuilder := mockcontentbuilders.NewMockVerifyEmailContentBuilder(ctrl)
	forgetPasswordBuilder := mockcontentbuilders.NewMockForgetPasswordContentBuilder(ctrl)
	ticketUpdatedBuilder := mockcontentbuilders.NewMockTicketUpdatedContentBuilder(ctrl)
	ticketDeletedBuilder := mockcontentbuilders.NewMockTicketDeletedContentBuilder(ctrl)
	emailSender := mocksenders.NewMockEmailSender(ctrl)

	contentBuilders := interfaces.ContentBuilders{
		VerifyEmail:    verifyEmailBuilder,
		ForgetPassword: forgetPasswordBuilder,
		TicketUpdated:  ticketUpdatedBuilder,
		TicketDeleted:  ticketDeletedBuilder,
	}
	senders := interfaces.Senders{
		Email: emailSender,
	}

	useCases := New(
		emailsService,
		ssoService,
		toysService,
		ticketsService,
		contentBuilders,
		senders,
	)

	testCases := []struct {
		name       string
		pagination *entities.Pagination
		userID     uint64
		setupMocks func(
			emailsService *mockservices.MockEmailsService,
			ssoService *mockservices.MockSsoService,
			toysService *mockservices.MockToysService,
			ticketsService *mockservices.MockTicketsService,
			verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
			forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
			ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
			ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
			emailSender *mocksenders.MockEmailSender,
		)
		expected      uint64
		errorExpected bool
	}{
		{
			name:   "success",
			userID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				emailsService.
					EXPECT().
					CountUserCommunications(gomock.Any(), uint64(1)).
					Return(uint64(1), nil).
					Times(1)
			},
			expected:      1,
			errorExpected: false,
		},
		{
			name:   "error",
			userID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				emailsService.
					EXPECT().
					CountUserCommunications(gomock.Any(), uint64(1)).
					Return(uint64(0), errors.New("not found")).
					Times(1)
			},
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(
					emailsService,
					ssoService,
					toysService,
					ticketsService,
					verifyEmailBuilder,
					forgetPasswordBuilder,
					ticketUpdatedBuilder,
					ticketDeletedBuilder,
					emailSender,
				)
			}

			actual, err := useCases.CountUserEmailCommunications(context.Background(), tc.userID)
			if tc.errorExpected {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, actual)
			}
		})
	}
}

func TestUseCases_SendVerifyEmailCommunication(t *testing.T) {
	ctrl := gomock.NewController(t)
	emailsService := mockservices.NewMockEmailsService(ctrl)
	ssoService := mockservices.NewMockSsoService(ctrl)
	toysService := mockservices.NewMockToysService(ctrl)
	ticketsService := mockservices.NewMockTicketsService(ctrl)
	verifyEmailBuilder := mockcontentbuilders.NewMockVerifyEmailContentBuilder(ctrl)
	forgetPasswordBuilder := mockcontentbuilders.NewMockForgetPasswordContentBuilder(ctrl)
	ticketUpdatedBuilder := mockcontentbuilders.NewMockTicketUpdatedContentBuilder(ctrl)
	ticketDeletedBuilder := mockcontentbuilders.NewMockTicketDeletedContentBuilder(ctrl)
	emailSender := mocksenders.NewMockEmailSender(ctrl)

	contentBuilders := interfaces.ContentBuilders{
		VerifyEmail:    verifyEmailBuilder,
		ForgetPassword: forgetPasswordBuilder,
		TicketUpdated:  ticketUpdatedBuilder,
		TicketDeleted:  ticketDeletedBuilder,
	}
	senders := interfaces.Senders{
		Email: emailSender,
	}

	useCases := New(
		emailsService,
		ssoService,
		toysService,
		ticketsService,
		contentBuilders,
		senders,
	)

	testCases := []struct {
		name       string
		userID     uint64
		setupMocks func(
			emailsService *mockservices.MockEmailsService,
			ssoService *mockservices.MockSsoService,
			toysService *mockservices.MockToysService,
			ticketsService *mockservices.MockTicketsService,
			verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
			forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
			ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
			ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
			emailSender *mocksenders.MockEmailSender,
		)
		expected      uint64
		errorExpected bool
	}{
		{
			name:   "success",
			userID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				user := entities.User{ID: 1, Email: "test@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(1)).
					Return(&user, nil).
					Times(1)

				verifyEmailBuilder.
					EXPECT().
					Subject().
					Return("Verify Email").
					Times(1)

				verifyEmailBuilder.
					EXPECT().
					Body(user).
					Return("Verify Email Body").
					Times(2)

				emailSender.
					EXPECT().
					Send(gomock.Any(), "Verify Email", "Verify Email Body", []string{"test@example.com"}).
					Return(nil).
					Times(1)

				emailsService.
					EXPECT().
					SaveCommunication(gomock.Any(), gomock.Any()).
					Return(uint64(1), nil).
					Times(1)
			},
			expected:      1,
			errorExpected: false,
		},
		{
			name:   "user not found",
			userID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(1)).
					Return(nil, errors.New("not found")).
					Times(1)
			},
			expected:      0,
			errorExpected: true,
		},
		{
			name:   "send error",
			userID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				user := entities.User{ID: 1, Email: "test@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(1)).
					Return(&user, nil).
					Times(1)

				verifyEmailBuilder.
					EXPECT().
					Subject().
					Return("Verify Email").
					Times(1)

				verifyEmailBuilder.
					EXPECT().
					Body(user).
					Return("Verify Email Body").
					Times(1)

				emailSender.
					EXPECT().
					Send(gomock.Any(), "Verify Email", "Verify Email Body", []string{"test@example.com"}).
					Return(errors.New("send failed")).
					Times(1)
			},
			expected:      0,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(
					emailsService,
					ssoService,
					toysService,
					ticketsService,
					verifyEmailBuilder,
					forgetPasswordBuilder,
					ticketUpdatedBuilder,
					ticketDeletedBuilder,
					emailSender,
				)
			}

			actual, err := useCases.SendVerifyEmailCommunication(context.Background(), tc.userID)
			if tc.errorExpected {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, actual)
			}
		})
	}
}

func TestUseCases_SendForgetPasswordEmailCommunication(t *testing.T) {
	ctrl := gomock.NewController(t)
	emailsService := mockservices.NewMockEmailsService(ctrl)
	ssoService := mockservices.NewMockSsoService(ctrl)
	toysService := mockservices.NewMockToysService(ctrl)
	ticketsService := mockservices.NewMockTicketsService(ctrl)
	verifyEmailBuilder := mockcontentbuilders.NewMockVerifyEmailContentBuilder(ctrl)
	forgetPasswordBuilder := mockcontentbuilders.NewMockForgetPasswordContentBuilder(ctrl)
	ticketUpdatedBuilder := mockcontentbuilders.NewMockTicketUpdatedContentBuilder(ctrl)
	ticketDeletedBuilder := mockcontentbuilders.NewMockTicketDeletedContentBuilder(ctrl)
	emailSender := mocksenders.NewMockEmailSender(ctrl)

	contentBuilders := interfaces.ContentBuilders{
		VerifyEmail:    verifyEmailBuilder,
		ForgetPassword: forgetPasswordBuilder,
		TicketUpdated:  ticketUpdatedBuilder,
		TicketDeleted:  ticketDeletedBuilder,
	}
	senders := interfaces.Senders{
		Email: emailSender,
	}

	useCases := New(
		emailsService,
		ssoService,
		toysService,
		ticketsService,
		contentBuilders,
		senders,
	)

	testCases := []struct {
		name       string
		userID     uint64
		setupMocks func(
			emailsService *mockservices.MockEmailsService,
			ssoService *mockservices.MockSsoService,
			toysService *mockservices.MockToysService,
			ticketsService *mockservices.MockTicketsService,
			verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
			forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
			ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
			ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
			emailSender *mocksenders.MockEmailSender,
		)
		expected      uint64
		errorExpected bool
	}{
		{
			name:   "success",
			userID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				user := entities.User{ID: 1, Email: "test@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(1)).
					Return(&user, nil).
					Times(1)

				forgetPasswordBuilder.
					EXPECT().
					Subject().
					Return("Forget Password").
					Times(1)

				forgetPasswordBuilder.
					EXPECT().
					Body(user).
					Return("Forget Password Body").
					Times(2)

				emailSender.
					EXPECT().
					Send(gomock.Any(), "Forget Password", "Forget Password Body", []string{"test@example.com"}).
					Return(nil).
					Times(1)

				emailsService.
					EXPECT().
					SaveCommunication(gomock.Any(), gomock.Any()).
					Return(uint64(1), nil).
					Times(1)
			},
			expected:      1,
			errorExpected: false,
		},
		{
			name:   "user not found",
			userID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(1)).
					Return(nil, errors.New("not found")).
					Times(1)
			},
			expected:      0,
			errorExpected: true,
		}, {
			name:   "send error",
			userID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				user := entities.User{ID: 1, Email: "test@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(1)).
					Return(&user, nil).
					Times(1)

				forgetPasswordBuilder.
					EXPECT().
					Subject().
					Return("Forget Password").
					Times(1)

				forgetPasswordBuilder.
					EXPECT().
					Body(user).
					Return("Forget Password Body").
					Times(1)

				emailSender.
					EXPECT().
					Send(gomock.Any(), "Forget Password", "Forget Password Body", []string{"test@example.com"}).
					Return(errors.New("send failed")).
					Times(1)
			},
			expected:      0,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(
					emailsService,
					ssoService,
					toysService,
					ticketsService,
					verifyEmailBuilder,
					forgetPasswordBuilder,
					ticketUpdatedBuilder,
					ticketDeletedBuilder,
					emailSender,
				)
			}

			actual, err := useCases.SendForgetPasswordEmailCommunication(context.Background(), tc.userID)
			if tc.errorExpected {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, actual)
			}
		})
	}
}

func TestUseCases_SendTicketUpdatedEmailCommunication(t *testing.T) {
	ctrl := gomock.NewController(t)
	emailsService := mockservices.NewMockEmailsService(ctrl)
	ssoService := mockservices.NewMockSsoService(ctrl)
	toysService := mockservices.NewMockToysService(ctrl)
	ticketsService := mockservices.NewMockTicketsService(ctrl)
	verifyEmailBuilder := mockcontentbuilders.NewMockVerifyEmailContentBuilder(ctrl)
	forgetPasswordBuilder := mockcontentbuilders.NewMockForgetPasswordContentBuilder(ctrl)
	ticketUpdatedBuilder := mockcontentbuilders.NewMockTicketUpdatedContentBuilder(ctrl)
	ticketDeletedBuilder := mockcontentbuilders.NewMockTicketDeletedContentBuilder(ctrl)
	emailSender := mocksenders.NewMockEmailSender(ctrl)

	contentBuilders := interfaces.ContentBuilders{
		VerifyEmail:    verifyEmailBuilder,
		ForgetPassword: forgetPasswordBuilder,
		TicketUpdated:  ticketUpdatedBuilder,
		TicketDeleted:  ticketDeletedBuilder,
	}
	senders := interfaces.Senders{
		Email: emailSender,
	}

	useCases := New(
		emailsService,
		ssoService,
		toysService,
		ticketsService,
		contentBuilders,
		senders,
	)

	testCases := []struct {
		name       string
		ticketID   uint64
		setupMocks func(
			emailsService *mockservices.MockEmailsService,
			ssoService *mockservices.MockSsoService,
			toysService *mockservices.MockToysService,
			ticketsService *mockservices.MockTicketsService,
			verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
			forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
			ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
			ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
			emailSender *mocksenders.MockEmailSender,
		)
		expected      []uint64
		errorExpected bool
	}{
		{
			name:     "success",
			ticketID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				ticket := entities.RawTicket{ID: 1}
				ticketsService.
					EXPECT().
					GetTicketByID(gomock.Any(), uint64(1)).
					Return(&ticket, nil).
					Times(1)

				respond := entities.Respond{MasterID: 2}
				ticketsService.
					EXPECT().
					GetTicketResponds(gomock.Any(), uint64(1)).
					Return([]entities.Respond{respond}, nil).
					Times(1)

				master := entities.Master{ID: 2, UserID: 3}
				toysService.
					EXPECT().
					GetMasterByID(gomock.Any(), uint64(2)).
					Return(&master, nil).
					Times(1)

				user := entities.User{ID: 3, Email: "master@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(3)).
					Return(&user, nil).
					Times(1)

				ticketUpdatedBuilder.
					EXPECT().
					Subject(ticket).
					Return("Update Ticket").
					Times(1)

				ticketUpdatedBuilder.
					EXPECT().
					Body(ticket, user).
					Return("Update Ticket Body").
					Times(2)

				emailSender.
					EXPECT().
					Send(gomock.Any(), "Update Ticket", "Update Ticket Body", []string{"master@example.com"}).
					Return(nil).
					Times(1)

				emailsService.
					EXPECT().
					SaveCommunication(gomock.Any(), gomock.Any()).
					Return(uint64(1), nil).
					Times(1)
			},
			expected:      []uint64{1},
			errorExpected: false,
		},
		{
			name:     "ticket not found",
			ticketID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				ticketsService.
					EXPECT().
					GetTicketByID(gomock.Any(), uint64(1)).
					Return(nil, errors.New("not found")).
					Times(1)
			},
			expected:      nil,
			errorExpected: true,
		}, {
			name:     "responds fetch error",
			ticketID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				ticket := entities.RawTicket{ID: 1}
				ticketsService.
					EXPECT().
					GetTicketByID(gomock.Any(), uint64(1)).
					Return(&ticket, nil).
					Times(1)

				ticketsService.
					EXPECT().
					GetTicketResponds(gomock.Any(), uint64(1)).
					Return(nil, errors.New("fetch failed")).
					Times(1)
			},
			expected:      nil,
			errorExpected: true,
		},
		{
			name:     "master not found",
			ticketID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				ticket := entities.RawTicket{ID: 1}
				ticketsService.
					EXPECT().
					GetTicketByID(gomock.Any(), uint64(1)).
					Return(&ticket, nil).
					Times(1)

				respond := entities.Respond{MasterID: 2}
				ticketsService.
					EXPECT().
					GetTicketResponds(gomock.Any(), uint64(1)).
					Return([]entities.Respond{respond}, nil).
					Times(1)

				toysService.
					EXPECT().
					GetMasterByID(gomock.Any(), uint64(2)).
					Return(nil, errors.New("not found")).
					Times(1)
			},
			expected:      nil,
			errorExpected: true,
		},
		{
			name:     "user not found",
			ticketID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				ticket := entities.RawTicket{ID: 1}
				ticketsService.
					EXPECT().
					GetTicketByID(gomock.Any(), uint64(1)).
					Return(&ticket, nil).
					Times(1)

				respond := entities.Respond{MasterID: 2}
				ticketsService.
					EXPECT().
					GetTicketResponds(gomock.Any(), uint64(1)).
					Return([]entities.Respond{respond}, nil).
					Times(1)

				master := entities.Master{ID: 2, UserID: 3}
				toysService.
					EXPECT().
					GetMasterByID(gomock.Any(), uint64(2)).
					Return(&master, nil).
					Times(1)

				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(3)).
					Return(nil, errors.New("not found")).
					Times(1)
			},
			expected:      nil,
			errorExpected: true,
		},
		{
			name:     "send error",
			ticketID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				ticket := entities.RawTicket{ID: 1}
				ticketsService.
					EXPECT().
					GetTicketByID(gomock.Any(), uint64(1)).
					Return(&ticket, nil).
					Times(1)

				respond := entities.Respond{MasterID: 2}
				ticketsService.
					EXPECT().
					GetTicketResponds(gomock.Any(), uint64(1)).
					Return([]entities.Respond{respond}, nil).
					Times(1)

				master := entities.Master{ID: 2, UserID: 3}
				toysService.
					EXPECT().
					GetMasterByID(gomock.Any(), uint64(2)).
					Return(&master, nil).
					Times(1)

				user := entities.User{ID: 3, Email: "master@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(3)).
					Return(&user, nil).
					Times(1)

				ticketUpdatedBuilder.
					EXPECT().
					Subject(ticket).
					Return("Update Ticket").
					Times(1)

				ticketUpdatedBuilder.
					EXPECT().
					Body(ticket, user).
					Return("Update Ticket Body").
					Times(1)

				emailSender.
					EXPECT().
					Send(gomock.Any(), "Update Ticket", "Update Ticket Body", []string{"master@example.com"}).
					Return(errors.New("send failed")).
					Times(1)
			},
			expected:      nil,
			errorExpected: true,
		},
		{
			name:     "save communication error",
			ticketID: 1,
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				ticket := entities.RawTicket{ID: 1}
				ticketsService.
					EXPECT().
					GetTicketByID(gomock.Any(), uint64(1)).
					Return(&ticket, nil).
					Times(1)

				respond := entities.Respond{MasterID: 2}
				ticketsService.
					EXPECT().
					GetTicketResponds(gomock.Any(), uint64(1)).
					Return([]entities.Respond{respond}, nil).
					Times(1)

				master := entities.Master{ID: 2, UserID: 3}
				toysService.
					EXPECT().
					GetMasterByID(gomock.Any(), uint64(2)).
					Return(&master, nil).
					Times(1)

				user := entities.User{ID: 3, Email: "master@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(3)).
					Return(&user, nil).
					Times(1)

				ticketUpdatedBuilder.
					EXPECT().
					Subject(ticket).
					Return("Update Ticket").
					Times(1)

				ticketUpdatedBuilder.
					EXPECT().
					Body(ticket, user).
					Return("Update Ticket Body").
					Times(2)

				emailSender.
					EXPECT().
					Send(gomock.Any(), "Update Ticket", "Update Ticket Body", []string{"master@example.com"}).
					Return(nil).
					Times(1)

				emailsService.
					EXPECT().
					SaveCommunication(gomock.Any(), gomock.Any()).
					Return(uint64(0), errors.New("save failed")).
					Times(1)
			},
			expected:      nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(
					emailsService,
					ssoService,
					toysService,
					ticketsService,
					verifyEmailBuilder,
					forgetPasswordBuilder,
					ticketUpdatedBuilder,
					ticketDeletedBuilder,
					emailSender,
				)
			}

			actual, err := useCases.SendTicketUpdatedEmailCommunication(context.Background(), tc.ticketID)
			if tc.errorExpected {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, actual)
			}
		})
	}
}

func TestUseCases_SendTicketDeletedEmailCommunication(t *testing.T) {
	ctrl := gomock.NewController(t)
	emailsService := mockservices.NewMockEmailsService(ctrl)
	ssoService := mockservices.NewMockSsoService(ctrl)
	toysService := mockservices.NewMockToysService(ctrl)
	ticketsService := mockservices.NewMockTicketsService(ctrl)
	verifyEmailBuilder := mockcontentbuilders.NewMockVerifyEmailContentBuilder(ctrl)
	forgetPasswordBuilder := mockcontentbuilders.NewMockForgetPasswordContentBuilder(ctrl)
	ticketUpdatedBuilder := mockcontentbuilders.NewMockTicketUpdatedContentBuilder(ctrl)
	ticketDeletedBuilder := mockcontentbuilders.NewMockTicketDeletedContentBuilder(ctrl)
	emailSender := mocksenders.NewMockEmailSender(ctrl)

	contentBuilders := interfaces.ContentBuilders{
		VerifyEmail:    verifyEmailBuilder,
		ForgetPassword: forgetPasswordBuilder,
		TicketUpdated:  ticketUpdatedBuilder,
		TicketDeleted:  ticketDeletedBuilder,
	}
	senders := interfaces.Senders{
		Email: emailSender,
	}

	useCases := New(
		emailsService,
		ssoService,
		toysService,
		ticketsService,
		contentBuilders,
		senders,
	)

	testCases := []struct {
		name       string
		ticketData dto.TicketDeletedDTO
		setupMocks func(
			emailsService *mockservices.MockEmailsService,
			ssoService *mockservices.MockSsoService,
			toysService *mockservices.MockToysService,
			ticketsService *mockservices.MockTicketsService,
			verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
			forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
			ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
			ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
			emailSender *mocksenders.MockEmailSender,
		)
		expected      []uint64
		errorExpected bool
	}{
		{
			name: "success",
			ticketData: dto.TicketDeletedDTO{
				TicketOwnerID:       1,
				RespondedMastersIDs: []uint64{2},
			},
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				owner := entities.User{ID: 1, Email: "owner@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(1)).
					Return(&owner, nil).
					Times(1)

				master := entities.Master{ID: 2, UserID: 3}
				toysService.
					EXPECT().
					GetMasterByID(gomock.Any(), uint64(2)).
					Return(&master, nil).
					Times(1)

				respondOwner := entities.User{ID: 3, Email: "master@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(3)).
					Return(&respondOwner, nil).
					Times(1)

				ticketData := dto.TicketDeletedDTO{TicketOwnerID: 1, RespondedMastersIDs: []uint64{2}}
				ticketDeletedBuilder.
					EXPECT().
					Subject(ticketData).
					Return("Delete Ticket").
					Times(1)

				ticketDeletedBuilder.
					EXPECT().
					Body(ticketData, owner, respondOwner).
					Return("Delete Ticket Body").
					Times(2)

				emailSender.
					EXPECT().
					Send(gomock.Any(), "Delete Ticket", "Delete Ticket Body", []string{"master@example.com"}).
					Return(nil).
					Times(1)

				emailsService.
					EXPECT().
					SaveCommunication(gomock.Any(), gomock.Any()).
					Return(uint64(1), nil).
					Times(1)
			},
			expected:      []uint64{1},
			errorExpected: false,
		},
		{
			name: "owner not found",
			ticketData: dto.TicketDeletedDTO{
				TicketOwnerID: 1,
			},
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(1)).
					Return(nil, errors.New("not found")).
					Times(1)
			},
			expected:      nil,
			errorExpected: true,
		}, {
			name: "master not found",
			ticketData: dto.TicketDeletedDTO{
				TicketOwnerID:       1,
				RespondedMastersIDs: []uint64{2},
			},
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				owner := entities.User{ID: 1, Email: "owner@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(1)).
					Return(&owner, nil).
					Times(1)

				toysService.
					EXPECT().
					GetMasterByID(gomock.Any(), uint64(2)).
					Return(nil, errors.New("not found")).
					Times(1)
			},
			expected:      nil,
			errorExpected: true,
		},
		{
			name: "respond owner not found",
			ticketData: dto.TicketDeletedDTO{
				TicketOwnerID:       1,
				RespondedMastersIDs: []uint64{2},
			},
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				owner := entities.User{ID: 1, Email: "owner@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(1)).
					Return(&owner, nil).
					Times(1)

				master := entities.Master{ID: 2, UserID: 3}
				toysService.
					EXPECT().
					GetMasterByID(gomock.Any(), uint64(2)).
					Return(&master, nil).
					Times(1)

				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(3)).
					Return(nil, errors.New("not found")).
					Times(1)
			},
			expected:      nil,
			errorExpected: true,
		},
		{
			name: "send error",
			ticketData: dto.TicketDeletedDTO{
				TicketOwnerID:       1,
				RespondedMastersIDs: []uint64{2},
			},
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				owner := entities.User{ID: 1, Email: "owner@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(1)).
					Return(&owner, nil).
					Times(1)

				master := entities.Master{ID: 2, UserID: 3}
				toysService.
					EXPECT().
					GetMasterByID(gomock.Any(), uint64(2)).
					Return(&master, nil).
					Times(1)

				respondOwner := entities.User{ID: 3, Email: "master@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(3)).
					Return(&respondOwner, nil).
					Times(1)

				ticketData := dto.TicketDeletedDTO{TicketOwnerID: 1, RespondedMastersIDs: []uint64{2}}
				ticketDeletedBuilder.
					EXPECT().
					Subject(ticketData).
					Return("Delete Ticket").
					Times(1)

				ticketDeletedBuilder.
					EXPECT().
					Body(ticketData, owner, respondOwner).
					Return("Delete Ticket Body").
					Times(1)

				emailSender.
					EXPECT().
					Send(gomock.Any(), "Delete Ticket", "Delete Ticket Body", []string{"master@example.com"}).
					Return(errors.New("send failed")).
					Times(1)
			},
			expected:      nil,
			errorExpected: true,
		},
		{
			name: "save communication error",
			ticketData: dto.TicketDeletedDTO{
				TicketOwnerID:       1,
				RespondedMastersIDs: []uint64{2},
			},
			setupMocks: func(
				emailsService *mockservices.MockEmailsService,
				ssoService *mockservices.MockSsoService,
				toysService *mockservices.MockToysService,
				ticketsService *mockservices.MockTicketsService,
				verifyEmailBuilder *mockcontentbuilders.MockVerifyEmailContentBuilder,
				forgetPasswordBuilder *mockcontentbuilders.MockForgetPasswordContentBuilder,
				ticketUpdatedBuilder *mockcontentbuilders.MockTicketUpdatedContentBuilder,
				ticketDeletedBuilder *mockcontentbuilders.MockTicketDeletedContentBuilder,
				emailSender *mocksenders.MockEmailSender,
			) {
				owner := entities.User{ID: 1, Email: "owner@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(1)).
					Return(&owner, nil).
					Times(1)

				master := entities.Master{ID: 2, UserID: 3}
				toysService.
					EXPECT().
					GetMasterByID(gomock.Any(), uint64(2)).
					Return(&master, nil).
					Times(1)

				respondOwner := entities.User{ID: 3, Email: "master@example.com"}
				ssoService.
					EXPECT().
					GetUserByID(gomock.Any(), uint64(3)).
					Return(&respondOwner, nil).
					Times(1)

				ticketData := dto.TicketDeletedDTO{TicketOwnerID: 1, RespondedMastersIDs: []uint64{2}}
				ticketDeletedBuilder.
					EXPECT().
					Subject(ticketData).
					Return("Delete Ticket").
					Times(1)

				ticketDeletedBuilder.
					EXPECT().
					Body(ticketData, owner, respondOwner).
					Return("Delete Ticket Body").
					Times(2)

				emailSender.
					EXPECT().
					Send(gomock.Any(), "Delete Ticket", "Delete Ticket Body", []string{"master@example.com"}).
					Return(nil).
					Times(1)

				emailsService.
					EXPECT().
					SaveCommunication(gomock.Any(), gomock.Any()).
					Return(uint64(0), errors.New("save failed")).
					Times(1)
			},
			expected:      nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(
					emailsService,
					ssoService,
					toysService,
					ticketsService,
					verifyEmailBuilder,
					forgetPasswordBuilder,
					ticketUpdatedBuilder,
					ticketDeletedBuilder,
					emailSender,
				)
			}

			actual, err := useCases.SendTicketDeletedEmailCommunication(context.Background(), tc.ticketData)
			if tc.errorExpected {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, actual)
			}
		})
	}
}
