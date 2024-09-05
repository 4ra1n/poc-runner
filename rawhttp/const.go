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

import "time"

const (
	DefaultNoProxy = "NO-PROXY"
	DefaultTimeout = time.Second * 10
)

const (
	version     = "HTTP/1.1"
	respDelim   = '\n'
	space       = " "
	headerSep   = ":"
	lineSep     = "\r\n"
	querySep    = "?"
	equalSep    = "="
	andSep      = "&"
	plusSep     = "+"
	spaceEncode = "%20"
)

const (
	httpProtocol  = "http"
	tcpProtocol   = "tcp"
	httpsProtocol = "https"
)

const (
	none             = ""
	httpDefaultPort  = "80"
	httpsDefaultPort = "443"
)

var (
	defaultUA = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) " +
		"Chrome/128.0.0.0 Safari/537.36"
	teChunked = "chunked"
)

var (
	ctHeader = "content-type"
	clHeader = "content-length"
	teHeader = "transfer-encoding"
)

var (
	acceptHeaderStd = "Accept"
	hostHeaderStd   = "Host"
	connHeaderStd   = "Connection"
	uaHeaderStd     = "User-Agent"
	ctHeaderStd     = "Content-Type"
	clHeaderStd     = "Content-Length"
)

var (
	acceptVal = "*/*"
	ctFormVal = "application/x-www-form-urlencoded"
	ctJsonVal = "application/json"
	ctXmlVal  = "application/xml"
)
