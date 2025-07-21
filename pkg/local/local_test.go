package local

import "testing"

const (
	fileName = "demo.pdf"
)

func Test(t *testing.T) {

	lcCfg := NewLocalConfig("./exp/from", "./exp/to")
	lc := NewLocal(lcCfg)

	lc.Upload(fileName)
}
