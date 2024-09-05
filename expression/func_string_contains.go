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
	return "contains"
}
