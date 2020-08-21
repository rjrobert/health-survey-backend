package daos

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/config"
	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/models"
)

// SurveyDAO handles actual interaction with database
type SurveyDAO struct{}

// NewSurveyDAO creates a new UserDAO
func NewSurveyDAO() *SurveyDAO {
	return &SurveyDAO{}
}

// Get does the actual query
func (dao *SurveyDAO) Get(id uint) (*models.Survey, error) {
	var survey models.Survey
	var err error
	if id == 0 {
		err = config.Config.DB.Set("gorm:auto_preload", true).
			First(&survey).
			Error
	} else {
		err = config.Config.DB.Set("gorm:auto_preload", true).
			Where("id = ?", id).
			Find(&survey).
			Error
	}

	return &survey, err
}

func PreloadSurveyData() {
	jsonFile, err := os.Open("config/survey.json")
	if err != nil {
		panic(err)
	}
	j, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	var survey models.Survey
	json.Unmarshal(j, &survey)

	config.Config.DB.NewRecord(&survey)
	config.Config.DB.Create(&survey)
}
