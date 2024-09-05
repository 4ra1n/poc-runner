package expression

import (
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
	"reflect"
	"strings"
)

type eStringContains struct {
	v  EString
	ci bool
}

func (e eStringContains) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, xerr.Wrap(
			fmt.Errorf("expect one argument, got %d", len(args)))
	}
	arg := args[0]
	switch arg.(type) {
	case EString:
		str := string(e.v)
		substr := string(arg.(EString))
		if e.ci {
			return EBool(strings.Contains(strings.ToLower(str), strings.ToLower(substr))), nil
		} else {
			return EBool(strings.Contains(str, substr)), nil
		}
	default:
		return nil, xerr.Wrap(fmt.Errorf("expect EString, got %s", reflect.TypeOf(arg)))
	}
}

func (e eStringContains) ToString() EString {
	panic("implement me")
}
