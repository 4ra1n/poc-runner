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

package engine

import (
	"errors"
	"fmt"

	"github.com/4ra1n/poc-runner/expression"
	"github.com/4ra1n/poc-runner/xerr"
)

type RespValue struct {
	Status  expression.EInt
	Body    expression.EValue
	Headers *RespHeaderValue
}

func (r *RespValue) ToString() expression.EString {
	panic("implement me")
}

func (r *RespValue) Get(name string) (expression.EValue, error) {
	switch name {
	case "status":
		return r.Status, nil
	case "headers":
		return r.Headers, nil
	case "content_type":
		ctHeader, err := r.Headers.Get("Content-Type")
		if err != nil {
			return expression.EString(""), err
		} else {
			return ctHeader, nil
		}
	case "body":
		return r.Body, nil
	case "body_string":
		bodyBytes, ok := r.Body.(expression.EBytes)
		if !ok {
			return nil, xerr.Wrap(fmt.Errorf("body must be bytes"))
		}
		return expression.EString(bodyBytes), nil
	default:
		return nil, xerr.Wrap(errors.New("not support: " + name))
	}
}

func (r *RespValue) Keys() []string {
	panic("implement me")
}
