package fetcher

import (
	"net/http"
)

func Fetch(url string, method string, headers map[string]string) (*http.Response, error) {

	var httpResp *http.Response

	switch method {
	case "GET":
		resp, err := Get(url, headers)

		if err != nil {
			return resp, err
		}

		httpResp = resp
	}

	return httpResp, nil

}
