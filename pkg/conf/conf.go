package conf

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

var confPath = flag.String("conf", "local.json", "配置文件地址")

func New() {
	if !flag.Parsed() {
		flag.Parse()
	}
	viper.SetConfigFile(*confPath)
	if e := viper.ReadInConfig(); e != nil {
		panic(e)
	}
	fmt.Println(viper.GetString("foo"))
}
