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
	"fmt"

	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/reverse"
	"github.com/4ra1n/poc-runner/xerr"
)

type Reverse struct {
	rev reverse.Reverse
}

func NewReverse() (*Reverse, error) {
	r, err := reverse.NewReverse(client.Instance)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	reverse.Instance = r
	return &Reverse{
		rev: r,
	}, nil
}

func (r *Reverse) ToString() EString {
	return EString(fmt.Sprintf("reverse-%p", r.rev))
}

func (r *Reverse) Get(name string) (EValue, error) {
	switch name {
	case "url":
		return EString(r.rev.GetUrl()), nil
	case "rmi":
		return EString(r.rev.GetRmi()), nil
	case "ldap":
		return EString(r.rev.GetLdap()), nil
	case "domain":
		return EString(r.rev.GetDNS()), nil
	case "wait":
		return &eFunctionReverseWait{r.rev}, nil
	default:
		return nil, xerr.Wrap(errors.New("not support: " + name))
	}
}

func (r *Reverse) Keys() []string {
	return []string{}
}
