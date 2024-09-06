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

	"github.com/4ra1n/poc-runner/xerr"
)

type eFunctionRandomAlpha struct {
	upper bool
}

func (e eFunctionRandomAlpha) ToString() EString {
	if e.upper {
		return "randomUppercase"
	} else {
		return "randomLowercase"
	}
}

func (e eFunctionRandomAlpha) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, xerr.Wrap(fmt.Errorf("expect one argument, got %d", len(args)))
	}
	n, ok := args[0].(EInt)
	if !ok {
		return nil, xerr.Wrap(fmt.Errorf("argument 1 expect EInt, got %s", reflect.TypeOf(args[0])))
	}
	var value string
	if e.upper {
		value = randStringRunes(int(n), []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	} else {
		value = randStringRunes(int(n), []rune("abcdefghijklmnopqrstuvwxyz"))
	}
	return EString(value), nil
}
