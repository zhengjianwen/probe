package exec

import (
	"fmt"
	"github.com/miekg/dns"
	pb "github.com/rongyungo/probe/server/proto"
	"golang.org/x/net/idna"
	"net"
	"strings"
	"time"
)

func ProbeDns(t *pb.Task) *pb.TaskResult {
	if t.DnsSpec == nil {
		return nil
	}

	querySpec := t.DnsSpec
	now := time.Now().UnixNano()

	if querySpec.ServerDesigned {
		return Return(t.GetBasicInfo().GetId(), t.GetBasicInfo().GetType(), RemoteResolve(querySpec), now)
	} else {
		return Return(t.GetBasicInfo().GetId(), t.GetBasicInfo().GetType(), LocalResolve(querySpec), now)
	}
}

func LocalResolve(d *pb.DnsSpec, tp ...pb.DnsSpecType) error {
	domain, queryType := d.Domain, d.Type
	if len(tp) > 0 {
		queryType = tp[0]
	}

	switch queryType {
	case pb.DnsSpec_A:
		addrs, err := net.LookupHost(domain)
		if err != nil {
			return err
		}

		if d.IfMatchIp && len(d.MatchIps) > 0 {
			for _, ip := range d.MatchIps {
				if isStrArrMatchStr(addrs, ip) {
					return nil
				}
			}
			return ErrIpUnMatch
		}

	case pb.DnsSpec_NS:
		nss, err := net.LookupNS(domain)
		if err != nil {
			return err
		}

		if d.IfMatchDomain && len(d.MatchDomains) > 0 {
			for _, domain := range d.MatchDomains {
				if isNSArrMatchStr(nss, domain) {
					return nil
				}
			}
			return ErrDomainUnMatch
		}
	case pb.DnsSpec_MX:
		_, err := net.LookupMX(domain)
		return err
	case pb.DnsSpec_CNAME:
		_, err := net.LookupCNAME(domain)
		return err
	case pb.DnsSpec_TXT:
		_, err := net.LookupTXT(domain)
		return err
	case pb.DnsSpec_ANY:
		if err := LocalResolve(d, pb.DnsSpec_A); err != nil {
			return err
		}
		if err := LocalResolve(d, pb.DnsSpec_NS); err != nil {
			return err
		}
		if err := LocalResolve(d, pb.DnsSpec_MX); err != nil {
			return err
		}
		if err := LocalResolve(d, pb.DnsSpec_CNAME); err != nil {
			return err
		}
		if err := LocalResolve(d, pb.DnsSpec_TXT); err != nil {
			return err
		}
	}

	return nil
}

func RemoteResolve(d *pb.DnsSpec) error {
	domain, server, queryType := d.Domain, d.DNSServer, d.Type.String()

	if !strings.HasSuffix(domain, ".") {
		domain = fmt.Sprintf("%s.", domain)
	}

	var err error
	domain, err = idna.ToASCII(domain)
	if err != nil {
		return ErrParseDomain
	}

	if _, ok := dns.IsDomainName(domain); !ok {
		return ErrInvalidDomain
	}

	tp, ok := dns.StringToType[strings.ToUpper(queryType)]
	if !ok {
		return ErrInvalidType
	}

	m := new(dns.Msg)
	m.SetQuestion(domain, tp)
	m.MsgHdr.RecursionDesired = true

	c := new(dns.Client)

	msg, _, err := c.Exchange(m, server)
	if err != nil {
		return err
	}

	return matchHandler(d, msg)
}

func matchHandler(d *pb.DnsSpec, msg *dns.Msg) error {
	switch d.Type {
	case pb.DnsSpec_A:
		if d.IfMatchIp && len(d.MatchIps) > 0 {
			for _, ip := range d.MatchIps {
				if isAnswerMatchStr(msg.Answer, ip) {
					return nil
				}
			}
			return ErrIpUnMatch
		}
	case pb.DnsSpec_NS:
		if d.IfMatchDomain && len(d.MatchDomains) > 0 {
			for _, domain := range d.MatchDomains {
				if isAnswerMatchStr(msg.Answer, domain) {
					return nil
				}
			}
			return ErrDomainUnMatch
		}
	}
	return nil
}
