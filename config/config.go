package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	DB struct {
		Addr   string `json:"addr"`
		Port   int    `json:"port"`
		User   string `json:"user"`
		Passwd string `json:"passwd"`
		DBName string `json:"dbname"`
	} `json:"db"`
	Redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		Db       int    `json:"db"`
	} `json:"redis"`
	Admin struct {
		Name     string `json:"name"`
		Mail     string `json:"mail"`
		Password string `json:"password"`
	} `json:"admin"`
	Version   string `json:"version"`
	Port      string `json:"port"`
	LogPath   string `json:"logpath"`
	LogName   string `json:"logname"`
	Address   string `json:"address"`
	Lang      string `json:"lang"`
	Secretkey string `json:"secretkey"`
}

var conf *Configuration

/*Config ...
*
 */
func Config() *Configuration {
	if conf == nil {
		initConfig()
	}

	return conf
}

/*initConfig ...
*
 */
func initConfig() {
	var err error

	viper.SetConfigName("configuration")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("read in config file error: %s\n", err)
		os.Exit(1)
	}
	if err = viper.Unmarshal(&conf); err != nil {
		fmt.Println("unmarshal config file error:", err)
		os.Exit(1)
	}

	fmt.Println("Configuration.conf", conf)
}
