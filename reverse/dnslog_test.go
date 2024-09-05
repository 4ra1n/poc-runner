package reverse

import (
	"fmt"
	"github.com/4ra1n/poc-runner/client"
	"github.com/4ra1n/poc-runner/rawhttp"
	"testing"
)

func TestDNSLOG(t *testing.T) {
	c, err := client.NewHttpClient(rawhttp.DefaultNoProxy, rawhttp.DefaultTimeout, false)
	if err != nil {
		panic(err)
	}
	rev, err := NewReverse(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(rev)
	rev.Wait(20)
}
