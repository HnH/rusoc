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
	CallMethod(string, url.Values) ([]byte, int, error)
	GenerateSignature(url.Values, string) (string)

	NewClient(url.Values) (Client, error)
}
