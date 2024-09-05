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
	"fmt"

	"github.com/4ra1n/poc-runner/base"

	"github.com/4ra1n/poc-runner/expression"
	"github.com/4ra1n/poc-runner/xerr"
)

type Rule map[string]expression.EFunction

func (r Rule) GetValue(_ *expression.Environment, key string) (expression.EValue, error) {
	if rule, ok := r[key]; ok {
		return rule, nil
	}
	return nil, xerr.Wrap(fmt.Errorf("'%s' undefined", key))
}

type RuleFunction struct {
	poc  *base.POC
	key  string
	rule *base.Rule
}

func (f *RuleFunction) ToString() expression.EString {
	return "RuleFunction"
}

func (f *RuleFunction) Call(_ []expression.EValue) (expression.EValue, error) {
	return execRule(f.poc, f.key, f.rule)
}
