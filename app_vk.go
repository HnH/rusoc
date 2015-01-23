package rusoc

import (
	"fmt"
	"strconv"
	"net/url"
)

func NewAppVk(apiId, apiSec string) App {
	return &AppVk{
		server:	"https://api.vk.com/method/%s?%s",
		apiId:	apiId,
		apiSec:	apiSec,
	}
}

// Структура приложения ВКонтакте
type AppVk struct {
	server	string
	apiId	string
	apiSec	string
}

// Социальная сеть текущего приложения
// Метод необходим для идентефикации при работе через интерфейс
func (self *AppVk) GetSocial() string {
	return VK
}

// api_id приложения
func (self *AppVk) GetKey() string {
	return self.apiId
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
func (self *AppVk) NewClient(req url.Values) (Client, error) {
	socialId, _	:= strconv.ParseUint(req.Get("social_id"), 10, 64)
	authKey		:= req.Get("auth_key")

	if socialId == 0 || len(authKey) == 0 {
		return nil, errInvalidParams
	}

	return &ClientVk{
		app:		self,
		socialId:	socialId,
		authKey:	authKey,
	}, nil
}
