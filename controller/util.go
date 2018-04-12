package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

type Env struct {
}

type Handler struct {
	*Env
	Ctx *Context
	//H func(e *Env, w http.ResponseWriter, r *http.Request) error
	Get  func()
	Post func()
	//	todo more method
}

func (h *Handler) SetCtx(ctx *Context) {
	h.Ctx = ctx
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

//func (h Handler) innerFunc(e *Env, w http.ResponseWriter, r *http.Request) (er error) {
//	defer func() {
//		if err := recover(); err != nil {
//			log.Println(err)
//			er = errors.New("inner error")
//		}
//	}()
//	er = h.H(e, w, r)
//	return
//}

func (h *Handler) H(env *Env, w http.ResponseWriter, r *http.Request) (err error) {
	defer func() {
		if er := recover(); er != nil {
			err = StatusError{Code: 500, Err: errors.New(fmt.Sprintf("%s", er))}
			//err = errors.New("inner error")
		}
	}()
	u := strings.Split(r.URL.String(), "?")
	if handlerInfo, found := MapHandler[u[0]]; !found {
		http.Error(w, http.StatusText(http.StatusNotFound),
			http.StatusNotFound)
		return nil
	} else {
		ctx := Context{}
		ctx.SetCtxRw(w)
		ctx.SetCtxReq(r)
		(*handlerInfo).SetCtx(&ctx)
		(*handlerInfo).Get()
		return nil
	}

	//method := r.Method
	//switch method {
	//case "GET":
	//	h.Get()
	//	return nil
	//default:
	//	http.Error(w, http.StatusText(http.StatusMethodNotAllowed),
	//	http.StatusMethodNotAllowed)
	//}
	//return nil
}

type BaseHandler struct {
	Handler
}

func (b BaseHandler) Get() {
	fmt.Println("mid do get")
}

type HandlerInterface interface {
	Get()
	SetCtx(*Context)
}

//type HandlerInterface interface {
//	Get() func()
//	Post() func()
//WriteResponse() func(interface{})
//	todo more method
//}
//func FirstHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
//	fmt.Println(r.Method)
//	r.ParseForm()
//	fmt.Println(r.Form)
//	//w.Write([]byte("ack ok"))
//	return StatusError{502, errors.New("eee")}
//}

var MapHandler map[string]*HandlerInterface
