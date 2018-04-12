package goserver

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"goser/controller"
	"log"
	"net/http"
	"strings"
	"time"
)

func middlewareHandler(next http.Handler) http.Handler {
	//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	//    // 执行handler之前的逻辑
	//    next.ServeHTTP(w, r)
	//    // 执行完毕handler后的逻辑
	//})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		dur := time.Since(start)
		u := strings.Split(r.URL.String(), "?")
		fmt.Println(u)
		//panic("error")
		log.Printf("%s\t%s\t%s\n", r.Method, u[0], dur)
	})
}

func AddRouter(url string, handler controller.HandlerInterface) {
	//	valid url
	//	http.Handle(url, middlewareHandler(http.HandlerFunc(handler)))
	env := &controller.Env{}
	controller.MapHandler[url] = &handler
	http.Handle(url, &controller.Handler{env, nil, handler.Get, handler.Post})
}

func Run() {
	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		log.Println("load conf file", err)
		return
	}
	controller.MapHandler = make(map[string]*controller.HandlerInterface)
	for i := 0; i < len(allRouters); i++ {
		AddRouter(allRouters[i].url, allRouters[i].handler)
	}
	val := cfg.MustInt("", "httpport")
	log.Println("port: ", val)
	host := fmt.Sprintf("127.0.0.1:%d", val)
	http.ListenAndServe(host, nil)
}
