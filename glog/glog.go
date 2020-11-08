package glog

import (
	"Go-Banana/utils"
	"fmt"
)

/*Info ...
*
 */
func Info(agrs ...string) {
	fmt.Println("[INFO][", utils.GetCurTime(), "]", agrs)
	return
}

/*Debug ...
*
 */
func Debug(agrs ...string) {
	fmt.Println("[DEBUG][", utils.GetCurTime(), "]", agrs)
	return
}

/*Warn ...
*
 */
func Warn(agrs ...string) {
	fmt.Println("[WARN][", utils.GetCurTime(), "]", agrs)
	return
}

/*Error ...
*
 */
func Error(agrs ...string) {
	fmt.Println("[ERROR][", utils.GetCurTime(), "]", agrs)
	return
}
