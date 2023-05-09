package api

import (
	"net/http"
	"studentAPI/data"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)


func AddStudentData(c *gin.Context) {
	glog.Info("POST /student")
	glog.Info(c.Request.Body)
}


func GetStudentData(c *gin.Context){
	glog.Info("GET /student")
	
	studentData, err := data.ReadData()
	if err != nil {
		glog.Error(err)
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(err.Error()))
	}
	
	c.IndentedJSON(http.StatusOK, studentData)
}