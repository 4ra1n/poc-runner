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
	"strconv"
	"strings"
	"time"

	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/rawhttp"
	"github.com/4ra1n/poc-runner/xerr"
)

func setHttpClient() error {
	proxy = strings.TrimSpace(proxy)
	timeout = strings.TrimSpace(timeout)
	if proxy == "" {
		proxy = rawhttp.DefaultNoProxy
	}
	var timeoutVal time.Duration
	if timeout == "" {
		timeoutVal = rawhttp.DefaultTimeout
	} else {
		t, err := strconv.Atoi(timeout)
		if err != nil {
			return xerr.Wrap(err)
		}
		if t < 1 || t > 61 {
			return xerr.Wrap(errors.New("timeout invalid"))
		}
		timeoutVal = time.Duration(t) * time.Second
	}
	c, err := client.NewHttpClient(proxy, timeoutVal, debug)
	if err != nil {
		return xerr.Wrap(err)
	}
	client.Instance = c
	return nil
}
