package validators

import "net/url"

func IsValidUrl(token string) bool {
	_, err := url.ParseRequestURI(token)
	if err != nil {
		return false
	}

	u, err := url.Parse(token)
	if err != nil || u.Host == "" {
		return false
	}
	return true
}
