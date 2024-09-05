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

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/4ra1n/poc-runner/engine"
	"github.com/4ra1n/poc-runner/log"
	"github.com/4ra1n/poc-runner/rawhttp"
)

var (
	version = "UNKNOWN"
	now     = "UNKNOWN"
	logo    = `
__________      _________   __________                                  
\______   \____ \_   ___ \  \______   \__ __  ____   ____   ___________ 
 |     ___/  _ \/    \  \/   |       _/  |  \/    \ /    \_/ __ \_  __ \
 |    |  (  <_> )     \____  |    |   \  |  /   |  \   |  \  ___/|  | \/
 |____|   \____/ \______  /  |____|_  /____/|___|  /___|  /\___  >__|   
                        \/          \/           \/     \/     \/`
	desc = "XRAY YAML POC RUNNER (Open Source Version) @ 4ra1n\n" +
		"PROJECT URL: https://github.com/4ra1n/poc-runner"
	info       = "VERSION: %s   BUILD-TIME: %s"
	updateUrl  = "https://poc-runner.oss-cn-hangzhou.aliyuncs.com/version.txt"
	releaseUrl = "https://github.com/4ra1n/poc-runner/releases/latest"
)

func initLog() {
	logo = fmt.Sprintf(logo)
	logo = strings.TrimSpace(logo)
	if !log.IsSupported() {
		log.DisableColor()
	} else {
		log.EnableColor()
	}
	log.GreenPrintln(logo)
	log.YellowPrintln(desc)
	log.BluePrintf(info, version, now)
	fmt.Println()
	// DEFAULT LEVEL
	log.SetLevel(log.InfoLevel)
	checkUpdate()
}

func checkUpdate() {
	// ignore dev code
	if version == "UNKNOWN" {
		return
	}
	// new raw http client
	client := &rawhttp.HTTPClient{
		Timeout:      3 * time.Second,
		ProxyAddr:    rawhttp.DefaultNoProxy,
		Debug:        false,
		MaxRedirects: 5,
	}
	resp, err := client.Get(updateUrl)
	if err != nil {
		log.Warn("check update error")
		return
	}
	ver := string(resp.Body)
	log.Debug("the latest version from remote:", ver)
	if ver != version {
		log.Info("new version:", ver)
		log.Info("download url:", releaseUrl)
	}
}

func wait(ctx context.Context) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-sigChan:
			log.Info("ctrl+c stop")
			return
		case <-ctx.Done():
			log.Info("run finish")
			return
		}
	}
}

/**
 *
 * poc-runner project @ 4ra1n
 *
 * 　　 へ　　　　　／|
 * 　　/＼7　　　 ∠＿/
 * 　 /　│　　 ／　／
 * 　│　Z ＿,＜　／　　 /`ヽ
 * 　│　　　　　ヽ　　 /　　〉
 * 　 Y　　　　　`　 /　　/
 * 　?●　?　●　　??〈　　/
 * 　()　 へ　　　　|　＼〈
 * 　　>? ?_　 ィ　 │ ／／
 * 　 / へ　　 /　?＜| ＼＼
 * 　 ヽ_?　　(_／　 │／／
 * 　　7　　　　　　　|／
 * 　　＞―r￣￣~∠--|
 */
func main() {
	// init log and print logo
	initLog()
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	// start
	go engine.Start(ctx, cancel)
	// wait
	wait(ctx)
}
