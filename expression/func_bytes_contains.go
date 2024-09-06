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
	"bytes"
	"fmt"
	"reflect"

	"github.com/4ra1n/poc-runner/xerr"
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
	return "bcontains"
}
