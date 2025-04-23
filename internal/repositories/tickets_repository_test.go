package repositories

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/DKhorkov/hmtm-tickets/api/protobuf/generated/go/tickets"
	"github.com/DKhorkov/libs/pointers"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	mockclients "github.com/DKhorkov/hmtm-notifications/mocks/clients"
)

func TestTicketsRepository_GetTicketByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketsClient := mockclients.NewMockTicketsClient(ctrl)
	repo := NewTicketsRepository(ticketsClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name           string
		id             uint64
		setupMocks     func(ticketsClient *mockclients.MockTicketsClient)
		expectedTicket *entities.RawTicket
		errorExpected  bool
	}{
		{
			name: "success",
			id:   1,
			setupMocks: func(ticketsClient *mockclients.MockTicketsClient) {
				ticketsClient.
					EXPECT().
					GetTicket(
						gomock.Any(),
						&tickets.GetTicketIn{ID: 1},
					).
					Return(&tickets.GetTicketOut{
						ID:          1,
						UserID:      1,
						CategoryID:  2,
						Name:        "Test Ticket",
						Description: "Description",
						Price:       pointers.New[float32](100),
						Quantity:    5,
						CreatedAt:   timestamppb.New(now),
						UpdatedAt:   timestamppb.New(now),
						TagIDs:      []uint32{1, 2},
						Attachments: []*tickets.Attachment{
							{
								ID:        1,
								TicketID:  1,
								Link:      "link1",
								CreatedAt: timestamppb.New(now),
								UpdatedAt: timestamppb.New(now),
							},
						},
					}, nil).
					Times(1)
			},
			expectedTicket: &entities.RawTicket{
				ID:          1,
				UserID:      1,
				CategoryID:  2,
				Name:        "Test Ticket",
				Description: "Description",
				Price:       pointers.New[float32](100),
				Quantity:    5,
				CreatedAt:   now,
				UpdatedAt:   now,
				TagIDs:      []uint32{1, 2},
				Attachments: []entities.TicketAttachment{
					{ID: 1, TicketID: 1, Link: "link1", CreatedAt: now, UpdatedAt: now},
				},
			},
			errorExpected: false,
		},
		{
			name: "error",
			id:   1,
			setupMocks: func(ticketsClient *mockclients.MockTicketsClient) {
				ticketsClient.
					EXPECT().
					GetTicket(
						gomock.Any(),
						&tickets.GetTicketIn{ID: 1},
					).
					Return(nil, errors.New("ticket not found")).
					Times(1)
			},
			expectedTicket: nil,
			errorExpected:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ticketsClient)
			}

			ticket, err := repo.GetTicketByID(context.Background(), tc.id)
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

func TestTicketsRepository_GetAllTickets(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketsClient := mockclients.NewMockTicketsClient(ctrl)
	repo := NewTicketsRepository(ticketsClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name            string
		setupMocks      func(ticketsClient *mockclients.MockTicketsClient)
		expectedTickets []entities.RawTicket
		errorExpected   bool
	}{
		{
			name: "success",
			setupMocks: func(ticketsClient *mockclients.MockTicketsClient) {
				ticketsClient.
					EXPECT().
					GetTickets(
						gomock.Any(),
						&emptypb.Empty{},
					).
					Return(&tickets.GetTicketsOut{
						Tickets: []*tickets.GetTicketOut{
							{
								ID:          1,
								UserID:      1,
								CategoryID:  2,
								Name:        "Ticket1",
								Description: "Desc1",
								Price:       pointers.New[float32](100),
								Quantity:    5,
								CreatedAt:   timestamppb.New(now),
								UpdatedAt:   timestamppb.New(now),
							},
							{
								ID:          2,
								UserID:      2,
								CategoryID:  3,
								Name:        "Ticket2",
								Description: "Desc2",
								Price:       pointers.New[float32](100),
								Quantity:    10,
								CreatedAt:   timestamppb.New(now),
								UpdatedAt:   timestamppb.New(now),
							},
						},
					}, nil).
					Times(1)
			},
			expectedTickets: []entities.RawTicket{
				{
					ID:          1,
					UserID:      1,
					CategoryID:  2,
					Name:        "Ticket1",
					Description: "Desc1",
					Price:       pointers.New[float32](100),
					Quantity:    5,
					CreatedAt:   now,
					UpdatedAt:   now,
					Attachments: make([]entities.TicketAttachment, 0),
				},
				{
					ID:          2,
					UserID:      2,
					CategoryID:  3,
					Name:        "Ticket2",
					Description: "Desc2",
					Price:       pointers.New[float32](100),
					Quantity:    10,
					CreatedAt:   now,
					UpdatedAt:   now,
					Attachments: make([]entities.TicketAttachment, 0),
				},
			},
			errorExpected: false,
		},
		{
			name: "error",
			setupMocks: func(ticketsClient *mockclients.MockTicketsClient) {
				ticketsClient.
					EXPECT().
					GetTickets(
						gomock.Any(),
						&emptypb.Empty{},
					).
					Return(nil, errors.New("fetch failed")).
					Times(1)
			},
			expectedTickets: nil,
			errorExpected:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ticketsClient)
			}

			ticketsList, err := repo.GetAllTickets(context.Background())
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, ticketsList)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedTickets, ticketsList)
			}
		})
	}
}

