package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	mocklogging "github.com/DKhorkov/libs/logging/mocks"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	mockrepositories "github.com/DKhorkov/hmtm-notifications/mocks/repositories"
)

func TestTicketsService_GetTicketByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketsRepository := mockrepositories.NewMockTicketsRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewTicketsService(ticketsRepository, logger)

	price := float32(100.0)
	now := time.Now()
	testCases := []struct {
		name           string
		id             uint64
		setupMocks     func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger)
		expectedTicket *entities.RawTicket
		errorExpected  bool
	}{
		{
			name: "success",
			id:   1,
			setupMocks: func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger) {
				ticketsRepository.
					EXPECT().
					GetTicketByID(gomock.Any(), uint64(1)).
					Return(&entities.RawTicket{
						ID:          1,
						UserID:      1,
						CategoryID:  1,
						Name:        "Test Ticket",
						Description: "Test Description",
						Price:       &price,
						Quantity:    1,
						CreatedAt:   now,
						UpdatedAt:   now,
						TagIDs:      []uint32{1, 2},
						Attachments: []entities.TicketAttachment{{ID: 1, TicketID: 1, Link: "link1"}},
					}, nil).
					Times(1)
			},
			expectedTicket: &entities.RawTicket{
				ID:          1,
				UserID:      1,
				CategoryID:  1,
				Name:        "Test Ticket",
				Description: "Test Description",
				Price:       &price,
				Quantity:    1,
				CreatedAt:   now,
				UpdatedAt:   now,
				TagIDs:      []uint32{1, 2},
				Attachments: []entities.TicketAttachment{{ID: 1, TicketID: 1, Link: "link1"}},
			},
			errorExpected: false,
		},
		{
			name: "error",
			id:   1,
			setupMocks: func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger) {
				ticketsRepository.
					EXPECT().
					GetTicketByID(gomock.Any(), uint64(1)).
					Return(nil, errors.New("not found")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedTicket: nil,
			errorExpected:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ticketsRepository, logger)
			}

			ticket, err := service.GetTicketByID(context.Background(), tc.id)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, ticket)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedTicket, ticket)
			}
		})
	}
}

func TestTicketsService_GetAllTickets(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketsRepository := mockrepositories.NewMockTicketsRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewTicketsService(ticketsRepository, logger)

	price := float32(100.0)
	now := time.Now()
	testCases := []struct {
		name            string
		setupMocks      func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger)
		expectedTickets []entities.RawTicket
		errorExpected   bool
	}{
		{
			name: "success",
			setupMocks: func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger) {
				ticketsRepository.
					EXPECT().
					GetAllTickets(gomock.Any()).
					Return([]entities.RawTicket{
						{
							ID:          1,
							UserID:      1,
							CategoryID:  1,
							Name:        "Test Ticket",
							Description: "Test Description",
							Price:       &price,
							Quantity:    1,
							CreatedAt:   now,
							UpdatedAt:   now,
							TagIDs:      []uint32{1, 2},
							Attachments: []entities.TicketAttachment{{ID: 1, TicketID: 1, Link: "link1"}},
						},
					}, nil).
					Times(1)
			},
			expectedTickets: []entities.RawTicket{
				{
					ID:          1,
					UserID:      1,
					CategoryID:  1,
					Name:        "Test Ticket",
					Description: "Test Description",
					Price:       &price,
					Quantity:    1,
					CreatedAt:   now,
					UpdatedAt:   now,
					TagIDs:      []uint32{1, 2},
					Attachments: []entities.TicketAttachment{{ID: 1, TicketID: 1, Link: "link1"}},
				},
			},
			errorExpected: false,
		},
		{
			name: "error",
			setupMocks: func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger) {
				ticketsRepository.
					EXPECT().
					GetAllTickets(gomock.Any()).
					Return(nil, errors.New("fetch failed")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedTickets: nil,
			errorExpected:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ticketsRepository, logger)
			}

			tickets, err := service.GetAllTickets(context.Background())
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, tickets)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedTickets, tickets)
			}
		})
	}
}

func TestTicketsService_GetUserTickets(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketsRepository := mockrepositories.NewMockTicketsRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewTicketsService(ticketsRepository, logger)

	price := float32(100.0)
	now := time.Now()
	testCases := []struct {
		name            string
		userID          uint64
		setupMocks      func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger)
		expectedTickets []entities.RawTicket
		errorExpected   bool
	}{
		{
			name:   "success",
			userID: 1,
			setupMocks: func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger) {
				ticketsRepository.
					EXPECT().
					GetUserTickets(gomock.Any(), uint64(1)).
					Return([]entities.RawTicket{
						{
							ID:          1,
							UserID:      1,
							CategoryID:  1,
							Name:        "Test Ticket",
							Description: "Test Description",
							Price:       &price,
							Quantity:    1,
							CreatedAt:   now,
							UpdatedAt:   now,
							TagIDs:      []uint32{1, 2},
							Attachments: []entities.TicketAttachment{{ID: 1, TicketID: 1, Link: "link1"}},
						},
					}, nil).
					Times(1)
			},
			expectedTickets: []entities.RawTicket{
				{
					ID:          1,
					UserID:      1,
					CategoryID:  1,
					Name:        "Test Ticket",
					Description: "Test Description",
					Price:       &price,
					Quantity:    1,
					CreatedAt:   now,
					UpdatedAt:   now,
					TagIDs:      []uint32{1, 2},
					Attachments: []entities.TicketAttachment{{ID: 1, TicketID: 1, Link: "link1"}},
				},
			},
			errorExpected: false,
		},
		{
			name:   "error",
			userID: 1,
			setupMocks: func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger) {
				ticketsRepository.
					EXPECT().
					GetUserTickets(gomock.Any(), uint64(1)).
					Return(nil, errors.New("fetch failed")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedTickets: nil,
			errorExpected:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ticketsRepository, logger)
			}

			tickets, err := service.GetUserTickets(context.Background(), tc.userID)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, tickets)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedTickets, tickets)
			}
		})
	}
}

