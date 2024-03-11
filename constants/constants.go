package constants

type MQType string

const (
	MQTypeRocketMQ MQType = "rocketmq"
	MQTypeKafka    MQType = "kafka"
	MQTypeAwsIot   MQType = "aws-iot"
	MQTypeMQTT     MQType = "mqtt"
)
