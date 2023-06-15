package client

import (
	"github.com/cherish-chat/xxim-server/common/pb"
	"github.com/cherish-chat/xxim-server/common/utils"
	"testing"
	"time"
)

// UserRegister 注册用户
func TestHttpClient_UserRegister(t *testing.T) {
	userId := "3"
	phone := "13600000003"
	//client := getHttpClient(t, nil)
	client := getWsClient(t, nil)
	time.Sleep(1 * time.Second)
	userRegisterResp, err := client.UserRegister(&pb.UserRegisterReq{
		Header:       &pb.RequestHeader{},
		UserId:       userId,
		RegisterTime: nil,
		Nickname:     utils.AnyPtr("用户" + userId),
		Avatar:       utils.AnyPtr("https://www.baidu.com"),
		AccountMap: map[string]string{
			pb.AccountTypeUsername:     "user" + userId,
			pb.AccountTypePassword:     utils.Md5("123456"),
			pb.AccountTypePasswordSalt: utils.Snowflake.String(),
			pb.AccountTypePhone:        phone,
			pb.AccountTypePhoneCode:    "86",
			pb.AccountTypeEmail:        phone + "@xxim.com",
		},
		ProfileMap: map[string]string{
			"birthday": "2020-01-01",
			"sex":      "1",
		},
		ExtraMap: map[string]string{
			"extra1": "extra1",
		},
		VerifyMap: map[string]string{
			pb.VerifyTypeSmsCode:     "123456",
			pb.VerifyTypeEmailCode:   "123456",
			pb.VerifyTypeCaptchaId:   "123456",
			pb.VerifyTypeCaptchaCode: "123456",
		},
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("%s", utils.Json.MarshalToString(userRegisterResp))
}

// UserAccessToken 获取用户访问令牌
func TestHttpClient_UserAccessToken(t *testing.T) {
	//client := getHttpClient(t, nil)
	client := getWsClient(t, nil)
	time.Sleep(1 * time.Second)
	userAccessTokenResp, err := client.UserAccessToken(&pb.UserAccessTokenReq{
		Header: &pb.RequestHeader{
			Platform:  1,
			InstallId: "3",
			UserToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ4NDI1Mjk5MjQsImp0aSI6IjMifQ.z_3mr5G3U_F3-XZF45lYrkCBE_eq7Qd5kuPsVCVFU1k",
		},
		AccountMap: map[string]string{
			pb.AccountTypeUsername: "user3",
			pb.AccountTypePassword: utils.Md5("123456"),
		},
		VerifyMap:  nil,
		ExpireTime: nil,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("%s", utils.Json.MarshalToString(userAccessTokenResp))
}
