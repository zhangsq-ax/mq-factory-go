package aws_msk_producer

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	helper "github.com/zhangsq-ax/aws-msk-helper-go"
	"github.com/zhangsq-ax/mq-factory-go/options"
)

type AwsMskProducer struct {
	topic     string
	partition *int
	writer    *kafka.Writer
}

func NewAwsMskProducer(opts *options.ProducerOptions) (*AwsMskProducer, error) {
	writer, err := helper.NewKafkaWriter(&helper.NewKafkaWriterOptions{
		Brokers:      opts.Endpoints,
		Topic:        opts.Topic,
		BatchSize:    opts.BatchSize,
		BatchTimeout: opts.BatchTimeout,
		Username:     opts.AccessKey,
		Password:     opts.AccessSecret,
	})

	if err != nil {
		return nil, err
	}

	return &AwsMskProducer{
		topic:     opts.Topic,
		partition: opts.Partition,
		writer:    writer,
	}, nil
}

func (p *AwsMskProducer) PublishWithTopic(ctx context.Context, topic string, msg []byte, properties map[string]interface{}) error {
	if topic == "" {
		topic = p.topic
	}
	if topic == "" {
		return fmt.Errorf("missing target topic")
	}
	key := []byte("")
	if properties != nil && properties["key"] != nil {
		if _, ok := properties["key"].([]byte); ok {
			key = properties["key"].([]byte)
		}
	}
	return helper.WriteMessage(ctx, p.writer, key, msg, topic, p.partition)
}

func (p *AwsMskProducer) Publish(ctx context.Context, msg []byte, properties map[string]interface{}) error {
	return p.PublishWithTopic(ctx, p.topic, msg, properties)
}
