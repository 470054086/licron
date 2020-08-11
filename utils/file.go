package utils

import (
	"os"
)

// 目录或者文件是否存在
func PathExist(dirPath string) error {
	stat, err := os.Stat(dirPath)
	if err != nil {
		_ = stat
		return err
	}
	if os.IsNotExist(err) {
		return nil
	}
	return nil
}

// 判断是文件还是目录
func IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

func MkdirPath(dirPath string) error {
	return nil
}
