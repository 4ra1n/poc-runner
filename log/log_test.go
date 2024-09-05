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
	"testing"
)

func TestSetLevel(t *testing.T) {
	fmt.Println("test debug level")
	SetLevel(DebugLevel)
	Debug("test info")
	Debugf("test infof: %s", "test")
	Info("test error")
	Infof("test errorf: %s", "test")
	Warn("test warn")
	Warnf("test warnf: %s", "test")
	Error("test error")
	Errorf("test errorf: %s", "test")

	fmt.Println("test info level")
	SetLevel(InfoLevel)
	Debug("test info")
	Debugf("test infof: %s", "test")
	Info("test error")
	Infof("test errorf: %s", "test")
	Warn("test warn")
	Warnf("test warnf: %s", "test")
	Error("test error")
	Errorf("test errorf: %s", "test")

	fmt.Println("test warn level")
	SetLevel(WarnLevel)
	Debug("test info")
	Debugf("test infof: %s", "test")
	Info("test error")
	Infof("test errorf: %s", "test")
	Warn("test warn")
	Warnf("test warnf: %s", "test")
	Error("test error")
	Errorf("test errorf: %s", "test")

	fmt.Println("test error level")
	SetLevel(ErrorLevel)
	Debug("test info")
	Debugf("test infof: %s", "test")
	Info("test error")
	Infof("test errorf: %s", "test")
	Warn("test warn")
	Warnf("test warnf: %s", "test")
	Error("test error")
	Errorf("test errorf: %s", "test")

	fmt.Println("test disabled level")
	SetLevel(Disabled)
	Debug("test info")
	Debugf("test infof: %s", "test")
	Info("test error")
	Infof("test errorf: %s", "test")
	Warn("test warn")
	Warnf("test warnf: %s", "test")
	Error("test error")
	Errorf("test errorf: %s", "test")
}
