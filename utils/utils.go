package utils

import (
	"encoding/json"
	"os"
)

var version = "0.0.1"

func GetUtilsVersion() string {
	return version
}

func InitUtils() int {
	//init utils
	return 0
}

/*
* 参数检查
 */
func ArgsCheck(s ...interface{}) int {
	if len(s) == 0 {
		return 0
	}

	for _, v := range s {
		if v == nil {
			return -1
		}
	}

	return 0
}

/*
* Json 编码接口
 */
func JsonEncode(filepath string, v interface{}) {
	if ArgsCheck(filepath) != 0 {
		return
	}

	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	err = json.NewEncoder(file).Encode(&v)
	if err != nil {
		return
	}
}

/*
* Json 解码接口
 */
func JsonDecode(filepath string, v interface{}) {
	if ArgsCheck(filepath) != 0 {
		return
	}
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	err = json.NewDecoder(file).Decode(&v)
	if err != nil {
		return
	}
}
