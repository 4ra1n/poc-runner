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
	"fmt"
	"html"
	"time"

	"github.com/4ra1n/poc-runner/xerr"
)

func BuildHTMLReport(result *Result) (string, error) {
	var rulesHtml string
	for k, v := range result.Details {
		req, err := base64.StdEncoding.DecodeString(v.Req)
		if err != nil {
			return "", xerr.Wrap(err)
		}
		resp, err := base64.StdEncoding.DecodeString(v.Resp)

		reqStr := html.EscapeString(string(req))
		respStr := html.EscapeString(string(resp))

		ruleHtml := fmt.Sprintf(ruleTemplate, k, reqStr, respStr, v.Result)
		rulesHtml += ruleHtml
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	dataHtml := fmt.Sprintf(dataTemplate, result.PocName, result.Target, t, rulesHtml)
	allHtml := fmt.Sprintf(baseTemplate, dataHtml)
	return allHtml, nil
}
