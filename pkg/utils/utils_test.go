package utils

import (
	"fmt"
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