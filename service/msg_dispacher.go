package service

import (
	"encoding/json"
)

// DispatherMsg 解析消息并完成消息调度
func DispatherMsg(jsonstr string) error {

	var tmp map[string]interface{}
	err := json.Unmarshal([]byte(jsonstr), &tmp)
	if nil != err {
		return err
	}
	h := tmp["H"]
	f := h.(map[string]interface{})
	fcode := f["F"]
	if fcode == FRegCode {

	}else if fcode ==
	return nil
}
