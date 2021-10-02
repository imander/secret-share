package app

import (
	"net/http"
	"time"

	"secret-share/cache"

	"github.com/gin-gonic/gin"
)

var secretCache = cache.New(time.Hour)

type Secret struct {
	ID   string `uri:"id" json:"id"`
	Data string `json:"secret"`
	TTL  int    `json:"ttl"`
}

func SecretPost(c *gin.Context) {
	s := &Secret{}
	if err := c.ShouldBindJSON(s); err != nil {
		badRequest(c, err)
		return
	}
	if s.TTL > 0 {
		ttl := time.Duration(s.TTL) * time.Second
		s.ID = secretCache.AddWithTTL(s, ttl)
	} else {
		s.ID = secretCache.Add(s)
	}
	c.JSON(http.StatusOK, gin.H{"id": s.ID})
}

func SecretGet(c *gin.Context) {
	s := &Secret{}
	if err := c.ShouldBindUri(s); err != nil {
		badRequest(c, err)
		return
	}
	secret, ok := secretCache.Get(s.ID)
	if !ok {
		secretExipred(c)
		return
	}
	secretCache.Delete(s.ID)
	c.HTML(http.StatusOK, "get-secret.html", gin.H{
		"title": "Share Secret",
		"data":  secret.(*Secret).Data,
	})
}
