package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "post-secret.html", gin.H{
		"title": "Share Secret",
	})
}
