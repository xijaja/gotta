package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"
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

// TraverseFileList 遍历指定目录的文件
func (gf *GottaFiles) TraverseFileList(path string) (newPathList []string) {
	// 增加后缀
	if path[len(path)-1:] != "/" {
		path += "/"
	}
	// 开始遍历
	iterateOverFiles(path, func(newPath string) {
		newPathList = append(newPathList, newPath)
	})
	// 等待遍历结束
RE:
	length := len(newPathList)
	time.Sleep(2 * time.Second)
	for len(newPathList) != length {
		goto RE
	}
	return newPathList
}

// iterateOverFiles 遍历指定路径的文件
func iterateOverFiles(path string, findOne func(file string)) {
	fs, _ := ioutil.ReadDir(path) // 获取路径下的文件列表
	for _, file := range fs {
		if file.IsDir() {
			// 遇到文件夹时就开启一个并发递归
			go iterateOverFiles(path+file.Name()+"/", findOne)
		} else {
			filePath := path + file.Name()
			fmt.Printf("扫描文件: \"%s\"\n", filePath)
			findOne(filePath) // 调用函数参数
		}
	}
}
