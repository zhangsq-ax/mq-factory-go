package options

import "github.com/zhangsq-ax/mq-factory-go/constants"

type ConsumerOptions struct {
	MQType       constants.MQType // mq type
	Endpoints    []string         // name server list for rocketmq or brokers for kafka
	GroupID      string           // consumer group id
	Topic        string           // topic
	GroupTopics  []string         // group topics for kafka
	AccessKey    string           // access key for rocketmq or username for kafka
	AccessSecret string           // secret key for rocketmq or password for kafka
	Partition    int              // partition for kafka
	InstanceID   string           // instance id for rocketmq
	Namespace    string           // namespace for rocketmq
}
