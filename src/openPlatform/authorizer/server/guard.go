package server

import (
	"github.com/yrzs/wechat-work/kernel"
)

type Guard struct {
	*kernel.ServerGuard
}

func NewGuard(app *kernel.ApplicationInterface) *Guard {

	guard := &Guard{
		kernel.NewServerGuard(app),
	}

	guard.OverrideGetToken()

	return guard

}

// Override Validate
func (guard *Guard) OverrideGetToken() {

	guard.GetToken = func() string {

		encryptor := (*guard.App).GetComponent("Encryptor").(*kernel.Encryptor)

		return encryptor.GetToken()
	}
}
