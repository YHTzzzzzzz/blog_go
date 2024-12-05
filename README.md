# blog_go
go后台版本的blog网站，blog_site_go的先行版本

## 项目结构
```plain_text
/blog_go
├── cmd
│   ├── migrate           # 数据库迁移命令
│   │   └── migrate.go    # 数据库迁移命令实现
│   ├── serve             # 启动服务器的命令
│   │   └── server.go     # 启动服务器的命令实现
│   └── root.go           # 根命令
├── config
│   ├── database.go       # 数据库配置
│   ├── general.go        # 全局配置
│   ├── server.go         # 服务器配置
│   ├── settings.yml      # 配置文件
│   └── validation.go     # 验证相关配置
├── global
│   └── global.go         # 全局配置或全局变量
├── middleware
│   ├── auth.go           # 认证中间件
│   ├── logging.go        # 日志记录中间件
│   └── validation.go     # 验证请求参数的中间件
├── models
│   ├── request           # 请求结构体
│   │   ├── login_request.go    # 登录请求结构体
│   │   └── register_request.go # 注册请求结构体
│   ├── article.go        # 文章数据模型
│   └── user.go           # 用户数据模型
├── pkg
│   ├── logger            # 日志工具
│   ├── utils             # 工具函数模块
│   └── validator         # 自定义验证工具
├── router
│   └── routes.go         # 路由初始化
├── types
│   ├── constants         # 常用状态码常量
│   │   └── status_code.go # 常用状态码常量定义
│   ├── enums             # 枚举类型定义
│   │   └── user_role.go  # 用户角色枚举类型
│   ├── errors            # 自定义错误类型
│   │   └── custom_errors.go # 自定义错误类型
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
