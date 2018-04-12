package controller

import (
	"fmt"
	"net/http"
)

func FirstHandler2(env *Env, w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()
	x, _ := r.Form["x"]
	y, _ := r.Form["y"]
	fmt.Println(x)
	fmt.Println(y)
	w.Write([]byte("rec ok"))
	return nil
}

type FirstHandler struct {
	BaseHandler
}

func (h *FirstHandler) Get() {
	h.Ctx.Req.ParseForm()
	x := 12
	y := 45
	fmt.Println(x / y)
	fmt.Println("read do get")
	h.Ctx.Rw.Write([]byte("rec ok 1111111"))
}
