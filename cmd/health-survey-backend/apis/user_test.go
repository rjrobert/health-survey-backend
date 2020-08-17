package apis

import (
	"net/http"
	"testing"

	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/test_data"
)

func TestUser(t *testing.T) {
	path := test_data.GetTestCaseFolder()
	runAPITests(t, []apiTestCase{
		{"t1 - get a User", "GET", "/users/:id", "/users/1", "", GetUser, http.StatusOK, path + "/user_t1.json"},
		{"t2 - get a User not Present", "GET", "/users/:id", "/users/9999", "", GetUser, http.StatusNotFound, ""},
	})
}
