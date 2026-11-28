package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupAPIRoute(r *gin.Engine) {
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	subs := r.Group("/subs")
	subs.Use(authMiddleWare())

	subs.GET("/domains", fetchTotalDomains)
	subs.GET("/domains/count", totalDomainsCount)
	subs.GET("/fetch", fetchSubs)
}

func fetchSubs(c *gin.Context) {
	domain := c.Query("domain")
	id := c.Query("id")

	var subFind SubFinder
	if id != "" {
		if err := databaseConn.Select("domain", "subs", "tag").Where("id=(?)", id).First(&subFind).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
				return
			}
			writeServerError(c)
			errL.Println(err)
			return
		}
	} else if domain != "" {
		if err := databaseConn.Select("domain", "subs", "tag").Where("domain=(?)", domain).First(&subFind).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
				return
			}
			writeServerError(c)
			errL.Println(err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"domain": subFind.Domain, "subs": subFind.Subs, "tag": subFind.Tag})
}

func fetchTotalDomains(c *gin.Context) {
	var data []SubFinder
	if err := databaseConn.Select("domain", "id", "tag").Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
			return
		}
		writeServerError(c)
		errL.Println(err)
		return
	}

	type domains struct {
		Domain string `json:"domain"`
		ID     int    `json:"id"`
		Tag    string `json:"tag"`
	}

	var response []domains

	for _, d := range data {
		response = append(response, domains{
			Domain: d.Domain,
			ID:     d.ID,
			Tag:    d.Tag,
		})
	}

	c.JSON(http.StatusOK, response)
}

func totalDomainsCount(c *gin.Context) {
	var count int64
	if err := databaseConn.Model(&SubFinder{}).Count(&count).Error; err != nil {
		writeServerError(c)
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

func writeServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
}
