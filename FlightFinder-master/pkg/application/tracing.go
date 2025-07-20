package application

import "github.com/gin-gonic/gin"

type TracingClient = gin.HandlerFunc

func NullTracingClient(c *gin.Context) {
	// no operation
}
