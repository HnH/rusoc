package rusoc

import (
	"fmt"
	"strconv"
	"net/url"
)

// Контструктор тестового приложения
func NewAppTest(appId uint64, appKey, secKey string) App {
	return &AppTest{
		server:	"https://test.api/%s?%s",
		appId:	appId,
		appKey:	appKey,
		secKey:	secKey,
	}
}

// Структура тестового приложения
type AppTest struct {
	server	string
	appId	uint64
	appKey	string
	secKey	string
}

// ID приложения
func (self *AppTest) GetId() uint64 {
	return self.appId
}

// Социальная сеть текущего приложения
// Метод необходим для идентефикации при работе через интерфейс
func (self *AppTest) GetSocial() string {
	return TT
}

// Публичный ключ приложения
func (self *AppTest) GetKey() string {
	return self.appKey
}

// Секретный ключ приложения
func (self *AppTest) GetSecretKey() string {
	return self.secKey
}

// Полный URL для вызова метода
func (self *AppTest) GetUrl(method string, params url.Values) string {
	return fmt.Sprintf(self.server, method, params.Encode())
}

// Конструктор клиента текущего приложения
func (self *AppTest) NewClient(req url.Values) (Client, error) {
	socialId, e	:= strconv.ParseUint(req.Get("social_id"), 10, 64)

	if e != nil || socialId == 0 {
		return nil, errInvalidParams
	}

	return &ClientTest{
		app:		self,
		socialId:	socialId,
	}, nil
}
