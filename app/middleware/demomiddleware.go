package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type DemoMiddleware struct {
}

func NewDemoMiddleware() *DemoMiddleware {
	return &DemoMiddleware{}
}

func (m *DemoMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("this is demo middleware")
	}
}
