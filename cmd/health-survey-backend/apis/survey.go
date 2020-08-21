package apis

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/daos"
	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/services"
)

// GetSurveys godoc
// @Summary Retrieves survey based on given ID
// @Produce json
// @Param id path integer true "Survey ID"
// @Success 200 {object} models.Survey
// @Router /surveys/{id} [get]
// @Security ApiKeyAuth
func GetSurveys(c *gin.Context) {
	s := services.NewSurveyService(daos.NewSurveyDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
