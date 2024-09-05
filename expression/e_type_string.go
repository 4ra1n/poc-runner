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

import (
	"errors"
	"github.com/4ra1n/poc-runner/xerr"
)

type EString string

func (v EString) Keys() []string {
	return []string{}
}

func (v EString) ToString() EString {
	return v
}

func (v EString) Get(name string) (EValue, error) {
	switch name {
	case "bsubmatch":
		return &eStringSubmatch{expr: string(v)}, nil
	case "submatch":
		return &eStringSubmatch{expr: string(v)}, nil
	case "bmatches":
		return &eStringMatches{expr: string(v)}, nil
	case "matches":
		return &eStringMatches{expr: string(v)}, nil
	case "contains":
		return &eStringContains{v: v}, nil
	case "icontains":
		return &eStringContains{v: v, ci: true}, nil
	}
	return nil, xerr.Wrap(errors.New("unsupported " + name))
}
