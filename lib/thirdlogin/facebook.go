package thirdlogin

import (
	"fmt"
	"github.com/huandu/facebook"
	"github.com/spf13/cast"
)

type FbUserInfo struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Picture   string `json:"picture"`
}

func GetFbUserInfo(uid, accessToken string) (FbUserInfo, error) {
	result, err := facebook.Get(fmt.Sprintf("/%s", uid), facebook.Params{
		"fields":       "id,first_name,last_name,name,picture,email",
		"access_token": accessToken,
	})
	if err != nil {
		return FbUserInfo{}, err
	}
	var user FbUserInfo
	user.Id = cast.ToString(result.Get("id"))
	user.FirstName = cast.ToString(result.Get("first_name"))
	user.LastName = cast.ToString(result.Get("last_name"))
	user.Name = cast.ToString(result.Get("name"))
	user.Email = cast.ToString(result.Get("email"))
	user.Picture = cast.ToString(result.Get("picture.data.url"))
	return user, nil
}
