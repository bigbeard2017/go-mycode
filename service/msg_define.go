package service

const F_REG_CODE string = "100001"

const PROTOCOL_VERSION string = "V1.0.1"

/**
* 完整的消息
 */
type Msg struct {
	H MsgHead
	B interface{}
}

/**
* 发送的消息头部
 */
type MsgHead struct {
	F string
	V string
	K string
	R bool
	M string
	S uint64
	T int
	I int
}

/**
*注册消息体
 */
type MsgRegBody struct {
	ServiceCode string
}

/**
*磁盘空间消息
**/
type DiskSpaceInfo struct {
	Path        string
	TotoalSpace int
	FreeSpace   int
}

/**
*进程信息消息
**/
type ProcessInfo struct {
	PID         int
	ProcessName string
	ProcessPath string
	CPU         float64
	MEM         float64
	ThreadCount int
	HanderCount int
	NetWork     int
	PortInfo    string
}

/**
*服务器整体信息
 */
type MsgServerInfoBody struct {
	ServerIP      string
	ServerName    string
	CollectTime   string
	ProcessCount  int
	CPU           float64
	MEM           float64
	ThreadCount   int
	HanderCount   int
	NetWork       int
	DiskFreeSpace []DiskSpaceInfo
	ProcessInfo   []ProcessInfo
}
