package main

import (
	config "Go-Banana/config"   //config
	dbproxy "Go-Banana/dbproxy" //db proxy
	utils "Go-Banana/utils"     //utils
	datastruct "Go-Banana/datastruct"  //data struct
	"fmt"
	//utils and ti
	//gin framework
)

type ModuleInitList type {
	Module package
	InitModule func
	DestroyModule func
}
var g_ModuleInitList ModuleInitList

func main() {
	Init()
	fmt.Println("Go Banana!")
}

func Init() int {
	g_ModuleInitList = []ModuleInitList



	ret := utils.InitUtils()
	if ret != 0 {
		fmt.Println("Init utils module fail! ret = ", ret)
	}

	ret = dbproxy.InitDBProxy()
	if ret != 0 {
		fmt.Println("Init db proxy module fail! ret = ", ret)
	}
	ret = config.InitConf()
	if ret != 0 {
		fmt.Println("Init conf module fail! ret = ", ret)
	}

	return 0
}
