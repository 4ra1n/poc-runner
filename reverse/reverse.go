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

package reverse

import (
	"errors"
	"fmt"
	"strings"

	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/xerr"
)

var (
	Type     string
	Instance Reverse
)

const (
	DNSLogCN   = "dnslog.cn"
	InteractSH = "interact.sh"
)

// Reverse
// 通用反连接口
// 目前是 dnslog.cn
// 后续可以拓展
type Reverse interface {
	// GetUrl
	// HTTP 反连
	GetUrl() string
	// GetRmi
	// RMI 反连
	GetRmi() string
	// GetLdap
	// LDAP 反连
	GetLdap() string
	// GetDNS
	// DNS 反连
	GetDNS() string
	// Wait
	// 等待
	Wait(int) bool
	// Close
	// 关闭
	Close()
}

func NewReverse(c *client.HttpClient) (Reverse, error) {
	if Type == DNSLogCN {
		resp, err := c.DoReq(&client.TheRequest{
			Target:         dnsLogCnUrl,
			Method:         "GET",
			Path:           "/getdomain.php",
			FollowRedirect: false,
			Body:           "",
			Headers:        make(map[string]string),
			IsFromPoC:      false,
		})
		if err != nil {
			return nil, xerr.Wrap(err)
		}
		r := randUpper(8)
		newDomain := strings.TrimSpace(string(resp.Body))
		if newDomain == "" {
			return nil, xerr.Wrap(errors.New("dnslog.cn error"))
		}
		session, ok := resp.Headers["Set-Cookie"]
		if !ok {
			return nil, xerr.Wrap(errors.New("dnslog.cn session error"))
		}
		return &DnsLogCn{
			c:         c,
			baseUrl:   dnsLogCnUrl,
			newDomain: fmt.Sprintf("%s.%s", r, newDomain),
			session:   session[0],
		}, nil
	}
	if Type == InteractSH {
		return NewInteract(c, "")
	}
	return nil, xerr.Wrap(errors.New("reverse type error"))
}
