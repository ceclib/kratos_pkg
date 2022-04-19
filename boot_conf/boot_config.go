package boot_conf

// NacosConfig 启动的时候依赖的nacos配置
type NacosConfig struct {
	Address     string
	Port        int
	Username    string
	Password    string
	Group       string
	DataID      string
	NameSpaceID string
}

type LogConfig struct {
	LogPath string
	AppName string
	Level   int
	IsDev   bool //是否在开发环境
}

// BootConfig 启动依赖的基础配置
type BootConfig struct {
	NacosConfig NacosConfig
	LogConfig   LogConfig
}
