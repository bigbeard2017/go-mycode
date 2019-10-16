package main

import (
	"fmt"
	"regexp"

	"go-mycode/HealthDetect/process"
	linuxmac "go-mycode/HealthDetect/system/linuxMac"
)

func test() {
	d := new(linuxmac.SysUsedInfo)
	px, _ := d.GetSystemUsedInfo()
	fmt.Printf("process count :%d,Cpu Used:%f,MemAll:%d KiB,MemFree:%d KiB\n", px.ProcessCount, px.CPUUsed, px.MemAll, px.MemFree)
}

func main() {
	//test()
	v := testIsDigit("aa df :de")
	println(v)
	d := testIsDigit("111.1")
	println(d)
}

func testIsDigit(val string) bool {
	pattern := " ^(\\-|\\+)?\\d+(\\.\\d+)?$" //反斜杠要转义
	result, ee := regexp.MatchString(pattern, val)
	fmt.Println(result)
	if nil != ee {
		fmt.Printf("%v\n", ee)
	}
	return result
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
