package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 配置文件映射 根节点
type serverConfig struct {
	AppConfigInstance  appConfig  `mapstructure:"server" json:"server" yaml:"server" toml:"server"`
	AuthorInfoInstance authorInfo `mapstructure:"author" json:"author" yaml:"author" toml:"author"`
}

// 应用信息
type appConfig struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env" toml:"env"`
	Port    string `mapstructure:"port" json:"port" yaml:"port" toml:"port"`
	Name    string `mapstructure:"name" json:"name" yaml:"name" toml:"name"`
	Version string `mapstructure:"version" json:"version" yaml:"version" toml:"version"`
}

type authorInfo struct {
	Name   string `mapstructure:"name" json:"name" yaml:"name" toml:"name"`
	Email  string `mapstructure:"email" json:"email" yaml:"email" toml:"email"`
	GitHub string `mapstructure:"github" json:"github" yaml:"github" toml:"github"`
}

var (
	// ServerConfigInstance 配置文件全局实例
	ServerConfigInstance serverConfig
)

// LoadConfiguration 初始化配置
func LoadConfiguration(configFile string) error {
	//// 初始化配置
	//configFile = "pkg/config/settings.yml"

	// 初始化viper
	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config file: %s", err))
	}

	// 配置监听 支持热更新
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&ServerConfigInstance); err != nil {
			fmt.Println(err)
		}
	})

	// 配置赋值给全局变量
	if err := v.Unmarshal(&ServerConfigInstance); err != nil {
		panic(fmt.Errorf("unmarshal config file: %s", err))
	}

	return nil
}
