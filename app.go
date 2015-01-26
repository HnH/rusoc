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
	GetUrl(string, string) (string)
	CallMethod(string, string) ([]byte, error)

	NewClient(url.Values) (Client, error)
}
