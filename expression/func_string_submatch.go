package expression

import (
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
	"reflect"
	"regexp"
)

type eStringSubmatch struct {
	expr string
}

func (e eStringSubmatch) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, xerr.Wrap(fmt.Errorf("expect one argument, got %d", len(args)))
	}
	regex, err := regexp.Compile(e.expr)
	if err != nil {
		return nil, err
	}
	arg := args[0]
	var str string
	switch arg.(type) {
	case EBytes:
		str = string(arg.(EBytes))
	case EString:
		str = string(arg.(EString))
	default:
		return nil, xerr.Wrap(fmt.Errorf("expect EBytes,EString, got %s", reflect.TypeOf(arg)))
	}
	obj := &MapObject{}
	match := regex.FindStringSubmatch(str)
	if match != nil {
		for i, name := range regex.SubexpNames() {
			if i != 0 && name != "" {
				obj.Set(name, EString(match[i]))
			}
		}
	}
	for _, key := range regex.SubexpNames() {
		if _, err := obj.Get(key); err != nil {
			obj.Set(key, EString(""))
		}
	}
	return obj, nil
}

func (e eStringSubmatch) ToString() EString {
	panic("implement me")
}
