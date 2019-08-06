package main

import (
	"fmt"
	"os"
	"os/signal"
	"socialbot/internal/common/app"
	"socialbot/internal/web"
	"socialbot/internal/web/console"
	"socialbot/internal/web/orm"
	"socialbot/internal/web/setting"
	"socialbot/internal/web/wblogger"
	"syscall"
)

func init() {
	err := setting.LoadAppPath()
	if err != nil {
		fmt.Printf("[ERROR] %+v \n", err)
		os.Exit(1)
	}

	setting.LoadFlag()

	err = setting.LoadConf()
	if err != nil {
		fmt.Printf("[ERROR] %+v \n", err)
		os.Exit(1)
	}

	err = wblogger.LoadLogger()
	if err != nil {
		fmt.Printf("[ERROR] %+v \n", err)
		os.Exit(1)
	}

	err = orm.LoadModel()
	if err != nil {
		fmt.Printf("[ERROR] %+v \n", err)
		os.Exit(1)
	}

	console.LoadConsole()

}

func main() {
	web.Run()
	HandleSig()
	app.RunShutdownCallback()
}

func HandleSig() {
	sig := make(chan os.Signal)
	signal.Notify(sig, /*syscall.SIGHUP,*/ syscall.SIGINT, syscall.SIGTERM, /*syscall.SIGQUIT*/)
	<-sig
}
