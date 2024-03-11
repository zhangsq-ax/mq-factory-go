package aws_iot_consumer

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/zhangsq-ax/mq-factory-go/options"
	mqtt_helper "github.com/zhangsq-ax/mqtt-helper-go"
)

type AwsIotConsumer struct {
	topic  string
	client *mqtt.Client
}

func NewAwsIotConsumer(opts *options.ConsumerOptions) (*AwsIotConsumer, error) {
	client, err := mqtt_helper.NewMQTTClient(&mqtt_helper.NewMQTTClientOptions{
		UsernameTransformer: mqtt_helper.NewAwsIotUsernameTransformer(opts.CustomAuthName),
		PasswordEncryptor:   mqtt_helper.NewHmacSha1Encryptor(opts.EncryptKey),
		Brokers:             opts.Endpoints,
		ClientID:            opts.ClientID,
		Username:            opts.AccessKey,
		Password:            opts.AccessSecret,
	})

	if err != nil {
		return nil, err
	}

	return &AwsIotConsumer{
		topic:  opts.Topic,
		client: client,
	}, nil
}

func (c *AwsIotConsumer) Start(ctx context.Context, msgHandler func(msg []byte), properties map[string]interface{}) error {
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
