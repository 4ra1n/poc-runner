package expression

import (
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
	"reflect"
)

type eFunctionToBytes struct {
}

func (e eFunctionToBytes) ToString() EString {
	return "bytes"
}

func (e eFunctionToBytes) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, xerr.Wrap(fmt.Errorf("expect one argument, got %d", len(args)))
	}
	arg := args[0]
	switch arg.(type) {
	case EString:
		return EBytes(arg.(EString)), nil
	default:
		return nil, xerr.Wrap(fmt.Errorf("expect EString, got %s", reflect.TypeOf(arg)))
	}
}
