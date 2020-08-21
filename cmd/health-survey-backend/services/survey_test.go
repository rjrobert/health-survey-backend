package services

import (
	"errors"
	"testing"

	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/models"
	"github.com/stretchr/testify/assert"
)

func TestNewSurveyService(t *testing.T) {
	dao := newmockSurveyDAO()
	s := NewSurveyService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestSurveyService_Get(t *testing.T) {
	s := NewSurveyService(newmockSurveyDAO())
	survey, err := s.Get(2)
	if assert.Nil(t, err) && assert.NotNil(t, survey) {
		assert.Equal(t, "Survey2", survey.Title)
		assert.Equal(t, true, survey.ShowCompletedPage)
	}

	survey, err = s.Get(100)
	assert.NotNil(t, err)
}

func newmockSurveyDAO() surveyDAO {
	return &mockSurveyDAO{
		records: []models.Survey{
			{Model: models.Model{ID: 1}, Title: "Survey1", ShowCompletedPage: false, NavigateToURL: "/dummyURL"},
			{Model: models.Model{ID: 2}, Title: "Survey2", ShowCompletedPage: true, NavigateToURL: "/dummyURL"},
		},
	}
}

// Mock Get function that replaces real Survey DAO
func (m *mockSurveyDAO) Get(id uint) (*models.Survey, error) {
	for _, record := range m.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

type mockSurveyDAO struct {
	records []models.Survey
}
