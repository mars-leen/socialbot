package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"math/rand"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const TimeLine = "2006-01-02 15:04:05"

var (
	ImgExtList   = []string{".BMP", ".JPG", ".JPEG", ".PNG", ".GIF"}
	VideoExtList = []string{".MP4", ".MTV", ".DAT", ".WMV", ".AVI", ".3GP", ".RMVB", ".AMV", ".DMV", ".FLV"}
)

// os
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func AppAbsPath() (appPath string, err error) {
	if IsWindows() && filepath.IsAbs(os.Args[0]) {
		appPath = filepath.Clean(os.Args[0])
	} else {
		appPath, err = exec.LookPath(os.Args[0])
	}
	if err != nil {
		return appPath, errors.Wrap(err, fmt.Sprintf("exec.LookPath error"))
	}

	appPath, err = filepath.Abs(filepath.Dir(appPath))
	if err != nil {
		return appPath, errors.Wrap(err, fmt.Sprintf("filepath.Abs error"))
	}
	return strings.Replace(appPath, "\\", "/", -1), nil
}

// time
func RandSleep(min, max int) {
	rn := GetRandInt(min, max)
	time.Sleep(time.Duration(rn) * time.Millisecond)
}

// rand
func GetRandInt(min, max int, identity ...int64) int {
	seed := time.Now().UnixNano()
	if len(identity) != 0 {
		seed = seed + 123456*identity[0]
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Intn(max-min)
}

func GetRandInt64(min, max int64, identity ...int64) int64 {
	seed := time.Now().UnixNano()
	if len(identity) != 0 {
		seed = seed + 123456*identity[0]
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Int63n(max-min)
}

// encrypt
func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func Sha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func SaltSha1(str, salt string) string {
	h := hmac.New(sha1.New, []byte(salt))
	h.Write([] byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func MixEncode(s string) string {
	return Md5(Sha1(s))
}

// string
func Trim(str, prefix, suffix string) string {
	str = TrimSpace(str)
	str = strings.TrimLeft(str, prefix)
	return strings.TrimRight(str, suffix)
}

func TrimSpace(str string) string {
	str = strings.TrimLeft(str, " ")
	str = strings.TrimRight(str, " ")
	return str
}

func Str2Int(str []string) []int {
	list := make([]int, len(str))
	for index,value := range str{
		list[index], _ = strconv.Atoi(value)
	}
	return list
}

func Str2Int64(str []string) []int64 {
	list := make([]int64, len(str))
	for index,value := range str{
		list[index], _ = strconv.ParseInt(value, 10, 64)
	}
	return list
}

func StringInArray(dst string, arr []string) bool {
	for _,v := range arr{
		if dst ==v {
			return true
		}
	}
	return false
}


func DeleteStringFromSlice(dst string, arr []string) []string{
	if len(arr) ==1 {
		if arr[0] == dst {
			return []string{}
		}else{
			return arr
		}
	}
	j := 0
	for _, val := range arr {
		if val == dst {
			arr[j] = val
			j++
		}
	}
	return arr[:j]
}


// file
func GetFileExt(file *multipart.FileHeader) (ext string, err error) {
	ext = filepath.Ext(file.Filename)
	if ext != "" {
		return ext, nil
	}
	conType := file.Header.Get("Content-Type")
	fileTypes, err := mime.ExtensionsByType(conType)
	if err != nil {
		return ext, errors.Wrap(err, fmt.Sprintf("mime.ExtensionsByType from conType(%s) failed", conType))
	}
	if len(fileTypes) > 0 {
		ext = fileTypes[0]
	}
	return ext, nil
}

func GetFileExtFromResp(resp *http.Response) (string, error) {
	conType := resp.Header.Get("Content-Type")
	fileTypes, err := mime.ExtensionsByType(conType)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("mime.ExtensionsByType from conType(%s) failed", conType))
	}
	if len(fileTypes) > 0 {
		return fileTypes[0], nil
	}
	return "", nil
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	dir := filepath.Dir(dst)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "MkdirAll failed")
	}

	src, err := file.Open()
	if err != nil {
		return errors.Wrap(err, "file.Open dist file failed")
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return errors.Wrap(err, "io.Copy failed")
	}
	return nil
}

func GetFileTypeByExt(ext string) string {
	ext = strings.ToUpper(ext)
	for _,m := range ImgExtList {
		if m == ext{
			return "IMG"
		}
	}
	for _,m := range VideoExtList {
		if m == ext{
			return "VIDEO"
		}
	}
	return ""
}

//email verify
func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
