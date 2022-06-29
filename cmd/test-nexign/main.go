package main

import (
	"github.com/halfrollboy/test-nexign/internal/api/handler"
	"go.uber.org/zap"
)

var logger *zap.Logger

// @title        Speller API
// @version      1.0
// @description  This service implements spalling texts.

// @host                     localhost:8080
// @BasePath                 /
// @query.collection.format  multi

// @x-extension-openapi  {"example": "value on a json format"}
func main() {
	logger, _ = zap.NewProduction()
	router := handler.NewRouter()
	logger.Fatal("Fatal error to start app", zap.Error(router.Run(":8080")))
	router.Run()
}
