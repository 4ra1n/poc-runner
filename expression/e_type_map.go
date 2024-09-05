package expression

import (
	"fmt"
	"github.com/4ra1n/poc-runner/util"
	"github.com/4ra1n/poc-runner/xerr"
)

type MapObject map[string]EValue

func (m MapObject) Keys() []string {
	return util.Keys(m)
}

func (m MapObject) ToString() EString {
	return "MapObject"
}

func (m MapObject) Set(name string, value EValue) {
	m[name] = value
}

func (m MapObject) Get(name string) (EValue, error) {
	if value, ok := m[name]; ok {
		return value, nil
	}
	return nil, xerr.Wrap(fmt.Errorf("'%s' undefined", name))
}
