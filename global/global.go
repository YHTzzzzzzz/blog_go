package global

import ut "github.com/go-playground/universal-translator"

// ServerConfigInstance 全局的服务器配置实例
var ServerConfigInstance *serverConfig

// TranslatorInstance 创建全局的翻译器实例
var TranslatorInstance ut.Translator
