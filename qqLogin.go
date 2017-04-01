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
	qqurl = "https://graph.qq.com/user/get_user_info"
)

type qqInfo struct {
	Ret       int    `json:"ret"`
	Msg       string `json:"msg"`
	Nickname  string `json:"nickname"`
	Gender    string `json:"gender"`
	AvatarUrl string `json:"figureurl_qq_1"`
}

func qqData(data ReqData) (string, error) {
	var resultInfo resultData
	info := new(qqInfo)
	err := info.getQQInfo(data)
	if err != nil {
		return "", err
	}

	resultInfo.Nickname = info.Nickname

	if info.Gender == "男" {
		resultInfo.Sex = 1
	} else if info.Gender == "女" {
		resultInfo.Sex = 2
	}
	resultInfo.Avatarurl = info.AvatarUrl

	b, _ := json.Marshal(resultInfo)

	return string(b), nil
}

func (c *qqInfo) getQQInfo(data ReqData) error {
	url := fmt.Sprintf("%s?access_token=%s&oauth_consumer_key=%s&openid=%s",
		qqurl, data.AccessToken, data.AppId, data.OpenId)

	resp, err := http.Get(url)
	if err != nil {
		return errors.New("[unknown error][" + err.Error() + "]")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("[read body failed][" + err.Error() + "]")
	}

	err = json.Unmarshal(body, c)
	if err != nil {
		return errors.New("[json unmarshal failed][" + err.Error() + "]")
	}

	if c.Ret != 0 {
		err = errors.New("[unknown error][code:" + strconv.Itoa(c.Ret) + ",msg:" + c.Msg + "]")
		return err
	}

	return nil
}
