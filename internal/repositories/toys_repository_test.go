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

	"github.com/DKhorkov/hmtm-toys/api/protobuf/generated/go/toys"
	"github.com/DKhorkov/libs/pointers"

	"github.com/DKhorkov/hmtm-notifications/internal/entities"
	mockclients "github.com/DKhorkov/hmtm-notifications/mocks/clients"
)

func TestToysRepository_GetAllToys(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysClient := mockclients.NewMockToysClient(ctrl)
	repo := NewToysRepository(toysClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name          string
		setupMocks    func(toysClient *mockclients.MockToysClient)
		expectedToys  []entities.Toy
		errorExpected bool
	}{
		{
			name: "success",
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetToys(
						gomock.Any(),
						&emptypb.Empty{},
					).
					Return(&toys.GetToysOut{
						Toys: []*toys.GetToyOut{
							{
								ID:          1,
								MasterID:    1,
								CategoryID:  2,
								Name:        "Toy1",
								Description: "Desc1",
								Price:       100,
								Quantity:    5,
								CreatedAt:   timestamppb.New(now),
								UpdatedAt:   timestamppb.New(now),
							},
						},
					}, nil).
					Times(1)
			},
			expectedToys: []entities.Toy{
				{
					ID:          1,
					MasterID:    1,
					CategoryID:  2,
					Name:        "Toy1",
					Description: "Desc1",
					Price:       100,
					Quantity:    5,
					CreatedAt:   now,
					UpdatedAt:   now,
					Tags:        make([]entities.Tag, 0),
					Attachments: make([]entities.ToyAttachment, 0),
				},
			},
			errorExpected: false,
		},
		{
			name: "error",
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetToys(
						gomock.Any(),
						&emptypb.Empty{},
					).
					Return(nil, errors.New("fetch failed")).
					Times(1)
			},
			expectedToys:  nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysClient)
			}

			toysList, err := repo.GetAllToys(context.Background())
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, toysList)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedToys, toysList)
			}
		})
	}
}

func TestToysRepository_GetMasterToys(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysClient := mockclients.NewMockToysClient(ctrl)
	repo := NewToysRepository(toysClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name          string
		masterID      uint64
		setupMocks    func(toysClient *mockclients.MockToysClient)
		expectedToys  []entities.Toy
		errorExpected bool
	}{
		{
			name:     "success",
			masterID: 1,
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetMasterToys(
						gomock.Any(),
						&toys.GetMasterToysIn{MasterID: 1},
					).
					Return(&toys.GetToysOut{
						Toys: []*toys.GetToyOut{
							{
								ID:          1,
								MasterID:    1,
								CategoryID:  2,
								Name:        "Toy1",
								Description: "Desc1",
								Price:       100,
								Quantity:    5,
								CreatedAt:   timestamppb.New(now),
								UpdatedAt:   timestamppb.New(now),
							},
						},
					}, nil).
					Times(1)
			},
			expectedToys: []entities.Toy{
				{
					ID:          1,
					MasterID:    1,
					CategoryID:  2,
					Name:        "Toy1",
					Description: "Desc1",
					Price:       100,
					Quantity:    5,
					CreatedAt:   now,
					UpdatedAt:   now,
					Tags:        make([]entities.Tag, 0),
					Attachments: make([]entities.ToyAttachment, 0),
				},
			},
			errorExpected: false,
		},
		{
			name:     "error",
			masterID: 1,
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetMasterToys(
						gomock.Any(),
						&toys.GetMasterToysIn{MasterID: 1},
					).
					Return(nil, errors.New("fetch failed")).
					Times(1)
			},
			expectedToys:  nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysClient)
			}

			toysList, err := repo.GetMasterToys(context.Background(), tc.masterID)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, toysList)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedToys, toysList)
			}
		})
	}
}

