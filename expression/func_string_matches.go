package expression

import (
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
	"reflect"
	"regexp"
)

type eStringMatches struct {
	expr string
}

func (e eStringMatches) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, xerr.Wrap(
			fmt.Errorf("expect one argument, got %d", len(args)))
	}
	regex, err := regexp.Compile(e.expr)
	if err != nil {
		return nil, err
	}
	arg := args[0]
	switch arg.(type) {
	case EBytes:
		return EBool(regex.Match(arg.(EBytes))), nil
	case EString:
		return EBool(regex.MatchString(string(arg.(EString)))), nil
	default:
		return nil, xerr.Wrap(fmt.Errorf("expect EBytes,EString, got %s", reflect.TypeOf(arg)))
	}
}

func (e eStringMatches) ToString() EString {
	panic("implement me")
}
