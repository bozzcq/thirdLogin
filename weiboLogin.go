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
	weiboUrl = "https://api.weibo.com/2/users/show.json"
)

type weiboInfo struct {
	ErrorCode int64  `json:"error_code,omitempty"`
	Error     string `json:"error,omitempty"`

	Uid         int64  `json:"id"`
	ScreenName  string `json:"screen_name"`
	Gender      string `json:"gender"`
	AvatarLarge string `json:"avatar_large"`
}

func weiboData(data ReqData) (string, error) {
	var resultInfo resultData
	info := new(weiboInfo)

	err := info.getWeiboInfo(data)
	if err != nil {
		return "", err
	}

	resultInfo.Nickname = info.ScreenName
	if info.Gender == "m" {
		resultInfo.Sex = 1
	} else if info.Gender == "f" {
		resultInfo.Sex = 2
	}

	resultInfo.Avatarurl = info.AvatarLarge

	b, _ := json.Marshal(resultInfo)
	return string(b), nil
}

func (c *weiboInfo) getWeiboInfo(data ReqData) error {
	url := fmt.Sprintf("%s?access_token=%s&uid=%s", weiboUrl, data.AccessToken, data.OpenId)
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

	if c.ErrorCode != 0 {
		err = errors.New("[unknown error][code:" + strconv.FormatInt(c.ErrorCode, 10) + ",msg:" + c.Error + "]")
		return err
	}

	return nil
}
