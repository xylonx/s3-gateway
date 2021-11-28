package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

var Config *Settings = new(Settings)

func SetupConfig(configPath string) error {
	v := viper.New()
	// set the default value of the config
	setDefault(v)
	v.SetConfigFile(configPath)

	// get config content
	if err := v.ReadInConfig(); err != nil {
		log.Printf("read config from path %v failed: %v\n", configPath, err)
		return err
	}
	// get config from env
	v.AutomaticEnv()

	if err := v.Unmarshal(Config); err != nil {
		log.Printf("unmarshal config failed: %v\n", err)
		return err
	}

	Config.Application.ReadTimeout = time.Second * time.Duration(Config.Application.ReadTimeoutSeconds)
	Config.Application.WriteTimeout = time.Second * time.Duration(Config.Application.WriteTimeoutSeconds)
	Config.Application.IdleTimeout = time.Second * time.Duration(Config.Application.IdleTimeoutSeconds)
	Config.Application.ShutdownWait = time.Second * time.Duration(Config.Application.ShutdownWaitSeconds)

	return nil
}

// TODO: set the default config
func setDefault(v *viper.Viper) {

}
