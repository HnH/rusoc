package rusoc

import (
	"fmt"
	"strconv"
	"net/url"
)

// Конструктор приложения ВКонтакте
func NewAppVk(apiId uint64, apiSec string) App {
	return &AppVk{
		server:	"https://api.vk.com/method/%s?%s",
		apiId:	apiId,
		apiSec:	apiSec,
	}
}

// Структура приложения ВКонтакте
type AppVk struct {
	server	string
	apiId	uint64
	apiSec	string
}

// ID приложения
func (self *AppVk) GetId() uint64 {
	return self.apiId
}

// Социальная сеть текущего приложения
// Метод необходим для идентефикации при работе через интерфейс
func (self *AppVk) GetSocial() string {
	return VK
}

// Публичный ключ приложения
// В Одноклассниках различаются id и ключ приложения
// ВКонтакте — возвращаем id приложения в виде строки
func (self *AppVk) GetKey() string {
	return strconv.FormatUint(self.apiId, 10)
}

// api_secret приложения
func (self *AppVk) GetSecretKey() string {
	return self.apiSec
}

// Полный URL для вызова метода
func (self *AppVk) GetUrl(method, params string) string {
	return fmt.Sprintf(self.server, method, params)
}

// Вызов метода с результатом в виде массива байтов
func (self *AppVk) CallMethod(method, params string) ([]byte, error) {
	return getHTTP(self.GetUrl(method, params))
}

// Конструктор клиента текущего приложения
// @see http://vk.com/dev/apps_init
func (self *AppVk) NewClient(req url.Values) (Client, error) {
	socialId, e	:= strconv.ParseUint(req.Get("viewer_id"), 10, 64)
	sessId		:= req.Get("sid")
	sessSecret	:= req.Get("secret")
	authKey		:= req.Get("auth_key")

	if e != nil || socialId == 0 || len(sessId) == 0 || len(sessSecret) == 0 || len(authKey) == 0 {
		return nil, errInvalidParams
	}

	return &ClientVk{
		app:		self,
		socialId:	socialId,
		sessId:		sessId,
		sessSecret:	sessSecret,
		authKey:	authKey,
	}, nil
}
