package expression

import (
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
	"math/rand"
	"reflect"
)

type eFunctionRandomInt struct {
}

func (e eFunctionRandomInt) ToString() EString {
	return "randomInt"
}

func (e eFunctionRandomInt) Call(args []EValue) (EValue, error) {
	if len(args) != 2 {
		return nil, xerr.Wrap(fmt.Errorf("expect two argument, got %d", len(args)))
	}
	minVal, ok := args[0].(EInt)
	if !ok {
		return nil, xerr.Wrap(fmt.Errorf("argument 1 expect EInt, got %s", reflect.TypeOf(args[0])))
	}
	maxVal, ok := args[1].(EInt)
	if !ok {
		return nil, xerr.Wrap(fmt.Errorf("argument 2 expect EInt, got %s", reflect.TypeOf(args[1])))
	}
	return EInt(rand.Intn(int(maxVal-minVal)+1) + 1), nil
}
