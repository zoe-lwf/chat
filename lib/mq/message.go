package mq

import (
	"chat/model"
	"chat/pkg/mq"
	"fmt"
	"github.com/wagslane/go-rabbitmq"
)

const (
	MessageQueue        = "message.queue"
	MessageRoutingKey   = "message.routing.key"
	MessageExchangeName = "message.exchange.name"
)

var (
	MessageMQ *mq.Conn
)

func InitMessageMQ(url string) {
	mq.InitRabbitMQ(url, MessageCreateHandler, MessageQueue, MessageRoutingKey, MessageExchangeName)
	fmt.Println("rabbit mq is ok!")
}

func MessageCreateHandler(d rabbitmq.Delivery) rabbitmq.Action {
	messageModels := model.ProtoMarshalToMessage(d.Body)
	if messageModels == nil {
		fmt.Println("空的")
		return rabbitmq.NackDiscard
	}
	err := model.CreateMessage(messageModels...)
	if err != nil {
		fmt.Println("[MessageCreateHandler] model.CreateMessage 失败，err:", err)
		return rabbitmq.NackDiscard
	}
	//fmt.Println("处理完消息：", string(d.Body))
	return rabbitmq.Ack
}
