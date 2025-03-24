# blog_go
go后台版本的blog网站，blog_site_go的先行版本

## 项目结构
```plain_text
/blog_go
├── cmd
│   ├── root.go          # Cobra 命令入口文件
│   └── serve
│       └── server.go    # 服务器启动文件
├── config
│   ├── general.go      # 配置加载逻辑
│   ├── settings.yml    # 配置文件
│   └── validation.go   # 验证相关配置
├── global
│   └── global.go       # 全局配置或全局变量
├── handler             # 处理请求的逻辑
│   └── example_handler.go # 示例请求处理逻辑
├── middleware
│   └── validation.go   # 验证请求参数的中间件
├── models
│   ├── request         # 存放请求结构体
│   │   └── example_request.go # 示例请求结构体
│   └── response        # 存放响应结构体
│       └── example_response.go # 示例响应结构体
├── pkg
│   └── task
│       ├── task_register.go  # 定时任务注册器
│       └── task_example.go   # 示例任务实现
├── router
│   └── routes.go       # 路由初始化
├── service
│   └── example_service.go # 示例业务逻辑实现
├── types
│   ├── constants       # 存放常量的 .go 文件
│   │   └── example_constants.go # 示例常量文件
│   ├── custom_errors.go  # 自定义错误类型
│   └── response.go       # 通用返回值结构定义
├── main.go               # 应用入口文件
└── go.mod                # Go 模块文件
```

## 项目启动

> 帮助命令
> ```shell
> go run main.go -h
> ```
> 启动命令
> ```shell
> go run main.go server -c ./config/settings.yml
> ```
> > config(-c) 参数可选，默认配置文件路径 ./config/settings.yml
