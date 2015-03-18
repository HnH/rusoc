package rusoc

import (
	"fmt"
	"strconv"
	"net/url"
)

// Контструктор приложения Одноклассников
func NewAppOk(appId uint64, appKey, secKey string) App {
	return &AppOk{
		server:	"https://api.ok.ru/api/%s?%s",
		appId:	appId,
		appKey:	appKey,
		secKey:	secKey,
	}
}

// Структура приложения Одноклассников
type AppOk struct {
	server	string
	appId	uint64
	appKey	string
	secKey	string
}

// ID приложения
func (self *AppOk) GetId() uint64 {
	return self.appId
}

// Социальная сеть текущего приложения
// Метод необходим для идентефикации при работе через интерфейс
func (self *AppOk) GetSocial() string {
	return OK
}

// Публичный ключ приложения
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
	return GetHTTP(self.GetUrl(method, params))
}

// Конструктор клиента текущего приложения
// @see http://apiok.ru/wiki/pages/viewpage.action?pageId=46137373#APIДокументация(Русский)-Параметрыприложения
func (self *AppOk) NewClient(req url.Values) (Client, error) {
	socialId, e	:= strconv.ParseUint(req.Get("logged_user_id"), 10, 64)
	sessKey		:= req.Get("session_key")
	sessSecKey	:= req.Get("session_secret_key")
	authSig		:= req.Get("auth_sig")

	if e != nil || socialId == 0 || len(sessKey) == 0 || len(sessSecKey) == 0 || len(authSig) == 0 {
		return nil, errInvalidParams
	}

	return &ClientOk{
		app:		self,
		socialId:	socialId,
		sessKey:	sessKey,
		sessSecKey:	sessSecKey,
		authSig:	authSig,
	}, nil
}
