package pincore

import (
	"fmt"
	"github.com/pkg/errors"
	"socialbot/pkg/pinbot/pinmodels"
)

type AuthError struct {
	error
}

type AccountError struct {
	error
}

type IpBannedError struct {
	error
}


type CodeError struct {
	error
}


func CheckRespErrors(responseError *pinmodels.ResponseError) error {
	if responseError == nil {
		return nil
	}
	var err error
	switch responseError.APIErrorCode {
		case ApiErrorAccountInvalidPassword:
			return errors.Wrap(AccountError{errors.New("AccountError")}, "password error, check your login  password")
		case ApiErrorAccountNotFound :
			return errors.Wrap(AccountError{errors.New("AccountError")}, "account is no exist, check your login  account")
		case ApiErrorIpBanned :
			return errors.Wrap(IpBannedError{errors.New("IpBannedError")}, "your ip has been banned")
		case ApiAuthError :
			return errors.Wrap(AuthError{errors.New("ApiAuthError")}, "auth error, need reLogin")
		default:
			err =  errors.Wrap(CodeError{errors.New("CodeError")}, fmt.Sprintf("code error, detail %#v", responseError))
	}
	return err
}

func CheckHttpsError(code int) error {
	switch code {
		case 200:
		default:
			return errors.Errorf("http code(%d) error", code)
	}
	return nil
}

func CheckIsAuthError(err error)  bool {
	if err == nil {
		return false
	}
	switch errors.Cause(err).(type) {
		case AuthError:
			return true
		default:
			return false
	}
}
