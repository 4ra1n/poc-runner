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
	"fmt"
	"net/url"
	"time"

	"github.com/4ra1n/poc-runner/xerr"
)

type Request struct {
	ID string
	// 目标信息
	Protocol       string
	Domain         string
	IP             string
	Port           string
	Path           string
	Method         string
	Headers        map[string]string
	FollowRedirect bool
	// 请求处理信息
	RawBody       []byte
	RawHeader     []byte
	RawRequest    []byte
	redirectCount int
}

func NewRequest(target string, method string) (*Request, error) {
	req := &Request{
		ID:            fmt.Sprintf("%d", time.Now().Unix()),
		Method:        method,
		Headers:       make(map[string]string),
		redirectCount: 0,
	}
	parsedURL, err := url.Parse(target)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	req.Protocol = parsedURL.Scheme
	host := parsedURL.Hostname()
	hostType := determineHostType(host)
	if hostType == IPv4 || hostType == IPv6 {
		req.Domain = ""
		req.IP = host
	} else if hostType == Domain {
		var ip string
		ip, err = resolveDomain(host)
		if err != nil {
			return nil, xerr.Wrap(err)
		}
		req.Domain = host
		req.IP = ip
	}
	req.Port = parsedURL.Port()
	if req.Port == none {
		if req.Protocol == httpsProtocol {
			req.Port = httpsDefaultPort
		} else if req.Protocol == httpProtocol {
			req.Port = httpDefaultPort
		}
	}
	req.Path = parsedURL.RequestURI()
	return req, nil
}

func (r *Request) SetHeader(key string, value string) {
	r.Headers[key] = value
}

func (r *Request) SetBody(body string) {
	r.RawBody = []byte(body)
}

func (r *Request) SetBytesBody(body []byte) {
	r.RawBody = body
}

func (r *Request) SetFollowRedirect(redirect bool) {
	r.FollowRedirect = redirect
}
