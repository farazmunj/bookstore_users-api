package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping will deal with ping
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
