package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupAPIRoute(r *gin.Engine) {
	subs := r.Group("/subs")

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
