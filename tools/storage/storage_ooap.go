package storage

import (
	"bufio"
	"errors"
	"log"
	"os"
)

//文件管理IO

type File struct {
	isExist bool
	path    string
	isFile  bool
	//file    *os.File
}

const W_APPEN = int(os.O_APPEND | os.O_CREATE | os.O_RDWR)
const W_NEW = int(os.O_CREATE | os.O_RDWR)

// 其实不是打开文件
// 而且检查文件或目录的属性
// 新建结构体
func OpenFile(path string) File {
	f := File{}
	//判断文件或目录是否存在
	sts, err := os.Stat(path)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		f.isExist = false
	} else {
		f.isExist = true
	}
	f.path = path

	// 如果文件或目录不存在
	if !f.isExist {
		return f
	}

	//判断是文件还是目录
	if sts.IsDir() {
		f.isFile = false
	} else {
		f.isFile = true
	}
	return f
}

// 读文件
func (f *File) Read() ([]byte, error) {
	ff, err := os.OpenFile(f.path, os.O_RDONLY, os.FileMode(0777))
	if err != nil {
		return nil, err
	}
	defer ff.Close()
	r := bufio.NewReader(ff)
	bufRes := make([]byte, 0)

	bufTemp := make([]byte, 1024)
	for {
		n, err := r.Read(bufTemp)
		if err != nil {
			return nil, err
		}
		if n == 0 {
			break
		} else {
			bufRes = append(bufRes, bufTemp[:n]...)
		}
	}
	return bufRes, nil
}

// 文件不存在会创建
// 根据writeType确定是覆盖写还是append写
func (f *File) Write(writeType int, b []string) error {
	ff, err := os.OpenFile(f.path, writeType, os.FileMode(0777))
	if err != nil {
		return err
	}
	defer ff.Close()

	for _, v := range b {
		n, err := ff.WriteString(v)
		if err != nil {
			log.Println(err)
		}
		if n < len(v) {
			log.Println("write byte num error")
		}
	}
	return nil
}

// 创建目录
func (f *File) CreateDir() error {
	err := os.MkdirAll(f.path, os.ModePerm)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// 删除
func (f *File) Delete() {
	if !f.isExist {
		return
	}
	if f.isFile {
		err := os.Remove(f.path)
		if err != nil {
			log.Panic(err)
		}
		f.refresh(f.path)
		return
	} else {
		err := os.RemoveAll(f.path)
		if err != nil {
			log.Panic(err)
		}
		f.refresh(f.path)
		return
	}

}

func (f *File) refresh(path string) {
	*f = OpenFile(path)
}
