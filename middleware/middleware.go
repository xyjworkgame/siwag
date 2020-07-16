package middleware

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
)



func ReadBody(req *http.Request) *string {
	save := req.Body
	var err error
	if req.Body == nil {
		req.Body = nil
	} else {
		save, req.Body, err = drainBody(req.Body)
		if err != nil {
			return nil
		}
	}
	b := bytes.NewBuffer([]byte(""))
	chunked := len(req.TransferEncoding) > 0 && req.TransferEncoding[0] == "chunked"
	if req.Body == nil {
		return nil
	}
	var dest io.Writer = b
	if chunked {
		dest = httputil.NewChunkedWriter(dest)
	}
	_, err = io.Copy(dest, req.Body)
	if chunked {
		dest.(io.Closer).Close()
	}
	req.Body = save
	body := b.String()
	return &body
}
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, nil, err
	}
	if err = b.Close(); err != nil {
		return nil, nil, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

var reqWriteExcludeHeaderDump = map[string]bool{
	"Host":              true, // not in Header map anyway
	"Content-Length":    true,
	"Transfer-Encoding": true,
	"Trailer":           true,
	"Accept-Encoding":   false,
	"Accept-Language":   false,
	"Cache-Control":     false,
	"Connection":        false,
	"Origin":            false,
	"User-Agent":        false,
}

func ReadPostForm(req *http.Request) map[string]string {
	postForm := map[string]string{}
	for _, param := range strings.Split(*ReadBody(req), "&") {
		value := strings.Split(param, "=")
		postForm[value[0]] = value[1]
	}
	return postForm
}

func ReadHeaders(req *http.Request) map[string]string {
	b := bytes.NewBuffer([]byte(""))
	err := req.Header.WriteSubset(b, reqWriteExcludeHeaderDump)
	if err != nil {
		return map[string]string{}
	}
	headers := map[string]string{}
	for _, header := range strings.Split(b.String(), "\n") {
		values := strings.Split(header, ":")
		if strings.EqualFold(values[0], "") {
			continue
		}
		headers[values[0]] = values[1]
	}
	return headers
}
