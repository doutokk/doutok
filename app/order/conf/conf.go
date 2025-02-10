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
	Env      string
	Kitex    Kitex    `yaml:"kitex"`
	MySQL    MySQL    `yaml:"mysql"`
	Redis    Redis    `yaml:"redis"`
	Registry Registry `yaml:"registry"`
}

type MySQL struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Kitex struct {
	Service  string `yaml:"service"`
	Address  string `yaml:"address"`
	LogLevel string `yaml:"log_level"`
}

type Registry struct {
	RegistryAddress []string `yaml:"registry_address"`
	Username        string   `yaml:"username"`
	Password        string   `yaml:"password"`
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
	viper.AutomaticEnv()

	// Set environment variable keys to match the configuration keys
	viper.SetEnvPrefix("APP") // Optional: set a prefix for environment variables
	viper.BindEnv("kitex.service", "APP_KITEX_SERVICE")
	viper.BindEnv("kitex.address", "APP_KITEX_ADDRESS")
	viper.BindEnv("kitex.log_level", "APP_KITEX_LOG_LEVEL")
	viper.BindEnv("mysql.host", "APP_MYSQL_HOST")
	viper.BindEnv("mysql.port", "APP_MYSQL_PORT")
	viper.BindEnv("mysql.username", "APP_MYSQL_USERNAME")
	viper.BindEnv("mysql.password", "APP_MYSQL_PASSWORD")
	viper.BindEnv("redis.address", "APP_REDIS_ADDRESS")
	viper.BindEnv("redis.username", "APP_REDIS_USERNAME")
	viper.BindEnv("redis.password", "APP_REDIS_PASSWORD")
	viper.BindEnv("redis.db", "APP_REDIS_DB")
	viper.BindEnv("registry.registry_address", "APP_REGISTRY_REGISTRY_ADDRESS")
	viper.BindEnv("registry.username", "APP_REGISTRY_USERNAME")
	viper.BindEnv("registry.password", "APP_REGISTRY_PASSWORD")

	err = viper.Unmarshal(conf)
	if err != nil {
		panic(err)
	}

	// Manually parse the registry address environment variable
	registryAddress := viper.GetString("registry.registry_address")
	if registryAddress != "" {
		conf.Registry.RegistryAddress = strings.Split(registryAddress, ",")
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
