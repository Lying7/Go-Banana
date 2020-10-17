package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var version = "0.0.1"

var svr_conf_file = "conf/svr_conf.json"

func GetConfVersion() string {
	return version
}

func InitConf() int {
	jsonParse := JsonStruct{}
	svrConf := SvrConfStruct{}

	jsonParse.Load(svr_conf_file, &svrConf)

	fmt.Println("Init Config Version:", svrConf.Version)

	return 0
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
