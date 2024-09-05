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
	"bufio"
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/4ra1n/poc-runner/proxy"

	"github.com/4ra1n/poc-runner/xerr"
)

type HTTPClient struct {
	Timeout      time.Duration
	ProxyAddr    string
	Debug        bool
	MaxRedirects int
}

func NewHTTPClient(timeout time.Duration, proxyAddr string) (*HTTPClient, error) {
	var err error
	if timeout > time.Minute {
		return nil, xerr.Wrap(errors.New("timeout too long"))
	}
	if timeout < time.Second {
		return nil, xerr.Wrap(errors.New("timeout too short"))
	}
	proxyAddr, err = cleanSocksProxy(proxyAddr)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	return &HTTPClient{
		Timeout:      timeout,
		ProxyAddr:    proxyAddr,
		MaxRedirects: 5,
	}, nil
}

func (c *HTTPClient) SetDebug(debug bool) {
	c.Debug = debug
}

func (c *HTTPClient) buildConn(req *Request) (net.Conn, error) {
	var (
		err  error
		conn net.Conn
	)
	addr := net.JoinHostPort(req.IP, req.Port)
	if c.ProxyAddr != DefaultNoProxy {
		dialer, err := proxy.SOCKS5(tcpProtocol, c.ProxyAddr, nil,
			&net.Dialer{
				Timeout: c.Timeout,
			})
		if err != nil {
			return nil, xerr.Wrap(err)
		}
		conn, err = dialer.Dial(tcpProtocol, addr)
		if err != nil {
			return nil, xerr.Wrap(err)
		}
		if req.Protocol == httpsProtocol {
			tlsConn := tls.Client(conn, &tls.Config{
				InsecureSkipVerify: true,
				ServerName:         req.Domain,
			})
			if err = tlsConn.Handshake(); err != nil {
				return nil, xerr.Wrap(err)
			}
			conn = tlsConn
		}
	} else {
		if req.Protocol == httpsProtocol {
			tlsConfig := &tls.Config{
				InsecureSkipVerify: true,
			}
			conn, err = tls.DialWithDialer(&net.Dialer{
				Timeout: c.Timeout,
			}, tcpProtocol, addr, tlsConfig)
		} else if req.Protocol == httpProtocol {
			conn, err = net.DialTimeout(tcpProtocol, addr, c.Timeout)
		} else {
			return nil, xerr.Wrap(errors.New("unsupported protocol: " + req.Protocol))
		}
		if err != nil {
			return nil, xerr.Wrap(err)
		}
	}
	return conn, err
}

func (c *HTTPClient) buildResp(conn net.Conn) (*Response, error) {
	reader := bufio.NewReader(conn)
	var rawResponseBuffer bytes.Buffer

	statusLine, err := reader.ReadString(respDelim)
	if err != nil {
		_ = conn.Close()
		return nil, xerr.Wrap(err)
	}
	rawResponseBuffer.WriteString(statusLine)
	statusLine = strings.TrimSpace(statusLine)
	parts := strings.SplitN(statusLine, space, 3)
	if len(parts) < 2 {
		_ = conn.Close()
		return nil, xerr.Wrap(fmt.Errorf("invalid status line: %s", statusLine))
	}
	statusCode := parts[1]
	statusInt, err := strconv.Atoi(statusCode)
	if err != nil {
		_ = conn.Close()
		return nil, xerr.Wrap(err)
	}

	headers := make(map[string][]string)
	var contentLength int
	var chunked bool
	for {
		line, err := reader.ReadString(respDelim)
		if err != nil {
			_ = conn.Close()
			return nil, xerr.Wrap(err)
		}
		rawResponseBuffer.WriteString(line)
		line = strings.TrimSpace(line)
		if line == none {
			break
		}
		partsIn := strings.SplitN(line, headerSep, 2)
		if len(partsIn) == 2 {
			key := strings.TrimSpace(partsIn[0])
			value := strings.TrimSpace(partsIn[1])

			headers[key] = append(headers[key], value)

			if strings.ToLower(key) == clHeader {
				contentLength, err = strconv.Atoi(value)
				if err != nil {
					_ = conn.Close()
					return nil, xerr.Wrap(fmt.Errorf("invalid content-length: %s", value))
				}
			}

			if strings.ToLower(key) == teHeader && strings.ToLower(value) == teChunked {
				chunked = true
			}
		}
	}

	var bodyBuffer bytes.Buffer
	if chunked {
		for {
			chunkSizeLine, err := reader.ReadString(respDelim)
			if err != nil {
				_ = conn.Close()
				return nil, xerr.Wrap(err)
			}
			rawResponseBuffer.WriteString(chunkSizeLine)
			chunkSizeLine = strings.TrimSpace(chunkSizeLine)
			chunkSize, err := strconv.ParseInt(chunkSizeLine, 16, 64)
			if err != nil {
				_ = conn.Close()
				return nil, xerr.Wrap(fmt.Errorf("invalid chunk size: %s", chunkSizeLine))
			}

			if chunkSize == 0 {
				break
			}

			chunkData := make([]byte, chunkSize)
			_, err = io.ReadFull(reader, chunkData)
			if err != nil {
				_ = conn.Close()
				return nil, xerr.Wrap(err)
			}
			bodyBuffer.Write(chunkData)
			rawResponseBuffer.Write(chunkData)

			chunkEndLine, err := reader.ReadString(respDelim)
			if err != nil {
				_ = conn.Close()
				return nil, xerr.Wrap(err)
			}
			rawResponseBuffer.WriteString(chunkEndLine)
		}
		trailerLine, err := reader.ReadString(respDelim)
		if err != nil {
			_ = conn.Close()
			return nil, xerr.Wrap(err)
		}
		rawResponseBuffer.WriteString(trailerLine)
	} else if contentLength > 0 {
		body := make([]byte, contentLength)
		_, err = io.ReadFull(reader, body)
		if err != nil && err != io.EOF {
			_ = conn.Close()
			return nil, xerr.Wrap(err)
		}
		bodyBuffer.Write(body)
		rawResponseBuffer.Write(body)
	} else {
		body, err := io.ReadAll(reader)
		if err != nil && err != io.EOF {
			_ = conn.Close()
			return nil, xerr.Wrap(err)
		}
		bodyBuffer.Write(body)
		rawResponseBuffer.Write(body)
	}

	return &Response{
		StatusCode:  statusInt,
		Status:      statusLine,
		Headers:     headers,
		Body:        bodyBuffer.Bytes(),
		RawResponse: rawResponseBuffer.Bytes(),
	}, nil
}
