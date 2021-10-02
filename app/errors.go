package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func notFound(c *gin.Context) {
	fmt.Printf("AppRoot = %+v\n", AppRoot)
	c.Redirect(http.StatusFound, AppRoot+"/")
}

func secretExipred(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"title": "Secret Expired",
	})
}

func badRequest(c *gin.Context, err error) {
	status := http.StatusBadRequest
	if c.IsAborted() {
		status = c.Writer.Status()
	}
	c.JSON(status, gin.H{"error": err.Error()})
}
