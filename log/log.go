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
	"log"
	"os"
	"sync"
)

var (
	errorLog *log.Logger
	infoLog  *log.Logger
	debugLog *log.Logger
	warnLog  *log.Logger
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex

	Error  func(...interface{})
	Errorf func(string, ...interface{})
	Warn   func(...interface{})
	Warnf  func(string, ...interface{})
	Info   func(...interface{})
	Infof  func(string, ...interface{})
	Debug  func(...interface{})
	Debugf func(string, ...interface{})
)

func init() {
	refresh()
}

func refresh() {
	errorLog = log.New(os.Stdout, Red+"[ERR]"+Reset+" ", log.Ltime|log.Lshortfile)
	infoLog = log.New(os.Stdout, Green+"[INF]"+Reset+" ", log.Ltime|log.Lshortfile)
	debugLog = log.New(os.Stdout, Blue+"[DBG]"+Reset+" ", log.Ltime|log.Lshortfile)
	warnLog = log.New(os.Stdout, Yellow+"[WRN]"+Reset+" ", log.Ltime|log.Lshortfile)
	loggers = []*log.Logger{errorLog, infoLog}
	mu = sync.Mutex{}

	Error = errorLog.Println
	Errorf = errorLog.Printf
	Warn = warnLog.Println
	Warnf = warnLog.Printf
	Info = infoLog.Println
	Infof = infoLog.Printf
	Debug = debugLog.Println
	Debugf = debugLog.Printf
}
