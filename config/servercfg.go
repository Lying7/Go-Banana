package config

import (
	"Go-Banana/glog"
	"encoding/json"
	"os"
)

var ServerCfg ServerConfig
var svrCfgPath string = "./config/serverconfig.json"

/*ServerConfig ...
*
 */
type ServerConfig struct {
	Version string
	Port    int16
	LogPath string
}

func init() {
	cfg, ok := getConfig(svrCfgPath).(ServerConfig)
	if !ok {
		glog.Error("Get Server Config Fail!")
	}
	ServerCfg = cfg
}

/*getConfig ...
*
 */
func getConfig(path string) interface{} {
	fp, err := os.OpenFile(path, os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		glog.Error(err.Error())
	}
	data := make([]byte, 100)
	n, err := fp.Read(data)
	if err != nil {
		glog.Error(err.Error())
	}
	glog.Info(string(data[:n]))

	var svrCfg ServerConfig
	err = json.Unmarshal(data[:n], &svrCfg)
	if err != nil {
		glog.Error(err.Error())
	}
	return svrCfg
}
