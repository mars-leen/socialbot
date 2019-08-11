package console

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"socialbot/internal/common/app"
	"socialbot/internal/web/service/serverService"
)

func LoadConsole()  {
	gocron.Every(60).Seconds().Do(func() {
		_ = serverService.UpdateListRobotServerInfo()
	})
	app.RegShutdownCallback(func() {
		fmt.Println("[INFO] clear cron start")
		gocron.Clear()
		fmt.Println("[INFO]  clear cron end")
	})
	go func() {
		<-gocron.Start()
	}()
}


