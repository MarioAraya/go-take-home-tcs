package main

import (
	"github.com/MarioAraya/go-take-home-tcs/internal/dashboard"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/dashboard/:id", dashboard.DashboardHandler)
	r.Run(":8080")
}
