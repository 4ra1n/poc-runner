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
	"errors"
	"fmt"
	"strings"
)

func Wrap(err error, message ...string) error {
	var finalMessage string
	if len(message) > 0 && message[0] != "" {
		finalMessage = fmt.Sprintf("ERR MSG -> %s\n%v", message[0], err)
	} else {
		if strings.HasPrefix(err.Error(), "ERR ROOT MSG ->") {
			finalMessage = err.Error()
		} else {
			finalMessage = fmt.Sprintf("ERR ROOT MSG -> %v", err)
		}
	}
	var xErr *XError
	if errors.As(err, &xErr) {
		return &XError{
			message: finalMessage,
			stack:   xErr.stack,
			err:     err,
		}
	}
	return &XError{
		message: finalMessage,
		stack:   captureStackTrace(3),
		err:     err,
	}
}
