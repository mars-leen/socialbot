package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	SystemError       = ERROR(1000, "system error")
	ParamError        = ERROR(1001, "param error")
	AuthError         = ERROR(1002, "auth error")
	DataIsNotExist    = ERROR(1003, "data is not exist")
	RepeatOperation   = ERROR(1004, "repeat operation")
	ReadFileFailed    = ERROR(1005, "read file failed!")
	FrequentOperation = ERROR(1006, "Frequent operation")
	InvalidOperation  = ERROR(1007, "Invalid operation")
	UploadFailed      = ERROR(1008, "upload failed")
	InvalidFileFormat = ERROR(1009, "invalid file type")

	AccountNotRegError     = ERROR(1010, "account not register")
	InValidAccountError    = ERROR(1011, "account auth failed")
	RegisteredAccountError = ERROR(1012, "account is registered")
	InvalidEmailError      = ERROR(1013, "invalid email format")

	DownLoadError = ERROR(1020, "download error")

	// media
	TagsIsEmptyError    = ERROR(1030, "tags is empty")
	ExceedImgLimitError = ERROR(1031, "img number limits are exceeded")
	UploadFileIsEmpty   = ERROR(1032, "empty upload")
	UploadFileNotFound  = ERROR(1033, "upload file not found")

	// user
	NicknameEmpty = ERROR(1100, "Nickname is needed")
)

type Result interface {
	Out(c *gin.Context)
	Errorf(f string, v ...interface{}) Result
}

type ErrorResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time int64  `json:"time"`
}

func (e ErrorResult) Out(c *gin.Context) {
	status := http.StatusOK
	if e.Code == 1002 {
		status = http.StatusUnauthorized
	}
	c.JSON(status, e)
}

func (e ErrorResult) Errorf(f string, v ...interface{}) Result {
	e.Msg = fmt.Sprintf(f, v...)
	return e
}

type SuccessResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Time int64       `json:"time"`
	Data interface{} `json:"data"`
}

func (e SuccessResult) Out(c *gin.Context) {
	c.JSON(http.StatusOK, e)
}

func (e SuccessResult) Errorf(f string, v ...interface{}) Result {
	e.Msg = fmt.Sprintf(f, v...)
	return e
}

func ERROR(code int, msg string) Result {
	return ErrorResult{
		Code: code,
		Msg:  msg,
		Time: time.Now().Unix(),
	}
}

func SUCCESS(data interface{}) Result {
	result := SuccessResult{
		Code: 0,
		Msg:  "success",
		Time: time.Now().Unix(),
		Data: data,
	}
	if data != nil {
		return result
	}
	result.Data = struct{}{}
	return result
}

//  Return success result , exist args return obj otherwise array
func SUCCESSARR(data interface{}) Result {
	result := SuccessResult{
		Code: 0,
		Msg:  "success",
		Time: time.Now().Unix(),
		Data: data,
	}
	if data != nil {
		return result
	}

	result.Data = []string{}
	return result
}
