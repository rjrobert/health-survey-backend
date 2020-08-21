package services

import "github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/models"

type surveyDAO interface {
	Get(id uint) (*models.Survey, error)
}

type SurveyService struct {
	dao surveyDAO
}

// NewSurveyService creates a new SurveyService with the given surveyDAO
func NewSurveyService(dao surveyDAO) *SurveyService {
	return &SurveyService{dao}
}

// Get just retrieves survey using Survey DAO
func (s *SurveyService) Get(id uint) (*models.Survey, error) {
	return s.dao.Get(id)
}
