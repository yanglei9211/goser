package main

import (
	"fmt"
	"goser/goserver"
	"net/http"
)

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("aaaaa")
	r.ParseForm()
	w.Write([]byte("123"))
}

func SecHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/a", 302)
	//fmt.Println("bbbbbb")
	//r.ParseForm()
	//w.Write([]byte("234"))
}

func main() {
	//http.HandleFunc("/", FirstHandler)
	//http.ListenAndServe("127.0.0.1:8088", nil)
	//cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	//if err != nil {
	//	log.Println("load conf file", err)
	//	return
	//}
	//val := cfg.MustInt("", "httpport")
	//fmt.Println(val)
	//goserver.AddRouter("/a", FirstHandler)
	//goserver.AddRouter("/b", SecHandler)
	goserver.Run()
}
