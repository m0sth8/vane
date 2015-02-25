package utils

import (
	"strings"
)

func AddTrailingSlash(url string) string {
	if !strings.HasSuffix(url, "/") {
		return url + "/"
	}
	return url
}

func AddHttpProtocol(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}
	return "http://" + url
}
