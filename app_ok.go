package rusoc

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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
func (self *AppOk) GetUrl(method string, params url.Values) string {
	return fmt.Sprintf(self.server, method, params.Encode())
}

// Вызов метода с результатом в виде массива байтов
func (self *AppOk) CallMethod(method string, params url.Values) ([]byte, int, error) {
	params.Set(KEY_SIG, self.GenerateSignature(params, self.GetSecretKey()))
	return GetHTTP(self.GetUrl(method, params))
}

// Генерация подписи для API OK
// @see http://apiok.ru/wiki/pages/viewpage.action?pageId=42476522
func (self *AppOk) GenerateSignature(request url.Values, secret string) (signature string) {
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
