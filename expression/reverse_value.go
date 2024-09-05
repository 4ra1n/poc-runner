package expression

import (
	"errors"
	"fmt"
	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/rawhttp"
	"github.com/4ra1n/poc-runner/reverse"
	"github.com/4ra1n/poc-runner/xerr"
)

type Reverse struct {
	rev reverse.Reverse
}

func NewReverse() (*Reverse, error) {
	cl, err := client.NewHttpClient(rawhttp.DefaultNoProxy, rawhttp.DefaultTimeout, false)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	r, err := reverse.NewReverse(cl)
	if err != nil {
		return nil, xerr.Wrap(err)
	}
	return &Reverse{
		rev: r,
	}, nil
}

func (r *Reverse) ToString() EString {
	return EString(fmt.Sprintf("reverse-%p", r.rev))
}

func (r *Reverse) Get(name string) (EValue, error) {
	switch name {
	case "url":
		return EString(r.rev.GetUrl()), nil
	case "rmi":
		return EString(r.rev.GetRmi()), nil
	case "ldap":
		return EString(r.rev.GetLdap()), nil
	case "domain":
		return EString(r.rev.GetDNS()), nil
	case "wait":
		return &eFunctionReverseWait{r.rev}, nil
	default:
		return nil, xerr.Wrap(errors.New("not support: " + name))
	}
}

func (r *Reverse) Keys() []string {
	return []string{}
}
