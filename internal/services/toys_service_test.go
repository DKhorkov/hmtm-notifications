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

func TestToysService_GetAllToys(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysRepository := mockrepositories.NewMockToysRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewToysService(toysRepository, logger)

	now := time.Now()
	testCases := []struct {
		name          string
		setupMocks    func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger)
		expectedToys  []entities.Toy
		errorExpected bool
	}{
		{
			name: "success",
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetAllToys(gomock.Any()).
					Return([]entities.Toy{
						{
							ID:          1,
							MasterID:    1,
							CategoryID:  1,
							Name:        "Test Toy",
							Description: "Test Description",
							Price:       100.0,
							Quantity:    10,
							CreatedAt:   now,
							UpdatedAt:   now,
							Tags:        []entities.Tag{{ID: 1, Name: "tag1"}},
							Attachments: []entities.ToyAttachment{{ID: 1, ToyID: 1, Link: "link1"}},
						},
					}, nil).
					Times(1)
			},
			expectedToys: []entities.Toy{
				{
					ID:          1,
					MasterID:    1,
					CategoryID:  1,
					Name:        "Test Toy",
					Description: "Test Description",
					Price:       100.0,
					Quantity:    10,
					CreatedAt:   now,
					UpdatedAt:   now,
					Tags:        []entities.Tag{{ID: 1, Name: "tag1"}},
					Attachments: []entities.ToyAttachment{{ID: 1, ToyID: 1, Link: "link1"}},
				},
			},
			errorExpected: false,
		},
		{
			name: "error",
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetAllToys(gomock.Any()).
					Return(nil, errors.New("fetch failed")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedToys:  nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysRepository, logger)
			}

			toys, err := service.GetAllToys(context.Background())
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, toys)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedToys, toys)
			}
		})
	}
}

func TestToysService_GetMasterToys(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysRepository := mockrepositories.NewMockToysRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewToysService(toysRepository, logger)

	now := time.Now()
	testCases := []struct {
		name          string
		masterID      uint64
		setupMocks    func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger)
		expectedToys  []entities.Toy
		errorExpected bool
	}{
		{
			name:     "success",
			masterID: 1,
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetMasterToys(gomock.Any(), uint64(1)).
					Return([]entities.Toy{
						{
							ID:          1,
							MasterID:    1,
							CategoryID:  1,
							Name:        "Test Toy",
							Description: "Test Description",
							Price:       100.0,
							Quantity:    10,
							CreatedAt:   now,
							UpdatedAt:   now,
							Tags:        []entities.Tag{{ID: 1, Name: "tag1"}},
						},
					}, nil).
					Times(1)
			},
			expectedToys: []entities.Toy{
				{
					ID:          1,
					MasterID:    1,
					CategoryID:  1,
					Name:        "Test Toy",
					Description: "Test Description",
					Price:       100.0,
					Quantity:    10,
					CreatedAt:   now,
					UpdatedAt:   now,
					Tags:        []entities.Tag{{ID: 1, Name: "tag1"}},
				},
			},
			errorExpected: false,
		},
		{
			name:     "error",
			masterID: 1,
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetMasterToys(gomock.Any(), uint64(1)).
					Return(nil, errors.New("fetch failed")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedToys:  nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysRepository, logger)
			}

			toys, err := service.GetMasterToys(context.Background(), tc.masterID)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, toys)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedToys, toys)
			}
		})
	}
}

func TestToysService_GetUserToys(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysRepository := mockrepositories.NewMockToysRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewToysService(toysRepository, logger)

	now := time.Now()
	testCases := []struct {
		name          string
		userID        uint64
		setupMocks    func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger)
		expectedToys  []entities.Toy
		errorExpected bool
	}{
		{
			name:   "success",
			userID: 1,
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetUserToys(gomock.Any(), uint64(1)).
					Return([]entities.Toy{
						{
							ID:          1,
							MasterID:    1,
							CategoryID:  1,
							Name:        "Test Toy",
							Description: "Test Description",
							Price:       100.0,
							Quantity:    10,
							CreatedAt:   now,
							UpdatedAt:   now,
						},
					}, nil).
					Times(1)
			},
			expectedToys: []entities.Toy{
				{
					ID:          1,
					MasterID:    1,
					CategoryID:  1,
					Name:        "Test Toy",
					Description: "Test Description",
					Price:       100.0,
					Quantity:    10,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			errorExpected: false,
		},
		{
			name:   "error",
			userID: 1,
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetUserToys(gomock.Any(), uint64(1)).
					Return(nil, errors.New("fetch failed")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedToys:  nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysRepository, logger)
			}

			toys, err := service.GetUserToys(context.Background(), tc.userID)
			if tc.errorExpected {
				require.Error(t, err)
				require.Nil(t, toys)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedToys, toys)
			}
		})
	}
}

