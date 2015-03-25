package rusoc

import "net/url"

// Интерфейс клиента приложения социальной сети
type Client interface {
	GetApp() (App)
	GetSocial() (string)
	GetSocialId() (uint64)
	GenerateSignature(url.Values) (string)

	CallMethod(string, url.Values) ([]byte, error)
	CheckAuth() (bool)
}
