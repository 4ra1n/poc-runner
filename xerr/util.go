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
	"runtime"
	"strings"
)

func trimFilePath(fullPath string) string {
	lastSlash := strings.LastIndex(fullPath, "/")
	if lastSlash >= 0 {
		return fullPath[lastSlash+1:]
	}
	return fullPath
}

func shouldIncludeFrame(filePath string) bool {
	goRoot := runtime.GOROOT()
	goRoot = strings.ReplaceAll(goRoot, "\\\\", "\\")
	goRoot = strings.ReplaceAll(goRoot, "\\", "/")
	return !strings.HasPrefix(filePath, goRoot)
}
