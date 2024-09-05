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
	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/expression"
)

// PocContext
// 单个 POC 的 Context
// 这里的很多信息保存暂时未用到
// 考虑到未来做统计需要
type PocContext struct {
	// 包装后的 HTTP CLIENT
	// 不建议直接使用 RAW CLIENT
	Client *client.HttpClient
	// 适用于当前 POC 运行环境的本地变量
	Local *Map[string, expression.EValue]
	// 记录每一个 rule 的 request 信息
	AllRequests *Map[string, *client.TheRequest]
	// 记录每一个 rule 的 response 信息
	AllResponses *Map[string, *client.TheResponse]
	// 记录每一个 rule 的执行结果
	AllResults *Map[string, bool]
}

func NewPocContext(c *client.HttpClient) (*PocContext, error) {
	return &PocContext{
		Client:       c,
		Local:        NewMap[string, expression.EValue](),
		AllRequests:  NewMap[string, *client.TheRequest](),
		AllResponses: NewMap[string, *client.TheResponse](),
		AllResults:   NewMap[string, bool](),
	}, nil
}
