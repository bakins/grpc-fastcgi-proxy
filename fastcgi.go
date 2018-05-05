package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	fcgi "github.com/tomasen/fcgi_client"
)

type fastcgiResponse struct {
	code   int
	body   []byte
	header http.Header
}

func paramsFromRequest(r *http.Request) map[string]string {
	params := map[string]string{
		"REQUEST_METHOD":    r.Method,
		"SERVER_PROTOCOL":   fmt.Sprintf("HTTP/%d.%d", r.ProtoMajor, r.ProtoMinor),
		"HTTP_HOST":         r.Host,
		"CONTENT_LENGTH":    fmt.Sprintf("%d", r.ContentLength),
		"CONTENT_TYPE":      r.Header.Get("Content-Type"),
		"REQUEST_URI":       r.URL.Path,
		"SCRIPT_NAME":       r.URL.Path,
		"GATEWAY_INTERFACE": "CGI/1.1",
		"QUERY_STRING":      r.URL.RawQuery,
	}

	for k, v := range r.Header {
		params["HTTP_"+strings.Replace(strings.ToUpper(k), "-", "_", -1)] = v[0]
	}

	delete(params, "HTTP_PROXY")

	return params
}

func (s *Server) request(r *http.Request, params map[string]string) (*fastcgiResponse, error) {
	resp := &fastcgiResponse{}

	f, err := fcgi.DialTimeout(s.fastEndpoint.Scheme, s.fastEndpoint.Host, 3*time.Second)
	if err != nil {
		resp.code = 500
		return resp, errors.Wrap(err, "fastcgi dial failed")
	}

	response, err := f.Request(params, r.Body)
	if err != nil {
		resp.code = 500
		return resp, errors.Wrap(err, "failed to make fastcgi request")
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		resp.code = 500
		return resp, errors.Wrap(err, "failed to read response from fastcgi")
	}

	resp.code = response.StatusCode
	resp.body = content
	resp.header = response.Header
	resp.code, err = statusFromHeaders(resp.header)

	if err != nil {
		resp.code = 500
		return resp, errors.Wrap(err, "failed to get status")
	}

	return resp, nil
}
func statusFromHeaders(h http.Header) (int, error) {
	text := h.Get("Status")

	h.Del("Status")

	if text == "" {
		return 200, nil
	}

	return strconv.Atoi(text[0:3])
}