func TestToysRepository_GetUserToys(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysClient := mockclients.NewMockToysClient(ctrl)
	repo := NewToysRepository(toysClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name          string
		userID        uint64
		setupMocks    func(toysClient *mockclients.MockToysClient)
		expectedToys  []entities.Toy
		errorExpected bool
	}{
		{
			name:   "success",
			userID: 1,
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetUserToys(
						gomock.Any(),
						&toys.GetUserToysIn{UserID: 1},
					).
					Return(&toys.GetToysOut{
						Toys: []*toys.GetToyOut{
							{
								ID:          1,
								MasterID:    1,
								CategoryID:  2,
								Name:        "Toy1",
								Description: "Desc1",
								Price:       100,
								Quantity:    5,
								CreatedAt:   timestamppb.New(now),
								UpdatedAt:   timestamppb.New(now),
							},
						},
					}, nil).
					Times(1)
			},
			expectedToys: []entities.Toy{
				{
					ID:          1,
					MasterID:    1,
					CategoryID:  2,
					Name:        "Toy1",
					Description: "Desc1",
					Price:       100,
					Quantity:    5,
					CreatedAt:   now,
					UpdatedAt:   now,
					Tags:        make([]entities.Tag, 0),
					Attachments: make([]entities.ToyAttachment, 0),
				},
			},
			errorExpected: false,
		},
		{
			name:   "error",
			userID: 1,
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetUserToys(
						gomock.Any(),
						&toys.GetUserToysIn{UserID: 1},
					).
					Return(nil, errors.New("fetch failed")).
					Times(1)
			},
			expectedToys:  nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysClient)
			}

			toysList, err := repo.GetUserToys(context.Background(), tc.userID)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, toysList)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedToys, toysList)
			}
		})
	}
}

func TestToysRepository_GetToyByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysClient := mockclients.NewMockToysClient(ctrl)
	repo := NewToysRepository(toysClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name          string
		id            uint64
		setupMocks    func(toysClient *mockclients.MockToysClient)
		expectedToy   *entities.Toy
		errorExpected bool
	}{
		{
			name: "success",
			id:   1,
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetToy(
						gomock.Any(),
						&toys.GetToyIn{ID: 1},
					).
					Return(&toys.GetToyOut{
						ID:          1,
						MasterID:    1,
						CategoryID:  2,
						Name:        "Test Toy",
						Description: "Desc",
						Price:       100,
						Quantity:    5,
						CreatedAt:   timestamppb.New(now),
						UpdatedAt:   timestamppb.New(now),
						Tags: []*toys.GetTagOut{
							{
								ID:        1,
								Name:      "Tag1",
								CreatedAt: timestamppb.New(now),
								UpdatedAt: timestamppb.New(now),
							},
						},
						Attachments: []*toys.Attachment{
							{
								ID:        1,
								ToyID:     1,
								Link:      "test",
								CreatedAt: timestamppb.New(now),
								UpdatedAt: timestamppb.New(now),
							},
						},
					}, nil).
					Times(1)
			},
			expectedToy: &entities.Toy{
				ID:          1,
				MasterID:    1,
				CategoryID:  2,
				Name:        "Test Toy",
				Description: "Desc",
				Price:       100,
				Quantity:    5,
				CreatedAt:   now,
				UpdatedAt:   now,
				Tags: []entities.Tag{
					{
						ID:   1,
						Name: "Tag1",
					},
				},
				Attachments: []entities.ToyAttachment{
					{
						ID:        1,
						ToyID:     1,
						Link:      "test",
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
			},
			errorExpected: false,
		},
		{
			name: "error",
			id:   1,
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetToy(
						gomock.Any(),
						&toys.GetToyIn{ID: 1},
					).
					Return(nil, errors.New("not found")).
					Times(1)
			},
			expectedToy:   nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysClient)
			}

			toy, err := repo.GetToyByID(context.Background(), tc.id)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, toy)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedToy, toy)
			}
		})
	}
}

