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

package expression

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"net/url"
	"reflect"

	"github.com/4ra1n/poc-runner/xerr"
)

// 这是内置的一些函数
// 有一些实现不够优雅 计划未来重构
var builtin = map[string]EFunction{
	"bytes":           &eFunctionToBytes{},
	"string":          &eFunctionToString{},
	"substr":          &eFunctionSubstr{},
	"sleep":           &eFunctionSleep{},
	"randomInt":       &eFunctionRandomInt{},
	"randomLowercase": &eFunctionRandomAlpha{upper: false},
	"randomUppercase": &eFunctionRandomAlpha{upper: true},
	"md5": &eFunctionHash{h: func() hash.Hash {
		return md5.New()
	}},
	"sha1": &eFunctionHash{h: func() hash.Hash {
		return sha1.New()
	}},
	"sha256": &eFunctionHash{h: func() hash.Hash {
		return sha256.New()
	}},
	"base64": &eFunctionCodec{fun: func(value EValue) (EValue, error) {
		switch value.(type) {
		case EBytes:
			return EString(base64.StdEncoding.EncodeToString(value.(EBytes))), nil
		case EString:
			return EString(base64.StdEncoding.EncodeToString([]byte(value.(EString)))), nil
		default:
			return nil, xerr.Wrap(fmt.Errorf("expect EString,EBytes, got %s", reflect.TypeOf(value)))
		}
	}},
	"urldecode": &eFunctionCodec{fun: func(value EValue) (EValue, error) {
		var v string
		switch value.(type) {
		case EBytes:
			v = string(value.(EBytes))
		case EString:
			v = string(value.(EString))
		default:
			return nil, xerr.Wrap(fmt.Errorf("expect EString,EBytes, got %s", reflect.TypeOf(value)))
		}
		decodedValue, err := url.QueryUnescape(v)
		if err != nil {
			return nil, xerr.Wrap(err)
		}
		return EString(decodedValue), err
	}},
	"get404Path": &eFunctionGet404Path{},
	"newReverse": &eFunctionNewReverse{},
	"print":      &eFunctionPrint{},
}
