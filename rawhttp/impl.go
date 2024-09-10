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
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/4ra1n/poc-runner/log"
	"github.com/4ra1n/poc-runner/xerr"
)

func (c *HTTPClient) buildReqRaw(req *Request) []byte {
	buf := bytes.Buffer{}

	// 不能使用标准库的方法 会重排序
	// 实际的 PATH 应该仅编码 不改顺序
	var finalPath string
	if !strings.Contains(req.Path, querySep) {
		finalPath = req.Path
	} else {
		temp := strings.Split(req.Path, querySep)
		path := temp[0]
		if len(temp) < 2 {
			finalPath = path + querySep
		} else {
			var q string
			if strings.Contains(temp[1], equalSep) {
				q = encodeParams(temp[1], false)
			} else {
				// FIX CVE-2021-40438
				q = temp[1]
			}
			finalPath = path + querySep + q
		}
	}

	// FIX FINAL PATH
	finalPath = strings.TrimLeft(finalPath, "/")
	finalPath = "/" + finalPath
	req.Path = finalPath

	buf.WriteString(req.Method + space + finalPath + space + version + lineSep)

	var headers []string
	if _, ok := req.Headers[hostHeaderStd]; !ok {
		headers = append(headers, hostHeaderStd+": "+getTarget(req))
	}
	if _, ok := req.Headers[acceptHeaderStd]; !ok {
		headers = append(headers, acceptHeaderStd+": "+acceptVal)
	}
	if _, ok := req.Headers[connHeaderStd]; !ok {
		headers = append(headers, connHeaderStd+": close")
	}
	if _, ok := req.Headers[uaHeaderStd]; !ok {
		headers = append(headers, uaHeaderStd+": "+defaultUA)
	}

	// FIX CT HEADER
	if _, ok := req.Headers[ctHeaderStd]; !ok {
		// DEFAULT FORM WHEN BODY NOT NIL
		if req.RawBody != nil && len(req.RawBody) > 0 {
			headers = append(headers, ctHeaderStd+": "+ctFormVal)
		}
	}

	for k, v := range req.Headers {
		if strings.ToLower(k) == clHeader {
			// FIX BUG
			// NOT SUPPORT CL HEADER
			continue
		}
		headers = append(headers, fmt.Sprintf("%s: %s", k, v))
	}

	for _, header := range headers {
		buf.WriteString(header + lineSep)
	}

	if req.RawBody == nil || len(req.RawBody) < 1 {
		buf.WriteString(lineSep)
	} else {
		if req.Method == MethodGet {
			log.Warn("get method should not contains body")
		}

		contentType := req.Headers[ctHeaderStd]
		if strings.Contains(contentType, ctFormVal) {
			bodyStr := string(req.RawBody)
			// FIX SPACE
			bodyStr = strings.TrimSpace(bodyStr)
			// FIX ES CVE-2015-5531
			if !strings.HasPrefix(bodyStr, "{") &&
				!strings.HasPrefix(bodyStr, "<") {
				encodedBody := encodeParams(string(req.RawBody), true)
				req.RawBody = []byte(encodedBody)
			}
		}

		bodyLength := strconv.Itoa(len(req.RawBody))
		buf.WriteString(clHeaderStd + ": " + bodyLength + lineSep + lineSep)
		buf.Write(req.RawBody)
	}

	return buf.Bytes()
}

func (c *HTTPClient) DoReq(req *Request) (*Response, error) {
	conn, err := c.buildConn(req)
	if err != nil {
		return nil, xerr.Wrap(err)
	}

	buildReq := c.buildReqRaw(req)

	if c.Debug && req.IsFromPoC {
		log.BluePrintln("------------------- REQUEST DEBUG -------------------")
		log.YellowPrintln(formatMessage(buildReq))
		log.BluePrintln("-----------------------------------------------------")
	} else if c.Debug && !req.IsFromPoC {
		log.BluePrintln(fmt.Sprintf("[DEBUG] %s://%s:%s%s", req.Protocol, req.IP, req.Port, req.Path))
	}

	n, err := conn.Write(buildReq)
	if err != nil {
		_ = conn.Close()
		return nil, xerr.Wrap(err)
	}
	if n != len(buildReq) {
		_ = conn.Close()
		return nil, xerr.Wrap(errors.New("send http request fail"))
	}

	resp, err := c.buildResp(conn)
	if err != nil {
		_ = conn.Close()
		return nil, xerr.Wrap(err)
	}
	resp.RawRequest = buildReq

	if req.FollowRedirect {
		for isRedirect(resp.StatusCode) && req.redirectCount < c.MaxRedirects {
			req.redirectCount++

			locationHeaders := resp.Headers["Location"]
			var location string
			for _, h := range locationHeaders {
				location = h
				break
			}
			if location == "" {
				_ = conn.Close()
				return nil, xerr.Wrap(errors.New("received redirect response without location header"))
			}

			newProtocol, newIP, newPort, newPath := resolveURLManually(
				req.Protocol, req.IP, req.Port, req.Path, location)

			newReq := *req
			newReq.Protocol = newProtocol
			newReq.IP = newIP
			newReq.Port = newPort
			newReq.Path = newPath

			_ = conn.Close()
			return c.DoReq(&newReq)
		}

		if req.redirectCount >= c.MaxRedirects {
			return nil, xerr.Wrap(errors.New("too many redirects"))
		}
	}

	if c.Debug && req.IsFromPoC {
		log.BluePrintln("------------------- RESPONSE DEBUG -------------------")
		log.YellowPrintln(formatMessage(resp.RawResponse))
		log.BluePrintln("-----------------------------------------------------")
	}

	_ = conn.Close()
	return resp, nil
}
