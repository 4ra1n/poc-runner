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
	"os"
	"strconv"

	"github.com/4ra1n/poc-runner/base"
	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/expression"
	"github.com/4ra1n/poc-runner/xerr"

	"gopkg.in/yaml.v3"
)

func InitYamlPoCFromBytes(
	ctx context.Context,
	c *client.HttpClient,
	data []byte,
) (*base.POC, error) {
	var object interface{}
	err := yaml.Unmarshal(data, &object)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	root, ok := object.(map[string]interface{})
	if !ok {
		return nil, xerr.Wrap(errors.New("yaml is not a map"))
	}
	return parse(ctx, c, root)
}

func InitYamlPoC(s string) (map[string]interface{}, error) {
	data, err := os.ReadFile(s)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	var object interface{}
	err = yaml.Unmarshal(data, &object)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	root, ok := object.(map[string]interface{})
	if !ok {
		return nil, xerr.Wrap(errors.New("yaml is not a map"))
	}
	return root, nil
}

func InitYamlPoCFromInterface(
	ctx context.Context,
	c *client.HttpClient,
	object map[string]interface{},
) (*base.POC, error) {
	return parse(ctx, c, object)
}

func parse(ctx context.Context, c *client.HttpClient, root map[string]interface{}) (*base.POC, error) {
	var err error
	poc := &base.POC{
		Ctx: ctx,
	}
	poc.Env = expression.NewEnvironment(ctx, nil)
	poc.Context, err = base.NewPocContext(c)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	// ------------------------ NAME ------------------------
	poc.Name = root["name"].(string)
	// ------------------------ transport ------------------------
	poc.Transport = root["transport"].(string)
	// ------------------------ set ------------------------
	poc.Set = base.NewMap[string, *base.Expr]()
	if val, ok := root["set"]; ok {
		var sets map[string]interface{}
		sets, ok = val.(map[string]interface{})
		if !ok {
			return nil, xerr.Wrap(errors.New("set is not a map"))
		}
		for k, v := range sets {
			expr := &base.Expr{}
			expr.Env = poc.Env
			tempStr, isStr := v.(string)
			if isStr {
				expr.StrValue = tempStr
			}
			tempInt, isInt := v.(int)
			if isInt {
				expr.StrValue = strconv.Itoa(tempInt)
			}
			poc.Set.Set(k, expr)
		}
	}
	// ------------------------ payloads ------------------------
	poc.Payload = base.NewPayload()
	if val, ok := root["payloads"]; ok {
		mapVal, mapOK := val.(map[string]interface{})
		if !mapOK {
			return nil, xerr.Wrap(errors.New("payloads syntax error: outer payloads is not a map"))
		}
		innerVal, innerOK := mapVal["payloads"]
		if !innerOK {
			return nil, xerr.Wrap(errors.New("payloads syntax error: payloads must contains payloads"))
		}
		itemVal, itemOK := innerVal.(map[string]interface{})
		if !itemOK {
			return nil, xerr.Wrap(errors.New("payloads syntax error: inner payloads is not a map"))
		}
		for k, v := range itemVal {
			vv, vvOK := v.(map[string]interface{})
			if !vvOK {
				return nil, xerr.Wrap(errors.New("payloads syntax error: payloads items must be a map"))
			}
			itemMap := base.NewMap[string, *base.Expr]()
			for ik, iv := range vv {
				expr := &base.Expr{}
				expr.Env = poc.Env
				tempStr, isStr := iv.(string)
				if isStr {
					expr.StrValue = tempStr
				}
				tempInt, isInt := iv.(int)
				if isInt {
					expr.StrValue = strconv.Itoa(tempInt)
				}
				itemMap.Set(ik, expr)
			}
			poc.Payload.Set(k, itemMap)
		}
	}
	// 必须存在
	poc.Rules = base.NewMap[string, *base.Rule]()
	// ------------------------ rules ------------------------
	rules, ok := root["rules"]
	if !ok {
		return nil, xerr.Wrap(errors.New("rules is not exits"))
	}
	rulesVal, ok := rules.(map[string]interface{})
	if !ok {
		return nil, xerr.Wrap(errors.New("rules is not a map"))
	}
	for k, v := range rulesVal {
		rule := &base.Rule{}
		req := &base.Request{}
		expr := &base.Expr{}
		ruleItems, ruleOK := v.(map[string]interface{})
		if !ruleOK {
			return nil, xerr.Wrap(errors.New("invalid rule"))
		}
		ruleItem, reqOK := ruleItems["request"].(map[string]interface{})
		if !reqOK {
			return nil, xerr.Wrap(errors.New("rule must contain request"))
		}
		exprItem, exprOK := ruleItems["expression"].(string)
		if !exprOK {
			return nil, xerr.Wrap(errors.New("rule must contain expression"))
		}
		expr.Env = poc.Env
		expr.StrValue = exprItem
		req.Method = ruleItem["method"].(string)
		reqPath, pathExi := ruleItem["path"]
		if pathExi {
			req.Path = reqPath.(string)
		} else {
			req.Path = "/"
		}
		body, bodyExi := ruleItem["body"]
		if bodyExi {
			req.Body = body.(string)
		}
		cache, cacheExi := ruleItem["cache"]
		if cacheExi {
			req.Cache = cache.(bool)
		}
		fr, frExi := ruleItem["follow_redirect"]
		if frExi {
			req.FollowRedirect = fr.(bool)
		}
		headers, headersExi := ruleItem["headers"]
		if headersExi {
			headersVal, headersOK := headers.(map[string]interface{})
			if !headersOK {
				return nil, xerr.Wrap(errors.New("invalid headers"))
			}
			headersActual := base.NewMap[string, string]()
			for headerKey, headerValue := range headersVal {
				headersActual.Set(headerKey, headerValue.(string))
			}
			req.Headers = headersActual
		} else {
			req.Headers = base.NewMap[string, string]()
		}
		rule.Req = req
		rule.Expression = expr
		// OUTPUT
		rule.Output = &base.Output{}
		pocOutput, ok := ruleItems["output"]
		if ok {
			items, isMap := pocOutput.(map[string]interface{})
			if !isMap {
				return nil, xerr.Wrap(errors.New("output must be a map"))
			}
			search, searchOK := items["search"]
			if !searchOK {
				return nil, xerr.Wrap(errors.New("output must contain search"))
			}
			searchExpr, isStr := search.(string)
			if !isStr {
				return nil, xerr.Wrap(errors.New("search must be a string"))
			}
			rule.Output.Search = searchExpr
			delete(items, "search")
			for k, v := range items {
				rule.Output.ValueName = k
				vv, ok := v.(string)
				if !ok {
					return nil, xerr.Wrap(errors.New("search value must be a string"))
				}
				rule.Output.ValueValue = vv
			}
		}
		poc.Rules.Set(k, rule)
	}
	// GLOBAL EXPRESSION
	poc.Expression = &base.Expr{}
	expr, ok := root["expression"]
	if !ok {
		return nil, xerr.Wrap(errors.New("poc must have expression"))
	}
	exprStr, isStr := expr.(string)
	if !isStr {
		return nil, xerr.Wrap(errors.New("expression must be a string"))
	}
	poc.Expression.Env = poc.Env
	poc.Expression.StrValue = exprStr
	// ------------------------ detail ------------------------
	poc.Detail = base.NewMap[string, string]()
	return poc, nil
}
