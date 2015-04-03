// API популярных русских социальных сетей
package rusoc

import (
	"time"
	"errors"
	"net/http"
	"io/ioutil"
	"crypto/md5"
	"encoding/hex"
)

const (
	OK	= "OK"
	VK	= "VK"
	MM	= "MM"
	TT	= "TT"

	KEY_SIG = "sig"
)

var (
	errInvalidParams = errors.New("Invalid parameters")
)

// Считаем md5 хэш
func GetMD5(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

// HTTP запрос
func GetHTTP(u string) (body []byte, statusCode int, err error) {
	var (
		timeout	= time.Duration(5 * time.Second)
		client	= http.Client{Timeout: timeout}
		response *http.Response
	)

	if response, err = client.Get(u); err != nil {
		return
	}

	statusCode = response.StatusCode

	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)

	return
}
