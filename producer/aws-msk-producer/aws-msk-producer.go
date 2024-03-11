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

func (p *AwsMskProducer) Publish(ctx context.Context, key []byte, msg []byte, topic string) error {
	if topic == "" {
		topic = p.topic
	}
	if topic == "" {
		return fmt.Errorf("missing target topic")
	}
	return helper.WriteMessage(ctx, p.writer, key, msg, topic, p.partition)
}
