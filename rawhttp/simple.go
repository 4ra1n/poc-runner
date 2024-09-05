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
	"github.com/4ra1n/poc-runner/xerr"
)

func (c *HTTPClient) Get(url string) (*Response, error) {
	req, err := NewRequest(url, MethodGet)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	return c.DoReq(req)
}

func (c *HTTPClient) PostForm(url string, body []byte) (*Response, error) {
	req, err := NewRequest(url, MethodPost)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	req.SetBytesBody(body)
	req.SetHeader(ctHeader, ctFormVal)
	return c.DoReq(req)
}

func (c *HTTPClient) PostJson(url string, body []byte) (*Response, error) {
	req, err := NewRequest(url, MethodPost)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	req.SetBytesBody(body)
	req.SetHeader(ctHeader, ctJsonVal)
	return c.DoReq(req)
}

func (c *HTTPClient) PostXml(url string, body []byte) (*Response, error) {
	req, err := NewRequest(url, MethodPost)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	req.SetBytesBody(body)
	req.SetHeader(ctHeader, ctXmlVal)
	return c.DoReq(req)
}
