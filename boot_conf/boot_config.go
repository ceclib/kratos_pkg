package boot_conf

// NacosConfig 启动的时候依赖的nacos配置
type NacosConfig struct {
	NacosConfigAddress  string
	NacosConfigPort     int
	NacosConfigUsername string
	NacosConfigPassword string
	NacosConfigGroup    string
	NacosConfigDataID   string
}

// BootConfig 启动依赖的基础配置
type BootConfig struct {
	NacosConfig
}
