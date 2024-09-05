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
	"testing"
)

func funcD() error {
	return Wrap(errors.New("something went wrong in funcD"))
}

func funcC() error {
	err := funcD()
	if err != nil {
		return Wrap(err, "this is c message")
	}
	return nil
}

func funcB() error {
	err := funcC()
	if err != nil {
		return Wrap(err, "this is b message")
	}
	return nil
}

func funcA() error {
	err := funcB()
	if err != nil {
		return Wrap(err, "this is a message")
	}
	return nil
}

func TestXErr(t *testing.T) {
	err := funcA()
	if err != nil {
		var xe *XError
		if errors.As(err, &xe) {
			xe.PrintStack()
		}
	}
}
