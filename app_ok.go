package rusoc

import (
	"fmt"
	"strconv"
	"net/url"
)

// Контструктор приложения
func NewAppOk(appKey, secKey string) App {
	return &AppOk{
		server:	"https://api.ok.ru/api/%s?%s",
		appKey:	appKey,
		secKey:	secKey,
	}
}

// Структура приложения Одноклассников
type AppOk struct {
	server	string
	appKey	string
	secKey	string
}

// Социальная сеть текущего приложения
// Метод необходим для идентефикации при работе через интерфейс
func (self *AppOk) GetSocial() string {
	return OK
}

// Ключ приложения
func (self *AppOk) GetKey() string {
	return self.appKey
}

// Секретный ключ приложения
func (self *AppOk) GetSecretKey() string {
	return self.secKey
}

// Полный URL для вызова метода
func (self *AppOk) GetUrl(method, params string) string {
	return fmt.Sprintf(self.server, method, params)
}

// Вызов метода с результатом в виде массива байтов
func (self *AppOk) CallMethod(method, params string) ([]byte, error) {
	return getHTTP(self.GetUrl(method, params))
}

// Конструктор клиента текущего приложения
func (self *AppOk) NewClient(req url.Values) (Client, error) {
	socialId, _	:= strconv.ParseUint(req.Get("social_id"), 10, 64)
	sessKey		:= req.Get("sess_key")
	sessSecKey	:= req.Get("sess_sec_key")

	if socialId == 0 || len(sessKey) == 0 || len(sessSecKey) == 0 {
		return nil, errInvalidParams
	}

	return &ClientOk{
		app:		self,
		socialId:	socialId,
		sessKey:	sessKey,
		sessSecKey:	sessSecKey,
	}, nil
}
