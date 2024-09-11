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
	"fmt"
	"testing"
	"time"

	"github.com/4ra1n/poc-runner/log"
)

func TestAdvanceAPI(t *testing.T) {
	ctx := context.Background()
	// NEW ADVANCE POC RUNNER
	runner, err := NewPocRunnerEx(
		ctx,                        // CONTEXT
		"socks5://127.0.0.1:10808", // SOCKS PROXY
		time.Second*10,             // TIMEOUT
		true,                       // DEBUG MODE
		"dnslog.cn",                // REVERSE CONFIG (dnslog.cn | interact.sh)
		log.DebugLevel,             // LOG LEVEL
	)
	if err != nil {
		return
	}
	// RUN POC
	report, err := runner.Run([]byte(poc), "https://example.com")
	if err != nil {
		return
	}
	fmt.Println(report)
}
