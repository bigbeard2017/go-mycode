package main

import (
	"go-mycode/HealthDetect/process"
)

/**
说明:
	第一步启动一个线程,定时采集,时间间隔从配置文件中获取.

**/
func main() {

}

func readConfig() {

}

func startDetect(f func(json string)) {

	var p process.Process
	result, err := p.GetAllProcess()
	if nil != err {

	} else {
		for _, r := range result {

		}
	}
}
