package rusoc

import (
	"net/url"
)

// Структура клиента приложения Моего Мира
type ClientMM struct {
	app			App
	socialId	uint64
	sessKey		string
	authSig		bool
}

func (self *ClientMM) GetApp() App {
	return self.app
}

func (self *ClientMM) GetSocial() string {
	return self.GetApp().GetSocial()
}

func (self *ClientMM) GetSocialId() uint64 {
	return self.socialId
}

// Вызов метода с результатом в виде массива байтов
func (self *ClientMM) CallMethod(method string, params url.Values) ([]byte, int, error) {
	return self.GetApp().CallMethod(method, params)
}

// Проверка авторизации пользователя на сервере ММ
func (self *ClientMM) CheckAuth() bool {
	return self.authSig
}
