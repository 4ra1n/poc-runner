package expression

import (
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
	"reflect"
	"time"
)

type eFunctionSleep struct {
}

func (e eFunctionSleep) ToString() EString {
	return "sleep"
}

func (e eFunctionSleep) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, xerr.Wrap(fmt.Errorf("expect one argument, got %d", len(args)))
	}
	second, ok := args[0].(EInt)
	if !ok {
		return nil, xerr.Wrap(fmt.Errorf("argument 1 expect EInt, got %s", reflect.TypeOf(args[0])))
	}
	time.Sleep(time.Duration(second) * time.Second)
	return EBool(true), nil
}
