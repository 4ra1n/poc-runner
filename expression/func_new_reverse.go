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

package expression

import "github.com/4ra1n/poc-runner/xerr"

type eFunctionNewReverse struct {
}

func (e *eFunctionNewReverse) Call(_ []EValue) (EValue, error) {
	rev, err := NewReverse()
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	return rev, nil
}

func (e *eFunctionNewReverse) ToString() EString {
	return "newReverse"
}
