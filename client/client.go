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

package client

import (
	"github.com/4ra1n/poc-runner/rawhttp"
	"github.com/4ra1n/poc-runner/xerr"

	"time"
)

var Instance *HttpClient

// HttpClient
// 这是 RAW HTTP 的包装
// 不建议直接使用 RAW HTTP CLIENT
type HttpClient struct {
	client *rawhttp.HTTPClient
}

func NewHttpClient(proxy string, timeout time.Duration, debug bool) (*HttpClient, error) {
	rawClient, err := rawhttp.NewHTTPClient(timeout, proxy)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	rawClient.SetDebug(debug)
	return &HttpClient{
		client: rawClient,
	}, nil
}

func (c *HttpClient) DoReq(req *TheRequest) (*TheResponse, error) {
	u := req.Target + req.Path
	rawReq, err := rawhttp.NewRequest(u, req.Method)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	rawReq.IsFromPoC = req.IsFromPoC
	rawReq.SetBody(req.Body)
	rawReq.SetFollowRedirect(req.FollowRedirect)
	for k, v := range req.Headers {
		rawReq.SetHeader(k, v)
	}
	resp, err := c.client.DoReq(rawReq)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	theResp := &TheResponse{}
	theResp.Code = resp.StatusCode
	bodyBytes := resp.Body
	theResp.Body = bodyBytes
	theResp.Headers = resp.Headers
	theResp.All = resp.RawResponse
	req.All = resp.RawRequest
	return theResp, nil
}

func (c *HttpClient) Get(u string) (*TheResponse, error) {
	resp, err := c.client.Get(u)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	theResp := &TheResponse{}
	theResp.Code = resp.StatusCode
	bodyBytes := resp.Body
	theResp.Body = bodyBytes
	theResp.Headers = resp.Headers
	theResp.All = resp.RawResponse
	// ignore request raw
	return theResp, nil
}
