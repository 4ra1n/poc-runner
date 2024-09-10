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

package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Picker struct {
	choices []string
	used    map[int]bool
}

func NewPicker(choices []string) *Picker {
	return &Picker{
		choices: choices,
		used:    make(map[int]bool),
	}
}

func (p *Picker) RandomPick() (string, bool) {
	if len(p.used) == len(p.choices) {
		return "", false
	}
	var index int
	for {
		index = rand.Intn(len(p.choices))
		if !p.used[index] {
			break
		}
	}
	p.used[index] = true
	return p.choices[index], true
}
