package task

import (
	"fmt"
	"testing"
	"time"
)

// 任务示例
func task1() {
	fmt.Println("任务1 执行时间:", time.Now()) // 你可以在这里添加日志记录功能
	LogToFile("任务1 执行时间: " + time.Now().String())
}

func task2() {
	fmt.Println("任务2 执行时间:", time.Now())
	LogToFile("任务2 执行时间: " + time.Now().String())
}

func TestTask(t *testing.T) {
	// 创建任务注册器
	taskRegister := NewTaskRegistry()

	// 注册定时任务，按 cron 表达式执行
	entryID1, err := taskRegister.RegisterTask("@every 5s", task1) // 每 5 秒执行一次 task1
	if err != nil {
		fmt.Println("注册任务1失败:", err)
		return
	}

	_, err = taskRegister.RegisterTask("@every 10s", task2) // 每 10 秒执行一次 task2
	if err != nil {
		fmt.Println("注册任务2失败:", err)
		return
	}

	// 启动所有任务
	taskRegister.Start()

	// 模拟任务执行一段时间后移除任务
	time.Sleep(20 * time.Second)

	// 移除任务1
	taskRegister.UnregisterTask(entryID1)
	fmt.Println("已移除任务1")

	// 程序保持运行，等待定时任务执行
	select {}
}
