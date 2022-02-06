package fetcher

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var server *httptest.Server
var serverURL string
var fetchTest = []struct {
	method             string
	url                string
	headers            map[string]string
	expectedResp       *http.Response
	expectedRespStr    string
	expectedStatusCode int
	expectedErr        error
}{
	{"GET", "http://localhost", map[string]string{
		"X-API-Key": "6a79b4fd-ccd2-4f9f-976e-e0e2c15ab91a",
	}, &http.Response{}, "welcome", 200, nil},
}

func TestMain(m *testing.M) {

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		xApiKey := r.Header.Get("X-API-Key")

		if len(xApiKey) < 1 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if xApiKey != XAPIKey {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Write([]byte("welcome"))

	}))

	serverURL = server.URL

	m.Run()

}

func TestFecth(t *testing.T) {

	for _, v := range fetchTest {
		resp, err := Fetch(serverURL, v.method, v.headers)

		if err != v.expectedErr {
			t.Errorf("expected %v got %v", v.expectedErr, err)
		}

		if resp == nil {
			t.Errorf("expected http response not nil, got %v", resp)
		}

		//read the resp
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		respString := string(bodyBytes)
		if respString != v.expectedRespStr {
			t.Errorf("expected %v got %v", v.expectedRespStr, respString)
		}

	}

}
