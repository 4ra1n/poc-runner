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
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/4ra1n/poc-runner/base"
	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/log"
	"github.com/4ra1n/poc-runner/rawhttp"
	"github.com/4ra1n/poc-runner/xerr"
)

var (
	debug   bool
	quiet   bool
	path    string
	target  string
	proxy   string
	timeout string
	output  string
	xe      *xerr.XError
)

func Start(ctx context.Context, cancel context.CancelFunc) {
	flag.StringVar(&path, "r", "", "poc path (support a.yml | a.yaml | pocs/*.yaml | pocs/*.yml)")
	flag.StringVar(&target, "t", "", "poc target (support http(s)://target | target | target/path)")
	flag.BoolVar(&debug, "debug", false, "debug mode (open debug output and set debug log level)")
	flag.BoolVar(&quiet, "quiet", false, "quiet mode (only output error and set error log level)")
	flag.StringVar(&proxy, "proxy", "", "set socks5 proxy (socks5://ip:port)")
	flag.StringVar(&timeout, "timeout", "", "set http client timeout second (10)")
	flag.StringVar(&output, "output", stdoutOut, "set output type (support stdout | txt | json | html)")
	flag.Parse()

	path = strings.TrimSpace(path)
	target = strings.TrimSpace(target)

	if path == "" || target == "" {
		log.RedPrintln("poc-runner usage:")
		log.GreenPrintln("example: ./poc-runner -r test.yml -t http://localhost")
		fmt.Println("options:")
		flag.PrintDefaults()
		log.RedPrintln("press ctrl+c to exit / 按 ctrl+c 退出")
		return
	}

	if quiet {
		log.SetLevel(log.ErrorLevel)
	} else {
		if debug {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.InfoLevel)
		}
	}

	log.Info("start poc runner")

	proxy = strings.TrimSpace(proxy)
	timeout = strings.TrimSpace(timeout)
	if proxy == "" {
		proxy = rawhttp.DefaultNoProxy
	}
	var timeoutVal time.Duration
	if timeout == "" {
		timeoutVal = rawhttp.DefaultTimeout
	} else {
		t, err := strconv.Atoi(timeout)
		if err != nil {
			checkError(err)
			return
		}
		if t < 1 || t > 60 {
			checkError(errors.New("timeout invalid"))
			return
		}
		timeoutVal = time.Duration(t) * time.Second
	}

	// INIT RAW HTTP CLIENT
	c, err := client.NewHttpClient(proxy, timeoutVal, debug)
	if err != nil {
		checkError(err)
		return
	}

	poc, err := ParseYAMLFile(ctx, c, path)
	if err != nil {
		checkError(err)
		return
	}

	log.Info("parse yaml file success")
	log.Infof("run: %s", poc.Name)
	log.Infof("target: %s", target)

	// INIT GLOBAL CACHE
	globalCache := base.NewGlobalCache()
	poc.Caches = globalCache

	success, err := RunPOC(poc, target)
	if err != nil {
		checkError(err)
		return
	}

	if !success {
		log.Info("no vulnerability found")
	} else {
		// validate output type
		output = strings.TrimSpace(output)
		output = strings.ToLower(output)
		switch output {
		case stdoutOut:
			log.RedPrintln("#################### FOUND VULNERABILITY ####################")
			log.RedPrintln("POC    -> ", poc.Name)
			log.RedPrintln("TARGET -> ", poc.Target)
			log.RedPrintln("#############################################################")
		case txtOut:
			var j string
			j, err = NewResultTxt(poc)
			if err != nil {
				checkError(err)
				return
			}
			fileName := fmt.Sprintf("%s-%d.txt", poc.Name, time.Now().UnixMilli())
			log.Infof("output file: %s", fileName)
			err = os.WriteFile(fileName, []byte(j), 0644)
			if err != nil {
				checkError(err)
				return
			}
		case htmlOut:
			// TODO
		case jsonOut:
			var j string
			j, err = NewResultJson(poc)
			if err != nil {
				checkError(err)
				return
			}
			fileName := fmt.Sprintf("%s-%d.json", poc.Name, time.Now().UnixMilli())
			log.Infof("output file: %s", fileName)
			err = os.WriteFile(fileName, []byte(j), 0644)
			if err != nil {
				checkError(err)
				return
			}
		default:
			log.Error("not support output type:", output)
			return
		}
	}

	log.Info("poc runner finish")
	cancel()
}