func TestToysService_GetToyByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysRepository := mockrepositories.NewMockToysRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewToysService(toysRepository, logger)

	now := time.Now()
	testCases := []struct {
		name          string
		id            uint64
		setupMocks    func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger)
		expectedToy   *entities.Toy
		errorExpected bool
	}{
		{
			name: "success",
			id:   1,
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetToyByID(gomock.Any(), uint64(1)).
					Return(&entities.Toy{
						ID:          1,
						MasterID:    1,
						CategoryID:  1,
						Name:        "Test Toy",
						Description: "Test Description",
						Price:       100.0,
						Quantity:    10,
						CreatedAt:   now,
						UpdatedAt:   now,
						Tags:        []entities.Tag{{ID: 1, Name: "tag1"}},
					}, nil).
					Times(1)
			},
			expectedToy: &entities.Toy{
				ID:          1,
				MasterID:    1,
				CategoryID:  1,
				Name:        "Test Toy",
				Description: "Test Description",
				Price:       100.0,
				Quantity:    10,
				CreatedAt:   now,
				UpdatedAt:   now,
				Tags:        []entities.Tag{{ID: 1, Name: "tag1"}},
			},
			errorExpected: false,
		},
		{
			name: "error",
			id:   1,
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetToyByID(gomock.Any(), uint64(1)).
					Return(nil, errors.New("not found")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedToy:   nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysRepository, logger)
			}

			toy, err := service.GetToyByID(context.Background(), tc.id)
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

func TestToysService_GetAllMasters(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysRepository := mockrepositories.NewMockToysRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewToysService(toysRepository, logger)

	now := time.Now()
	info := "Master Info"
	testCases := []struct {
		name            string
		setupMocks      func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger)
		expectedMasters []entities.Master
		errorExpected   bool
	}{
		{
			name: "success",
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetAllMasters(gomock.Any()).
					Return([]entities.Master{
						{
							ID:        1,
							UserID:    1,
							Info:      &info,
							CreatedAt: now,
							UpdatedAt: now,
						},
					}, nil).
					Times(1)
			},
			expectedMasters: []entities.Master{
				{
					ID:        1,
					UserID:    1,
					Info:      &info,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			errorExpected: false,
		},
		{
			name: "error",
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetAllMasters(gomock.Any()).
					Return(nil, errors.New("fetch failed")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedMasters: nil,
			errorExpected:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysRepository, logger)
			}

			masters, err := service.GetAllMasters(context.Background())
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

func TestToysService_GetMasterByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysRepository := mockrepositories.NewMockToysRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewToysService(toysRepository, logger)

	now := time.Now()
	info := "Master Info"
	testCases := []struct {
		name           string
		id             uint64
		setupMocks     func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger)
		expectedMaster *entities.Master
		errorExpected  bool
	}{
		{
			name: "success",
			id:   1,
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetMasterByID(gomock.Any(), uint64(1)).
					Return(&entities.Master{
						ID:        1,
						UserID:    1,
						Info:      &info,
						CreatedAt: now,
						UpdatedAt: now,
					}, nil).
					Times(1)
			},
			expectedMaster: &entities.Master{
				ID:        1,
				UserID:    1,
				Info:      &info,
				CreatedAt: now,
				UpdatedAt: now,
			},
			errorExpected: false,
		},
		{
			name: "error",
			id:   1,
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetMasterByID(gomock.Any(), uint64(1)).
					Return(nil, errors.New("not found")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedMaster: nil,
			errorExpected:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysRepository, logger)
			}

			master, err := service.GetMasterByID(context.Background(), tc.id)
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

func TestToysService_GetAllCategories(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysRepository := mockrepositories.NewMockToysRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewToysService(toysRepository, logger)

	testCases := []struct {
		name               string
		setupMocks         func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger)
		expectedCategories []entities.Category
		errorExpected      bool
	}{
		{
			name: "success",
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetAllCategories(gomock.Any()).
					Return([]entities.Category{
						{ID: 1, Name: "Category 1"},
					}, nil).
					Times(1)
			},
			expectedCategories: []entities.Category{
				{ID: 1, Name: "Category 1"},
			},
			errorExpected: false,
		},
		{
			name: "error",
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetAllCategories(gomock.Any()).
					Return(nil, errors.New("fetch failed")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedCategories: nil,
			errorExpected:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysRepository, logger)
			}

			categories, err := service.GetAllCategories(context.Background())
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

func TestToysService_GetCategoryByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysRepository := mockrepositories.NewMockToysRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewToysService(toysRepository, logger)

	testCases := []struct {
		name             string
		id               uint32
		setupMocks       func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger)
		expectedCategory *entities.Category
		errorExpected    bool
	}{
		{
			name: "success",
			id:   1,
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetCategoryByID(gomock.Any(), uint32(1)).
					Return(&entities.Category{ID: 1, Name: "Category 1"}, nil).
					Times(1)
			},
			expectedCategory: &entities.Category{ID: 1, Name: "Category 1"},
			errorExpected:    false,
		},
		{
			name: "error",
			id:   1,
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetCategoryByID(gomock.Any(), uint32(1)).
					Return(nil, errors.New("not found")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedCategory: nil,
			errorExpected:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysRepository, logger)
			}

			category, err := service.GetCategoryByID(context.Background(), tc.id)
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

