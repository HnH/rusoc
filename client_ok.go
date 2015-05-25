package rusoc

import (
	"fmt"
	"net/url"
)

// Структура клиента приложения для Одноклассников
type ClientOk struct {
	app			App
	socialId	uint64
	sessKey		string
	sessSecKey	string
	authSig		string
}

func (self *ClientOk) GetApp() App {
	return self.app
}

func (self *ClientOk) GetSocial() string {
	return self.GetApp().GetSocial()
}

func (self *ClientOk) GetSocialId() uint64 {
	return self.socialId
}

// Вызов метода с результатом в виде массива байтов
func (self *ClientOk) CallMethod(method string, params url.Values) ([]byte, int, error) {
	params.Set(KEY_SIG, self.GetApp().GenerateSignature(params, self.sessSecKey))
	return GetHTTP(self.GetApp().GetUrl(method, params))
}

// Проверка авторизации пользователя с текущим session_key на сервере OK
func (self *ClientOk) CheckAuth() bool {
	return GetMD5(fmt.Sprintf(`%d%s%s`, self.GetSocialId(), self.sessKey, self.GetApp().GetSecretKey())) == self.authSig
}
