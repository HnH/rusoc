package rusoc

import (
	"fmt"
	"strconv"
	"net/url"
	"strings"
	"sort"
)

// Контструктор приложения Мой Мир
func NewAppMM(appId uint64, secKey string) App {
	return &AppMM{
		server:	"http://www.appsmail.ru/platform/api?%s",
		appId:	appId,
		secKey:	secKey,
	}
}

// Структура приложения Мой Мир
type AppMM struct {
	server	string
	appId	uint64
	secKey	string
}

// ID приложения
func (self *AppMM) GetId() uint64 {
	return self.appId
}

// Социальная сеть текущего приложения
// Метод необходим для идентефикации при работе через интерфейс
func (self *AppMM) GetSocial() string {
	return MM
}

// Публичный ключ приложения
// В Одноклассниках различаются id и ключ приложения
// В Моём Мире — возвращаем id приложения в виде строки
func (self *AppMM) GetKey() string {
	return strconv.FormatUint(self.appId, 10)
}

// secret_key приложения
func (self *AppMM) GetSecretKey() string {
	return self.secKey
}

// Полный URL для вызова метода
func (self *AppMM) GetUrl(method string, params url.Values) string {
	return fmt.Sprintf(self.server, params.Encode())
}

// Вызов метода с результатом в виде массива байтов
func (self *AppMM) CallMethod(method string, params url.Values) ([]byte, int, error) {
	params.Set("method", method)
	// Всегда используем схему сервер-сервер для вызов
	// @see http://api.mail.ru/docs/guides/restapi/#server
	params.Set("secure", "1")
	params.Set(KEY_SIG, self.GenerateSignature(params, self.GetSecretKey()))
	return GetHTTP(self.GetUrl(method, params))
}

// Генерация подписи для API MM
// @see http://api.mail.ru/docs/guides/restapi/#sig
func (self *AppMM) GenerateSignature(request url.Values, secret string) (signature string) {
	var reqArr = make([]string, len(request))
	for k, _ := range request {
		if len(request[k]) > 1 {
			reqArr = append(reqArr, k + "=" + strings.Join(request[k], ","))
		} else {
			reqArr = append(reqArr, k + "=" + request.Get(k))
		}
	}
	sort.Strings(reqArr)

	signature = strings.Join(reqArr, "") + secret
	signature = strings.ToLower(GetMD5(signature))
	return
}

// Конструктор клиента текущего приложения
// @see http://api.mail.ru/docs/guides/social-apps/#params
func (self *AppMM) NewClient(req url.Values) (Client, error) {
	socialId, e	:= strconv.ParseUint(req.Get("vid"), 10, 64)
	sessKey		:= req.Get("session_key")

	if e != nil || socialId == 0 || len(sessKey) == 0 {
		return nil, errInvalidParams
	}

	// т.к. не присылается доп. параметра для проверки авторизации на сервере API
	// Смотрим сходится ли подпись авторизационного запроса и сохраняем результат для checkAuth()
	c := &ClientMM{
		app:		self,
		socialId:	socialId,
		sessKey:	sessKey,
	}

	sig := req.Get(KEY_SIG)
	req.Del(KEY_SIG)
	if sig == self.GenerateSignature(req, self.GetSecretKey()) {
		c.authSig = true
	}

	return c, nil
}
