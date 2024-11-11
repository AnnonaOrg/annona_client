package utils

import (
	"errors"
	"os"
	"strings"
)

// 判断是否为系统目录的函数
func isSystemDirectory(pathName string) bool {
	// 定义一些常见的系统目录
	systemDirs := []string{
		"/sys", "/proc", "/dev", // Linux 系统目录
		"C:\\Windows", "C:\\Program Files", // Windows 系统目录
	}

	// 判断目标路径是否以某个系统目录为前缀
	for _, dir := range systemDirs {
		if strings.HasPrefix(pathName, dir) {
			return true
		}
	}

	return false
}
func Remove(name string) error {
	if isSystemDirectory(name) {
		return errors.New(name + " is a system directory")
	}
	return os.Remove(name)
}