func TestTicketsRepository_GetUserTickets(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketsClient := mockclients.NewMockTicketsClient(ctrl)
	repo := NewTicketsRepository(ticketsClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name            string
		userID          uint64
		setupMocks      func(ticketsClient *mockclients.MockTicketsClient)
		expectedTickets []entities.RawTicket
		errorExpected   bool
	}{
		{
			name:   "success",
			userID: 1,
			setupMocks: func(ticketsClient *mockclients.MockTicketsClient) {
				ticketsClient.
					EXPECT().
					GetUserTickets(
						gomock.Any(),
						&tickets.GetUserTicketsIn{UserID: 1},
					).
					Return(
						&tickets.GetTicketsOut{
							Tickets: []*tickets.GetTicketOut{
								{
									ID:          1,
									UserID:      1,
									CategoryID:  2,
									Name:        "Ticket1",
									Description: "Desc1",
									Price:       pointers.New[float32](100),
									Quantity:    5,
									CreatedAt:   timestamppb.New(now),
									UpdatedAt:   timestamppb.New(now),
								},
							},
						},
						nil,
					).
					Times(1)
			},
			expectedTickets: []entities.RawTicket{
				{
					ID:          1,
					UserID:      1,
					CategoryID:  2,
					Name:        "Ticket1",
					Description: "Desc1",
					Price:       pointers.New[float32](100),
					Quantity:    5,
					CreatedAt:   now,
					UpdatedAt:   now,
					Attachments: make([]entities.TicketAttachment, 0),
				},
			},
			errorExpected: false,
		},
		{
			name:   "error",
			userID: 1,
			setupMocks: func(ticketsClient *mockclients.MockTicketsClient) {
				ticketsClient.
					EXPECT().
					GetUserTickets(
						gomock.Any(),
						&tickets.GetUserTicketsIn{UserID: 1},
					).
					Return(nil, errors.New("fetch failed")).
					Times(1)
			},
			expectedTickets: nil,
			errorExpected:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ticketsClient)
			}

			ticketsList, err := repo.GetUserTickets(context.Background(), tc.userID)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, ticketsList)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedTickets, ticketsList)
			}
		})
	}
}

func TestTicketsRepository_GetRespondByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketsClient := mockclients.NewMockTicketsClient(ctrl)
	repo := NewTicketsRepository(ticketsClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name            string
		id              uint64
		setupMocks      func(ticketsClient *mockclients.MockTicketsClient)
		expectedRespond *entities.Respond
		errorExpected   bool
	}{
		{
			name: "success",
			id:   1,
			setupMocks: func(ticketsClient *mockclients.MockTicketsClient) {
				ticketsClient.
					EXPECT().
					GetRespond(
						gomock.Any(),
						&tickets.GetRespondIn{ID: 1},
					).
					Return(&tickets.GetRespondOut{
						ID:        1,
						MasterID:  2,
						TicketID:  1,
						Price:     200,
						Comment:   pointers.New("Test Comment"),
						CreatedAt: timestamppb.New(now),
						UpdatedAt: timestamppb.New(now),
					}, nil).
					Times(1)
			},
			expectedRespond: &entities.Respond{
				ID:        1,
				MasterID:  2,
				TicketID:  1,
				Price:     200,
				Comment:   pointers.New("Test Comment"),
				CreatedAt: now,
				UpdatedAt: now,
			},
			errorExpected: false,
		},
		{
			name: "error",
			id:   1,
			setupMocks: func(ticketsClient *mockclients.MockTicketsClient) {
				ticketsClient.
					EXPECT().
					GetRespond(
						gomock.Any(),
						&tickets.GetRespondIn{ID: 1},
					).
					Return(nil, errors.New("respond not found")).
					Times(1)
			},
			expectedRespond: nil,
			errorExpected:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ticketsClient)
			}

			respond, err := repo.GetRespondByID(context.Background(), tc.id)
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

