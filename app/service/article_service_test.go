package service

import (
	"blog-api/dto"
	"blog-api/model"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) GetList(ctx context.Context, preloads []string) ([]model.Article, []error) {
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

func TestGetArticleDetail(t *testing.T) {
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

	repo := new(MockRepo)
	service := NewArticleService(repo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo.On("GetDetail", mock.Anything, model.GetDetailArticleQueryParams{
				Id: &tt.inputID,
			}).Return(tt.mockReturn, tt.mockError)

			article, errs := service.GetDetail(context.TODO(), dto.GetArticleDetailRequest{
				Id: tt.inputID,
			})

			if tt.expectErr {
				for _, err := range errs {
					assert.Error(t, err)
				}
				assert.Nil(t, article)
			} else {
				assert.Zero(t, len(errs))
				assert.NotNil(t, article)
				assert.Equal(t, tt.mockReturn.Title, article.Title)
			}

			repo.AssertExpectations(t)
		})
	}
}
