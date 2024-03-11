package rocket_mq_producer

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/zhangsq-ax/mq-factory-go/options"
	rmqHelper "github.com/zhangsq-ax/rocketmq-helper-go"
)

type RocketMqProducer struct {
	topic    string
	producer rocketmq.Producer
}

func NewRocketMqProducer(opts *options.ProducerOptions) (*RocketMqProducer, error) {
	helper := rmqHelper.NewRocketMQHelper(&rmqHelper.RocketMQHelperOptions{
		NameServers: opts.Endpoints,
		InstanceId:  opts.InstanceID,
		Namespace:   opts.Namespace,
		AccessKey:   opts.AccessKey,
		SecretKey:   opts.AccessSecret,
		ProducerOptions: &rmqHelper.ProducerOptions{
			Group: opts.GroupID,
		},
	})

	producer, err := helper.CreateProducer(nil)
	if err != nil {
		return nil, err
	}

	return &RocketMqProducer{
		topic:    opts.Topic,
		producer: producer,
	}, nil
}

func (p *RocketMqProducer) PublishWithTopic(ctx context.Context, topic string, msg []byte, properties map[string]interface{}) error {
	if topic == "" {
		topic = p.topic
	}
	if topic == "" {
		return fmt.Errorf("missing target topic")
	}

	message := &primitive.Message{
		Topic: topic,
		Body:  msg,
	}
	if properties != nil {
		if _, ok := properties["keys"].([]string); ok {
			message.WithKeys(properties["keys"].([]string))
		}
		if _, ok := properties["tags"].(string); ok {
			message.WithTag(properties["tags"].(string))
		}
	}

	_, err := p.producer.SendSync(ctx, message)
	return err
}

func (p *RocketMqProducer) Publish(ctx context.Context, msg []byte, properties map[string]interface{}) error {
	return p.PublishWithTopic(ctx, p.topic, msg, properties)
}
