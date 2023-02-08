package kafka

import "github.com/opensourceways/kafka-lib/mq"

var (
	DefaultMQ = NewMQ()
)

func Init(opts ...mq.Option) error {
	return DefaultMQ.Init(opts...)
}

func Connect() error {
	return DefaultMQ.Connect()
}

func Disconnect() error {
	return DefaultMQ.Disconnect()
}

func Publish(topic string, msg *mq.Message, opts ...mq.PublishOption) error {
	return DefaultMQ.Publish(topic, msg, opts...)
}

func Subscribe(topic, group string, handler mq.Handler) (mq.Subscriber, error) {
	return DefaultMQ.Subscribe(topic, group, handler)
}

func String() string {
	return DefaultMQ.String()
}
