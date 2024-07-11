package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mlchain/mlchain-sandbox/internal/middleware"
	"github.com/mlchain/mlchain-sandbox/internal/static"
)

func Setup(eng *gin.Engine) {
	eng.Use(middleware.Auth())

	eng.POST(
		"/v1/sandbox/run",
		middleware.MaxRequest(static.GetMlchainSandboxGlobalConfigurations().MaxRequests),
		middleware.MaxWorker(static.GetMlchainSandboxGlobalConfigurations().MaxWorkers),
		RunSandboxController,
	)
	eng.GET(
		"/v1/sandbox/dependencies",
		GetDependencies,
	)

	eng.POST(
		"/v1/sandbox/dependencies/update",
		UpdateDependencies,
	)
}
