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
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/4ra1n/poc-runner/base"
	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/log"
	"github.com/4ra1n/poc-runner/reverse"
	"github.com/4ra1n/poc-runner/xerr"
)

var xe *xerr.XError

func Start(ctx context.Context, cancel context.CancelFunc) {
	// PARSE FLAG
	err := parseFlag()
	if err != nil {
		checkError(err)
		return
	}

	// PARSE TARGET
	targetList, err := getTargetList()
	if err != nil {
		checkError(err)
		return
	}

	configLog()

	log.Info("start poc runner")

	// SET HTTP CLIENT
	err = setHttpClient()
	if err != nil {
		checkError(err)
		return
	}

	// INIT GLOBAL CACHE
	globalCache := base.NewGlobalCache()

	// INIT REVERSE
	reverse.Type = rev

	// INIT RESULT
	resultList := base.NewList[*base.POC]()

	// PARSE YAML FILE
	poc, err := InitYamlPoC(path)
	if err != nil {
		checkError(err)
		return
	}
	log.Info("parse yaml file success")

	var success bool
	for _, theTarget := range targetList {
		// INIT POC
		runningPoC, err := InitYamlPoCFromInterface(ctx, client.Instance, poc)
		if err != nil {
			checkError(err)
			continue
		}
		// INIT POC
		log.Infof("run: %s", runningPoC.Name)
		log.Infof("target: %s", target)

		runningPoC.Caches = globalCache

		log.Infof("scan target: %s", theTarget)
		// CHECK TARGET
		if check {
			var ok bool
			ok, theTarget, err = checkTarget(theTarget)
			if err != nil {
				checkError(err)
				return
			}
			if !ok {
				log.Errorf("target %s is not available")
				log.Warn("please check network/proxy config")
				continue
			}
		}
		// RUN POC
		success, err = RunPOC(runningPoC, theTarget)
		if err != nil {
			checkError(err)
			continue
		}
		if success {
			resultList.Add(runningPoC)
		}
	}

	if !success {
		log.Info("no vulnerability found")
	} else {
		// validate output type
		output = strings.TrimSpace(output)
		output = strings.ToLower(output)
		switch output {
		case stdoutOut:
			baseOutput(resultList)
		case txtOut:
			log.Info("output type: txt")
			baseOutput(resultList)
			var j string
			j, err = NewResultTxt(resultList)
			if err != nil {
				checkError(err)
				return
			}
			fileName := fmt.Sprintf("%s-%d.txt", "result-txt", time.Now().UnixMilli())
			log.Infof("output file: %s", fileName)
			err = os.WriteFile(fileName, []byte(j), 0644)
			if err != nil {
				checkError(err)
				return
			}
		case htmlOut:
			log.Info("output type: html")
			baseOutput(resultList)
			var j string
			j, err = NewResultHTML(resultList)
			if err != nil {
				checkError(err)
				return
			}
			fileName := fmt.Sprintf("%s-%d.html", "result-html", time.Now().UnixMilli())
			log.Infof("output file: %s", fileName)
			err = os.WriteFile(fileName, []byte(j), 0644)
			if err != nil {
				checkError(err)
				return
			}
		case jsonOut:
			log.Info("output type: json")
			baseOutput(resultList)
			var j string
			j, err = NewResultJson(resultList)
			if err != nil {
				checkError(err)
				return
			}
			fileName := fmt.Sprintf("%s-%d.json", "result-json", time.Now().UnixMilli())
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

	if reverse.Instance != nil {
		reverse.Instance.Close()
	}

	log.Info("poc runner finish")
	cancel()
}
