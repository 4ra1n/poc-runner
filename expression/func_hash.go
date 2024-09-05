package expression

import (
	"encoding/hex"
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
	"hash"
	"reflect"
)

type eFunctionHash struct {
	h func() hash.Hash
}

func (e eFunctionHash) ToString() EString {
	return "hash"
}

func (e eFunctionHash) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, xerr.Wrap(fmt.Errorf("expect one argument, got %d", len(args)))
	}
	arg := args[0]
	h := e.h()
	h.Write([]byte(arg.ToString()))
	switch arg.(type) {
	case EString:
		return EString(hex.EncodeToString(h.Sum(nil))), nil
	case EBytes:
		return EString(hex.EncodeToString(h.Sum(nil))), nil
	default:
		return nil, xerr.Wrap(fmt.Errorf("expect EString,EBytes, got %s", reflect.TypeOf(arg)))
	}
}
