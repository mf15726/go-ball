package interfaces

import (
	"net/http"

	"github.com/mf15726/go-ball/enum"
)

type IRequestHandler interface {
	HandleRequest(enum.HttpMethod, http.ResponseWriter, *http.Request, map[enum.HttpMethod]func(http.ResponseWriter, *http.Request))
}
