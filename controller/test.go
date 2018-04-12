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
	var x, y int
	MustInt(&x, h.Query("x"))
	MustInt(&y, h.Query("y"))
	h.writeResponse(map[string]interface{}{
		"rec": x + y,
	})
}
