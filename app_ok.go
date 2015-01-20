package rusoc

import (
	"strconv"
	"net/url"
)

func NewOkApp(appKey, secKey string) App {
	return &AppOk{
		server:	"https://api.ok.ru/api/",
		appKey:	appKey,
		secKey:	secKey,
	}
}

type AppOk struct {
	server	string
	appKey	string
	secKey	string
}

func (self *AppOk) GetSocial() string {
	return OK
}

func (self *AppOk) GetKey() string {
	return self.appKey
}

func (self *AppOk) GetSecretKey() string {
	return self.secKey
}

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
