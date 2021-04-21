package handler

import "github.com/gin-gonic/gin"
import "github.com/yagoazedias/dog-escaper/respository"

type PortHandler struct {
	r respository.PortRepositoryInterface
}

func (h *PortHandler) GetLastStatus(c *gin.Context) {
	status, err := h.r.GetLastStatus()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
	}

	c.JSON(200, gin.H{
		"status": status,
	})
}

func NewPortHandler(r respository.PortRepositoryInterface) *PortHandler {
	return &PortHandler{r}
}