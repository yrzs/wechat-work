package handlers

import (
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/kernel/contract"
	"net/http"
)

type Authorized struct {
	contract.EventHandlerInterface

	App *kernel.ApplicationInterface
}

func NewAuthorized(app *kernel.ApplicationInterface) *Authorized {
	handler := &Authorized{
		App: app,
	}
	return handler
}

func (handler *Authorized) Handle(request *http.Request, header contract.EventInterface, content interface{}) interface{} {

	return nil
}
