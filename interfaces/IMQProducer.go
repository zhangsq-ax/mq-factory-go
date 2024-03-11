package interfaces

import "context"

type IMQProducer interface {
	Publish(ctx context.Context, msg []byte, properties map[string]interface{}) error
	PublishWithTopic(ctx context.Context, topic string, msg []byte, properties map[string]interface{}) error
}
