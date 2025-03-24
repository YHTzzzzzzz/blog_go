package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

// Registry 定时任务注册中心
type Registry struct {
	CronScheduler *cron.Cron
}

// NewTaskRegistry 创建一个新的 TaskRegister 实例
func NewTaskRegistry() *Registry {
	return &Registry{
		CronScheduler: cron.New(),
	}
}

// RegisterTask 注册任务，任务会按照 cron 表达式运行
func (tr *Registry) RegisterTask(cronExpr string, task func()) (cron.EntryID, error) {
	entryID, err := tr.CronScheduler.AddFunc(cronExpr, task)
	if err != nil {
		return 0, fmt.Errorf("无法注册定时任务: %v", err)
	}
	return entryID, nil
}

// UnregisterTask 取消注册任务
func (tr *Registry) UnregisterTask(entryID cron.EntryID) {
	tr.CronScheduler.Remove(entryID)
}

// Start 启动所有注册的任务
func (tr *Registry) Start() {
	tr.CronScheduler.Start()
}

// Stop 停止定时任务调度
func (tr *Registry) Stop() {
	tr.CronScheduler.Stop()
}

// LogToFile 记录日志到文件
func LogToFile(message string) {
	logFile, err := os.OpenFile("task.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("无法打开日志文件", err)
		return
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			fmt.Println("无法关闭日志文件", err)
		}
	}(logFile)

	logger := log.New(logFile, "TASK_LOG: ", log.LstdFlags)
	logger.Println(message)
}
