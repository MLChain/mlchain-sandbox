package main

import "github.com/mlchain/mlchain-sandbox/internal/core/lib/nodejs"
import "C"

//export MlchainSeccomp
func MlchainSeccomp(uid int, gid int, enable_network bool) {
	nodejs.InitSeccomp(uid, gid, enable_network)
}

func main() {}
