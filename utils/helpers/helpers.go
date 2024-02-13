package helpers

import "strings"

// RemoveDomainError removes all the commonly found
// prefixes from URL such as http, https, www
// then checks of the remaining string is the DOMAIN itself
func RemoveDomainError(url string, ourDomainUrl string) bool {

	if url == ourDomainUrl {
		return true
	}
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	return newURL == ourDomainUrl
}

func EnforceHTTP(url string) string {
	// make every url http and https
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}
