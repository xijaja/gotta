package files

import (
	"bufio"
	"fmt"
	"os"
)

type GottaFiles byte

// WriteSth 读取文件并写入文件, 如果文件不存在则创建文件, 如果文件存在则覆盖
func (gf *GottaFiles) WriteSth(filePath string, sth string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	// 及时关闭file句柄
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("文件关闭失败", err)
		}
	}(file)
	// 写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	_, err = write.WriteString(sth)
	if err != nil {
		fmt.Println("写入文件失败", err)
		return
	}
	// Flush将缓存的文件真正写入到文件中
	err = write.Flush()
	if err != nil {
		fmt.Println("写入文件失败", err)
		return
	}
}
