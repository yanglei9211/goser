package goserver

import (
	"goser/controller"
)

type UrlHandler struct {
	url     string
	handler controller.HandlerInterface
}

var allRouters = []UrlHandler{
	UrlHandler{url: "/test", handler: &controller.FirstHandler{}},
}
