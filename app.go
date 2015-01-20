package rusoc

import (
	"net/url"
)

type App interface {
	GetSocial() (string)
	GetKey() (string)
	GetSecretKey() (string)

	NewClient(url.Values) (Client, error)
}