func TestToysRepository_GetAllMasters(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysClient := mockclients.NewMockToysClient(ctrl)
	repo := NewToysRepository(toysClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name            string
		setupMocks      func(toysClient *mockclients.MockToysClient)
		expectedMasters []entities.Master
		errorExpected   bool
	}{
		{
			name: "success",
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetMasters(
						gomock.Any(),
						&emptypb.Empty{},
					).
					Return(&toys.GetMastersOut{
						Masters: []*toys.GetMasterOut{
							{
								ID:        1,
								UserID:    1,
								Info:      pointers.New("Master Info"),
								CreatedAt: timestamppb.New(now),
								UpdatedAt: timestamppb.New(now),
							},
						},
					},
						nil,
					).
					Times(1)
			},
			expectedMasters: []entities.Master{
				{
					ID:        1,
					UserID:    1,
					Info:      pointers.New("Master Info"),
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			errorExpected: false,
		},
		{
			name: "error",
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetMasters(
						gomock.Any(),
						&emptypb.Empty{},
					).
					Return(nil, errors.New("fetch failed")).
					Times(1)
			},
			expectedMasters: nil,
			errorExpected:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysClient)
			}

			masters, err := repo.GetAllMasters(context.Background())
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, masters)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedMasters, masters)
			}
		})
	}
}

func TestToysRepository_GetMasterByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysClient := mockclients.NewMockToysClient(ctrl)
	repo := NewToysRepository(toysClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name           string
		id             uint64
		setupMocks     func(toysClient *mockclients.MockToysClient)
		expectedMaster *entities.Master
		errorExpected  bool
	}{
		{
			name: "success",
			id:   1,
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetMaster(
						gomock.Any(),
						&toys.GetMasterIn{ID: 1},
					).
					Return(
						&toys.GetMasterOut{
							ID:        1,
							UserID:    1,
							Info:      pointers.New("Master Info"),
							CreatedAt: timestamppb.New(now),
							UpdatedAt: timestamppb.New(now),
						},
						nil,
					).
					Times(1)
			},
			expectedMaster: &entities.Master{
				ID:        1,
				UserID:    1,
				Info:      pointers.New("Master Info"),
				CreatedAt: now,
				UpdatedAt: now,
			},
			errorExpected: false,
		},
		{
			name: "error",
			id:   1,
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetMaster(
						gomock.Any(),
						&toys.GetMasterIn{ID: 1},
					).
					Return(nil, errors.New("not found")).
					Times(1)
			},
			expectedMaster: nil,
			errorExpected:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysClient)
			}

			master, err := repo.GetMasterByID(context.Background(), tc.id)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, master)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedMaster, master)
			}
		})
	}
}

func TestToysRepository_GetAllCategories(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysClient := mockclients.NewMockToysClient(ctrl)
	repo := NewToysRepository(toysClient)

	testCases := []struct {
		name               string
		setupMocks         func(toysClient *mockclients.MockToysClient)
		expectedCategories []entities.Category
		errorExpected      bool
	}{
		{
			name: "success",
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetCategories(
						gomock.Any(),
						&emptypb.Empty{},
					).
					Return(&toys.GetCategoriesOut{
						Categories: []*toys.GetCategoryOut{
							{ID: 1, Name: "Category1"},
						},
					}, nil).
					Times(1)
			},
			expectedCategories: []entities.Category{
				{ID: 1, Name: "Category1"},
			},
			errorExpected: false,
		},
		{
			name: "error",
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetCategories(
						gomock.Any(),
						&emptypb.Empty{},
					).
					Return(nil, errors.New("fetch failed")).
					Times(1)
			},
			expectedCategories: nil,
			errorExpected:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysClient)
			}

			categories, err := repo.GetAllCategories(context.Background())
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, categories)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedCategories, categories)
			}
		})
	}
}

func TestToysRepository_GetCategoryByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysClient := mockclients.NewMockToysClient(ctrl)
	repo := NewToysRepository(toysClient)

	testCases := []struct {
		name             string
		id               uint32
		setupMocks       func(toysClient *mockclients.MockToysClient)
		expectedCategory *entities.Category
		errorExpected    bool
	}{
		{
			name: "success",
			id:   1,
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetCategory(
						gomock.Any(),
						&toys.GetCategoryIn{ID: 1},
					).
					Return(&toys.GetCategoryOut{
						ID:   1,
						Name: "Test Category",
					}, nil).
					Times(1)
			},
			expectedCategory: &entities.Category{
				ID:   1,
				Name: "Test Category",
			},
			errorExpected: false,
		},
		{
			name: "error",
			id:   1,
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetCategory(
						gomock.Any(),
						&toys.GetCategoryIn{ID: 1},
					).
					Return(nil, errors.New("not found")).
					Times(1)
			},
			expectedCategory: nil,
			errorExpected:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysClient)
			}

			category, err := repo.GetCategoryByID(context.Background(), tc.id)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, category)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedCategory, category)
			}
		})
	}
}

