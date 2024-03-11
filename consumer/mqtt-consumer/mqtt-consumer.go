package mqtt_consumer

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/zhangsq-ax/mq-factory-go/options"
	mqtt_helper "github.com/zhangsq-ax/mqtt-helper-go"
)

type MQTTConsumer struct {
	topic  string
	client *mqtt.Client
}

func NewMQTTConsumer(opts *options.ConsumerOptions) (*MQTTConsumer, error) {
	client, err := mqtt_helper.NewMQTTClient(&mqtt_helper.NewMQTTClientOptions{
		Brokers:  opts.Endpoints,
		ClientID: opts.ClientID,
		Username: opts.AccessKey,
		Password: opts.AccessSecret,
	})
	if err != nil {
		return nil, err
	}

	return &MQTTConsumer{
		topic:  opts.Topic,
		client: client,
	}, nil
}

func (c *MQTTConsumer) Start(ctx context.Context, msgHandler func(msg []byte), properties map[string]interface{}) error {
	qos := byte(0)
	if properties != nil {
		if _, ok := properties["qos"].(byte); ok {
			qos = properties["qos"].(byte)
		}
	}

	go func(ctx context.Context) {
		<-ctx.Done()
		(*c.client).Unsubscribe(c.topic)
	}(ctx)

	token := (*c.client).Subscribe(c.topic, qos, func(client mqtt.Client, msg mqtt.Message) {
		go msgHandler(msg.Payload())
		msg.Ack()
	})
	token.Wait()
	return token.Error()
}
