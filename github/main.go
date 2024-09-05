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

package main

import (
	"os"
	"strings"

	"github/color"
	"github/util"
)

var (
	useProxy   = false
	socksProxy = "127.0.0.1:10808"
)

func usage() {
	color.RedPrintln("USAGE: go run .\\github\\main.go proxy|no-proxy")
}

func main() {
	if !color.IsSupported() {
		color.DisableColor()
	}
	args := os.Args
	if len(args) != 2 {
		usage()
		return
	}
	if args[1] == "proxy" {
		color.GreenPrintln("OPTIONS: USE PROXY")
		useProxy = true
	} else if args[1] == "no-proxy" {
		color.GreenPrintln("OPTIONS: NOT USE PROXY")
		useProxy = false
	} else {
		usage()
		return
	}

	tokenBytes, err := os.ReadFile("token.txt")
	if err != nil {
		color.RedPrintln(err)
		return
	}

	token := strings.TrimSpace(string(tokenBytes))
	repoOwner := "4ra1n"
	repoName := "poc-runner"

	color.GreenPrintln("--- START GITHUB CACHES CLEAN ---")
	util.CleanCache(token, useProxy, socksProxy, repoOwner, repoName)

	color.GreenPrintln("--- START GITHUB ACTION BUILDS CLEAN ---")
	util.CleanAction(token, useProxy, socksProxy, repoOwner, repoName, "build.yml")

	color.GreenPrintln("--- START GITHUB ACTION LEAKS CLEAN ---")
	util.CleanAction(token, useProxy, socksProxy, repoOwner, repoName, "leak.yml")

	color.GreenPrintln("--- START GITHUB ACTION CHECK CLEAN ---")
	util.CleanAction(token, useProxy, socksProxy, repoOwner, repoName, "check.yml")

	color.GreenPrintln("--- START GITHUB ACTION TRUFFLEHOG CLEAN ---")
	util.CleanAction(token, useProxy, socksProxy, repoOwner, repoName, "trufflehog.yml")

	color.GreenPrintln("GITHUB CLEAN FINISH")
}
