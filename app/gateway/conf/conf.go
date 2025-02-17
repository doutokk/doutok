package conf

import (
	"bytes"
	_ "embed"
	"strings"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/joho/godotenv"
	"github.com/kr/pretty"
	"github.com/spf13/viper"
	"gopkg.in/validator.v2"
)

var (
	//go:embed conf.yaml
	configFile []byte
	conf       *Config
	once       sync.Once
)

type Config struct {
	Kitex struct {
		Service  string `mapstructure:"service"`
		Address  string `mapstructure:"address"`
		LogLevel string `mapstructure:"log_level"`
		OtlpAddr string `mapstructure:"otlp_address"`
	} `mapstructure:"kitex"`

	Registry struct {
		RegistryAddress []string `mapstructure:"registry_address"`
		Username        string   `mapstructure:"username"`
		Password        string   `mapstructure:"password"`
	} `mapstructure:"registry"`

	MySQL struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"mysql"`

	Redis struct {
		Address  string `mapstructure:"address"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		klog.Warn("Error loading .env file")
	}

	conf = new(Config)
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(configFile))
	if err != nil {
		panic(err)
	}

	// Enable automatic environment variable reading
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.Unmarshal(conf)
	if err != nil {
		panic(err)
	}

	if err := validator.Validate(conf); err != nil {
		klog.Error("validate config error - %v", err)
		panic(err)
	}
	pretty.Printf("%+v\n", conf)
}

func LogLevel() klog.Level {
	level := GetConf().Kitex.LogLevel
	switch level {
	case "trace":
		return klog.LevelTrace
	case "debug":
		return klog.LevelDebug
	case "info":
		return klog.LevelInfo
	case "notice":
		return klog.LevelNotice
	case "warn":
		return klog.LevelWarn
	case "error":
		return klog.LevelError
	case "fatal":
		return klog.LevelFatal
	default:
		return klog.LevelInfo
	}
}
