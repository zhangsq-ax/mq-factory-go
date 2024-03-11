package options

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/zhangsq-ax/mq-factory-go/constants"
	"time"
)

type ProducerOptions struct {
	MQType                constants.MQType                    // mq type
	Endpoints             []string                            // name server list for rocketmq or brokers for kafka
	AccessKey             string                              // access key for rocketmq; username for kafka; username for MQTT
	AccessSecret          string                              // secret key for rocketmq; password for kafka; password for MQTT
	GroupID               string                              // consumer group id for rocketmq
	Topic                 string                              // topic
	Partition             *int                                // partition for kafka
	BatchSize             int                                 // batch size for kafka
	BatchTimeout          time.Duration                       // batch timeout for kafka
	InstanceID            string                              // instance id for rocketmq
	Namespace             string                              // namespace for rocketmq
	CustomAuthName        string                              // custom auth name for AWS IoT Core (MQTT)
	EncryptKey            string                              // encrypt key for AWS IoT Core (MQTT)
	ConnectionLostHandler func(client mqtt.Client, err error) // connection lost handler for MQTT
	ALPN                  []string                            // alpn for MQTT
	ClientID              string                              // client id for MQTT
}