func TestToysRepository_GetAllTags(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysClient := mockclients.NewMockToysClient(ctrl)
	repo := NewToysRepository(toysClient)

	testCases := []struct {
		name          string
		setupMocks    func(toysClient *mockclients.MockToysClient)
		expectedTags  []entities.Tag
		errorExpected bool
	}{
		{
			name: "success",
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetTags(
						gomock.Any(),
						&emptypb.Empty{},
					).
					Return(&toys.GetTagsOut{
						Tags: []*toys.GetTagOut{
							{ID: 1, Name: "Tag1"},
						},
					}, nil).
					Times(1)
			},
			expectedTags: []entities.Tag{
				{ID: 1, Name: "Tag1"},
			},
			errorExpected: false,
		},
		{
			name: "error",
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetTags(
						gomock.Any(),
						&emptypb.Empty{},
					).
					Return(nil, errors.New("fetch failed")).
					Times(1)
			},
			expectedTags:  nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysClient)
			}

			tags, err := repo.GetAllTags(context.Background())
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, tags)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedTags, tags)
			}
		})
	}
}

func TestToysRepository_GetTagByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysClient := mockclients.NewMockToysClient(ctrl)
	repo := NewToysRepository(toysClient)

	testCases := []struct {
		name          string
		id            uint32
		setupMocks    func(toysClient *mockclients.MockToysClient)
		expectedTag   *entities.Tag
		errorExpected bool
	}{
		{
			name: "success",
			id:   1,
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetTag(
						gomock.Any(),
						&toys.GetTagIn{ID: 1},
					).
					Return(&toys.GetTagOut{
						ID:   1,
						Name: "Test Tag",
					}, nil).
					Times(1)
			},
			expectedTag: &entities.Tag{
				ID:   1,
				Name: "Test Tag",
			},
			errorExpected: false,
		},
		{
			name: "error",
			id:   1,
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetTag(
						gomock.Any(),
						&toys.GetTagIn{ID: 1},
					).
					Return(nil, errors.New("not found")).
					Times(1)
			},
			expectedTag:   nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysClient)
			}

			tag, err := repo.GetTagByID(context.Background(), tc.id)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, tag)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedTag, tag)
			}
		})
	}
}

func TestToysRepository_GetMasterByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysClient := mockclients.NewMockToysClient(ctrl)
	repo := NewToysRepository(toysClient)

	now := time.Now().UTC().Truncate(time.Second)

	testCases := []struct {
		name           string
		userID         uint64
		setupMocks     func(toysClient *mockclients.MockToysClient)
		expectedMaster *entities.Master
		errorExpected  bool
	}{
		{
			name:   "success",
			userID: 1,
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetMasterByUser(
						gomock.Any(),
						&toys.GetMasterByUserIn{UserID: 1},
					).
					Return(&toys.GetMasterOut{
						ID:        1,
						UserID:    1,
						Info:      pointers.New("Master Info"),
						CreatedAt: timestamppb.New(now),
						UpdatedAt: timestamppb.New(now),
					}, nil).
					Times(1)
			},
			expectedMaster: &entities.Master{
				ID:        1,
				UserID:    1,
				Info:      pointers.New("Master Info"),
				CreatedAt: now,
				UpdatedAt: now,
			},
			errorExpected: false,
		},
		{
			name:   "error",
			userID: 1,
			setupMocks: func(toysClient *mockclients.MockToysClient) {
				toysClient.
					EXPECT().
					GetMasterByUser(
						gomock.Any(),
						&toys.GetMasterByUserIn{UserID: 1},
					).
					Return(nil, errors.New("not found")).
					Times(1)
			},
			expectedMaster: nil,
			errorExpected:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysClient)
			}

			master, err := repo.GetMasterByUser(context.Background(), tc.userID)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, master)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedMaster, master)
			}
		})
	}
}
