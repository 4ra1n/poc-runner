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
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/4ra1n/poc-runner/xerr"
)

type HostType int

const (
	IPv4   = HostType(0x000f)
	IPv6   = HostType(0x00f0)
	Domain = HostType(0x0f00)
)

func determineHostType(hostname string) HostType {
	if ip := net.ParseIP(hostname); ip != nil {
		if strings.Contains(hostname, ".") {
			return IPv4
		} else if strings.Contains(hostname, ":") {
			return IPv6
		}
	}
	if strings.Contains(hostname, ":") && net.ParseIP(hostname) != nil {
		return IPv6
	}
	return Domain
}

func cleanSocksProxy(proxyAddr string) (string, error) {
	if strings.HasPrefix(proxyAddr, "http") {
		return none, xerr.Wrap(errors.New("only support socks proxy"))
	}
	if strings.HasPrefix(proxyAddr, "socks") {
		proxyAddr = strings.TrimPrefix(proxyAddr, "socks://")
		proxyAddr = strings.TrimPrefix(proxyAddr, "socks4://")
		proxyAddr = strings.TrimPrefix(proxyAddr, "socks5://")
		if _, _, err := net.SplitHostPort(proxyAddr); err != nil {
			return none, xerr.Wrap(errors.New("expected socks5://ip:port proxy addr"))
		}
	}
	return proxyAddr, nil
}

func getTarget(req *Request) string {
	if strings.TrimSpace(req.Domain) != "" {
		return req.Domain
	}
	if strings.TrimSpace(req.Port) == "" {
		return req.IP
	}
	return fmt.Sprintf("%s:%s", req.IP, req.Port)
}

func isRedirect(statusCode int) bool {
	return statusCode == 301 || statusCode == 302 || statusCode == 303 || statusCode == 307 || statusCode == 308
}

func resolveURLManually(protocol, ip, port, basePath, location string) (newProtocol, newIP, newPort, newPath string) {
	newProtocol = protocol
	newIP = ip
	newPort = port

	if isAbsoluteURL(location) {
		parts := splitAbsoluteURL(location)
		newProtocol = parts.scheme
		newIP = parts.host
		if parts.port != "" {
			newPort = parts.port
		} else {
			newPort = ""
		}
		newPath = parts.path
	} else {
		if location[0] == '/' {
			newPath = location
		} else {
			baseDir := basePath[:strings.LastIndex(basePath, "/")+1]
			newPath = baseDir + location
		}
	}

	return
}

func isAbsoluteURL(url string) bool {
	return strings.HasPrefix(url, "http")
}

func splitAbsoluteURL(rawURL string) (parts struct{ scheme, host, port, path string }) {
	schemeEnd := strings.Index(rawURL, "://")
	if schemeEnd != -1 {
		parts.scheme = rawURL[:schemeEnd]
		rawURL = rawURL[schemeEnd+3:]
	}

	var hostPort string
	pathStart := strings.Index(rawURL, "/")
	if pathStart != -1 {
		hostPort = rawURL[:pathStart]
		parts.path = rawURL[pathStart:]
	} else {
		hostPort = rawURL
		parts.path = "/"
	}

	portStart := strings.Index(hostPort, ":")
	if portStart != -1 {
		parts.host = hostPort[:portStart]
		parts.port = hostPort[portStart+1:]
	} else {
		parts.host = hostPort
		if parts.scheme == "https" {
			parts.port = "443"
		} else {
			parts.port = "80"
		}
	}
	return
}
