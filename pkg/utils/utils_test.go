package utils

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestAppAbsPath(t *testing.T) {
	ap,_ := AppAbsPath()
	f, _ := os.Open(filepath.Join(ap, "views/socailbot-brain/dist"))
	names, _ := f.Readdirnames(-1)
	defer f.Close()
	fmt.Println(names)
}

func TestDeleteStringFromSlice(t *testing.T) {
	//res := DeleteStringFromSlice("b", []string{"a","b", "c"})
	res := DeleteSlice2(3, []int{1,2,3,4,5})
	fmt.Println(res)
}

func DeleteSlice2(dst int, a []int) []int {
	j := 0
	for _, val := range a {
		if val == dst {
			a[j] = val
			j++
		}
	}
	return a[:j]
}

func TestEnDe(t *testing.T) {
	hostMapHeader := map[string]map[string]string{
		"qq":{
			"Referer":        "https://www.qq.com",
			"Sec-Fetch-Mode": "no-cors",
			"User-Agent":     "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36",
		},
	}
	rs, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(hostMapHeader)
	if err!= nil {
		log.Fatal(err)
	}

	fmt.Println(string(rs))
	temp := map[string]map[string]string{}

	err =  jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(rs, &temp)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(temp["qq"])
}