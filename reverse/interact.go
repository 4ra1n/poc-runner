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
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/log"
	"github.com/4ra1n/poc-runner/xerr"
	"strings"
	"time"
)

type InteractResult struct {
	Data   []string `json:"data"`
	Extra  string   `json:"extra"`
	AesKey string   `json:"aes_key"`
}

type InteractData struct {
	Protocol      string `json:"protocol"`
	UniqueID      string `json:"unique-id"`
	FullID        string `json:"full-id"`
	RawReq        string `json:"raw-request"`
	RawResp       string `json:"raw-response"`
	RemoteAddress string `json:"remote-address"`
	Timestamp     string `json:"timestamp"`
}

type RegisterRequest struct {
	PublicKey     string `json:"public-key"`
	SecretKey     string `json:"secret-key"`
	CorrelationID string `json:"correlation-id"`
}

var defaultServers = []string{
	"oast.pro",
	"oast.live",
	"oast.site",
	"oast.online",
	"oast.fun",
	"oast.me",
}

type Interact struct {
	c             *client.HttpClient
	server        string
	correlationID string
	domain        string
	secretKey     string
	privateKey    *rsa.PrivateKey
	publicKey     *rsa.PublicKey
	pubKeyData    string
}

func NewInteract(c *client.HttpClient, server string) (*Interact, error) {
	if server == "" {
		server = randomPick(defaultServers)
	}
	correlationID := randLower(20)
	secretKey := randomUUID()

	pri, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	pubKey := &pri.PublicKey

	pubKeyData, err := encodePublicKey(pubKey)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	reg := RegisterRequest{
		PublicKey:     pubKeyData,
		SecretKey:     secretKey,
		CorrelationID: correlationID,
	}
	data, err := json.Marshal(reg)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	regReq := &client.TheRequest{
		Target:         "http://" + server,
		Method:         "POST",
		Path:           "/register",
		FollowRedirect: false,
		Body:           string(data),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	reResp, err := c.DoReq(regReq)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	log.Infof("interact register response %d", reResp.Code)

	tt := randLower(13)
	domain := fmt.Sprintf("%s%s.%s", correlationID, tt, server)
	log.Infof("interact domain %s", domain)

	i := &Interact{
		c:             c,
		server:        server,
		correlationID: correlationID,
		secretKey:     secretKey,
		domain:        domain,
		privateKey:    pri,
		publicKey:     pubKey,
		pubKeyData:    pubKeyData,
	}
	return i, nil
}

func (i *Interact) Close() {
	de := &RegisterRequest{
		SecretKey:     i.secretKey,
		CorrelationID: i.correlationID,
	}
	data, err := json.Marshal(de)
	if err != nil {
		log.Error(err)
		return
	}
	deReq := &client.TheRequest{
		Target:         "http://" + i.server,
		Method:         "POST",
		Path:           "/deregister",
		FollowRedirect: false,
		Body:           string(data),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	deResp, err := i.c.DoReq(deReq)
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("interact deregister response %d", deResp.Code)
}

func (i *Interact) GetUrl() string {
	return fmt.Sprintf("%s://%s.%s", "http", "http", i.domain)
}

func (i *Interact) GetRmi() string {
	return fmt.Sprintf("%s://%s.%s", "rmi", "rmi", i.domain)
}

func (i *Interact) GetLdap() string {
	return fmt.Sprintf("%s://%s.%s", "ldap", "ldap", i.domain)
}

func (i *Interact) GetDNS() string {
	return fmt.Sprintf("%s.%s", "dbs", i.domain)
}

func (i *Interact) Wait(w int) bool {
	pollPath := fmt.Sprintf("/poll?id=%s&secret=%s", i.correlationID, i.secretKey)
	pollReq := &client.TheRequest{
		Target:         "http://" + i.server,
		Method:         "GET",
		Path:           pollPath,
		FollowRedirect: false,
		Body:           "",
		Headers:        make(map[string]string),
	}
	var respAesKey string
	t := time.Duration(w) * time.Second
	for start := time.Now(); time.Since(start) < t; time.Sleep(time.Millisecond * 1000) {
		pollResp, err := i.c.DoReq(pollReq)
		if err != nil {
			log.Error(err)
			continue
		}
		var result *InteractResult
		err = json.Unmarshal(pollResp.Body, &result)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info("parse interact response data success")
		respAesKey = result.AesKey
		for _, d := range result.Data {
			var jData []byte
			jData, err = decryptMessage(respAesKey, d, i.privateKey)
			if err != nil {
				log.Error(err)
				continue
			}
			var revData *InteractData
			err = json.Unmarshal(jData, &revData)
			if err != nil {
				log.Error(err)
				continue
			}
			log.Infof("receive data %s", revData.FullID)
			temp := strings.TrimRight(i.domain, i.server)
			remote := strings.ToLower(revData.FullID)
			local := strings.ToLower(temp)
			if strings.Contains(remote, local) {
				return true
			}
		}
	}
	return false
}
