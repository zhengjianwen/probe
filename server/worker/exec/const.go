package exec

import "errors"

// dns error
var (
	ErrInvalidType   error = errors.New("invalid dns query type")
	ErrParseDomain   error = errors.New("domain could not be parsed")
	ErrInvalidDomain error = errors.New("invalid domain name")
	ErrIpUnMatch     error = errors.New("dns A query, ip not match")
	ErrDomainUnMatch error = errors.New("dns NS query, domain not match")
)

var (
	ErrStatusCodeUnMatch   error = errors.New("http response status code un match")
	ErrResponseHeadUnMatch error = errors.New("http response head un match")
	ErrResponseBodyUnMatch error = errors.New("http response body un match")
)
