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

package base

import (
	"errors"

	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/xerr"
)

type CacheItem struct {
	PocName  string
	RuleName string
	Req      *client.TheRequest
	Resp     *client.TheResponse
}

type GlobalCache struct {
	caches *List[*CacheItem]
}

func NewGlobalCache() *GlobalCache {
	return &GlobalCache{
		caches: NewList[*CacheItem](),
	}
}

func (g *GlobalCache) SetCache(
	pocName string,
	rule string,
	req *client.TheRequest,
	resp *client.TheResponse) error {
	cacheItem := &CacheItem{
		PocName:  pocName,
		RuleName: rule,
		Req:      req,
		Resp:     resp,
	}
	if g.caches == nil {
		return xerr.Wrap(errors.New("must init global cache first"))
	}
	// 没有并发问题
	// 这个 List 是安全的
	g.caches.Add(cacheItem)
	return nil
}

func (g *GlobalCache) GetCache(req *client.TheRequest) (*client.TheResponse, error) {
	if g.caches == nil {
		return nil, xerr.Wrap(errors.New("must init global cache first"))
	}
	for _, v := range g.caches.Items() {
		loadReq := v.Req
		// EQUALS
		if req.Equals(loadReq) {
			return v.Resp, nil
		}
	}
	// 找不到返回 err 也是 nil
	// 不要带来多余的麻烦
	return nil, nil
}
