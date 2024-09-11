package utils

import (
	"net/http"

	"github.com/mf15726/go-ball/enum"
)

type RequestHandler struct {
}

func (requestHandler RequestHandler) HandleRequest(httpMethod enum.HttpMethod, w http.ResponseWriter, r *http.Request, functionMap map[enum.HttpMethod]func(w http.ResponseWriter, r *http.Request)) {
	if f, ok := functionMap[httpMethod]; ok {
		f(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
