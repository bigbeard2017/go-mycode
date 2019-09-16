package main

import (
	"fmt"
	"go-mycode/HealthDetect/process"
)

func main() {
	var a process.Process
	x := &a
	result, err := x.GetAllProcess()
	if err != nil {
		println(err)
		return
	}
	for p := range result {
		o := result[p]
		fmt.Printf("cpu:%f,mem:%f,pid:%d,processName:%s,cmd:%s\r", o.CPU, o.Memory, o.Pid, o.ProcessName, o.ProcessPath)
	}
}
