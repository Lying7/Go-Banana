package utils

import (
	"time"
)

/*GetCurTime ...
*
 */
func GetCurTime() string {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	return timeStr
}
