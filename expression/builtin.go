package expression

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/4ra1n/poc-runner/xerr"
	"hash"
	"net/url"
	"reflect"
)

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
}
