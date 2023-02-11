package tool

import "os"

// 判断文件/文件夹是否存在
func IsFileExisted(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if os.IsNotExist(err) {
		return false
	}
	return true
}
