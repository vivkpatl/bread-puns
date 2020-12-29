package impl

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

//https://my-bao-server.herokuapp.com/api/breadpuns
func GetPun(c *gin.Context) {
	resp, err := http.Get("https://my-bao-server.herokuapp.com/api/breadpuns")
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not get bread pun :( sadge",
		})
		return
	}

	var breadPun string
	err = json.NewDecoder(resp.Body).Decode(&breadPun)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not process bread pun response :( sadge",
		})
		return
	}

	c.JSON(http.StatusOK, breadPun)
}