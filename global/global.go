package global

import (
	"blog_go/pkg/task"
	ut "github.com/go-playground/universal-translator"
)

// ServerConfigInstance 全局的服务器配置实例
var ServerConfigInstance *serverConfig

// TranslatorInstance 创建全局的翻译器实例
var TranslatorInstance ut.Translator

// TaskRegistryInstance 全局的任务注册实例
var TaskRegistryInstance *task.Registry
