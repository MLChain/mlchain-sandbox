package main

import (
	"fmt"

	"github.com/mlchain/dify-sandbox/internal/core/runner/python"
	"github.com/mlchain/dify-sandbox/internal/core/runner/types"
	"github.com/mlchain/dify-sandbox/internal/service"
	"github.com/mlchain/dify-sandbox/internal/static"
)

func main() {
	static.InitConfig("conf/config.yaml")
	python.PreparePythonDependenciesEnv()
	resp := service.RunPython3Code(`import json;print(json.dumps({"hello": "world"}))`,
		``,
		&types.RunnerOptions{
			EnableNetwork: true,
		})

	fmt.Println(resp.Data)
}