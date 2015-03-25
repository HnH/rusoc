package rusoc

import (
	"fmt"
	"strconv"
	"net/url"
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
	if sig == c.GenerateSignature(req) {
		c.authSig = true
	}

	return c, nil
}
