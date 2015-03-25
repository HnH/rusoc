package rusoc

import (
	"sort"
	"strings"
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

// Генерация подписи для API MM
// @see http://api.mail.ru/docs/guides/restapi/#sig
func (self *ClientMM) GenerateSignature(request url.Values) (signature string) {
	reqArr := make([]string, len(request))

	for k, _ := range request {
		if len(request[k]) > 1 {
			reqArr = append(reqArr, k + "=" + strings.Join(request[k], ","))
		} else {
			reqArr = append(reqArr, k + "=" + request.Get(k))
		}
	}
	sort.Strings(reqArr)
	signature = strings.Join(reqArr, "")
	signature += self.GetApp().GetSecretKey()
	signature = GetMD5(signature)
	signature = strings.ToLower(signature)

	return
}

// Вызов метода с результатом в виде массива байтов
func (self *ClientMM) CallMethod(method string, params url.Values) ([]byte, error) {
	params.Set("method", method)
	params.Set(KEY_SIG, self.GenerateSignature(params))
	return GetHTTP(self.GetApp().GetUrl(method, params))
}

// Проверка авторизации пользователя на сервере ММ
func (self *ClientMM) CheckAuth() bool {
	return self.authSig
}
