package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "admin login here.",
	})
}