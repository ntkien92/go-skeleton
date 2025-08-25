package suites

import (
	"api-tests/config"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ArticleSuite struct {
	suite.Suite
	client *resty.Client
}

func (s *ArticleSuite) SetupSuite() {
	s.client = resty.New()
}

func TestArticleSuite(t *testing.T) {
	suite.Run(t, new(ArticleSuite))
}

func (s *ArticleSuite) TestGetListArticles() {
	resp, err := s.client.R().
		Get(config.Global.BaseURL + "/articles")

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, resp.StatusCode())
}
