package userLogic

import (
	"fmt"
	"socialbot/internal/web/common"
	"socialbot/internal/web/model"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/wblogger"
	"socialbot/pkg/jwtauth"
	"socialbot/pkg/utils"
	"strings"
)

func Register(accountForm model.AccountForm, ip string) common.Result {
	userModel := &model.User{}
	if accountForm.Type != common.AccountTypeEmail {
		return common.ParamError
	}
	// verify email
	if isEmail := utils.VerifyEmailFormat(accountForm.Account); !isEmail {
		return common.InvalidEmailError
	}
	has, err := userModel.GetOneByEmail(accountForm.Account)
	userModel.Email = accountForm.Account
	userModel.RegisterIp = ip
	if err != nil {
		fmt.Printf("%+v \n", err)
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if has {
		return common.RegisteredAccountError
	}

	//default value
	userModel.Nickname = getDefaultNickName(userModel.Email)
	userModel.Password = utils.MixEncode(accountForm.Password)

	_, err = userModel.Insert()
	if err != nil {
		fmt.Printf("%+v \n", err)
		wblogger.Log.Error(err)
		return common.SystemError
	}

	// gen jwt token
	token, err := jwtauth.GenerateToken(userModel.Id)
	if err != nil {
		fmt.Printf("%+v \n", err)
		wblogger.Log.Error(err)
		return common.SystemError
	}

	result := map[string]interface{}{
		"token":    token,
		"userUri":  userModel.Uri,
		"nickname": userModel.Nickname,
		"intro":    "",
		"avatar":   "",
		"identity": userModel.Identity,
	}
	return common.SUCCESS(result)
}

func Login(accountForm model.AccountForm) common.Result {
	var has bool
	var err error

	// get account info
	userModel := &model.User{}
	if accountForm.Type != common.AccountTypeEmail {
		return common.ParamError
	}
	has, err = userModel.GetOneByEmail(accountForm.Account)
	userModel.Email = accountForm.Account
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !has {
		return common.AccountNotRegError
	}

	// check password
	sctPassword := utils.MixEncode(accountForm.Password)
	if sctPassword != userModel.Password {
		return common.InValidAccountError
	}

	// gen jwt token
	token, err := jwtauth.GenerateToken(userModel.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	result := map[string]interface{}{
		"token":    token,
		"userUri":  userModel.Uri,
		"nickname": userModel.Nickname,
		"intro":    userModel.Intro,
		"avatar":   configService.GetUploadFullUrl(userModel.Avatar),
		"identity": userModel.Identity,
	}
	return common.SUCCESS(result)
}

func getDefaultNickName(email string) string {
	ns := strings.Split(email, "@")
	return configService.GetHostName() + "-" + ns[0]
}


func AdminLogin(accountForm model.AccountForm) common.Result {
	var has bool
	var err error

	// get account info
	userModel := &model.User{}
	if accountForm.Type != common.AccountTypeEmail {
		return common.ParamError
	}
	has, err = userModel.GetOneByEmail(accountForm.Account)
	userModel.Email = accountForm.Account
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if !has {
		return common.AccountNotRegError
	}

	if userModel.Identity != common.SuperAdmin {
		return common.AuthError
	}

	// check password
	sctPassword := utils.MixEncode(accountForm.Password)
	if sctPassword != userModel.Password {
		return common.InValidAccountError
	}

	// gen jwt token
	token, err := jwtauth.GenerateToken(userModel.Id)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	result := map[string]interface{}{
		"token":    token,
		"userUri":  userModel.Uri,
		"nickname": userModel.Nickname,
		"intro":    userModel.Intro,
		"avatar":   configService.GetUploadFullUrl(userModel.Avatar),
		"identity": userModel.Identity,
	}
	return common.SUCCESS(result)
}