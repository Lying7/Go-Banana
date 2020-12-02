package utils

import (
	"time"
)

/*GetTime ...
*
 */
func GetTime() time.Time {
	return time.Now()
}

/*GetTimeFormat ...
*
 */
func GetTimeFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

/* GetMillisec,MS,msec,MSEL,[计量] milliscond
*
 */
func GetMillisec() int64 {
	return (time.Now().UnixNano() / 1e6)
}
