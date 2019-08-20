package userLogic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/uploadLogic"
	"socialbot/internal/web/model"
	"socialbot/internal/web/service/configService"
	"socialbot/internal/web/service/storageService"
	"socialbot/internal/web/service/userService"
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
	}
	return common.SUCCESS(result)
}

func EditProfile(c *gin.Context,nickname, intro string) common.Result {
	user, err := userService.MustGetTokenUser(c)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	editCols := make([]string, 0, 3)
	editCols = append(editCols, "nickname")
	if intro!= "" && intro != user.Intro {
		user.Intro = intro
		editCols = append(editCols, "intro")
	}
	if nickname != user.Nickname{
		user.Nickname = nickname
		editCols = append(editCols, "nickname")
	}


	filename, err := uploadLogic.UploadAvatar(c)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}
	if filename != "" {
		editCols = append(editCols, "avatar")
		if  user.Avatar != ""{
			storageService.SyncRemoveSigFile("local", filepath.Join(configService.GetStorageUploadPath(), user.Avatar))
		}
		user.Avatar = filename
	}

	if len(editCols) == 0{
		return common.SUCCESS(nil)
	}

	_, err = user.UpdateColsById(user.Id, editCols...)
	if err != nil {
		wblogger.Log.Error(err)
		return common.SystemError
	}

	result := map[string]interface{}{
		"userUri":  user.Uri,
		"nickname": user.Nickname,
		"intro":    user.Intro,
		"avatar":   configService.GetUploadFullUrl(user.Avatar),
	}
	return common.SUCCESS(result)
}


func getDefaultNickName(email string) string {
	ns := strings.Split(email, "@")
	w,_:= configService.GetWebsite()
	return w.HostName + "-" + ns[0]
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