package validators

import (
	"net"
	"net/url"
	"strings"
)

func IsValidUrl(str string) bool {
	linkurl, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}

	address := net.ParseIP(linkurl.Host)

	if address == nil {
		return strings.Contains(linkurl.Host, ".")
	}

	return true
}
