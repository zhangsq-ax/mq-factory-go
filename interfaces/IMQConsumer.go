package interfaces

import "context"

type IMQConsumer interface {
	Start(ctx context.Context, msgHandler func(msg []byte), properties map[string]interface{}) error
}
