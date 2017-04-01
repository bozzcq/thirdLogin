package thirdLogin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	weixinUrl = "https://api.weixin.qq.com/sns/userinfo"
)

type wxUserInfo struct {
	Errcode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	Nickname   string `json:"nickname"`
	Sex        int    `json:"sex"`
	Headimgurl string `json:"headimgurl"`
}

func wechatData(data ReqData) (string, error) {
	var result resultData
	userinfo := new(wxUserInfo)
	err := userinfo.getWechatUserInfo(data)
	if err != nil {
		return "", err
	}

	result.Nickname = userinfo.Nickname
	result.Sex = userinfo.Sex
	result.Avatarurl = userinfo.Headimgurl

	b, _ := json.Marshal(result)

	return string(b), nil
}

func (c *wxUserInfo) getWechatUserInfo(data ReqData) error {
	url := fmt.Sprintf("%s?access_token=%s&openid=%s", weixinUrl, data.AccessToken, data.OpenId)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, c)
	if err != nil {
		return errors.New("[json unmarshal failed][" + err.Error() + "]")
	}

	if c.Errcode != 0 {
		return errors.New("[unknown failed][code:" + strconv.Itoa(c.Errcode) + ",msg:" + c.ErrMsg + "]")
	}

	return nil
}
