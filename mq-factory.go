package mq_factory

import (
	"fmt"
	"github.com/zhangsq-ax/mq-factory-go/constants"
	aws_iot_consumer "github.com/zhangsq-ax/mq-factory-go/consumer/aws-iot-consumer"
	aws_msk_consumer "github.com/zhangsq-ax/mq-factory-go/consumer/aws-msk-consumer"
	mqtt_consumer "github.com/zhangsq-ax/mq-factory-go/consumer/mqtt-consumer"
	rocket_mq_consumer "github.com/zhangsq-ax/mq-factory-go/consumer/rocket-mq-consumer"
	"github.com/zhangsq-ax/mq-factory-go/interfaces"
	"github.com/zhangsq-ax/mq-factory-go/options"
	aws_iot_producer "github.com/zhangsq-ax/mq-factory-go/producer/aws-iot-producer"
	aws_msk_producer "github.com/zhangsq-ax/mq-factory-go/producer/aws-msk-producer"
	mqtt_producer "github.com/zhangsq-ax/mq-factory-go/producer/mqtt-producer"
	rocket_mq_producer "github.com/zhangsq-ax/mq-factory-go/producer/rocket-mq-producer"
)

func NewMQProducer(opts *options.ProducerOptions) (interfaces.IMQProducer, error) {
	switch opts.MQType {
	case constants.MQTypeRocketMQ:
		return rocket_mq_producer.NewRocketMqProducer(opts)
	case constants.MQTypeKafka:
		return aws_msk_producer.NewAwsMskProducer(opts)
	case constants.MQTypeAwsIot:
		return aws_iot_producer.NewAwsIotProducer(opts)
	case constants.MQTypeMQTT:
		return mqtt_producer.NewMQTTProducer(opts)
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
	case constants.MQTypeAwsIot:
		return aws_iot_consumer.NewAwsIotConsumer(opts)
	case constants.MQTypeMQTT:
		return mqtt_consumer.NewMQTTConsumer(opts)
	default:
		return nil, fmt.Errorf("unknown mq type: %s", opts.MQType)
	}
}
