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

package engine

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/4ra1n/poc-runner/log"
	"github.com/4ra1n/poc-runner/reverse"
)

var (
	debug   bool
	quiet   bool
	check   bool
	path    string
	target  string
	proxy   string
	timeout string
	output  string
	rev     string
)

func parseFlag() error {
	flag.StringVar(&path, "r", "", "poc path (support a.yml | a.yaml | pocs/*.yaml | pocs/*.yml)")
	flag.StringVar(&target, "t", "", "poc target (support http(s)://target | target | url.txt )")
	flag.BoolVar(&debug, "debug", false, "debug mode (open debug output and set debug log level)")
	flag.BoolVar(&quiet, "quiet", false, "quiet mode (only output error and set error log level)")
	flag.BoolVar(&check, "check", false, "check target is valid before run poc")
	flag.StringVar(&proxy, "proxy", "", "set socks5 proxy (socks5://ip:port)")
	flag.StringVar(&timeout, "timeout", "", "set http client timeout second (10)")
	flag.StringVar(&output, "output", stdoutOut, "set output type (support stdout | txt | json | html)")
	flag.StringVar(&rev, "reverse", reverse.DNSLogCN, "set reverse type (support dnslog.cn | interact.sh)")
	flag.Parse()

	path = strings.TrimSpace(path)
	target = strings.TrimSpace(target)

	if path == "" || target == "" {
		log.RedPrintln("poc-runner usage:")
		log.GreenPrintln("example: ./poc-runner -r test.yml -t http://localhost")
		fmt.Println("options:")
		flag.PrintDefaults()
		log.RedPrintln("press ctrl+c to exit / 按 ctrl+c 退出")
		return errors.New("input error")
	}
	return nil
}
