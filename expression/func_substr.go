package expression

import (
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
	"reflect"
)

type eFunctionSubstr struct {
}

func (e eFunctionSubstr) ToString() EString {
	return "substr"
}

func (e eFunctionSubstr) Call(args []EValue) (EValue, error) {
	if len(args) != 3 {
		return nil, xerr.Wrap(fmt.Errorf("expect 3 argument, got %d", len(args)))
	}
	value, ok := args[0].(EString)
	if !ok {
		return nil, xerr.Wrap(fmt.Errorf("argument 1 expect EString, got %s", reflect.TypeOf(args[0])))
	}
	start, ok := args[1].(EInt)
	if !ok {
		return nil, xerr.Wrap(fmt.Errorf("argument 2 expect EInt, got %s", reflect.TypeOf(args[1])))
	}
	end, ok := args[2].(EInt)
	if !ok {
		return nil, xerr.Wrap(fmt.Errorf("argument 3 expect EInt, got %s", reflect.TypeOf(args[2])))
	}
	return EString(string(value)[start:end]), nil
}
