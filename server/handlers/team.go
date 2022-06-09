package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"query/server/conf"
	"query/server/external"
)

func GetTeamById(c *gin.Context) {
	client := resty.New().
		SetBaseURL(conf.ScheduleHosts).
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", c.GetHeader("Authorization"))
	team, err := external.GetTeamById(c.Param("teamId"), client)

	if err != nil {
		log.Fatalln("ERROR: " + err.Error())
	}

	c.JSON(http.StatusOK, team)
}
