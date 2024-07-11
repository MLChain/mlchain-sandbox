package service

import (
	"errors"

	"github.com/mlchain/mlchain-sandbox/internal/core/runner/types"
	"github.com/mlchain/mlchain-sandbox/internal/static"
)

var (
	ErrNetworkDisabled = errors.New("network is disabled, please enable it in the configuration")
)

func checkOptions(options *types.RunnerOptions) error {
	configuration := static.GetMlchainSandboxGlobalConfigurations()

	if options.EnableNetwork && !configuration.EnableNetwork {
		return ErrNetworkDisabled
	}

	return nil
}
