package interfaces

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Execute(c *gin.Context)
}
