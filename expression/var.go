/*
 * poc-runner project
 * Copyright (C) 2024 4ra1n
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package expression

import (
	"fmt"
	"strings"

	"github.com/4ra1n/poc-runner/xerr"
)

type Vars interface {
	GetValue(env *Environment, key string) (EValue, error)
}

type MapVars map[string]EValue

func (m MapVars) GetValue(_ *Environment, key string) (EValue, error) {
	if value, ok := m[key]; ok {
		return value, nil
	}
	return nil, xerr.Wrap(fmt.Errorf("'%s' undefined", key))
}

func (m MapVars) SetValue(key string, value EValue) {
	m[key] = value
}

type combinedVars []Vars

func (cv combinedVars) GetValue(env *Environment, key string) (EValue, error) {
	maxLen := len(cv) - 1
	for i := maxLen; i >= 0; i-- {
		item := cv[i]
		if item == nil {
			continue
		}
		value, err := item.GetValue(env, key)
		if err == nil {
			return value, nil
		}
		if i == 0 || !strings.Contains(err.Error(), "undefined") {
			return nil, xerr.Wrap(err)
		}
	}
	return nil, xerr.Wrap(fmt.Errorf("'%s' undefined", key))
}
