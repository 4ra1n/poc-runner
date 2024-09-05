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
}
