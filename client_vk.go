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

func (self *ClientVk) GetSocialId() uint64 {
	return self.socialId
}

// Генерация подписи для API ВКонтакте
// @see https://vk.com/page-1_2369497
func (self *ClientVk) GenerateSignature(request url.Values) (signature string) {
	reqArr := make([]string, len(request))

	i := 0
	for k, _ := range request {
		reqArr[i] = k + "=" + request.Get(k)
		i++
	}
	sort.Strings(reqArr)
	signature = strings.Join(reqArr, "")
	signature += self.GetApp().GetSecretKey()
	signature = getMD5(signature)

	return
}

// Проверка авторизации пользователя на сервере ВКонтакте
// @see: https://vk.com/dev.php?method=auth_key
func (self *ClientVk) CheckAuth() bool {
	return getMD5(fmt.Sprintf(`%s_%d_%s`, self.GetApp().GetKey(), self.GetSocialId(), self.GetApp().GetSecretKey())) == self.authKey
}
