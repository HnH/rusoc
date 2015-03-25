package rusoc

import (
	"net/url"
)

// Интерфейс приложения в социальной сети
type App interface {
	GetId() (uint64)
	GetSocial() (string)
	GetKey() (string)
	GetSecretKey() (string)
	GetUrl(string, url.Values) (string)

	NewClient(url.Values) (Client, error)
}
