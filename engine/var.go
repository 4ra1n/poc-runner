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

package engine

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/4ra1n/poc-runner/expression"
)

func replaceVar(input string, variables map[string]expression.EValue) string {
	var newVar string
	if strings.Contains(input, "{{") &&
		strings.Contains(input, "}}") {
		temp := input
		newVar = parsePath(temp, variables)
	} else {
		newVar = input
	}
	return newVar
}

func parsePath(path string, variables map[string]expression.EValue) string {
	re := regexp.MustCompile(`\{\{\s*(\w+)\s*}}`)
	parsedPath := re.ReplaceAllStringFunc(path, func(match string) string {
		varName := re.FindStringSubmatch(match)[1]
		if val := variables[varName]; val != nil {
			es, ok := val.(expression.EString)
			if ok {
				return string(es)
			}
			ei, ok := val.(expression.EInt)
			if ok {
				return strconv.Itoa(int(ei))
			}
		}
		return match
	})
	return parsedPath
}
