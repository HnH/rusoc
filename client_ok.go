package rusoc

import (
	"sort"
	"strings"
	"net/url"
	"fmt"
)

type ClientOk struct {
	app			App
	socialId	uint64
	sessKey		string
	sessSecKey	string
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
		signature += self.app.GetSecretKey()
	}

	signature = getMD5(signature)
	signature = strings.ToLower(signature)

	return
}

// Проверка авторизации пользователя с текущим session_key на сервере OK
func (self *ClientOk) CheckAuth() bool {
	request := url.Values{}
	request.Add("application_key", self.app.GetKey())
	request.Add("session_key", self.sessKey)
	request.Add("sig", self.GenerateSignature(request))

	// URL для запроса
	// @see http://apiok.ru/wiki/display/api/users.getLoggedInUser+ru
	var (
		url		= "https://api.ok.ru/api/users/getLoggedInUser?" + request.Encode()
		body	[]byte
		err		error
	)

	if body, err = getHTTP(url); err != nil {
		return false
	}

	return string(body) == fmt.Sprintf(`"%d"`, self.socialId)
}
