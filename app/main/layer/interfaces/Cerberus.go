package interfaces

import "github.com/gin-gonic/gin"

type Cerberus interface {
	Execute() gin.HandlerFunc
}
