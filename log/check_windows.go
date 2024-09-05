//go:build windows

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

package log

import (
	"syscall"
	"unsafe"
)

// Windows 系统修改 ENABLE_VIRTUAL_TERMINAL_PROCESSING
func isWindowsColorSupported() bool {
	var (
		kernel32           = syscall.NewLazyDLL("kernel32.dll")
		procGetConsoleMode = kernel32.NewProc("GetConsoleMode")
		procSetConsoleMode = kernel32.NewProc("SetConsoleMode")
	)
	var mode uint32
	stdoutHandle := uintptr(syscall.Stdout)
	r, _, _ := procGetConsoleMode.Call(stdoutHandle, uintptr(unsafe.Pointer(&mode)))
	if r == 0 {
		return false
	}
	// ENABLE_VIRTUAL_TERMINAL_PROCESSING
	newMode := mode | 0x0004
	r, _, _ = procSetConsoleMode.Call(stdoutHandle, uintptr(newMode))
	if r == 0 {
		return false
	}
	return true
}
