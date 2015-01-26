package rusoc

import "net/url"

// Интерфейс клиента приложения социальной сети
type Client interface {
	GetApp() (App)
	GetSocialId() (uint64)
	GenerateSignature(url.Values) (string)

	CheckAuth() (bool)
}
