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

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/log"
)

const dnsLogCnUrl = "http://dnslog.cn/"

type DnsLogCnData struct {
}

type DnsLogCn struct {
	c         *client.HttpClient
	baseUrl   string
	newDomain string
	session   string
}

func (d *DnsLogCn) GetUrl() string {
	return fmt.Sprintf("%s://%s.%s", "http", "http", d.newDomain)
}

func (d *DnsLogCn) GetRmi() string {
	return fmt.Sprintf("%s://%s.%s", "rmi", "rmi", d.newDomain)
}

func (d *DnsLogCn) GetLdap() string {
	return fmt.Sprintf("%s://%s.%s", "ldap", "ldap", d.newDomain)
}

func (d *DnsLogCn) GetDNS() string {
	return fmt.Sprintf("%s.%s", "dns", d.newDomain)
}

func (d *DnsLogCn) Wait(i int) bool {
	t := time.Duration(i) * time.Second
	for start := time.Now(); time.Since(start) < t; time.Sleep(time.Millisecond * 500) {
		if d.waitInternal() {
			return true
		}
	}
	return false
}

func (d *DnsLogCn) Close() {
}

func (d *DnsLogCn) waitInternal() bool {
	resp, err := d.c.DoReq(&client.TheRequest{
		Target:         dnsLogCnUrl,
		Method:         "GET",
		Path:           "/getrecords.php",
		FollowRedirect: false,
		Body:           "",
		Headers: map[string]string{
			"Cookie": d.session,
		},
		IsFromPoC: false,
	})
	if err != nil {
		log.Error(err)
		return false
	}
	var data interface{}
	err = json.Unmarshal(resp.Body, &data)
	if err != nil {
		log.Error(err)
		return false
	}
	if val, ok := data.([]interface{}); ok {
		for _, v := range val {
			if vi, viOk := v.([]interface{}); viOk {
				if len(vi) != 3 {
					continue
				}
				l := vi[0]
				up := strings.ToUpper(l.(string))
				if strings.Contains(up, strings.ToUpper(d.newDomain)) {
					log.Info("reverse receive request")
					return true
				}
			}
		}
	}
	return false
}
