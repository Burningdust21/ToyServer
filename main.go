package main

import (
	"fmt"
	"net/http"
	"local/logger"
)
/*
Requirements: 用 Golang 开发一个程序，实现一个 HTTP API，返回系统的信息
  [D] OS 信息
  [D] CPU 使用量前 5 的进程
  [D] 内存使用量前 5 的进程
  [D] 支持多种数据源，包括 1）Linux 作业中的日志文件 2）Golang 程序自己采集的数据
  [D] 必须使用到 Interface、goroutine & channel
  [D] 至少 1 种设计模式，并注释说明
  [] API 满足 REST 规范，并提供 API 文档
  [] 单测覆盖率达 20% 以上
  [D] 使用 go mod 管理第三方库
*/

func main() {
	fmt.Println("This is my server!", http.TimeFormat)
	if err := http.ListenAndServe(":2021", http.HandlerFunc(GetSystemInfo)); err != nil {
		panic(err)
	}
}


func GetSystemInfo(w http.ResponseWriter, req *http.Request){
	if req.URL.Path == "/"{
		if _, err := fmt.Fprintln(w, logger.FromGo()); err != nil {
			panic(err)
		}
	} else if req.URL.Path == "/log" {
		if _, err := fmt.Fprintln(w, logger.FromLog()); err != nil {
			panic(err)
		}
	}
}