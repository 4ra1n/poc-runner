package expression

import (
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
	"reflect"
)

type eFunctionRandomAlpha struct {
	upper bool
}

func (e eFunctionRandomAlpha) ToString() EString {
	if e.upper {
		return "randomUppercase"
	} else {
		return "randomLowercase"
	}
}

func (e eFunctionRandomAlpha) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, xerr.Wrap(fmt.Errorf("expect one argument, got %d", len(args)))
	}
	n, ok := args[0].(EInt)
	if !ok {
		return nil, xerr.Wrap(fmt.Errorf("argument 1 expect EInt, got %s", reflect.TypeOf(args[0])))
	}
	var value string
	if e.upper {
		value = randStringRunes(int(n), []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	} else {
		value = randStringRunes(int(n), []rune("abcdefghijklmnopqrstuvwxyz"))
	}
	return EString(value), nil
}
