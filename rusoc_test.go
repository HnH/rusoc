package rusoc

import (
	"fmt"
	"testing"
	"net/url"
)

func TestVK(t *testing.T) {
	var (
		app = NewAppVk(100500, "secretKey")
		err error
	)

	if app.GetSocial() != VK || app.GetId() != 100500 || app.GetKey() != "100500" || app.GetSecretKey() != "secretKey" {
		t.Errorf("VK app failed")
	}

	var cReq = url.Values{}
	cReq.Set("viewer_id", "100500")
	cReq.Set("sid", "hash")
	cReq.Set("secret", "hash")
	cReq.Set("auth_key", GetMD5(fmt.Sprintf("%s_%d_%s", app.GetKey(), 100500, app.GetSecretKey())))

	if app.GenerateSignature(cReq, app.GetSecretKey()) != "991095660ee01c11e22338f882d6cea9" {
		t.Errorf("VK app failed")
	}

	var c Client
	if c, err = app.NewClient(cReq); err != nil {
		t.Errorf("VK client error: %+v:", err)
	}

	if c.GetSocial() != VK || c.GetSocialId() != 100500 || !c.CheckAuth() {
		t.Errorf("VK client failed")
	}
}

func TestOK(t *testing.T) {
	var (
		app = NewAppOk(100500, "key", "secretKey")
		err error
	)

	if app.GetSocial() != OK || app.GetId() != 100500 || app.GetKey() != "key" || app.GetSecretKey() != "secretKey" {
		t.Errorf("OK app failed")
	}

	var cReq = url.Values{}
	cReq.Set("logged_user_id", "100500")
	cReq.Set("session_key", "zzz")
	cReq.Set("session_secret_key", "zzzSecret")
	cReq.Set("auth_sig", GetMD5(fmt.Sprintf("%d%s%s", 100500, "zzz", app.GetSecretKey())))

	if app.GenerateSignature(cReq, app.GetSecretKey()) != "665f580d7673e88405740b7055141652" {
		t.Errorf("OK app failed")
	}

	var c Client
	if c, err = app.NewClient(cReq); err != nil {
		t.Errorf("OK client error: %+v:", err)
	}

	if c.GetSocial() != OK || c.GetSocialId() != 100500 || !c.CheckAuth() {
		t.Errorf("OK client failed")
	}
}

func TestMM(t *testing.T) {
	var (
		app = NewAppMM(100500, "secretKey")
		err error
	)

	if app.GetSocial() != MM || app.GetId() != 100500 || app.GetKey() != "100500" || app.GetSecretKey() != "secretKey" {
		t.Errorf("MM app failed")
	}

	var cReq = url.Values{}
	cReq.Set("vid", "100500")
	cReq.Set("session_key", "zzz")
	cReq.Set(KEY_SIG, app.GenerateSignature(cReq, app.GetSecretKey()))

	var c Client
	if c, err = app.NewClient(cReq); err != nil {
		t.Errorf("MM client error: %+v:", err)
	}

	if c.GetSocial() != MM || c.GetSocialId() != 100500 || !c.CheckAuth() {
		t.Errorf("MM client failed")
	}
}
