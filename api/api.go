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

package api

import (
	"context"
	"errors"
	"time"

	"github.com/4ra1n/poc-runner/base"
	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/engine"
	"github.com/4ra1n/poc-runner/log"
	"github.com/4ra1n/poc-runner/rawhttp"
	"github.com/4ra1n/poc-runner/reverse"
)

type PocRunner struct {
	ctx         context.Context
	client      *client.HttpClient
	proxy       string
	timeout     time.Duration
	reverseType string
	logLevel    int
	debug       bool
}

func NewPocRunner(ctx context.Context) (*PocRunner, error) {
	return NewPocRunnerEx(
		ctx,
		rawhttp.DefaultNoProxy,
		rawhttp.DefaultTimeout,
		false,
		"dnslog.cn",
		log.InfoLevel,
	)
}

func NewPocRunnerEx(
	ctx context.Context,
	proxy string,
	timeout time.Duration,
	debug bool,
	reverseType string,
	logLevel int,
) (*PocRunner, error) {
	log.SetLevel(logLevel)
	c, err := client.NewHttpClient(proxy, timeout, debug)
	if err != nil {
		return nil, err
	}
	return &PocRunner{
		client:      c,
		ctx:         ctx,
		proxy:       proxy,
		timeout:     timeout,
		debug:       debug,
		reverseType: reverseType,
		logLevel:    logLevel,
	}, nil
}

func (p *PocRunner) Run(input []byte, target string) (string, error) {
	reverse.Type = p.reverseType
	poc, err := engine.InitYamlPoCFromBytes(p.ctx, p.client, input)
	pocList := base.NewList[*base.POC]()
	pocList.Add(poc)
	globalCache := base.NewGlobalCache()
	poc.Caches = globalCache
	if err != nil {
		return "", err
	}
	success, err := engine.RunPOC(poc, target)
	if err != nil {
		return "", err
	}
	if success {
		return engine.NewResultJson(pocList)
	} else {
		return "", errors.New("no result")
	}
}
