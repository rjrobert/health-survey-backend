package daos

import (
	"testing"

	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/config"
	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/test_data"
	"github.com/stretchr/testify/assert"
)

func TestSurveyDAO_Get(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewSurveyDAO()

	survey, err := dao.Get(1)

	expected := map[string]string{"title": "Survey1", "navigate_to_url": "/dummyURL"}

	assert.Nil(t, err)
	assert.Equal(t, expected["title"], survey.Title)
	assert.Equal(t, expected["navigate_to_url"], survey.NavigateToURL)
}

func TestSurveyDAO_GetNotPresent(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewSurveyDAO()

	survey, err := dao.Get(9999)

	assert.NotNil(t, err)
	assert.Equal(t, "", survey.Title)
	assert.Equal(t, "", survey.NavigateToURL)
}
