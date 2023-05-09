package main

import (
	"fmt"
	"net/http"
	"studentAPI/api"

	"flag"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func noRoute(c *gin.Context) {
	glog.Errorf("route %s not found", c.Request.URL)
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"code": 404,
		"message": fmt.Sprintf("route %s not found", c.Request.URL),
	})
}


func main() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Lookup("stderrthreshold").Value.Set("ERROR")

	glog.Info("Starting server...")

	router := gin.Default()
	router.NoRoute(noRoute)
	router.GET("/student", api.GetStudentData)
	router.POST("/student", api.AddStudentData)
	router.Run("localhost:8080")
}	