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

type POC struct {
	Env        *expression.Environment
	Context    *PocContext
	Caches     *GlobalCache
	Target     string
	Name       string
	Transport  string
	Set        *Map[string, *Expr]
	Payload    *Payload
	Rules      *Map[string, *Rule]
	Expression *Expr
	Detail     *Map[string, string]
}

func (p *POC) DoReq(req *client.TheRequest) (*client.TheResponse, error) {
	return p.Context.Client.DoReq(req)
}

func (p *POC) Locals() map[string]expression.EValue {
	return p.Context.Local.ToMap()
}

func (p *POC) SetLocal(key string, value expression.EValue) {
	p.Context.Local.Set(key, value)
}

func (p *POC) GetLocal(key string) expression.EValue {
	return p.Context.Local.Get(key)
}

func (p *POC) SetReq(rule string, req *client.TheRequest) {
	p.Context.AllRequests.Set(rule, req)
}

func (p *POC) GetReq(rule string) *client.TheRequest {
	return p.Context.AllRequests.Get(rule)
}

func (p *POC) SetResp(rule string, resp *client.TheResponse) {
	p.Context.AllResponses.Set(rule, resp)
}

func (p *POC) GetResp(rule string) *client.TheResponse {
	return p.Context.AllResponses.Get(rule)
}

func (p *POC) SetResult(rule string, result bool) {
	p.Context.AllResults.Set(rule, result)
}

func (p *POC) GetResult(rule string) bool {
	return p.Context.AllResults.Get(rule)
}

func (p *POC) GetAllResults() map[string]bool {
	return p.Context.AllResults.ToMap()
}
