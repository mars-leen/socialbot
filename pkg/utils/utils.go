package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const TimeLine = "2006-01-02 15:04:05"

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
	return  min + r.Intn(max-min)
}

func GetRandInt64(min, max int64, identity ...int64) int64 {
	seed := time.Now().UnixNano()
	if len(identity) != 0 {
		seed = seed + 123456*identity[0]
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return  min + r.Int63n(max-min)
}


// encrypt
func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
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
