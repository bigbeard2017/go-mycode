package process

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

/**
* 获取系统的进程信息
**/
type Process struct {
	ProcessName string
	ProcessPath string
	CPU         float64
	Memory      float64
	ThreadCount int
	Pid         int
}

/**
* 	功能:
*		获取系统的所有进程信息
**/
func (p *Process) GetAllProcess() ([]Process, error) {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	processes := make([]Process, 0)
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		}

		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		pid, err := strconv.Atoi(ft[1])
		if err != nil {
			continue
		}
		cpu, err := strconv.ParseFloat(ft[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		mem, err := strconv.ParseFloat(ft[5], 64)

		cmd := ft[10]

		processes = append(processes, Process{Pid: pid, CPU: cpu, Memory: mem, ProcessPath: cmd})
	}
	return processes, nil
}
