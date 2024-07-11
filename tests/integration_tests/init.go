package integrationtests_test

import (
	"github.com/mlchain/mlchain-sandbox/internal/core/runner/python"
	"github.com/mlchain/mlchain-sandbox/internal/static"
	"github.com/mlchain/mlchain-sandbox/internal/utils/log"
)

func init() {
	static.InitConfig("conf/config.yaml")

	// Test case for sys_fork
	err := python.PreparePythonDependenciesEnv()
	if err != nil {
		log.Panic("failed to initialize python dependencies sandbox: %v", err)
	}
}
