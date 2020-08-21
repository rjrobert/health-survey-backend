package apis

import (
	"net/http"
	"testing"

	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/test_data"
)

func TestSurvey(t *testing.T) {
	path := test_data.GetTestCaseFolder()
	runAPITests(t, []apiTestCase{
		{"t1 - get a Survey", "GET", "/surveys/:id", "/surveys/1", "", GetSurveys, http.StatusOK, path + "/survey_t1.json"},
		{"t2 - get a Survey not Present", "GET", "/surveys/:id", "/surveys/9999", "", GetSurveys, http.StatusNotFound, ""},
	})
}
