package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"path/filepath"
)

var Conf = new(Config)

type Config struct {
	EmailConfig `mapstructure:"email"`
	SmsConfig   `mapstructure:"sms"`
}

type EmailConfig struct {
	SendName string   `mapstructure:"send_name"`
	EmailQQ  EmailQQ  `mapstructure:"email_QQ"`
	Email163 Email163 `mapstructure:"email_163"`
}

type EmailQQ struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	AuthCode string `mapstructure:"auth_code"`
}

type Email163 struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	AuthCode string `mapstructure:"auth_code"`
}

type SmsConfig struct {
	ApiKey string `mapstructure:"api_key"`
	TplId  TplId  `mapstructure:"tpl_id"`
}

type TplId struct {
	CodeIssue string `mapstructure:"code_issue"`
}

func init() {
	absPath, err := filepath.Abs(".")
	if err != nil {
		panic(fmt.Errorf("get abs path failed, err: %v", err))
	}
	cPath := filepath.Join(absPath, "config/config.yaml")
	log.Printf("==========> config path: %v", cPath)
	viper.SetConfigFile(cPath)

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Infoln("配置文件修改啦...")
		if err := viper.Unmarshal(&Conf); err != nil {
			log.Errorf("load config failed, err: %v", err)
		}
	})

	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	if err = viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}

	// 加载环境变量
	viper.AutomaticEnv()

	return
}
