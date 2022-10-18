package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"supabase_server/pkg/util/environment"
)

type config struct {
	Database struct {
		DSN string `json:"dsn"`
	}
	Server struct {
		Address    string `json:"address"`
		ApiAddress string `json:"apiAddress"`
		JWKSUrl    string `json:"jwksUrl"`
	}
}

var C config

// ReadConfigOption is a config option
type ReadConfigOption struct {
	AppEnv string
}

func ReadConfig(option ReadConfigOption) {
	Config := &C

	if environment.IsDev() {
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config")
	} else if environment.IsTest() || (option.AppEnv == environment.Test) {
		fmt.Println(rootDir())
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config.test")
	} else if environment.IsE2E() || (option.AppEnv == environment.E2E) {
		fmt.Println(rootDir())
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config.e2e")
	} else {
		// production configuration here
		fmt.Println("production configuration here")
	}

	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
