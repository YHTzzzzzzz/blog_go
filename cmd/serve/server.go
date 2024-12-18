package serve

import (
	"blog_go/global"
	"blog_go/router"
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// NewServerCmd 初始化并返回 serverCmd
func NewServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Start the HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			// 启动 Gin HTTP 服务器
			fmt.Println("Starting server...")
			startServer()
		},
	}
}

// 更换了目录结构，直接 暴露给 root.go 使用 修改为 NewServerCmd
//var serverCmd = &cobra.Command{
//	Use:   "server",
//	Short: "Start the HTTP server",
//	Run: func(cmd *cobra.Command, args []string) {
//		// 启动 Gin HTTP 服务器
//		fmt.Println("Starting server...")
//		startServer()
//	},
//}

func startServer() {
	// 路由绑定
	r := router.InitRoutes()

	port := global.ServerConfigInstance.AppConfigInstance.Port
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
