package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setupAPIRoute(r *gin.Engine) {
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	subs := r.Group("/subs")
	subs.Use(authMiddleWare())

	subs.GET("/domains/count", totalDomainsCount)
	subs.GET("/fetch", fetchSubs)
}

func fetchSubs(c *gin.Context) {
	id, exists := c.GetQuery("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	var subFind SubFinder
	if err := databaseConn.Select("domain", "subs", "tag").Where("id=(?)", id).First(&subFind).Error; err != nil {
		errL.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"domain": subFind.Domain, "subs": subFind.Subs, "tag": subFind.Tag})
}

func totalDomainsCount(c *gin.Context) {
	var count int64
	if err := databaseConn.Model(&SubFinder{}).Count(&count).Error; err != nil {
		errL.Print(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		expected := os.Getenv("API_KEY")

		// skip authentication if `API_KEY` env variable not set
		if expected == "" {
			c.Next()
			return
		}

		if apiKey != expected {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid API Key"})
		}
		c.Next()
	}
}
