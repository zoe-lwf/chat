package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var GlobalConfig *Configuration

type Configuration struct {
	MySQL struct {
		DNS string `yaml:"dns" json:"dns"`
	} `mapstructure:"mysql"`
}

func (c Configuration) String() string {

	return fmt.Sprintf(c.MySQL.DNS)
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
