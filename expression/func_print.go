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
	"strconv"
	"strings"

	"github.com/4ra1n/poc-runner/log"
	"github.com/4ra1n/poc-runner/xerr"
)

type eFunctionPrint struct {
}

func (e *eFunctionPrint) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return EBool(false), xerr.Wrap(fmt.Errorf("expect one argument, got %d", len(args)))
	}
	arg := args[0]
	var output string

	switch val := arg.(type) {
	case EBool:
		// 处理布尔类型
		if val {
			output = "true"
		} else {
			output = "false"
		}
	case EString:
		// 处理字符串类型
		output = string(val)
	case EInt:
		// 处理整数类型
		output = strconv.Itoa(int(val))
	case EFloat:
		// 处理浮点数类型
		output = strconv.FormatFloat(float64(val), 'E', -1, 32)
	case EBytes:
		// 处理字节数组类型
		output = hex.EncodeToString(val)
	case MapObject:
		// 处理 MapObject 类型
		output = fmt.Sprintf("MAP OBJECT [%s]", strings.Join(val.Keys(), ","))
	default:
		// 默认调用 ToString
		output = string(arg.ToString())
	}

	log.BluePrintf("[YAML PRINT] %s\n", output)
	return EBool(true), nil
}

func (e *eFunctionPrint) ToString() EString {
	return "print"
}
