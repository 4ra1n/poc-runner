package expression

import (
	"errors"
	"github.com/4ra1n/poc-runner/xerr"
)

type EString string

func (v EString) Keys() []string {
	panic("implement me")
}

func (v EString) ToString() EString {
	return v
}

func (v EString) Get(name string) (EValue, error) {
	switch name {
	case "bsubmatch":
		return &eStringSubmatch{expr: string(v)}, nil
	case "submatch":
		return &eStringSubmatch{expr: string(v)}, nil
	case "bmatches":
		return &eStringMatches{expr: string(v)}, nil
	case "matches":
		return &eStringMatches{expr: string(v)}, nil
	case "contains":
		return &eStringContains{v: v}, nil
	case "icontains":
		return &eStringContains{v: v, ci: true}, nil
	}
	return nil, xerr.Wrap(errors.New("unsupported " + name))
}
