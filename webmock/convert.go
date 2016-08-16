package webmock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elazarl/goproxy"
)

type connection struct {
	Request    *request  `json:"request"`
	Response   *response `json:"response"`
	RecordedAt string    `json:"recorded_at"`
}

type request struct {
	Header *header `json:"header"`
	String string  `json:"string"`
	Method string  `json:"method"`
	URL    string  `json:"url"`
}

type response struct {
	Status string  `json:"status"`
	Header *header `json:"header"`
	String string  `json:"string"`
}

type header struct {
	ContentType   string `json:"Content-Type"`
	ContentLength string `json:"Content-Length"`
}

type responseBody struct {
	Message string `json:"message"`
}

func structToJSON(v interface{}) (string, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	out := new(bytes.Buffer)
	json.Indent(out, jsonBytes, "", "    ")
	return out.String(), nil
}

func jsonToStruct(b []byte) *connection {
	var conn connection
	err := json.Unmarshal(b, &conn)
	if err != nil {
		// TODO: Add error handling
		fmt.Println("JSON Marshal Error: ")
		return &conn
	}
	return &conn
}

func createReqStruct(body string, ctx *goproxy.ProxyCtx) *request {
	contentType := ctx.Req.Header.Get("Content-Type")
	contentLength := ctx.Req.Header.Get("Content-Length")
	header := &header{contentType, contentLength}
	method := ctx.Req.Method
	host := ctx.Req.URL.Host
	path := ctx.Req.URL.Path

	return &request{header, body, method, host + path}
}

func createRespStruct(b []byte, ctx *goproxy.ProxyCtx) *response {
	contentType := ctx.Resp.Header.Get("Content-Type")
	contentLength := ctx.Resp.Header.Get("Content-Length")
	header := &header{contentType, contentLength}
	body := strings.TrimRight(string(b), "\n")

	return &response{ctx.Resp.Status, header, body}
}

func createErrorMessage(str string) (string, error) {
	mes := "Not found webmock-proxy cache. URL: " + str
	body := &responseBody{mes}
	return structToJSON(body)
}