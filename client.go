package rusoc

import "net/url"

// Интерфейс клиента приложения социальной сети
type Client interface {
	GetApp() (App)
	GetSocial() (string)
	GetSocialId() (uint64)

	CallMethod(string, url.Values) ([]byte, int, error)
	CheckAuth() (bool)
}
