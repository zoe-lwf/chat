package ws

import (
	"chat/common"
	"chat/protocol/pb"
	"fmt"
	"google.golang.org/protobuf/proto"
)

// GetOutputMsg 组装出下行消息

func GetOutputMsg(cmdType pb.CmdType, code int32, message proto.Message) ([]byte, error) {
	output := &pb.Output{
		Type:    cmdType,
		Code:    code,
		CodeMsg: common.GetErrorMessage(uint32(code), ""),
		Data:    nil,
	}
	if message != nil {
		msgBytes, err := proto.Marshal(message)
		if err != nil {
			fmt.Println("[GetOutputMsg] message marshal err:", err)
			return nil, err
		}
		output.Data = msgBytes
	}

	bytes, err := proto.Marshal(output)
	if err != nil {
		fmt.Println("[GetOutputMsg] output marshal err:", err)
		return nil, err
	}
	return bytes, nil

}
