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
	"fmt"
	"os"
	"runtime"
	"strings"
)

var (
	Red    string
	Green  string
	Yellow string
	Blue   string
	Reset  string
)

func DisableColor() {
	Red = ""
	Green = ""
	Yellow = ""
	Blue = ""
	Reset = ""
	refresh()
}

func EnableColor() {
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Blue = "\033[34m"
	Reset = "\033[0m"
	refresh()
}

func IsSupported() bool {
	if runtime.GOOS != "windows" {
		term, ok := os.LookupEnv("TERM")
		if !ok || term == "dumb" {
			return false
		}
		return true
	}
	// GOLAND IDE
	env, ok := os.LookupEnv("IDEA_INITIAL_DIRECTORY")
	if ok {
		env = strings.ToUpper(env)
		if strings.Contains(env, "GOLAND") &&
			strings.Contains(env, "JETBRAINS") {
			return true
		}
	}
	env, ok = os.LookupEnv("IJ_RESTARTER_LOG")
	if ok {
		env = strings.ToUpper(env)
		if strings.Contains(env, "GOLAND") &&
			strings.Contains(env, "JETBRAINS") {
			return true
		}
	}
	// enable color terminal in windows
	return isWindowsColorSupported()
}

func RedPrintln(a ...interface{}) {
	fmt.Println(Red + fmt.Sprint(a...) + Reset)
}

func GreenPrintln(a ...interface{}) {
	fmt.Println(Green + fmt.Sprint(a...) + Reset)
}

func YellowPrintln(a ...interface{}) {
	fmt.Println(Yellow + fmt.Sprint(a...) + Reset)
}

func BluePrintln(a ...interface{}) {
	fmt.Println(Blue + fmt.Sprint(a...) + Reset)
}

func RedPrintf(format string, a ...interface{}) {
	fmt.Printf(Red+format+Reset, a...)
}

func GreenPrintf(format string, a ...interface{}) {
	fmt.Printf(Green+format+Reset, a...)
}

func YellowPrintf(format string, a ...interface{}) {
	fmt.Printf(Yellow+format+Reset, a...)
}

func BluePrintf(format string, a ...interface{}) {
	fmt.Printf(Blue+format+Reset, a...)
}

func RedSprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(Red+format+Reset, a...)
}

func GreenSprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(Green+format+Reset, a...)
}

func YellowSprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(Yellow+format+Reset, a...)
}

func BlueSprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(Blue+format+Reset, a...)
}
