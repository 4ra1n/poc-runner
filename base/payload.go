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

package base

import (
	"regexp"
	"strconv"
)

// Payload
// XRAY 的 PAYLOAD 语法后续考虑支持
type Payload struct {
	values *Map[string, *Map[string, *Expr]]
}

var re = regexp.MustCompile(`(\D+)(\d*)`)

func naturalCompare(a, b string) bool {
	partsA := re.FindStringSubmatch(a)
	partsB := re.FindStringSubmatch(b)
	if partsA[1] != partsB[1] {
		return partsA[1] < partsB[1]
	}
	if len(partsA[2]) == 0 || len(partsB[2]) == 0 {
		return a < b
	}
	numA, _ := strconv.Atoi(partsA[2])
	numB, _ := strconv.Atoi(partsB[2])
	return numA < numB
}

func NewPayload() *Payload {
	return &Payload{
		values: NewMap[string, *Map[string, *Expr]](),
	}
}

func (p *Payload) Set(key string, ma *Map[string, *Expr]) {
	p.values.Set(key, ma)
}

func (p *Payload) Get() *Map[string, *Map[string, *Expr]] {
	p.values.Sort(func(a, b string) bool {
		return naturalCompare(a, b)
	})
	return p.values
}

func (p *Payload) Available() bool {
	return p.values.Length() > 0
}
