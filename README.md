# rusoc [![GoDoc](https://godoc.org/github.com/HnH/rusoc?status.svg)](https://godoc.org/github.com/HnH/rusoc)

Библиотека на Go (Golang) для работы с API российских социальных сетей (Одноклассники, ВКонтакте, Мой Мир). Работа со всеми платформами максимально унифицирована и по сути сводится к реализации структур двух типов для каждой социальной сети: [приложения](http://godoc.org/github.com/HnH/rusoc#App) и [клиента](http://godoc.org/github.com/HnH/rusoc#Client). Специфичные для каждой социальной сети вещи вынесены за рамки данной библиотеки.

# Пример использования

Приложение для ВКонтакте, получающее токен и отправляющее уведомление клиенту.

```go
package main

import (
	"fmt"
	"log"
	"errors"
	"net/url"
	"encoding/json"
	"github.com/HnH/rusoc"
)

// Приложение
type appVk struct {
	*rusoc.AppVk
	token string
}

// Обновляем access_token
func (self *appVk) updateToken() (err error) {
	var req = url.Values{}
	req.Add("client_id", self.GetKey())
	req.Add("client_secret", self.GetSecretKey())
	req.Add("grant_type", "client_credentials")

	var body []byte
	if body, _, err = rusoc.GetHTTP("https://oauth.vk.com/access_token?" + req.Encode()); err != nil {
		return
	}

	var rsp map[string]*json.RawMessage
	if err = json.Unmarshal(body, &rsp); err != nil {
		return
	}

	// Разбираем ответ сервера
	var ok bool
	if _, ok = rsp["error"]; ok {
		return errors.New("Не удалось получить access_token")
	}

	if _, ok = rsp["access_token"]; !ok {
		return errors.New("Не удалось получить access_token")
	}

	if err = json.Unmarshal(*rsp["access_token"], &self.token); err != nil {
		return
	}

	return nil
}

// Клиент
type clientVk struct {
	*rusoc.ClientVk
}

// Отправляем уведомление
func (self *clientVk) sendNotification(txt, token string) {
	var req = url.Values{}
	req.Add("client_id", self.GetApp().GetKey())
	req.Add("client_secret", self.GetApp().GetSecretKey())
	req.Add("user_id", fmt.Sprintf("%d", self.GetSocialId()))
	req.Add("message", txt)
	req.Add("access_token", token)

	self.CallMethod("secure.sendNotification", req)
}

func main() {
	var (
		app = appVk{
			AppVk: rusoc.NewAppVk(100500, "hash").(*rusoc.AppVk),
		}
		err error
	)

	if err = app.updateToken(); err != nil {
		log.Fatalf("Err: %+v:", err)
	}

	// В нормальном режиме эти данные передаются приложению социальной сетью
	var cReq = url.Values{}
	cReq.Set("viewer_id", "100500")
	cReq.Set("sid", "hash")
	cReq.Set("secret", "hash")
	cReq.Set("auth_key", "hash")

	var c rusoc.Client
	if c, err = app.NewClient(cReq); err != nil {
		log.Fatalf("Err: %+v:", err)
	}

	var client = clientVk{c.(*rusoc.ClientVk)}
	client.sendNotification("Проверка", app.token)
}
```
