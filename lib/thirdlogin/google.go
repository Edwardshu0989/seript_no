package thirdlogin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/yeeyuntech/yeego/yeeHttp"
)

// 根据AccessToken获取用户信息
// https://stackoverflow.com/questions/7130648/get-user-info-via-google-api
// https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=
// {
//  "id": "108580142136809228375",
//  "email": "agelinazf@gmail.com",
//  "verified_email": true,
//  "name": "张方",
//  "given_name": "方",
//  "family_name": "张",
//  "link": "https://plus.google.com/108580142136809228375",
//  "picture": "https://lh4.googleusercontent.com/-vJnC-KZyfGU/AAAAAAAAAAI/AAAAAAAAAAg/Z5qpCfk2NbQ/photo.jpg",
//  "gender": "male",
//  "locale": "zh-CN"
// }
// 根据IdToken获取token信息
// https://developers.google.com/identity/sign-in/web/backend-auth?authuser=2
// https://oauth2.googleapis.com/tokeninfo?id_token=
// {
//  "iss": "accounts.google.com",
//  "azp": "695774712749-nrt662s2cb6hqutqure9g31gg8m7p0rk.apps.googleusercontent.com",
//  "aud": "695774712749-nrt662s2cb6hqutqure9g31gg8m7p0rk.apps.googleusercontent.com",
//  "sub": "108580142136809228375",
//  "email": "agelinazf@gmail.com",
//  "email_verified": "true",
//  "at_hash": "gBz6jtpJc1XxYnuU_a8bLg",
//  "iat": "1548903702",
//  "exp": "1548907302",
//  "alg": "RS256",
//  "kid": "6fb05f742366ee4cf4bcf49f984c487e45c8c83d",
//  "typ": "JWT"
// }

var (
	userInfoApi  = "https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s"
	tokenInfoApi = "https://oauth2.googleapis.com/tokeninfo?id_token=%s"
)

type GoogleUserInfo struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Link          string `json:"link"`
	Picture       string `json:"picture"`
	Gender        string `json:"gender"`
	Locale        string `json:"locale"`
}

func GetGoogleUserInfo(accessToken string) (GoogleUserInfo, error) {
	u := fmt.Sprintf(userInfoApi, accessToken)
	data, err := yeeHttp.Get(u).Exec().ToBytes()
	if err != nil {
		return GoogleUserInfo{}, err
	}
	errCode, _ := jsonparser.GetInt(data, "error", "code")
	if errCode != 0 {
		msg, _ := jsonparser.GetString(data, "error", "message")
		return GoogleUserInfo{}, errors.New(msg)
	}
	var user GoogleUserInfo
	err = json.Unmarshal(data, &user)
	if err != nil {
		return GoogleUserInfo{}, err
	}
	return user, nil
}
