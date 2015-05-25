package rusoc

import (
	"fmt"
	"net/url"
)

// Структура клиента приложения ВКонтакте
type ClientVk struct {
	app			App
	socialId	uint64
	sessId		string
	sessSecret	string
	authKey		string
}

func (self *ClientVk) GetApp() App {
	return self.app
}

func (self *ClientVk) GetSocial() string {
	return self.GetApp().GetSocial()
}

func (self *ClientVk) GetSocialId() uint64 {
	return self.socialId
}

// Вызов метода с результатом в виде массива байтов
func (self *ClientVk) CallMethod(method string, params url.Values) ([]byte, int, error) {
	return self.GetApp().CallMethod(method, params)
}

// Проверка авторизации пользователя на сервере ВКонтакте
// @see: https://vk.com/dev.php?method=auth_key
func (self *ClientVk) CheckAuth() bool {
	return GetMD5(fmt.Sprintf(`%s_%d_%s`, self.GetApp().GetKey(), self.GetSocialId(), self.GetApp().GetSecretKey())) == self.authKey
}
