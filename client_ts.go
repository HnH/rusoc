package rusoc

import (
	"net/url"
)

// Структура клиента тестового приложения
type ClientTest struct {
	app			App
	socialId	uint64
}

func (self *ClientTest) GetApp() App {
	return self.app
}

func (self *ClientTest) GetSocialId() uint64 {
	return self.socialId
}

// Генерация подписи
func (self *ClientTest) GenerateSignature(request url.Values) (signature string) {
	return
}

// Проверка авторизации пользователя с текущим session_key на сервере OK
func (self *ClientTest) CheckAuth() bool {
	return true
}
