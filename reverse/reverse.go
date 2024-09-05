package reverse

// Reverse
// 通用反连接口
// 目前是 dnslog.cn
// 后续可以拓展
type Reverse interface {
	// GetUrl
	// HTTP 反连
	GetUrl() string
	// GetRmi
	// RMI 反连
	GetRmi() string
	// GetLdap
	// LDAP 反连
	GetLdap() string
	// GetDNS
	// DNS 反连
	GetDNS() string
	// Wait
	// 等待
	Wait(int) bool
}
