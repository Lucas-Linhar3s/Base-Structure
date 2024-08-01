package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// NewConfig returns a new Viper instance configured to read the given config file.
func NewViper(p string) *viper.Viper {
	envConf := os.Getenv("APP_CONF")
	if envConf == "" {
		envConf = p
	}
	fmt.Println("load conf file:", envConf)
	return getViper(envConf)
}

func getViper(dir string) *viper.Viper {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(currentDir, dir)
	conf := viper.New()
	conf.SetConfigFile(path)
	err = conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}

// LoadAttributes loads attributes from config file and returns a new Config instance.
func LoadAttributes(conf *viper.Viper) *Config {
	return &Config{
		Env: conf.GetString("env"),
		Http: &Http{
			Host: conf.GetString("http.host"),
			Port: conf.GetString("http.port"),
		},
		Security: &Security{
			ApiSign: &ApiSign{
				AppKey: conf.GetString("security.api_sign.app_key"),
				AppSecurity: conf.GetString("security.api_sign.app_security"),
			},
			Jwt: &Jwt{
				Key: conf.GetString("security.jwt.key"),
			},
		},
		Data: &Data{
			Db: &Db{
				User: &User{
					Driver: conf.GetString("data.db.user.driver"),
					Nick: conf.GetString("data.db.user.nick"),
					Name: conf.GetString("data.db.user.name"),
					Username: conf.GetString("data.db.user.username"),
					Password: conf.GetString("data.db.user.password"),
					Hostname: conf.GetString("data.db.user.hostname"),
					Port: conf.GetString("data.db.user.port"),
					MaxConn: conf.GetInt("data.db.user.max_conn"),
					MaxIdle: conf.GetInt("data.db.user.max_idle"),
					TransactionTimeout: conf.GetInt("data.db.user.transaction_timeout"),
					Dsn:     conf.GetString("data.db.user.dsn"),
				},
			},
		},
		Log: &Log{
			LogLevel: conf.GetString("log.log_level"),
			Encoding: conf.GetString("log.encoding"),
			LogFileName: conf.GetString("log.log_file_name"),
			MaxBackups: conf.GetInt("log.max_backups"),
			MaxAge: conf.GetInt("log.max_age"),
			MaxSize: conf.GetInt("log.max_size"),
			Compress: conf.GetBool("log.compress"),
		},
	}
}
