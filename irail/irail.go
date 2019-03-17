package irail

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	base      = "https://api.irail.be"
	cacheDir  = "ijzerweg-cache"
	userAgent = "Ijzerweg/0.1 (https://github.com/cimm/ijzerweg)"
	apiFormat = "json"
)

// Fetch builds the endpoint from the base, path, and params after
// which it loads the respnse body from this endpoint.
func Fetch(path string, params map[string]string) []byte {
	url := endpoint(path, params)
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", userAgent)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func endpoint(path string, params map[string]string) string {
	return base + path + "?" + paramsToQuery(params)
}

func paramsToQuery(params map[string]string) string {
	params["format"] = apiFormat
	tuples := []string{}
	for k, v := range params {
		t := strings.Builder{}
		t.WriteString(k)
		t.WriteString("=")
		t.WriteString(v)
		tuples = append(tuples, t.String())
	}
	return strings.Join(tuples, "&")
}
