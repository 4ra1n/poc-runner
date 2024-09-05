package expression

import (
	"fmt"
	"github.com/4ra1n/poc-runner/reverse"
	"github.com/4ra1n/poc-runner/xerr"
	"reflect"
)

type eFunctionReverseWait struct {
	rev reverse.Reverse
}

func (e eFunctionReverseWait) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, xerr.Wrap(fmt.Errorf("expect 1 argument, got %d", len(args)))
	}
	value, ok := args[0].(EInt)
	if !ok {
		return nil, xerr.Wrap(fmt.Errorf("argument 1 expect EInt, got %s", reflect.TypeOf(args[0])))
	}
	return EBool(e.rev.Wait(int(value))), nil
}

func (e eFunctionReverseWait) ToString() EString {
	return "wait"
}
