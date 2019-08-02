package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"socialbot/internal/common/app"
	"socialbot/internal/common/setting"
	"xorm.io/core"
)

var SocialBotOrm *xorm.Engine

func LoadModel() (err error) {
	SocialBotOrm, err = newOrm(setting.Cfg.Db.Socialbot)
	if err != nil {
		return err
	}
	app.RegShutdownCallback(func() {
		fmt.Println("[INFO] close db start")
		err := SocialBotOrm.Close()
		if err != nil {
			fmt.Printf("[ERROR] close db error %+v \n", err)
		}
		fmt.Println("[INFO] close db end")
	})
	return nil
}

func newOrm(c setting.DbConfig) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(c.Driver, c.DataSource)
	if err != nil {
		return nil, fmt.Errorf("xorm.NewEngine failed, param %#v, error: %v \n", c, err)
	}
	//engine.Logger().SetLevel(core.LOG_ERR)

	// ping
	err = engine.Ping()
	if err != nil {
		return nil, fmt.Errorf("xorm.NewEngine ping failed, param %#v, error: %v \n", c, err)
	}

	engine.SetMaxOpenConns(c.MaxOpenCon)
	engine.SetMaxIdleConns(c.MaxIdleCon)

	// set table mapper
	engine.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, c.Prefix))

	// set column mapper
	//engine.SetColumnMapper(ColumnMapper{})
	return engine, nil
}
