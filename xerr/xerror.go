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

package xerr

import (
	"strings"

	"github.com/4ra1n/poc-runner/log"
)

// XError
// 自定义错误类型
// 可以输出堆栈
// 本质是为了让错误变成 JAVA 的感觉
type XError struct {
	message string
	stack   string
	err     error
}

func (e *XError) Error() string {
	return e.message
}
func (e *XError) Unwrap() error {
	return e.err
}

func (e *XError) PrintStack() {
	maxLineLength := e.calculateMaxLineLength()
	errorText := " ERROR "
	errorTextLength := len(errorText)
	padding := (maxLineLength + 4 - errorTextLength) / 2
	startBorder := strings.Repeat("#", padding) + errorText +
		strings.Repeat("#", maxLineLength+4-padding-errorTextLength)
	endBorder := strings.Repeat("#", maxLineLength+4)
	log.RedPrintln(startBorder)
	log.BluePrintf("%s", e.message)
	log.YellowPrintf("%s", e.stack)
	log.RedPrintln(endBorder)
}

func (e *XError) calculateMaxLineLength() int {
	maxLength := 0
	allLines := append(strings.Split(e.message, "\n"), strings.Split(e.stack, "\n")...)
	for _, line := range allLines {
		lineLength := len(line)
		if lineLength > maxLength {
			maxLength = lineLength
		}
	}
	return maxLength
}
