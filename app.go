package rusoc

import (
	"net/url"
)

type App interface {
	GetSocial() (string)
	GetKey() (string)
	GetSecretKey() (string)
	GetUrl(string, string) (string)
	CallMethod(string, string) ([]byte, error)

	NewClient(url.Values) (Client, error)
}
