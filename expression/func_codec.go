package expression

import (
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
)

type eFunctionCodec struct {
	fun func(value EValue) (EValue, error)
}

func (e eFunctionCodec) ToString() EString {
	return "codec"
}

func (e eFunctionCodec) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, xerr.Wrap(fmt.Errorf("expect one argument, got %d", len(args)))
	}
	return e.fun(args[0])
}
