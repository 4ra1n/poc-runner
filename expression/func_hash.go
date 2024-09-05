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
	"encoding/hex"
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
	"hash"
	"reflect"
)

type eFunctionHash struct {
	h func() hash.Hash
}

func (e eFunctionHash) ToString() EString {
	return "hash"
}

func (e eFunctionHash) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, xerr.Wrap(fmt.Errorf("expect one argument, got %d", len(args)))
	}
	arg := args[0]
	h := e.h()
	h.Write([]byte(arg.ToString()))
	switch arg.(type) {
	case EString:
		return EString(hex.EncodeToString(h.Sum(nil))), nil
	case EBytes:
		return EString(hex.EncodeToString(h.Sum(nil))), nil
	default:
		return nil, xerr.Wrap(fmt.Errorf("expect EString,EBytes, got %s", reflect.TypeOf(arg)))
	}
}
