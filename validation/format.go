package validation

import (
	"fmt"
	"net/url"
	"strings"
)

func FormatUrl(url string) bool {

	// check if the length of the link is greater than 0
	if len(url) == 0 || len(url) >= 255 {
		fmt.Println("Please enter a valid link.")
		return false
	}

	// check the scheme
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return false
	}
	// check the url if the url is valid
	_, err := url.ParseRequestURI(url)
	if err != nil {
		return false
	}
	return true
}
