package vultr

import (
	"fmt"
	"testing"
)

func TestCreateScript(t *testing.T) {
	CreateScript("", "test", "aaa", "boot")
}

func TestUpdateOrCreateScriptByName(t *testing.T) {
	err := UpdateOrCreateScriptByName("", "aaa", "fdafdafasdf", "boot")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("success")
}