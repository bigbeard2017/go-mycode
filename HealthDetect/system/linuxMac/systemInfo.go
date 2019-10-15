package system

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

type MemStatus struct {
	All  uint32 `json:"all"`
	Used uint32 `json:"used"`
	Free uint32 `json:"free"`
	Self uint64 `json:"self"`
}

func MemStat() MemStatus {
	//自身占用
	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)
	mem := MemStatus{}
	mem.Self = memStat.Alloc

	// //系统占用,仅linux/mac下有效
	// //system memory usage
	// sysInfo := new(syscall.Sysinfo_t)
	// err := syscall.Sysinfo(sysInfo)
	// if err == nil {
	// 	mem.All = sysInfo.Totalram * uint32(syscall.Getpagesize())
	// 	mem.Free = sysInfo.Freeram * uint32(syscall.Getpagesize())
	// 	mem.Used = mem.All - mem.Free
	// }
	return mem
}

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

/**
* 系统信息
 */
type SysUsedInfo struct {
	CPUUsed      float64
	MemUsed      float64
	MemAll       float64
	NetDown      float32
	NetUp        float32
	ProcessCount int
}

func (p *SysUsedInfo) GetSystemUsedInfo() (*SysUsedInfo, error) {
	print("start")
	cmd := exec.Command("top", "-n0")
	print("start22")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	print("start333")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Printf("%v\n", out)
	var r *SysUsedInfo

	r.CPUUsed = 0.0
	lineAllProcess := readline(out)

	processTokens := strings.Split(lineAllProcess, " ")
	count, err := strconv.Atoi(processTokens[1])
	if nil != err {
		r.ProcessCount = count
	}
	cpuline := readline(out)
	cpuTokens := strings.Split(cpuline, " ")
	cpuFree := strings.Replace(cpuTokens[11], "%", "", -1)
	cpuf, _ := strconv.ParseFloat(cpuFree, 32)
	r.CPUUsed = 100.0 - cpuf

	return r, nil
}

func readline(buff bytes.Buffer) string {
	line, err := buff.ReadString('\n')
	if nil != err {
		return ""
	}
	return line
}
