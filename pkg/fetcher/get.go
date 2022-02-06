package fetcher

import (
	"net"
	"net/http"
	"strings"
	"time"
)

//Get makes an HTTP request and returns the HTTP response
func Get(url string, headers map[string]string) (*http.Response, error) {

	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConnsPerHost:   30000,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   120 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)

	req.Close = true

	for k, v := range headers {
		req.Header.Set(strings.TrimSpace(k), strings.TrimSpace(v))
	}

	if err != nil {

		return &http.Response{}, err
	}

	// ctx, cancel := context.WithTimeout(req.Context(), 120*time.Second)
	// defer cancel()
	// req = req.WithContext(ctx)

	res, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return res, nil

}
