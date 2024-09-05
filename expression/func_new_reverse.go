package expression

import "github.com/4ra1n/poc-runner/xerr"

type eFunctionNewReverse struct {
}

func (e *eFunctionNewReverse) Call(_ []EValue) (EValue, error) {
	rev, err := NewReverse()
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	return rev, nil
}

func (e *eFunctionNewReverse) ToString() EString {
	return "newReverse"
}
