package rusoc

import (
	"fmt"
	"strconv"
	"net/url"
	"strings"
	"sort"
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
func (self *AppVk) GetUrl(method string, params url.Values) string {
	return fmt.Sprintf(self.server, method, params.Encode())
}

// Вызов метода с результатом в виде массива байтов
func (self *AppVk) CallMethod(method string, params url.Values) ([]byte, int, error) {
	params.Set(KEY_SIG, self.GenerateSignature(params, self.GetSecretKey()))
	return GetHTTP(self.GetUrl(method, params))
}

// Генерация подписи для API ВКонтакте
// @see https://vk.com/page-1_2369497
func (self *AppVk) GenerateSignature(request url.Values, secret string) (signature string) {
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
	signature = GetMD5(signature)
	return
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
