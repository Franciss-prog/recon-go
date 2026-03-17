package validation

import (
	"fmt"
	"net/url"
	"strings"
)

func FormatUrl(link string) bool {

	// check if the length of the link is greater than 0
	if len(link) == 0 || len(link) >= 255 {
		fmt.Println("Please enter a valid link.")
		return false
	}

	// check the scheme of the url
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		// this condition is just incase cuz we wont know if the user will enter http or https
		link = "http://" + link
	}

	// check the url if the url is valid
	_, err := url.ParseRequestURI(link)
	if err != nil {
		fmt.Errorf("Invalid URL: %v", err)
		return false
	}
	return true
}
