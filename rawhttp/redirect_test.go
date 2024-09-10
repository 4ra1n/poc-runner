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
	"net/http"
	"testing"
	"time"
)

func TestHTTPClientRedirectTrue(t *testing.T) {
	go func() {
		http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/test", http.StatusFound)
		})
		http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://example.com", http.StatusFound)
		})
		err := http.ListenAndServe(":10302", nil)
		if err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)

	client, _ := NewHTTPClient(5*time.Second, DefaultNoProxy)
	req, err := NewRequest("http://127.0.0.1:10302/redirect", MethodGet)
	if err != nil {
		panic(err)
	}
	req.SetFollowRedirect(true)
	response, err := client.DoReq(req)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("HTTP Response:\n", string(response.RawResponse))
	}

	if response.StatusCode != 200 {
		panic("test error")
	}
}

func TestHTTPClientRedirectFalse(t *testing.T) {
	go func() {
		http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/test", http.StatusFound)
		})
		http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://example.com", http.StatusFound)
		})
		err := http.ListenAndServe(":10302", nil)
		if err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)

	client, _ := NewHTTPClient(5*time.Second, DefaultNoProxy)
	req, err := NewRequest("http://127.0.0.1:10302/redirect", MethodGet)
	if err != nil {
		panic(err)
	}
	req.SetFollowRedirect(false)
	response, err := client.DoReq(req)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("HTTP Response:\n", string(response.RawResponse))
	}

	if response.StatusCode != 302 {
		panic("test error")
	}
}
