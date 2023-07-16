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
		RPCPort        string `mapstructure:"rpc_port"`         //rpc RPC 服务器端口
		WorkerPoolSize uint32 `mapstructure:"worker_pool_size"` //队列数量
		MaxWorkerTask  uint32 `mapstructure:"max_worker_task"`  //业务 worker 对应负责的任务队列最大任务存储数量
	} `mapstructure:"app"`

	JWT struct {
		SignKey    string `mapstructure:"sign_key"`    // JWT 签名密钥
		ExpireTime int    `mapstructure:"expire_time"` // JWT 过期时间（小时）
	} `mapstructure:"jwt"`

	Redis struct {
		Addr     string `mapstructure:"addr"`     // Redis 地址
		Password string `mapstructure:"password"` // Redis 认证密码
	} `mapstructure:"redis"`
	RabbitMQ struct {
		URL string `mapstructure:"url"`
	}
}

func (c Configuration) String() string {

	return fmt.Sprintf(
		c.MySQL.DNS,
		c.App.Salt,
		c.App.IP,
		c.App.HTTPServerPort,
		c.App.RPCPort,
		c.App.WebsocketPort,
		c.App.WorkerPoolSize,
		c.App.MaxWorkerTask,
		c.JWT.SignKey,
		c.JWT.ExpireTime,
		c.Redis.Addr,
		c.Redis.Password,
		c.RabbitMQ.URL,
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
