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
	"fmt"
	"runtime"
	"strings"
)

func captureStackTrace(skip int) string {
	var sb strings.Builder
	sb.WriteString("\n")
	pc := make([]uintptr, 32)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		if shouldIncludeFrame(frame.File) {
			relativeFile := trimFilePath(frame.File)
			sb.WriteString(fmt.Sprintf("  at %s (%s:%d)\n",
				frame.Function, relativeFile, frame.Line))
		}
		if !more {
			break
		}
	}
	return sb.String()
}
