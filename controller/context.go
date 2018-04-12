package controller

import (
	"net/http"
	"sync"
)

type Context struct {
	mu  sync.Mutex
	Rw  http.ResponseWriter
	Req *http.Request
}

func (ctx *Context) SetCtxRw(rw http.ResponseWriter) {
	ctx.mu.Lock()
	defer func() {
		ctx.mu.Unlock()
	}()
	ctx.Rw = rw
}

func (ctx *Context) SetCtxReq(req *http.Request) {
	ctx.mu.Lock()
	defer func() {
		ctx.mu.Unlock()
	}()
	ctx.Req = req
}

func (ctx Context) GetCtxRw() http.ResponseWriter {
	ctx.mu.Lock()
	defer func() {
		ctx.mu.Unlock()
	}()
	return ctx.Rw
}

func (ctx Context) GetCtxReq() *http.Request {
	ctx.mu.Lock()
	defer func() {
		ctx.mu.Unlock()
	}()
	return ctx.Req
}
