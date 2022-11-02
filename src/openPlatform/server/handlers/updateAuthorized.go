package handlers

import (
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/kernel/contract"
	"net/http"
)

type UpdateAuthorized struct {
	contract.EventHandlerInterface

	App *kernel.ApplicationInterface
}

func NewUpdateAuthorized(app *kernel.ApplicationInterface) *UpdateAuthorized {
	handler := &UpdateAuthorized{
		App: app,
	}
	return handler
}

func (handler *UpdateAuthorized) Handle(request *http.Request, header contract.EventInterface, content interface{}) interface{} {

	return nil
}
