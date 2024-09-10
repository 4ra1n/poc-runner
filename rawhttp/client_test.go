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
	"testing"
	"time"
)

func TestRawHttpNoProxy(t *testing.T) {
	client, _ := NewHTTPClient(5*time.Second, DefaultNoProxy)
	client.SetDebug(true)
	response, err := client.Get("http://example.com")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("HTTP Response:\n", string(response.RawResponse))
	}
}

func TestRawHttpsNoProxy(t *testing.T) {
	client, _ := NewHTTPClient(5*time.Second, DefaultNoProxy)
	req, err := NewRequest("https://example.com", MethodGet)
	if err != nil {
		panic(err)
	}
	response, err := client.DoReq(req)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("HTTP Response:\n", string(response.RawResponse))
	}
}

func TestRawHttpProxy(t *testing.T) {
	client, err := NewHTTPClient(5*time.Second, "socks5://127.0.0.1:10808")
	if err != nil {
		panic(err)
	}
	req, err := NewRequest("http://example.com", MethodGet)
	if err != nil {
		panic(err)
	}
	response, err := client.DoReq(req)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("HTTP Response:\n", string(response.RawResponse))
	}
}

func TestRawHttpsProxy(t *testing.T) {
	client, err := NewHTTPClient(5*time.Second, "socks5://127.0.0.1:10808")
	if err != nil {
		panic(err)
	}
	req, err := NewRequest("https://example.com", MethodGet)
	if err != nil {
		panic(err)
	}
	response, err := client.DoReq(req)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("HTTP Response:\n", string(response.RawResponse))
	}
}
