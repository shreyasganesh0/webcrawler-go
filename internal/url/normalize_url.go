package url

import (
	neturl "net/url"
)


func NormalizeURL(inputURL string) (string, error) {


	url, err := neturl.Parse(inputURL)
	if err != nil {

		return "", err
	}

	normalizedURL := url.Host + url.Path

	return normalizedURL, nil
}