func TestTicketsService_GetRespondByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketsRepository := mockrepositories.NewMockTicketsRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewTicketsService(ticketsRepository, logger)

	price := float32(50.0)
	comment := "Test Comment"
	now := time.Now()
	testCases := []struct {
		name            string
		id              uint64
		setupMocks      func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger)
		expectedRespond *entities.Respond
		errorExpected   bool
	}{
		{
			name: "success",
			id:   1,
			setupMocks: func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger) {
				ticketsRepository.
					EXPECT().
					GetRespondByID(gomock.Any(), uint64(1)).
					Return(&entities.Respond{
						ID:        1,
						TicketID:  1,
						MasterID:  2,
						Price:     price,
						Comment:   &comment,
						CreatedAt: now,
						UpdatedAt: now,
					}, nil).
					Times(1)
			},
			expectedRespond: &entities.Respond{
				ID:        1,
				TicketID:  1,
				MasterID:  2,
				Price:     price,
				Comment:   &comment,
				CreatedAt: now,
				UpdatedAt: now,
			},
			errorExpected: false,
		},
		{
			name: "error",
			id:   1,
			setupMocks: func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger) {
				ticketsRepository.
					EXPECT().
					GetRespondByID(gomock.Any(), uint64(1)).
					Return(nil, errors.New("not found")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedRespond: nil,
			errorExpected:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ticketsRepository, logger)
			}

			respond, err := service.GetRespondByID(context.Background(), tc.id)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, respond)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedRespond, respond)
			}
		})
	}
}

func TestTicketsService_GetTicketResponds(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketsRepository := mockrepositories.NewMockTicketsRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewTicketsService(ticketsRepository, logger)

	price := float32(50.0)
	comment := "Test Comment"
	now := time.Now()
	testCases := []struct {
		name             string
		ticketID         uint64
		setupMocks       func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger)
		expectedResponds []entities.Respond
		errorExpected    bool
	}{
		{
			name:     "success",
			ticketID: 1,
			setupMocks: func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger) {
				ticketsRepository.
					EXPECT().
					GetTicketResponds(gomock.Any(), uint64(1)).
					Return([]entities.Respond{
						{
							ID:        1,
							TicketID:  1,
							MasterID:  2,
							Price:     price,
							Comment:   &comment,
							CreatedAt: now,
							UpdatedAt: now,
						},
					}, nil).
					Times(1)
			},
			expectedResponds: []entities.Respond{
				{
					ID:        1,
					TicketID:  1,
					MasterID:  2,
					Price:     price,
					Comment:   &comment,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			errorExpected: false,
		},
		{
			name:     "error",
			ticketID: 1,
			setupMocks: func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger) {
				ticketsRepository.
					EXPECT().
					GetTicketResponds(gomock.Any(), uint64(1)).
					Return(nil, errors.New("fetch failed")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedResponds: nil,
			errorExpected:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ticketsRepository, logger)
			}

			responds, err := service.GetTicketResponds(context.Background(), tc.ticketID)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, responds)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedResponds, responds)
			}
		})
	}
}

func TestTicketsService_GetUserResponds(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketsRepository := mockrepositories.NewMockTicketsRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewTicketsService(ticketsRepository, logger)

	price := float32(50.0)
	comment := "Test Comment"
	now := time.Now()
	testCases := []struct {
		name             string
		userID           uint64
		setupMocks       func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger)
		expectedResponds []entities.Respond
		errorExpected    bool
	}{
		{
			name:   "success",
			userID: 2,
			setupMocks: func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger) {
				ticketsRepository.
					EXPECT().
					GetUserResponds(gomock.Any(), uint64(2)).
					Return([]entities.Respond{
						{
							ID:        1,
							TicketID:  1,
							MasterID:  2,
							Price:     price,
							Comment:   &comment,
							CreatedAt: now,
							UpdatedAt: now,
						},
					}, nil).
					Times(1)
			},
			expectedResponds: []entities.Respond{
				{
					ID:        1,
					TicketID:  1,
					MasterID:  2,
					Price:     price,
					Comment:   &comment,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			errorExpected: false,
		},
		{
			name:   "error",
			userID: 2,
			setupMocks: func(ticketsRepository *mockrepositories.MockTicketsRepository, logger *mocklogging.MockLogger) {
				ticketsRepository.
					EXPECT().
					GetUserResponds(gomock.Any(), uint64(2)).
					Return(nil, errors.New("fetch failed")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedResponds: nil,
			errorExpected:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ticketsRepository, logger)
			}

			responds, err := service.GetUserResponds(context.Background(), tc.userID)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, responds)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedResponds, responds)
			}
		})
	}
}
