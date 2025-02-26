package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WriteFile(filePath string, byteFile []byte) {
	// 获取目录路径
	dirPath := filepath.Dir(filePath)

	// 检查目录是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 目录不存在，创建目录
		errr := os.MkdirAll(dirPath, 0755) // 0755 是权限
		if errr != nil {
			panic(errr)
		}
		fmt.Printf("目录 %s 创建成功\n", dirPath)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 如果文件不存在，则创建并写入嵌入的内容
		errr := ioutil.WriteFile(filePath, byteFile, 0644)
		if errr != nil {
			fmt.Println("无法写入文件：", errr)
			return
		}
		fmt.Println("文件已写入：", filePath)
	} else {
		fmt.Println("文件已存在：", filePath)
	}
}
