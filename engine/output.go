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
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/4ra1n/poc-runner/base"
	"github.com/4ra1n/poc-runner/xerr"
)

const (
	stdoutOut = "stdout"
	txtOut    = "txt"
	htmlOut   = "html"
	jsonOut   = "json"
)

type ResultDetail struct {
	Req    string `json:"req"`
	Resp   string `json:"resp"`
	Result bool   `json:"result"`
}

// Result
// 扫描结果
type Result struct {
	PocName string                   `json:"poc_name"`
	Target  string                   `json:"target"`
	Details map[string]*ResultDetail `json:"details"`
}

func newResultInternal(poc *base.POC) *Result {
	result := &Result{
		PocName: poc.Name,
		Target:  poc.Target,
		Details: make(map[string]*ResultDetail),
	}
	ruleKeys := poc.Context.AllRequests.Keys()
	for _, k := range ruleKeys {
		req := poc.Context.AllRequests.Get(k)
		resp := poc.Context.AllResponses.Get(k)
		res := poc.Context.AllResults.Get(k)
		detail := &ResultDetail{
			Req:    base64.StdEncoding.EncodeToString(req.All),
			Resp:   base64.StdEncoding.EncodeToString(resp.All),
			Result: res,
		}
		result.Details[k] = detail
	}
	return result
}

func NewResultJson(poc *base.POC) (string, error) {
	res := newResultInternal(poc)
	data, err := json.Marshal(res)
	if err != nil {
		return "", xerr.Wrap(err)
	}
	return string(data), nil
}

func NewResultTxt(poc *base.POC) (string, error) {
	res := newResultInternal(poc)
	buf := &strings.Builder{}
	buf.WriteString(poc.Name)
	buf.WriteString("\n")
	buf.WriteString(poc.Target)
	buf.WriteString("\n")
	for k, v := range res.Details {
		buf.WriteString("rule:")
		buf.WriteString(k)
		buf.WriteString("\n")
		buf.WriteString("request:")
		buf.WriteString("\n")
		buf.WriteString(v.Req)
		buf.WriteString("\n")
		buf.WriteString("response:")
		buf.WriteString("\n")
		buf.WriteString(v.Resp)
		buf.WriteString("\n")
	}
	return buf.String(), nil
}

func NewResultHTML(poc *base.POC) (string, error) {
	res := newResultInternal(poc)
	data, err := BuildHTMLReport(res)
	if err != nil {
		return "", xerr.Wrap(err)
	}
	return data, nil
}