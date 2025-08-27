package service

import (
	"blog-api/dto"
	"blog-api/interfaces"
	"blog-api/model"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) GetList(ctx context.Context) ([]model.Article, []error) {
	return nil, nil
}

func (m *MockRepo) GetDetail(ctx context.Context, params model.GetDetailArticleQueryParams) (*model.Article, []error) {
	args := m.Called(ctx, params)

	var article *model.Article
	if v := args.Get(0); v != nil {
		article = v.(*model.Article)
	}

	var errs []error
	if v := args.Get(1); v != nil {
		errs = v.([]error)
	}

	return article, errs
}

func (m *MockRepo) Create(ctx context.Context, data model.Article) (string, []error) {
	return "", nil
}

type ArticleServiceTestSuite struct {
	suite.Suite
	repo    *MockRepo
	service interfaces.ArticleServiceInterface
}

func (s *ArticleServiceTestSuite) SetupTest() {
	s.repo = new(MockRepo)
	s.service = NewArticleService(s.repo)
}



func (s *ArticleServiceTestSuite) TestGetArticleDetail_Cases() {
	ids := []uuid.UUID{
		uuid.New(),
		uuid.New(),
	}
	tests := []struct {
		name       string
		inputID    string
		mockReturn *model.Article
		mockError  []error
		expectErr  bool
	}{
		{
			name:    "success",
			inputID: ids[0].String(),
			mockReturn: &model.Article{
				ModelUUID: model.ModelUUID{ID: ids[0]},
				Title:     "Mock Title",
			},
			mockError: nil,
			expectErr: false,
		},
		{
			name:       "not found",
			inputID:    ids[1].String(),
			mockReturn: nil,
			mockError:  []error{errors.New("not found")},
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.repo.On("GetDetail", context.TODO(), model.GetDetailArticleQueryParams{
				Id: &tt.inputID,
			}).Return(tt.mockReturn, tt.mockError)

			article, errs := s.service.GetDetail(context.TODO(), dto.GetArticleDetailRequest{
				Id: tt.inputID,
			})

			if tt.expectErr {
				for _, err := range errs {
					assert.Error(s.T(), err)
				}
				assert.Nil(s.T(), article)
			} else {
				assert.Zero(s.T(), len(errs))
				assert.NotNil(s.T(), article)
				assert.Equal(s.T(), tt.mockReturn.Title, article.Title)
			}

			s.repo.AssertExpectations(s.T())
		})
	}
}

// ----------------------
// Run suite
// ----------------------
func TestArticleServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ArticleServiceTestSuite))
}
