package web

import (
	"net/http"

	"github.com/SeerLink/seerlink/core/services/seerlink"

	"github.com/gin-gonic/gin"
)

// PingController has the ping endpoint.
type PingController struct {
	App seerlink.Application
}

// Show returns pong.
func (eic *PingController) Show(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
