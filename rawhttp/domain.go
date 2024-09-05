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

package rawhttp

import (
	"errors"
	"net"

	"github.com/4ra1n/poc-runner/xerr"
)

const (
	defaultDomainIP = "127.0.0.1"
)

func resolveDomain(domain string) (string, error) {
	addresses, err := net.LookupIP(domain)
	if err != nil {
		return defaultDomainIP, xerr.Wrap(err)
	}
	if len(addresses) == 0 {
		return defaultDomainIP, xerr.Wrap(errors.New("no ip for domain " + domain))
	}
	return addresses[0].String(), nil
}
