package cmd

import (
	"blog_go/pkg/config"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:          "blog",
	Short:        "blog site backend",
	Long:         "blog",
	SilenceUsage: true, // 防止重复打印帮助信息 默认false
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tips()
			return errors.New("缺少参数")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tips()
	},
}

// tips
func tips() {
	welcome := fmt.Sprintf("欢迎使用 %s %s，使用 `-h`或`--help` 获取更多帮助。",
		config.ServerConfigInstance.AppConfigInstance.Name,
		config.ServerConfigInstance.AppConfigInstance.Version)
	author := fmt.Sprintf("作者: %s", config.ServerConfigInstance.AuthorInfoInstance.Name)
	email := fmt.Sprintf("邮箱: %s", config.ServerConfigInstance.AuthorInfoInstance.Email)
	github := fmt.Sprintf("Github: %s", config.ServerConfigInstance.AuthorInfoInstance.GitHub)
	fmt.Println(welcome)
	fmt.Println(author)
	fmt.Println(email)
	fmt.Println(github)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// 可以在项目启动前做一些初始化配置
	cobra.OnInitialize(initConfig)

	// 定义子命令共享标签 定义标签  config 缩写 -c 默认值 pkg/config/settings.yml 描述信息： Path to the configuration file
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "pkg/config/settings.yml", "Path to the configuration file")
	//fmt.Println(configFile)

	rootCmd.AddCommand(serverCmd)
}

func initConfig() {
	// 读取配置文件
	if err := config.LoadConfiguration(configFile); err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	// 初始化 validator 翻译器
	config.InitValidator()
}
