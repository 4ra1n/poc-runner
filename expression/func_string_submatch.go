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
	"reflect"
	"regexp"

	"github.com/4ra1n/poc-runner/xerr"
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
	return "submatch"
}
