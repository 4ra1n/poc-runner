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
	"math/rand"
	"reflect"
)

type eFunctionRandomInt struct {
}

func (e eFunctionRandomInt) ToString() EString {
	return "randomInt"
}

func (e eFunctionRandomInt) Call(args []EValue) (EValue, error) {
	if len(args) != 2 {
		return nil, xerr.Wrap(fmt.Errorf("expect two argument, got %d", len(args)))
	}
	minVal, ok := args[0].(EInt)
	if !ok {
		return nil, xerr.Wrap(fmt.Errorf("argument 1 expect EInt, got %s", reflect.TypeOf(args[0])))
	}
	maxVal, ok := args[1].(EInt)
	if !ok {
		return nil, xerr.Wrap(fmt.Errorf("argument 2 expect EInt, got %s", reflect.TypeOf(args[1])))
	}
	return EInt(rand.Intn(int(maxVal-minVal)+1) + 1), nil
}
