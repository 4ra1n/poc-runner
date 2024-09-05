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
	"strings"

	"github.com/4ra1n/poc-runner/expression"
)

type RespHeaderValue struct {
	Headers map[string][]string
}

func (r *RespHeaderValue) Get(name string) (expression.EValue, error) {
	for k, v := range r.Headers {
		all := &strings.Builder{}
		// 处理多个头的情况
		// 最常见的多头是 Set-Cookie
		// response.headers["Set-Cookie"].contains("ANY")
		for _, vv := range v {
			all.WriteString(vv)
			all.WriteString(" ")
		}
		valueStr := strings.TrimSpace(all.String())
		// 目前没有 POC 用到 Set-Cookie 和 == 判断
		// 所以拼接操作是不存在逻辑问题的
		if strings.ToLower(k) == strings.ToLower(name) {
			return expression.EString(valueStr), nil
		}
	}
	return expression.EString(""), nil
}

func (r *RespHeaderValue) Keys() []string {
	return []string{}
}

func (r *RespHeaderValue) ToString() expression.EString {
	return "headers"
}
