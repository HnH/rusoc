package rusoc

import "net/url"

type Client interface {
	GetApp() (App)
	GetSocialId() (uint64)
	GenerateSignature(url.Values) (string)

	CheckAuth() (bool)
}
