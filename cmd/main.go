package main

import "github.com/sahilrana7582/go-storage/pkg/local"

func main() {
	lcCfg := local.NewLocalConfig("./exp/from", "./exp/to")
	lc := local.NewLocal(lcCfg)

	lc.Upload("demo.pdf")
}
