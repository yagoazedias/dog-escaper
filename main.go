package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yagoazedias/dog-escaper/handler"
	"github.com/yagoazedias/dog-escaper/infraestruture"
	"github.com/yagoazedias/dog-escaper/mqtt"
	"github.com/yagoazedias/dog-escaper/respository"
)

func main() {
	router := gin.Default()
	PortHandler := handler.NewPortHandler(respository.PortRepository{})
	router.Use(infraestruture.CORSMiddleware())

	v1 := router.Group("/v1")
	{
		v1.GET("/port", PortHandler.GetLastStatus)
	}

	mqtt.ConfigureMQQT()
	_ = router.Run(":8000")
}