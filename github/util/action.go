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

package util

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/google/go-github/v57/github"
	"golang.org/x/net/proxy"
	"golang.org/x/oauth2"

	"github/color"
)

func CleanAction(token string, useProxy bool, socks string, repoOwner string, repoName string, action string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	if useProxy {
		dialer, err := proxy.SOCKS5("tcp",
			socks, nil, proxy.Direct)
		if err != nil {
			color.RedPrintln(err)
			return
		}
		t := tc.Transport.(*oauth2.Transport)
		t.Base = &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return dialer.Dial(network, addr)
			},
		}
		tc.Transport = t
	}
	client := github.NewClient(tc)

	runs, _, err := client.Actions.ListWorkflowRunsByFileName(
		ctx, repoOwner, repoName, action, &github.ListWorkflowRunsOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, run := range runs.WorkflowRuns {
		req, err := http.NewRequest("DELETE", run.GetURL(), nil)
		if err != nil {
			fmt.Println(err)
			continue
		}
		req.Header.Set("Authorization", "token "+token)
		req.Header.Set("Accept", "application/vnd.github.v3+json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		_ = resp.Body.Close()

		if resp.StatusCode == http.StatusNoContent {
			fmt.Printf("run %d got deleted\n", run.GetID())
		}
	}
}
