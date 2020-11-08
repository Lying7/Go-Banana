package glog

import (
	"fmt"
)

/*Info ...
*
 */
func Info(agrs ...string) {
	fmt.Println(agrs)
	return
}

/*Debug ...
*
 */
func Debug(agrs ...string) {
	fmt.Println(agrs)
	return
}

/*Warn ...
*
 */
func Warn(agrs ...string) {
	fmt.Println(agrs)
	return
}

/*Error ...
*
 */
func Error(agrs ...string) {
	fmt.Println(agrs)
	return
}
