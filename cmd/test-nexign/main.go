package main

import (
	"github.com/halfrollboy/test-nexign/internal/api/handler"
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	logger, _ = zap.NewProduction()
	router := handler.NewRouter()
	logger.Fatal("Fatal error to start app", zap.Error(router.Run(":8080")))
	router.Run()
}
