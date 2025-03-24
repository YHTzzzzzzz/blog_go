package config

import (
	"blog_go/global"
	"blog_go/pkg/task"
)

func NewTaskRegister() {
	global.TaskRegistryInstance = task.NewTaskRegistry()
}
