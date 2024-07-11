package main

import (
	"github.com/mlchain/mlchain-sandbox/internal/core/lib/python"
)
import "C"

//export MlchainSeccomp
func MlchainSeccomp(uid int, gid int, enable_network bool) {
	python.InitSeccomp(uid, gid, enable_network)
}

func main() {}
