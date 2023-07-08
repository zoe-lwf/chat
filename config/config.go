package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var GlobalConfig *Configuration

type Configuration struct {
	MySQL struct {
		DNS string `mapstructure:"dns"`
	} `mapstructure:"mysql"`

	App struct {
		Salt           string `mapstructure:"salt"`
		IP             string `mapstructure:"ip"`               // 应用程序 IP 地址
		HTTPServerPort string `mapstructure:"http_server_port"` // HTTP 服务器端口
		WebsocketPort  string `mapstructure:"websocket_server_port"`
	} `mapstructure:"app"`

	JWT struct {
		SignKey    string `mapstructure:"sign_key"`    // JWT 签名密钥
		ExpireTime int    `mapstructure:"expire_time"` // JWT 过期时间（小时）
	} `mapstructure:"jwt"`
}

func (c Configuration) String() string {

	return fmt.Sprintf(
		c.MySQL.DNS,
		c.App.Salt,
		c.App.IP,
		c.App.HTTPServerPort,
		c.App.WebsocketPort,
		c.JWT.SignKey,
		c.JWT.ExpireTime,
	)
}

func InitConfig(configPath string) {
	//导入配置文件
	viper.SetConfigFile(configPath)
	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	GlobalConfig = new(Configuration) //值
	err = viper.Unmarshal(GlobalConfig)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	//TODO log and reload

}
