package handler

import (
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	controllers "github.com/halfrollboy/test-nexign/internal/api/controllers"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func LogMiddleware() gin.HandlerFunc {
	defer logger.Sync()

	return func(c *gin.Context) {
		logger.Info("log to fetch URL",
			// zap.String("code", code),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
		c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		c.Next()
	}
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(LogMiddleware())
	pprof.Register(router)
	router.POST("/", controllers.CheckCorrect)

	// router.GET("/swagger/*any", swaggo.WrapHandler(swaggerFiles.Handler))

	return router
}
