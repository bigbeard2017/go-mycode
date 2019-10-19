package service

import (
	tools "go-mycode/Tools"
	"sync"
)

var msgIndex uint64 = 0

func getMsgIndex() uint64 {
	var mutex sync.Mutex
	mutex.Lock()
	msgIndex++
	defer mutex.Unlock()
	return msgIndex
}

/**
* 创建注册用的json消息
**/
func CreateRegisterMsg(serviceCode string) string {

	msgH := createHead(F_REG_CODE, "")
	var msg Msg
	msg.H = msgH
	var msgb MsgRegBody
	msgb.ServiceCode = serviceCode

	msg.B = msgb
	json, err := tools.ConvertToJson(msg)
	if nil != err {
		return ""
	}
	return json
}

func createHead(code, token string) MsgHead {
	var m MsgHead
	m.F = code
	m.V = PROTOCOL_VERSION
	m.R = true
	m.K = token
	m.S = getMsgIndex() //strconv.Itoa(getMsgIndex())
	m.T = 1
	m.I = 1
	m.M = ""
	return m
}

/**
* 创建检测消息
**/
func CreateDetectMsg() string {

	json := ""
	return json
}
