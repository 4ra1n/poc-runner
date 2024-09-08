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
	"errors"
	"strings"
	"unicode"

	"github.com/4ra1n/poc-runner/base"
	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/expression"
	"github.com/4ra1n/poc-runner/log"
	"github.com/4ra1n/poc-runner/xerr"
)

func RunPOC(poc *base.POC, t string) (bool, error) {
	poc.Target = t
	// 执行 set 部分设置环境变量
	err := execSet(poc)
	if err != nil {
		return false, xerr.Wrap(err)
	}
	log.Info("run set expr finish")
	return execGlobalRules(poc)
}

func execSet(poc *base.POC) error {
	pocSet := poc.Set

	// FIX BUG
	// GOLANG YAML 是无序的
	// 但是 newReverse 一定要在 reverse.url 之前执行
	setMap := base.NewMap[string, *base.Expr]()
	reverse := pocSet.Get("reverse")
	if reverse != nil {
		setMap.Set("reverse", reverse)
	}
	for _, k := range pocSet.Keys() {
		setMap.Set(k, pocSet.Get(k))
	}

	for _, k := range setMap.Keys() {
		v := setMap.Get(k)
		expr := v.StrValue
		vars := expression.MapVars{}
		for kk, vv := range poc.Context.Local.ToMap() {
			vars[kk] = vv
		}
		res, err := poc.Env.EvalWithVars(expr, vars)
		if err != nil {
			return xerr.Wrap(err)
		}
		if res != nil {
			log.Infof("set %s -> %s", k, res.ToString())
			poc.Context.Local.Set(k, res)
		}
	}
	return nil
}

func execRule(poc *base.POC, key string, rule *base.Rule) (expression.EValue, error) {
	log.Infof("run rule: %s", key)
	ruleReq := rule.Req
	reqPath := replaceVar(ruleReq.Path, poc.Locals())
	method := ruleReq.Method

	headers := ruleReq.Headers
	headersMap := make(map[string]string)
	for _, k := range headers.Keys() {
		headersMap[k] = replaceVar(headers.Get(k), poc.Locals())
	}

	ruleReq.Body = strings.TrimRightFunc(ruleReq.Body, unicode.IsSpace)
	ruleReq.Body = replaceVar(ruleReq.Body, poc.Locals())

	theReq := &client.TheRequest{
		Target:         poc.Target,
		Method:         method,
		Path:           reqPath,
		FollowRedirect: ruleReq.FollowRedirect,
		Body:           ruleReq.Body,
		Headers:        headersMap,
	}
	poc.SetReq(key, theReq)

	var (
		err        error
		cachedResp *client.TheResponse
		theResp    *client.TheResponse
	)
	// if user use cache
	if ruleReq.Cache {
		// try to get cache
		cachedResp, err = poc.Caches.GetCache(theReq)
		if err != nil {
			return nil, xerr.Wrap(err)
		}
		if cachedResp == nil {
			// if cache is null
			// do request
			theResp, err = poc.DoReq(theReq)
			if err != nil {
				return nil, xerr.Wrap(err)
			}
		} else {
			theResp = cachedResp
		}
		// set cache
		err = poc.Caches.SetCache(poc.Name, key, theReq, theResp)
		if err != nil {
			return nil, xerr.Wrap(err)
		}
	} else {
		// directly do request
		theResp, err = poc.DoReq(theReq)
		if err != nil {
			return nil, xerr.Wrap(err)
		}
	}
	if theResp == nil {
		// this should not happen
		return nil, xerr.Wrap(errors.New("response is null"))
	}
	poc.SetResp(key, theResp)

	respValue := &RespValue{
		Status: expression.EInt(theResp.Code),
		Body:   expression.EBytes(theResp.Body),
		Headers: &RespHeaderValue{
			Headers: theResp.Headers,
		},
	}
	vars := expression.MapVars{"response": respValue}

	for k, v := range poc.Context.Local.ToMap() {
		vars[k] = v
	}

	value, err := poc.Env.EvalWithVars(rule.Expression.StrValue, vars)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	result, ok := value.(expression.EBool)
	if !ok {
		return nil, xerr.Wrap(errors.New("rule result must be a bool"))
	}
	resultBool := bool(result)
	log.Infof("rule [%s] result [%v]", key, resultBool)
	poc.SetResult(key, resultBool)

	// EVAL OUTPUT
	expr := rule.Output.Search
	if expr == "" {
		return expression.EBool(resultBool), nil
	}
	value, err = poc.Env.EvalWithVars(expr, vars)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	searchResult := expression.MapVars{"search": value}
	searchRet, err := poc.Env.EvalWithVars(rule.Output.ValueValue, searchResult)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	log.Infof("output [%s] -> [%s]", rule.Output.ValueName, searchRet)
	poc.SetLocal(rule.Output.ValueName, searchRet)

	return expression.EBool(resultBool), nil
}

func execGlobalRulesInternal(poc *base.POC) (bool, error) {
	var rules = make(Rule)
	for _, ruleKey := range poc.Rules.Keys() {
		rules[ruleKey] = &RuleFunction{
			poc:  poc,
			key:  ruleKey,
			rule: poc.Rules.Get(ruleKey),
		}
	}
	value, err := poc.Env.EvalWithVars(poc.Expression.StrValue, rules)
	if err != nil {
		return false, xerr.Wrap(err)
	}
	result, ok := value.(expression.EBool)
	if !ok {
		return false, xerr.Wrap(errors.New("rule result must be a bool"))
	}
	resultBool := bool(result)
	log.Infof("rule [global] result [%v]", resultBool)
	return resultBool, nil
}

func execGlobalRules(poc *base.POC) (bool, error) {
	if poc.Payload.Available() {
		pMap := poc.Payload.Get()
		for _, key := range pMap.Keys() {
			v := pMap.Get(key)
			for _, innerKey := range v.Keys() {
				innerVal := v.Get(innerKey)
				expr := innerVal.StrValue
				vars := expression.MapVars{}
				for kk, vv := range poc.Context.Local.ToMap() {
					vars[kk] = vv
				}
				res, err := poc.Env.EvalWithVars(expr, vars)
				if err != nil {
					return false, xerr.Wrap(err)
				}
				if res != nil {
					log.Infof("payloads [%s] %s -> %s", key, innerKey, res.ToString())
					poc.Context.Local.Set(innerKey, res)
				}
			}
			success, err := execGlobalRulesInternal(poc)
			if err != nil {
				return false, xerr.Wrap(err)
			}
			log.Infof("payloads [%s] result [%v]", key, success)
			if success {
				return true, nil
			}
		}
		return false, nil
	} else {
		success, err := execGlobalRulesInternal(poc)
		if err != nil {
			return false, xerr.Wrap(err)
		}
		return success, nil
	}
}
