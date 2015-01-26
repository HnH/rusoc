package rusoc

import (
	"fmt"
	"sort"
	"strings"
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

func (self *ClientOk) GetSocialId() uint64 {
	return self.socialId
}

// Генерация подписи для API OK
// @see http://apiok.ru/wiki/pages/viewpage.action?pageId=42476522
func (self *ClientOk) GenerateSignature(request url.Values) (signature string) {
	reqArr := make([]string, len(request))

	for k, _ := range request {
		reqArr = append(reqArr, k + "=" + request.Get(k))
	}
	sort.Strings(reqArr)
	signature = strings.Join(reqArr, "")

	// APP secret key для запросов без session_key
	// SESS secret key для запросов с session_key
	if _, chk := request["session_key"]; chk {
		signature += self.sessSecKey
	} else {
		signature += self.GetApp().GetSecretKey()
	}

	signature = getMD5(signature)
	signature = strings.ToLower(signature)

	return
}

// Проверка авторизации пользователя с текущим session_key на сервере OK
func (self *ClientOk) CheckAuth() bool {
	return getMD5(fmt.Sprintf(`%d_%s_%s`, self.GetSocialId(), self.sessKey, self.GetApp().GetSecretKey())) == self.authSig
}
