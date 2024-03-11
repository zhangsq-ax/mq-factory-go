package mq_factory

import (
	"fmt"
	"github.com/zhangsq-ax/mq-factory-go/constants"
	aws_msk_consumer "github.com/zhangsq-ax/mq-factory-go/consumer/aws-msk-consumer"
	rocket_mq_consumer "github.com/zhangsq-ax/mq-factory-go/consumer/rocket-mq-consumer"
	"github.com/zhangsq-ax/mq-factory-go/interfaces"
	"github.com/zhangsq-ax/mq-factory-go/options"
	aws_msk_producer "github.com/zhangsq-ax/mq-factory-go/producer/aws-msk-producer"
	rocket_mq_producer "github.com/zhangsq-ax/mq-factory-go/producer/rocket-mq-producer"
)

func NewMQProducer(opts *options.ProducerOptions) (interfaces.IMQProducer, error) {
	switch opts.MQType {
	case constants.MQTypeRocketMQ:
		return rocket_mq_producer.NewRocketMqProducer(opts)
	case constants.MQTypeKafka:
		return aws_msk_producer.NewAwsMskProducer(opts)
	default:
		return nil, fmt.Errorf("unknown mq type: %s", opts.MQType)
	}
}

func NewMQConsumer(opts *options.ConsumerOptions) (interfaces.IMQConsumer, error) {
	switch opts.MQType {
	case constants.MQTypeRocketMQ:
		return rocket_mq_consumer.NewRocketMqConsumer(opts)
	case constants.MQTypeKafka:
		return aws_msk_consumer.NewAwsMskConsumer(opts)
	default:
		return nil, fmt.Errorf("unknown mq type: %s", opts.MQType)
	}
}
