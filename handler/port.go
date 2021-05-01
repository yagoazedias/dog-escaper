package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
import "github.com/yagoazedias/dog-escaper/respository"

type PortHandler struct {
	r respository.PortRepositoryInterface
}

func (h *PortHandler) GetLastStatus(c *gin.Context) {
	port, err := h.r.GetLastStatus()

	if err != nil {
		fmt.Println("error trying to get door status", err)
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"isOpen": port.IsOpen,
		"timestamp": port.Timestamp,
	})
}

func NewPortHandler(r respository.PortRepositoryInterface) *PortHandler {
	return &PortHandler{r}
}