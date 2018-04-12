package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
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

func (h Handler) Query(arg string) string {
	res := h.Ctx.Req.Form.Get(arg)
	return res
}

func (h *Handler) SetCtx(ctx *Context) {
	h.Ctx = ctx
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
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
	dur := time.Since(start)
	u := strings.Split(r.URL.String(), "?")
	log.Printf("%s\t%s\t%s\n", r.Method, u[0], dur)
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
			//v := reflect.ValueOf(er)
			if e, ok := er.(StatusError); ok {
				err = e
			} else {
				err = StatusError{Code: 500, Err: errors.New(fmt.Sprintf("%s", er))}
			}
			//err = errors.New("inner error")
		}
	}()
	u := strings.Split(r.URL.String(), "?")
	if handlerInfo, found := MapHandler[u[0]]; !found {
		http.Error(w, http.StatusText(http.StatusNotFound),
			http.StatusNotFound)
		return nil
	} else {
		r.ParseForm()
		ctx := Context{}
		ctx.SetCtxRw(w)
		ctx.SetCtxReq(r)
		(*handlerInfo).SetCtx(&ctx)
		method := r.Method
		switch method {
		case "GET":
			(*handlerInfo).Get()
			return nil
		case "POST":
			(*handlerInfo).Post()
			return nil
		default:
			http.Error(w, http.StatusText(http.StatusNotFound),
				http.StatusNotFound)
			return nil
		}
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
	err := StatusError{Code: 405, Err: errors.New("method not allowed")}
	panic(err)
	//panic("http error 405, method not allowed")
}

func (b BaseHandler) Post() {
	err := StatusError{Code: 405, Err: errors.New("method not allowed")}
	panic(err)
}

func (b BaseHandler) writeResponse(r map[string]interface{}) {
	response, err := json.Marshal(
		map[string]interface{}{
			"status": 1,
			"data":   r,
		})
	if err != nil {
		panic(err)
	}
	b.Ctx.Rw.Write([]byte(string(response)))
}

type HandlerInterface interface {
	Get()
	Post()
	SetCtx(*Context)
	//	todo more method
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
