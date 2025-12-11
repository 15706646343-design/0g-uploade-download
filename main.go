// 0g-storage-example 主程序入口
package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ququzone/0g-storage-example/cmd"
)

func main() {

	// 加载 .env 环境变量文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	cmd.Execute()//执行命令行程序
}

