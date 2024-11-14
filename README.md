# blog_go
go后台版本的blog网站，blog_site_go的先行版本

```plain_text
blog_go/
├── cmd/                    # 存放应用程序入口
│   └── sub.go              # 子程序入口
├── pkg/                    # 公共库和功能模块
│   ├── api/                # 存放 API 相关的代码
│   │   ├── handler/        # API 处理器
│   │   ├── router.go       # 路由配置
│   ├── config/             # 配置文件相关代码
│   ├── database/           # 数据库操作相关代码
│   ├── middleware/         # 中间件
│   ├── model/              # 数据模型
│   ├── service/            # 核心服务逻辑
│   ├── util/               # 工具类
│   └── validator/          # 请求参数验证
├── internal/               # 内部应用代码，通常是专有的业务逻辑代码，不希望外部使用
│   ├── auth/               # 鉴权相关的代码
│   └── job/                # 后台任务处理相关代码
├── scripts/                 # 一些管理脚本或部署脚本
│   └── init_db.sh          # 初始化数据库的脚本
├── migrations/             # 数据库迁移文件
│   └── 001_create_users.sql
├── api/                    # 生成的 OpenAPI/Protobuf 定义
│   └── swagger.json        # swagger 文档
├── web/                    # 前端相关文件（如果有）
│   └── static/             # 静态资源
├── test/                   # 测试代码
│   ├── api_test.go         # API 测试
│   └── service_test.go     # 核心服务测试
├── main.go                 # 程序启动入口
├── go.mod                  # Go module 配置文件
├── go.sum                  # Go module 校验文件
├── Dockerfile              # Docker 构建文件
└── README.md               # 项目说明文件
```
