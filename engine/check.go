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
	"errors"
	"fmt"
	"strings"

	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/log"
	"github.com/4ra1n/poc-runner/xerr"
)

func checkTarget(target string) (bool, string, error) {
	c := client.Instance
	target = strings.TrimSpace(target)
	if !strings.HasPrefix(target, "http") {
		target = fmt.Sprintf("http://%s", target)
	}
	resp, err := c.Get(target)
	if err != nil {
		return false, "", xerr.Wrap(err)
	}
	if resp != nil {
		log.Infof("check target response: %d", resp.Code)
		return true, target, err
	} else {
		return false, "", xerr.Wrap(errors.New("response is nil"))
	}
}
