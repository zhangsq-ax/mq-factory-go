package options

import (
	"github.com/zhangsq-ax/mq-factory-go/constants"
	"time"
)

type ProducerOptions struct {
	MQType       constants.MQType // mq type
	Endpoints    []string         // name server list for rocketmq or brokers for kafka
	AccessKey    string           // access key for rocketmq or username for kafka
	AccessSecret string           // secret key for rocketmq or password for kafka
	GroupID      string           // consumer group id for rocketmq
	Topic        string           // topic
	Partition    *int             // partition for kafka
	BatchSize    int              // batch size for kafka
	BatchTimeout time.Duration    // batch timeout for kafka
	InstanceID   string           // instance id for rocketmq
	Namespace    string           // namespace for rocketmq
}
