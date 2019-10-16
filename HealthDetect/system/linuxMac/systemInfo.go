package system

import (
	"bytes"
	"fmt"
	tools "go-mycode/Tools"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

/**
* 获取硬盘使用情况.
 */
type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

/**
* disk usage of path/disk.
 */
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return disk
}

/**
* 获取系统内存,CPU和进程数量信息.
 */
type SysUsedInfo struct {
	CPUUsed      float64
	MemFree      uint64
	MemAll       uint64
	NetDown      float32
	NetUp        float32
	ProcessCount int
}

/**
*获取系统的进程数量,CPU使用率,内存大小和空闲内存.
 */
func (p *SysUsedInfo) GetSystemUsedInfo() (*SysUsedInfo, error) {

	cmd := exec.Command("top", "-bn1")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		fmt.Printf(" Execute command  happen an  error :%v\n", err)
		return nil, err
	}

	var lines [5]string
	for i := 0; i < len(lines); i++ {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		} else {
			lines[i] = line
		}
	}

	var r SysUsedInfo

	processTokens := strings.Split(lines[1], " ")
	count, err := strconv.Atoi(processTokens[1])
	if nil == err {
		r.ProcessCount = count
	} else {
		fmt.Printf("Convert process count  happend an  error :%v\n", err)
		fmt.Println(lines[1])
	}

	cpuTokens := strings.Split(lines[2], " ")
	var cpuInfo [5]string
	i := 0
	for index := 0; index < len(cpuTokens); {
		r := tools.IsNumeric(cpuTokens[index])
		if r {
			cpuInfo[i] = cpuTokens[index]
			i++
		}
	}
	cpuFree := cpuInfo[3]
	cpuf, err := strconv.ParseFloat(cpuFree, 64)
	if nil == err {
		r.CPUUsed = 100.0 - cpuf
	} else {
		fmt.Printf(" Convert CPU used precentage happend an  error :%v\n", err)
		fmt.Println(lines[2])
	}

	memTokens := strings.Split(lines[3], " ")

	memall, err := strconv.ParseUint(memTokens[3], 10, 64)
	if nil == err {
		r.MemAll = memall
	} else {
		fmt.Printf(" Convert Mem All   (use index 3) happend an  error :%v\n", err)
		fmt.Println(lines[3])
	}
	memfree, err := strconv.ParseUint(memTokens[7], 10, 64)
	if nil == err {
		r.MemFree = memfree
	} else {
		fmt.Printf(" Convert Mem free (use index 7) happend an  error :%v\n", err)
		fmt.Println(lines[3])
	}
	memcatch, err := strconv.ParseUint(memTokens[13], 10, 64)
	if nil == err {
		r.MemFree = r.MemFree + memcatch
	} else {
		fmt.Printf(" Convert Mem catch (use index 13) happend an  error :%v\n", err)
		fmt.Println(lines[3])
	}

	return &r, nil
}
