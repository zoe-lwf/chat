package protoService

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestProtobuf(t *testing.T) {

	usermsg := &Userinfo{
		Name:  "123",
		Age:   int32(28),
		Hobby: []byte("qqq"),
	}
	//返回byte[],但这是符合protobuf格式的byte[]
	userData, err := proto.Marshal(usermsg)
	if err != nil {
		fmt.Println("Marshaling err:", err)
	}
	//[10 3 49 50 51 16 28 26 3 113 113 113]
	fmt.Println(userData)

	userInfo := new(Userinfo)
	//我们直接传入一个byte[]，比如：[]byte("qqq")，这个方法会报错，格式不对
	err = proto.Unmarshal(userData, userInfo)
	err = proto.Unmarshal(userData, userInfo)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userInfo)
}
