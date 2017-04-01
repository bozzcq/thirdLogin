获取第三方登录的用户信息，目前只有昵称，性别和头像

type ReqData struct {
	LoginType   int8 //1:wechat  2:qq   3:weibo
	AccessToken string
	OpenId      string
	AppId       string
	Secret      string
}

使用者需要传入以上的信息
simple:

repdata := ReqData{
    LoginType:1,
    AccessToken: "xxxxx",
    OpenId: "xxx",
    AppId: "xxx",
    Secret: "xxx",
}

result,err := repdata.GetResponseData()

result是string类型，使用者要自己进行序列化
