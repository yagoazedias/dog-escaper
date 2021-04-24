package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yagoazedias/dog-escaper/respository"
)
import "github.com/yagoazedias/dog-escaper/handler"

func main() {
	router := gin.Default()
	PortHandler := handler.NewPortHandler(respository.PortRepository{})
	v1 := router.Group("/v1")
	{
		v1.GET("/port", PortHandler.GetLastStatus)
	}
	router.Run(":8000")
}