package exec

import (
	"fmt"
	pb "github.com/ten-cloud/prober/server/proto"
	"net"
	"testing"
)

func TestRemoteResolve(t *testing.T) {
	tk := pb.Task_Dns{
		Type:      pb.Task_Dns_A,
		Domain:    "www.baidu.com",
		DNSServer: "114.114.114.114:53",
	}
	if err := RemoteResolve(&tk); err != nil {
		t.Fatal(err)
	}
}

func TestLocalResolve(t *testing.T) {
	tk := pb.Task_Dns{
		Type:   pb.Task_Dns_A,
		Domain: "www.baidu.com",
	}
	if err := LocalResolve(&tk); err != nil {
		t.Fatal(err)
	}
}

func TestLocalResolve_MatchIps(t *testing.T) {
	var domain string = "www.baidu.com"
	ips, err := net.LookupHost(domain)
	if err != nil {
		t.Fatal(err)
	}
	if len(ips) == 0 {
		t.Fatal("no ip found")
	}

	tk := pb.Task_Dns{
		Type:      pb.Task_Dns_A,
		Domain:    domain,
		IfMatchIp: true,
		MatchIps:  ips,
	}

	if err := LocalResolve(&tk); err != nil {
		t.Fatal(err)
	}

	tk.MatchIps = []string{"1.1.1.1"}
	if err := LocalResolve(&tk); err != ErrIpUnMatch {
		t.Fatal(err)
	}
}

func TestLocalResolve_MatchDomains(t *testing.T) {
	var domain string = "jiankongbao.com"
	domains, err := net.LookupNS(domain)
	if err != nil {
		t.Fatal(err)
	}
	if len(domains) == 0 {
		t.Fatal("no domain found")
	}

	var l []string
	for _, domain := range domains {
		fmt.Println(domain.Host)
		l = append(l, domain.Host)
	}

	tk := pb.Task_Dns{
		Type:          pb.Task_Dns_NS,
		Domain:        domain,
		IfMatchDomain: true,
		MatchDomains:  l,
	}

	if err := LocalResolve(&tk); err != nil {
		t.Fatal(err)
	}

	tk.MatchDomains = []string{"www.g.com"}
	if err := LocalResolve(&tk); err != ErrDomainUnMatch {
		t.Fatal(err)
	}
}
