package static

import (
	"os"
	"strconv"
	"strings"

	"github.com/mlchain/mlchain-sandbox/internal/types"
	"github.com/mlchain/mlchain-sandbox/internal/utils/log"
	"gopkg.in/yaml.v3"
)

var mlchainSandboxGlobalConfigurations types.MlchainSandboxGlobalConfigurations

func InitConfig(path string) error {
	mlchainSandboxGlobalConfigurations = types.MlchainSandboxGlobalConfigurations{}

	// read config file
	configFile, err := os.Open(path)
	if err != nil {
		return err
	}

	defer configFile.Close()

	// parse config file
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&mlchainSandboxGlobalConfigurations)
	if err != nil {
		return err
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err == nil {
		mlchainSandboxGlobalConfigurations.App.Debug = debug
	}

	max_workers := os.Getenv("MAX_WORKERS")
	if max_workers != "" {
		mlchainSandboxGlobalConfigurations.MaxWorkers, _ = strconv.Atoi(max_workers)
	}

	max_requests := os.Getenv("MAX_REQUESTS")
	if max_requests != "" {
		mlchainSandboxGlobalConfigurations.MaxRequests, _ = strconv.Atoi(max_requests)
	}

	port := os.Getenv("SANDBOX_PORT")
	if port != "" {
		mlchainSandboxGlobalConfigurations.App.Port, _ = strconv.Atoi(port)
	}

	timeout := os.Getenv("WORKER_TIMEOUT")
	if timeout != "" {
		mlchainSandboxGlobalConfigurations.WorkerTimeout, _ = strconv.Atoi(timeout)
	}

	api_key := os.Getenv("API_KEY")
	if api_key != "" {
		mlchainSandboxGlobalConfigurations.App.Key = api_key
	}

	python_path := os.Getenv("PYTHON_PATH")
	if python_path != "" {
		mlchainSandboxGlobalConfigurations.PythonPath = python_path
	}

	if mlchainSandboxGlobalConfigurations.PythonPath == "" {
		mlchainSandboxGlobalConfigurations.PythonPath = "/usr/local/bin/python3"
	}

	python_lib_path := os.Getenv("PYTHON_LIB_PATH")
	if python_lib_path != "" {
		mlchainSandboxGlobalConfigurations.PythonLibPaths = strings.Split(python_lib_path, ",")
	}

	if len(mlchainSandboxGlobalConfigurations.PythonLibPaths) == 0 {
		mlchainSandboxGlobalConfigurations.PythonLibPaths = DEFAULT_PYTHON_LIB_REQUIREMENTS
	}

	python_pip_mirror_url := os.Getenv("PIP_MIRROR_URL")
	if python_pip_mirror_url != "" {
		mlchainSandboxGlobalConfigurations.PythonPipMirrorURL = python_pip_mirror_url
	}
	nodejs_path := os.Getenv("NODEJS_PATH")
	if nodejs_path != "" {
		mlchainSandboxGlobalConfigurations.NodejsPath = nodejs_path
	}

	if mlchainSandboxGlobalConfigurations.NodejsPath == "" {
		mlchainSandboxGlobalConfigurations.NodejsPath = "/usr/local/bin/node"
	}

	enable_network := os.Getenv("ENABLE_NETWORK")
	if enable_network != "" {
		mlchainSandboxGlobalConfigurations.EnableNetwork, _ = strconv.ParseBool(enable_network)
	}

	allowed_syscalls := os.Getenv("ALLOWED_SYSCALLS")
	if allowed_syscalls != "" {
		strs := strings.Split(allowed_syscalls, ",")
		ary := make([]int, len(strs))
		for i := range ary {
			ary[i], err = strconv.Atoi(strs[i])
			if err != nil {
				return err
			}
		}
		mlchainSandboxGlobalConfigurations.AllowedSyscalls = ary
	}

	if mlchainSandboxGlobalConfigurations.EnableNetwork {
		log.Info("network has been enabled")
		socks5_proxy := os.Getenv("SOCKS5_PROXY")
		if socks5_proxy != "" {
			mlchainSandboxGlobalConfigurations.Proxy.Socks5 = socks5_proxy
		}

		if mlchainSandboxGlobalConfigurations.Proxy.Socks5 != "" {
			log.Info("using socks5 proxy: %s", mlchainSandboxGlobalConfigurations.Proxy.Socks5)
		}

		https_proxy := os.Getenv("HTTPS_PROXY")
		if https_proxy != "" {
			mlchainSandboxGlobalConfigurations.Proxy.Https = https_proxy
		}

		if mlchainSandboxGlobalConfigurations.Proxy.Https != "" {
			log.Info("using https proxy: %s", mlchainSandboxGlobalConfigurations.Proxy.Https)
		}

		http_proxy := os.Getenv("HTTP_PROXY")
		if http_proxy != "" {
			mlchainSandboxGlobalConfigurations.Proxy.Http = http_proxy
		}

		if mlchainSandboxGlobalConfigurations.Proxy.Http != "" {
			log.Info("using http proxy: %s", mlchainSandboxGlobalConfigurations.Proxy.Http)
		}
	}
	return nil
}

// avoid global modification, use value copy instead
func GetMlchainSandboxGlobalConfigurations() types.MlchainSandboxGlobalConfigurations {
	return mlchainSandboxGlobalConfigurations
}

type RunnerDependencies struct {
	PythonRequirements string
}

var runnerDependencies RunnerDependencies

func GetRunnerDependencies() RunnerDependencies {
	return runnerDependencies
}

func SetupRunnerDependencies() error {
	file, err := os.ReadFile("dependencies/python-requirements.txt")
	if err != nil {
		if err == os.ErrNotExist {
			return nil
		}
		return err
	}

	runnerDependencies.PythonRequirements = string(file)

	return nil
}
