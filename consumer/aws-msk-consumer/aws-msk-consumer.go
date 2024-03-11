package aws_msk_consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
	helper "github.com/zhangsq-ax/aws-msk-helper-go"
	"github.com/zhangsq-ax/mq-factory-go/options"
)

type AwsMskConsumer struct {
	reader *kafka.Reader
}

func NewAwsMskConsumer(opts *options.ConsumerOptions) (*AwsMskConsumer, error) {
	reader, err := helper.NewKafkaReader(&helper.NewKafkaReaderOptions{
		Brokers:     opts.Endpoints,
		GroupTopics: opts.GroupTopics,
		Topic:       opts.Topic,
		GroupID:     opts.GroupID,
		Partition:   opts.Partition,
		Username:    opts.AccessKey,
		Password:    opts.AccessSecret,
	})
	if err != nil {
		return nil, err

	}

	return &AwsMskConsumer{
		reader: reader,
	}, nil
}

func (c *AwsMskConsumer) Start(ctx context.Context, msgHandler func(msg []byte)) error {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := c.reader.ReadMessage(ctx)
				if err != nil {
					continue
				}
				msgHandler(msg.Value)
			}
		}
	}()

	return nil
}
