package rocket_mq_consumer

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/zhangsq-ax/mq-factory-go/options"
	rmqHelper "github.com/zhangsq-ax/rocketmq-helper-go"
)

type RocketMqConsumer struct {
	topic    string
	consumer rocketmq.PushConsumer
}

func NewRocketMqConsumer(opts *options.ConsumerOptions) (*RocketMqConsumer, error) {
	helper := rmqHelper.NewRocketMQHelper(&rmqHelper.RocketMQHelperOptions{
		NameServers: opts.Endpoints,
		InstanceId:  opts.InstanceID,
		Namespace:   opts.Namespace,
		AccessKey:   opts.AccessKey,
		SecretKey:   opts.AccessSecret,
		ConsumerOptions: &rmqHelper.ConsumerOptions{
			Group:        opts.GroupID,
			ConsumeModel: consumer.Clustering,
			ConsumeFrom: &rmqHelper.ConsumeFromOptions{
				Where: consumer.ConsumeFromLastOffset,
			},
		},
	})

	csm, err := helper.CreateConsumer(nil)
	if err != nil {
		return nil, err
	}

	return &RocketMqConsumer{
		topic:    opts.Topic,
		consumer: csm,
	}, nil
}

func (c *RocketMqConsumer) Start(ctx context.Context, msgHandler func(msg []byte), properties map[string]interface{}) error {
	return rmqHelper.ConsumeByConsumer(c.consumer, &rmqHelper.ConsumeOptions{
		Topic: c.topic,
		Selector: consumer.MessageSelector{
			Type:       consumer.TAG,
			Expression: "*",
		},
	}, func(msg *primitive.MessageExt) consumer.ConsumeResult {
		go msgHandler(msg.Body)
		return consumer.ConsumeSuccess
	})
}
