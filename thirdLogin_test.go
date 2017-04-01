package thirdLogin

import (
	"fmt"
	"testing"
)

func TestThirdLogin(t *testing.T) {
	req := ReqData{
		LoginType:   1,
		AccessToken: "afasdfasdf",
		OpenId:      "fasdfasdf",
		AppId:       "wx43ec0531247be1",
		Secret:      "2ba18f70265439ba1fd5c321d31a779d",
	}

	str, err := req.GetResponseData()

	fmt.Println(str, err)
}
