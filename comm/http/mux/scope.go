package mux

import (
	"bytes"
	"net/http"

	"golang.org/x/net/context"
)

func Scope(ctx context.Context) *RequestScope {
	return ctx.Value(RequestScopeKey).(*RequestScope)
}

type param struct {
	key, value []byte
}

var RequestScopeKey = "RequestScore"

type RequestScope struct {
	r          *http.Request
	w          http.ResponseWriter
	pathParams []*param

	reply []byte
}

func (rs *RequestScope) PathParam(key []byte) []byte {
	for _, p := range rs.pathParams {
		if bytes.Equal(p.key, key) {
			return p.value
		}
	}
	return nil
}

func (rs *RequestScope) SetPathParam(key, value []byte) {
	rs.pathParams = append(rs.pathParams, &param{key, value})
}

func (rs *RequestScope) SetReply(data []byte) {
	rs.reply = data
}

func (rs *RequestScope) Request() *http.Request {
	return rs.r
}

func (rs *RequestScope) ResponseWriter() http.ResponseWriter {
	rs.w.Header().Set("Content-Type", "application/json")
	return rs.w
}

func (rs *RequestScope) done() {
	if rs.w != nil && rs.reply != nil {
		rs.w.Write(rs.reply)
	}
}
