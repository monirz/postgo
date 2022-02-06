package fetcher

import (
	"testing"
)

const XAPIKey = "6a79b4fd-ccd2-4f9f-976e-e0e2c15ab91a"

var getTests = []struct {
	url               string
	headers           map[string]string
	expectedStausCode int
}{
	{url: "http://localhost:8090", headers: map[string]string{
		"X-API-Key": "6a79b4fd-ccd2-4f9f-976e-e0e2c15ab91a",
	}, expectedStausCode: 200,
	},
	{url: "http://localhost:8090", headers: map[string]string{}, expectedStausCode: 401},
}

func TestGet(t *testing.T) {

	for _, v := range getTests {
		resp, err := Get(server.URL, v.headers)

		if err != nil {
			t.Errorf("expected %v, got %v", nil, err)
		}

		if resp.StatusCode != v.expectedStausCode {
			t.Errorf("expected %v got, %v", v.expectedStausCode, resp.StatusCode)
		}

	}

}
