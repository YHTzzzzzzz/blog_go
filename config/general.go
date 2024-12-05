package config

import (
	"blog_go/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
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
		if err := v.Unmarshal(&global.ServerConfigInstance); err != nil {
			fmt.Println(err)
		}
	})

	// 配置赋值给全局变量
	if err := v.Unmarshal(&global.ServerConfigInstance); err != nil {
		panic(fmt.Errorf("unmarshal config file: %s", err))
	}

	return nil
}
