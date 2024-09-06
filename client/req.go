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

type TheRequest struct {
	Target         string
	Method         string
	Path           string
	FollowRedirect bool
	Body           string
	Headers        map[string]string
	All            []byte
}

// Equals
// 判断两个 TheRequest 对象是否完全一致
// 给 CACHE 功能使用
func (r *TheRequest) Equals(req *TheRequest) bool {
	if r.Target == req.Target &&
		r.Method == req.Method &&
		r.Path == req.Path &&
		r.FollowRedirect == req.FollowRedirect &&
		r.Body == req.Body {
		// HEADERS EQUALS
		for k, v := range r.Headers {
			vv, exi := req.Headers[k]
			if !exi {
				return false
			}
			if vv != v {
				return false
			}
		}
		return true
	}
	return false
}
