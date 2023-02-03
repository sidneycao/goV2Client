package storage

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

//配置文件IO管理

type File struct {
	isExist bool
	path    string
	isFile  bool
	file    *os.File
}

const W_APPEN = int(os.O_APPEND | os.O_CREATE | os.O_RDWR)
const W_NEW = int(os.O_CREATE | os.O_RDWR)

// 其实不是打开文件
// 而且检查文件或目录的属性
// 新建结构体
func Open(path string) File {
	f := File{}
	//判断文件或目录是否存在
	sts, err := os.Stat(path)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		f.isExist = false
	} else {
		f.isExist = true
	}
	fmt.Println(f)
	f.path = path
	//判断是文件还是目录
	if sts.IsDir() {
		f.isFile = false
	} else {
		f.isFile = true
	}
	return f
}

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

const configDir = "/usr/local/etc/goV2Config"

var configModule string = configDir + "/" + "config_module.json"

var module = `{
	{
		"outbounds": [
			{
				"protocol": "vmess",
				"tag": "default",
				"settings": {
					"vnext": [
					   {
						  "address": "{Add}",
						  "port": {Port},
						  "users": [
							  {
								  "encryption": "none",
								  "id": "{ID}",
								  "alterId": {Aid},
								  "security": "auto"
							  }
						  ]
					   }
					]
				},
				"streamSettings":{
					"sockopt": {
						"mark": 255
					}
				}
			}
		]
	}
}`

// 加载配置文件模板
func LoadConfigModule() string {
	// 检查配置文件目录是否存在
	d := Open(configDir)
	fmt.Println("here2...")
	if !d.isExist {
		log.Printf("the config dir %s is not exists, creating...\n", configDir)
		err := os.MkdirAll(d.path, os.ModePerm)
		if err != nil {
			log.Panic("failed to create the config dir... ")
		}
	}
	// 读取配置文件模板  如果不存在就创建
	f := Open(configModule)
	r, err := f.Read()
	if err == nil && string(r) != "" {
		return string(r)
	} else {
		CreateConfigModule()
		return module
	}
}

// 创建配置文件模板
func CreateConfigModule() {
	f := Open(configModule)
	err := f.Write(W_NEW, []string{module})
	if err != nil {
		log.Panic("failed to create config module file...")
	}
}