func TestTicketsRepository_GetTicketResponds(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketsClient := mockclients.NewMockTicketsClient(ctrl)
	repo := NewTicketsRepository(ticketsClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name             string
		ticketID         uint64
		setupMocks       func(ticketsClient *mockclients.MockTicketsClient)
		expectedResponds []entities.Respond
		errorExpected    bool
	}{
		{
			name:     "success",
			ticketID: 1,
			setupMocks: func(ticketsClient *mockclients.MockTicketsClient) {
				ticketsClient.
					EXPECT().
					GetTicketResponds(
						gomock.Any(),
						&tickets.GetTicketRespondsIn{TicketID: 1},
					).
					Return(&tickets.GetRespondsOut{
						Responds: []*tickets.GetRespondOut{
							{
								ID:        1,
								MasterID:  2,
								TicketID:  1,
								Price:     200,
								Comment:   pointers.New("Test Comment"),
								CreatedAt: timestamppb.New(now),
								UpdatedAt: timestamppb.New(now),
							},
						},
					},
						nil,
					).
					Times(1)
			},
			expectedResponds: []entities.Respond{
				{
					ID:        1,
					MasterID:  2,
					TicketID:  1,
					Price:     200,
					Comment:   pointers.New("Test Comment"),
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			errorExpected: false,
		},
		{
			name:     "error",
			ticketID: 1,
			setupMocks: func(ticketsClient *mockclients.MockTicketsClient) {
				ticketsClient.
					EXPECT().
					GetTicketResponds(
						gomock.Any(),
						&tickets.GetTicketRespondsIn{TicketID: 1},
					).
					Return(nil, errors.New("fetch failed")).
					Times(1)
			},
			expectedResponds: nil,
			errorExpected:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ticketsClient)
			}

			responds, err := repo.GetTicketResponds(context.Background(), tc.ticketID)
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

func TestTicketsRepository_GetUserResponds(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketsClient := mockclients.NewMockTicketsClient(ctrl)
	repo := NewTicketsRepository(ticketsClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name             string
		userID           uint64
		setupMocks       func(ticketsClient *mockclients.MockTicketsClient)
		expectedResponds []entities.Respond
		errorExpected    bool
	}{
		{
			name:   "success",
			userID: 1,
			setupMocks: func(ticketsClient *mockclients.MockTicketsClient) {
				ticketsClient.
					EXPECT().
					GetUserResponds(
						gomock.Any(),
						&tickets.GetUserRespondsIn{UserID: 1},
					).
					Return(&tickets.GetRespondsOut{
						Responds: []*tickets.GetRespondOut{
							{
								ID:        1,
								MasterID:  1,
								TicketID:  2,
								Price:     200,
								Comment:   pointers.New("Test Comment"),
								CreatedAt: timestamppb.New(now),
								UpdatedAt: timestamppb.New(now),
							},
						},
					},
						nil,
					).
					Times(1)
			},
			expectedResponds: []entities.Respond{
				{
					ID:        1,
					MasterID:  1,
					TicketID:  2,
					Price:     200,
					Comment:   pointers.New("Test Comment"),
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			errorExpected: false,
		},
		{
			name:   "error",
			userID: 1,
			setupMocks: func(ticketsClient *mockclients.MockTicketsClient) {
				ticketsClient.
					EXPECT().
					GetUserResponds(
						gomock.Any(),
						&tickets.GetUserRespondsIn{UserID: 1},
					).
					Return(nil, errors.New("fetch failed")).
					Times(1)
			},
			expectedResponds: nil,
			errorExpected:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(ticketsClient)
			}

			responds, err := repo.GetUserResponds(context.Background(), tc.userID)
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
