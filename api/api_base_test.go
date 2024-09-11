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

package api

import (
	"context"
	"fmt"
	"testing"
)

var poc = `name: poc-yaml-test
transport: http
set:
  rand: randomInt(200000000, 210000000)
rules:
  r0:
    request:
      cache: true
      method: POST
      path: /test
      headers:
        Content-Type: text/xml
      body: test
      follow_redirects: false
    expression: response.status == 404
expression: r0()`

func TestAPI(t *testing.T) {
	ctx := context.Background()
	runner, err := NewPocRunner(ctx)
	if err != nil {
		return
	}
	report, err := runner.Run([]byte(poc), "https://example.com")
	if err != nil {
		return
	}
	fmt.Println(report)
}
