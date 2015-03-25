package rusoc

import (
	"fmt"
	"sort"
	"strings"
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

// Генерация подписи для API ВКонтакте
// @see https://vk.com/page-1_2369497
func (self *ClientVk) GenerateSignature(request url.Values) (signature string) {
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

	return
}

// Вызов метода с результатом в виде массива байтов
func (self *ClientVk) CallMethod(method string, params url.Values) ([]byte, error) {
	params.Set(KEY_SIG, self.GenerateSignature(params))
	return GetHTTP(self.GetApp().GetUrl(method, params))
}

// Проверка авторизации пользователя на сервере ВКонтакте
// @see: https://vk.com/dev.php?method=auth_key
func (self *ClientVk) CheckAuth() bool {
	return GetMD5(fmt.Sprintf(`%s_%d_%s`, self.GetApp().GetKey(), self.GetSocialId(), self.GetApp().GetSecretKey())) == self.authKey
}
