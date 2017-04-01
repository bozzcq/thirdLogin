package thirdLogin

import (
	"errors"
)

const (
	WechatLogin = iota + 1
	QQLogin
	WeiboLogin
)

type ReqData struct {
	LoginType   int8 //1:wechat  2:qq   3:weibo
	AccessToken string
	OpenId      string
	AppId       string
	Secret      string
}

func (r ReqData) GetResponseData() (string, error) {
	var s string
	var err error
	switch r.LoginType {
	case WechatLogin:
		s, err = wechatData(r)
	case QQLogin:
		s, err = qqData(r)
	case WeiboLogin:
		s, err = weiboData(r)
	default:
		err = errors.New("unknown logintype")
	}

	return s, err
}
