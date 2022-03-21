package viper

import (
	"gitee.com/kratos_pkg/boot_conf"
	"github.com/spf13/viper"
)

func init() {
	_ = viper.BindEnv("nacos.adress", "NACOS_ADDRESS")
	_ = viper.BindEnv("nacos.port", "NACOS_PORT")
	_ = viper.BindEnv("nacos.username", "NACOS_USERNAME")
	_ = viper.BindEnv("nacos.password", "NACOS_PASSWORD")
	_ = viper.BindEnv("nacos.group", "NACOS_GROUP")
	_ = viper.BindEnv("nacos.dataid", "NACOS_DATAID")

}

// Config 初始化配置
func Config() *boot_conf.BootConfig {

	NacosConfig := boot_conf.NacosConfig{
		NacosConfigAddress:  viper.GetString("nacos.adress"),
		NacosConfigPort:     viper.GetInt("nacos.port"),
		NacosConfigUsername: viper.GetString("nacos.username"),
		NacosConfigPassword: viper.GetString("nacos.password"),
		NacosConfigGroup:    viper.GetString("nacos.group"),
		NacosConfigDataID:   viper.GetString("nacos.dataid"),
	}
	return &boot_conf.BootConfig{NacosConfig: NacosConfig}
}
