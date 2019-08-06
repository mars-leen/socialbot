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

func TestListServer(t *testing.T) {
	c := GetClient("")
	list, err :=  c.GetServers()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v \n", list)
}

func TestCreateServer(t *testing.T) {

}