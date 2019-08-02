package setting

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	j "github.com/json-iterator/go"
	"io/ioutil"
	"os"
	"path/filepath"
	"socialbot/pkg/utils"
)

var (
	AppVar     = "0.0.1"
	AppPath    string
	ConfigPath string
	Cfg  *Config
)
const defaultConfig = "config.json"

type Config struct {
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
	Server struct {
		Mode              string  `json:"mode"`
		Listen          string `json:"listen"`
		ReadTimeout     int    `json:"readTimeout"`
		WriteTimeout    int    `json:"writeTimeout"`
		IdleTimeout     int    `json:"idleTimeout"`
		ShutdownTimeout int    `json:"shutdownTimeout"`
	} `json:"server"`
	Db struct {
		Socialbot DbConfig `json:"socialbot"`
	} `json:"db"`
}

type DbConfig struct {
	Driver      string `json:"driver"`
	DataSource  string `json:"dataSource"`
	Prefix      string `json:"prefix"`
	MaxLifetime int    `json:"maxLifetime"`
	MaxOpenCon  int    `json:"maxOpenCon"`
	MaxIdleCon  int    `json:"maxIdleCon"`
}

func LoadFlag() {
	conf := flag.String("c", "", "--c config.json   [config path]", )
	showVersion := flag.Bool("v", false, "-v [show REPL version]", )
	help := flag.Bool("h", false, "-h [show help list]", )

	flag.Parse()
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
		return
	}
	if *showVersion {
		fmt.Println(AppVar)
		os.Exit(0)
		return
	}

	ConfigPath = *conf
	if ConfigPath == "" {
		ConfigPath = filepath.Join("config", defaultConfig)
	}
}

func LoadAppPath() error {
	path, err := utils.AppAbsPath()
	if err != nil {
		return err
	}
	AppPath = path
	return nil
}

func LoadConf() error {
	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		return errors.New("config file not found")
	}

	result, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		return errors.Wrap(err, "read config  failed!")
	}
	c := &Config{}
	err = j.ConfigCompatibleWithStandardLibrary.Unmarshal(result, c)
	if err != nil {
		return errors.Wrap(err, "decode config file failed")
	}
	Cfg = c
	return nil
}