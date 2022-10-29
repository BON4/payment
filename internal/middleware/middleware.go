package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Middleware struct {
	logger *logrus.Entry
}

func NewMiddleware(logger *logrus.Entry) *Middleware {
	return &Middleware{
		logger: logger,
	}
}

func (m *Middleware) CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		ctx.Next()
	}
}
