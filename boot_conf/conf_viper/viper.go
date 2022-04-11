package conf_viper

import (
	"fmt"
	"gitee.com/ceclib/kratos_pkg/boot_conf"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

func init() {
	_ = viper.BindEnv("nacos.adress", "NACOS_ADDR")
	_ = viper.BindEnv("nacos.port", "NACOS_PORT")
	_ = viper.BindEnv("nacos.username", "NACOS_USERNAME")
	_ = viper.BindEnv("nacos.password", "NACOS_PASSWORD")
	_ = viper.BindEnv("nacos.group", "NACOS_GROUP")
	_ = viper.BindEnv("nacos.dataid", "NACOS_DATAID")
	_ = viper.BindEnv("nacos.namespace_id", "NACOS_NAMESPACE")

}

// Config 初始化配置

func Config() *boot_conf.BootConfig {
	//原始地址
	sourceAddress := viper.GetString("nacos.adress")
	//原始端口
	port := viper.GetInt("nacos.port")
	//切分端口
	addressArray := strings.Split(sourceAddress, ":")

	var address string

	//如果在地址中带了端口的话，则用这个端口覆盖掉参数的端口
	if len(addressArray) == 2 {
		address = addressArray[0]
		port, _ = strconv.Atoi(addressArray[1])
	} else {
		address = sourceAddress
	}
	NacosConfig := boot_conf.NacosConfig{
		Address:     address,
		Port:        port,
		Username:    viper.GetString("nacos.username"),
		Password:    viper.GetString("nacos.password"),
		Group:       viper.GetString("nacos.group"),
		DataID:      viper.GetString("nacos.dataid"),
		NameSpaceID: viper.GetString("nacos.namespace_id"),
	}
	fmt.Printf("启动的配置文件为:%+v", NacosConfig)
	return &boot_conf.BootConfig{NacosConfig: NacosConfig}
}
