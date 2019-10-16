package main

import (
	"fmt"
	"time"

	"go-mycode/HealthDetect/process"
	linuxmac "go-mycode/HealthDetect/system/linuxMac"
)

func test() {
	d := new(linuxmac.SysUsedInfo)
	px, _ := d.GetSystemUsedInfo()
	fmt.Printf("process count :%d,Cpu Used:%f,MemAll:%d KiB,MemFree:%d KiB\n", px.ProcessCount, px.CPUUsed, px.MemAll, px.MemFree)
}

func main() {
	var flg bool
	flg = true
	for flg {
		time.Sleep(10)
		fmt.Printf("%v\n", time.Now.)
	}
}

func testProcess() {
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
