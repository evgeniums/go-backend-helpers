package pubsub_redis

import (
	"context"
	"fmt"
	"sync"

	"github.com/evgeniums/go-backend-helpers/pkg/app_context"
	"github.com/evgeniums/go-backend-helpers/pkg/config"
	"github.com/evgeniums/go-backend-helpers/pkg/config/object_config"
	"github.com/evgeniums/go-backend-helpers/pkg/logger"
	"github.com/evgeniums/go-backend-helpers/pkg/message"
	"github.com/evgeniums/go-backend-helpers/pkg/pubsub"
	"github.com/evgeniums/go-backend-helpers/pkg/validator"
	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Address  string `default:"localhost:6379" validate:"required" vmessage:"Address of Redis server not set"`
	Db       int
	Password string `mask:"true"`
}

type RedisClient struct {
	RedisConfig
	redisClient *redis.Client
	context     context.Context
}

func (r *RedisClient) Config() interface{} {
	return &r.RedisConfig
}

func (r *RedisClient) Init(cfg config.Config, log logger.Logger, vld validator.Validator, configPath ...string) error {

	err := object_config.LoadLogValidate(cfg, log, vld, r, "redis_pubsub", configPath...)
	if err != nil {
		return log.PushFatalStack("failed to init Redis client", err)
	}

	r.context = context.Background()
	r.redisClient = redis.NewClient(&redis.Options{
		Addr:     r.Address,
		Password: r.Password,
		DB:       r.Db,
	})
	err = r.redisClient.Ping(r.context).Err()
	if err != nil {
		return log.PushFatalStack("failed to connect to Redis server", err)
	}

	return nil
}

func (p *RedisClient) Shutdown() {
	p.redisClient.Close()
}

//---------------------------------------

type Publisher struct {
	RedisClient
	pubsub.PublisherBase
}

func NewPublisher(serializer ...message.Serializer) *Publisher {
	p := &Publisher{}
	p.Construct(serializer...)
	return p
}

func (p *Publisher) Publish(topicName string, obj interface{}) error {

	payload, err := p.Serialize(obj)
	if err != nil {
		return err
	}

	return p.redisClient.Publish(p.context, topicName, payload).Err()
}

//---------------------------------------

type Subscriber struct {
	RedisClient
	pubsub.SubscriberBase

	mutex    sync.RWMutex
	channels map[string]*redis.PubSub
}

func NewSubscriber(app app_context.Context, serializer ...message.Serializer) *Subscriber {
	s := &Subscriber{}
	s.Construct(app, serializer...)

	return s
}

func (s *Subscriber) Subscribe(topic pubsub.Topic) error {

	err := s.AddTopic(topic)
	if err != nil {
		return err
	}

	channel := s.redisClient.Subscribe(s.context, topic.Name())
	_, err = channel.Receive(s.context)
	if err != nil {
		return fmt.Errorf("failed to receive from redis pubsub channel: %s", err)
	}

	s.mutex.Lock()
	s.channels[topic.Name()] = channel
	s.mutex.Unlock()

	ch := channel.Channel()
	readMessages := func() {
		for msg := range ch {
			opCtx := s.NewOpContext(topic.Name())
			s.Handle(opCtx, topic.Name(), []byte(msg.Payload))
			opCtx.Close()
		}
	}
	go readMessages()

	return nil
}

func (s *Subscriber) Unsubscribe(topicName string) {

	s.mutex.Lock()
	channel, ok := s.channels[topicName]
	s.mutex.Unlock()
	if !ok {
		return
	}

	channel.Unsubscribe(s.context)

	s.mutex.Lock()
	delete(s.channels, topicName)
	s.mutex.Unlock()

	s.DeleteTopic(topicName)
}