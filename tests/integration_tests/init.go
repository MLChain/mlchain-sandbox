package integrationtests_test

import (
	"github.com/mlchain/dify-sandbox/internal/core/runner/python"
	"github.com/mlchain/dify-sandbox/internal/static"
	"github.com/mlchain/dify-sandbox/internal/utils/log"
)

func init() {
	static.InitConfig("conf/config.yaml")

	// Test case for sys_fork
	err := python.PreparePythonDependenciesEnv()
	if err != nil {
		log.Panic("failed to initialize python dependencies sandbox: %v", err)
	}
}