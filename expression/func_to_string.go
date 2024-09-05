package expression

import (
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
	"reflect"
)

type eFunctionToString struct {
}

func (e eFunctionToString) ToString() EString {
	return "string"
}

func (e eFunctionToString) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, xerr.Wrap(fmt.Errorf("expect one argument, got %d", len(args)))
	}
	arg := args[0]
	switch arg.(type) {
	case EInt, EFloat, EBool, EString, EBytes:
		return arg.ToString(), nil
	default:
		return nil, xerr.Wrap(fmt.Errorf("expect EValue, got %s", reflect.TypeOf(arg)))
	}
}