func TestToysService_GetAllTags(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysRepository := mockrepositories.NewMockToysRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewToysService(toysRepository, logger)

	testCases := []struct {
		name          string
		setupMocks    func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger)
		expectedTags  []entities.Tag
		errorExpected bool
	}{
		{
			name: "success",
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetAllTags(gomock.Any()).
					Return([]entities.Tag{
						{ID: 1, Name: "tag1"},
					}, nil).
					Times(1)
			},
			expectedTags: []entities.Tag{
				{ID: 1, Name: "tag1"},
			},
			errorExpected: false,
		},
		{
			name: "error",
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetAllTags(gomock.Any()).
					Return(nil, errors.New("fetch failed")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedTags:  nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysRepository, logger)
			}

			tags, err := service.GetAllTags(context.Background())
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

func TestToysService_GetTagByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysRepository := mockrepositories.NewMockToysRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewToysService(toysRepository, logger)

	testCases := []struct {
		name          string
		id            uint32
		setupMocks    func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger)
		expectedTag   *entities.Tag
		errorExpected bool
	}{
		{
			name: "success",
			id:   1,
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetTagByID(gomock.Any(), uint32(1)).
					Return(&entities.Tag{ID: 1, Name: "tag1"}, nil).
					Times(1)
			},
			expectedTag:   &entities.Tag{ID: 1, Name: "tag1"},
			errorExpected: false,
		},
		{
			name: "error",
			id:   1,
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetTagByID(gomock.Any(), uint32(1)).
					Return(nil, errors.New("not found")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedTag:   nil,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysRepository, logger)
			}

			tag, err := service.GetTagByID(context.Background(), tc.id)
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

func TestToysService_GetMasterByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	toysRepository := mockrepositories.NewMockToysRepository(ctrl)
	logger := mocklogging.NewMockLogger(ctrl)
	service := NewToysService(toysRepository, logger)

	now := time.Now()
	info := "Master Info"
	testCases := []struct {
		name           string
		userID         uint64
		setupMocks     func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger)
		expectedMaster *entities.Master
		errorExpected  bool
	}{
		{
			name:   "success",
			userID: 1,
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetMasterByUser(gomock.Any(), uint64(1)).
					Return(&entities.Master{
						ID:        1,
						UserID:    1,
						Info:      &info,
						CreatedAt: now,
						UpdatedAt: now,
					}, nil).
					Times(1)
			},
			expectedMaster: &entities.Master{
				ID:        1,
				UserID:    1,
				Info:      &info,
				CreatedAt: now,
				UpdatedAt: now,
			},
			errorExpected: false,
		},
		{
			name:   "error",
			userID: 1,
			setupMocks: func(toysRepository *mockrepositories.MockToysRepository, logger *mocklogging.MockLogger) {
				toysRepository.
					EXPECT().
					GetMasterByUser(gomock.Any(), uint64(1)).
					Return(nil, errors.New("not found")).
					Times(1)

				logger.
					EXPECT().
					ErrorContext(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedMaster: nil,
			errorExpected:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupMocks != nil {
				tc.setupMocks(toysRepository, logger)
			}

			master, err := service.GetMasterByUser(context.Background(), tc.userID)
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
