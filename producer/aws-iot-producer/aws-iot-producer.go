package aws_iot_producer

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/zhangsq-ax/mq-factory-go/options"
	mqtt_helper "github.com/zhangsq-ax/mqtt-helper-go"
)

type AwsIotProducer struct {
	topic  string
	client *mqtt.Client
}

func NewAwsIotProducer(opts *options.ProducerOptions) (*AwsIotProducer, error) {
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

	return &AwsIotProducer{
		topic:  opts.Topic,
		client: client,
	}, nil
}

func (p *AwsIotProducer) PublishWithTopic(ctx context.Context, topic string, msg []byte, properties map[string]interface{}) error {
	qos := byte(0)
	retained := false
	if properties != nil {
		if _, ok := properties["qos"].(byte); ok {
			qos = properties["qos"].(byte)
		}
		if _, ok := properties["retained"].(bool); ok {
			retained = properties["retained"].(bool)
		}
	}
	token := (*p.client).Publish(topic, qos, retained, msg)
	token.Wait()
	return token.Error()
}

func (p *AwsIotProducer) Publish(ctx context.Context, msg []byte, properties map[string]interface{}) error {
	return p.PublishWithTopic(ctx, p.topic, msg, properties)
}
