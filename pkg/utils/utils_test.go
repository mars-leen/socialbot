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
