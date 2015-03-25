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

func (self *ClientTest) GetSocial() string {
	return self.GetApp().GetSocial()
}

func (self *ClientTest) GetSocialId() uint64 {
	return self.socialId
}

// Генерация подписи
func (self *ClientTest) GenerateSignature(request url.Values) (signature string) {
	return
}

// Вызов метода с результатом в виде массива байтов
func (self *ClientTest) CallMethod(method string, params url.Values) ([]byte, error) {
	params.Set(KEY_SIG, self.GenerateSignature(params))
	return GetHTTP(self.GetApp().GetUrl(method, params))
}

// Проверка авторизации пользователя с текущим session_key на сервере OK
func (self *ClientTest) CheckAuth() bool {
	return true
}
