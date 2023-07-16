package mq

import (
	"fmt"
	"github.com/wagslane/go-rabbitmq"
	"log"
)

type Conn struct {
	conn         *rabbitmq.Conn
	consumer     *rabbitmq.Consumer
	publisher    *rabbitmq.Publisher
	queueName    string
	routingKey   string
	exchangeName string
}

// InitRabbitMQ 初始化连接 启动消费者、初始化生产者
func InitRabbitMQ(url string, f rabbitmq.Handler, queue, routingKey, exchangeName string) *Conn {
	conn, err := rabbitmq.NewConn(url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 消费者，注册时已经启动了
	consumer, err := rabbitmq.NewConsumer(
		conn,
		f,
		queue,
		rabbitmq.WithConsumerOptionsRoutingKey(routingKey),
		rabbitmq.WithConsumerOptionsExchangeName(exchangeName), // exchange 名称
		rabbitmq.WithConsumerOptionsExchangeDeclare,            // 声明交换器
	)
	if err != nil {
		panic(err)
	}
	// 生产者
	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsExchangeName(exchangeName), // exchange 名称
		rabbitmq.WithPublisherOptionsExchangeDeclare,            // 声明交换器
	)
	if err != nil {
		panic(err)
	}

	defer publisher.Close()
	// 连接被拒绝
	publisher.NotifyReturn(func(r rabbitmq.Return) {
		//log.Printf("message returned from server: %s", string(r.Body))
	})

	// 提交确认
	publisher.NotifyPublish(func(c rabbitmq.Confirmation) {
		//log.Printf("message confirmed from server. tag: %v, ack: %v", c.DeliveryTag, c.Ack)
	})
	return &Conn{
		conn:         conn,
		consumer:     consumer,
		publisher:    publisher,
		queueName:    queue,
		routingKey:   routingKey,
		exchangeName: exchangeName,
	}
}

// Publish 发送消息，该消息实际由执行 InitRabbitMQ 注册时传入的 f 消费

func (c *Conn) Publish(data []byte) error {
	if data == nil || len(data) == 0 {
		fmt.Println("data 为空，publish 不发送")
		return nil
	}
	return c.publisher.Publish(
		data,
		[]string{c.routingKey},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsPersistentDelivery,       // 消息持久化
		rabbitmq.WithPublishOptionsExchange(c.exchangeName), // 要发送的 exchange
	)
}
