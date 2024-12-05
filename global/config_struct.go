package global

// serverConfig 配置文件映射 根节点
type serverConfig struct {
	AppConfigInstance  appConfig  `mapstructure:"server" json:"server" yaml:"server" toml:"server"`
	AuthorInfoInstance authorInfo `mapstructure:"author" json:"author" yaml:"author" toml:"author"`
}

// appConfig 应用信息结构
type appConfig struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env" toml:"env"`
	Port    string `mapstructure:"port" json:"port" yaml:"port" toml:"port"`
	Name    string `mapstructure:"name" json:"name" yaml:"name" toml:"name"`
	Version string `mapstructure:"version" json:"version" yaml:"version" toml:"version"`
	Locale  string `mapstructure:"locale" json:"locale" yaml:"locale" toml:"locale"`
}

// authorInfo 作者信息结构
type authorInfo struct {
	Name   string `mapstructure:"name" json:"name" yaml:"name" toml:"name"`
	Email  string `mapstructure:"email" json:"email" yaml:"email" toml:"email"`
	GitHub string `mapstructure:"github" json:"github" yaml:"github" toml:"github"`
}
