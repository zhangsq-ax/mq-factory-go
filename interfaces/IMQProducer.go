package interfaces

import "context"

type IMQProducer interface {
	Publish(ctx context.Context, key []byte, msg []byte, topic string) error
}
