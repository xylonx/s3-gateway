package config

import "time"

type Settings struct {
	Application ApplicationConfig `mapstructure:"application"`
	Providers   []ProviderSettings
}

type ApplicationConfig struct {
	Host                string        `mapstructure:"host"`
	Port                string        `mapstructure:"port"`
	ReadTimeoutSeconds  int           `mapstructure:"read_timeout_seconds"`
	ReadTimeout         time.Duration `mapstructure:-"`
	WriteTimeoutSeconds int           `mapstructure:"write_timeout_seconds"`
	WriteTimeout        time.Duration `mapstructure:"-"`
	IdleTimeoutSeconds  int           `mapstructure:"idle_timeout_seconds"`
	IdleTimeout         time.Duration `mapstructure:"-"`
	ShutdownWaitSeconds int           `mapstructure:"shutdown_wait_seconds"`
	ShutdownWait        time.Duration `mapstructure:"-"`
}

type ProviderSettings struct {
	Type              string
	Endpoint          string
	AccessID          string
	AccessKey         string
	VirtualBucketname string
	Region            string
	BucketName        string
}
