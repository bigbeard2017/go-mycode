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

	for index := 0; index < len(result); index++ {
		for j := index; j < len(result); j++ {
			if result[index].Pid < result[j].Pid {
				x := result[index]
				result[index] = result[j]
				result[j] = x
			}
		}
	}
	println("==========================================")
	for p := range result {
		o := result[p]
		fmt.Printf("user:%s,\tstartTime:%s,\tcpu:%f,\tmem:%f,\tpid:%d,\tcmd:%s\r", o.User, o.StartTime, o.CPU, o.Memory, o.Pid, o.ProcessPath)
	}
}
