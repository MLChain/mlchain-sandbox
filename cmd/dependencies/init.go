package main

import (
	"github.com/mlchain/mlchain-sandbox/internal/core/runner/python"
	"github.com/mlchain/mlchain-sandbox/internal/static"
	"github.com/mlchain/mlchain-sandbox/internal/utils/log"
)

func main() {
	static.InitConfig("conf/config.yaml")

	err := python.PreparePythonDependenciesEnv()
	if err != nil {
		log.Panic("failed to initialize python dependencies sandbox: %v", err)
	}

	log.Info("Python dependencies initialized successfully")
}
