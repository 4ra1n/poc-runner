package expression

import (
	"bytes"
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
	"reflect"
)

type eBytesBcontains struct {
	b  EBytes
	ci bool
}

func (e eBytesBcontains) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, xerr.Wrap(
			fmt.Errorf("expect one argument, got %d", len(args)))
	}
	arg := args[0]
	switch arg.(type) {
	case EBytes:
		if e.ci {
			return EBool(bytes.Contains(bytes.ToLower(e.b), bytes.ToLower(arg.(EBytes)))), nil
		} else {
			return EBool(bytes.Contains(e.b, arg.(EBytes))), nil
		}
	default:
		return nil, xerr.Wrap(fmt.Errorf("expect EBytes, got %s", reflect.TypeOf(arg)))
	}
}

func (e eBytesBcontains) ToString() EString {
	panic("implement me")
}
