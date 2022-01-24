// https://github.com/moby/moby
package pubsubservice

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/moby/moby/pkg/pubsub"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	pub *pubsub.Publisher
}

var once sync.Once
var obj *Server

type PubSubMessage struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Payload              string   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubSubMessage) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PubSubMessage) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type PublishRequest struct {
	Pubsubmessage        *PubSubMessage `protobuf:"bytes,1,opt,name=pubsubmessage,proto3" json:"pubsubmessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

type PublishResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRequest) GetPubsubmessage() *PubSubMessage {
	if m != nil {
		return m.Pubsubmessage
	}
	return nil
}

type SubscribeRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func GetPubsubService() *Server {
	once.Do(func() {
		obj = new(Server)
		obj.pub = pubsub.NewPublisher(1000*time.Millisecond, 10)
	})
	return obj
}

func (s *Server) Publish(ctx context.Context, req PublishRequest) (*PublishResponse, error) {
	msg := req.GetPubsubmessage()
	s.pub.Publish(msg)
	fmt.Println("Publish Topic:", msg.GetTopic())
	return &PublishResponse{
		Result: true,
	}, nil
}

func (s *Server) Subscribe(req *SubscribeRequest) (chan interface{}, error) {
	ch := s.pub.SubscribeTopic(func(v interface{}) bool {
		data := v.(*PubSubMessage)
		return strings.HasPrefix(data.GetTopic(), req.GetTopic())
	})

	return ch, nil
}

func (s *Server) Ping() {
	go func() {
		time.Sleep(500 * time.Microsecond)
		// 發送測試 ping pone
		time.Sleep(1 * time.Second)
		ctx := context.Background()
		msg := &PubSubMessage{
			Topic:   "ping:",
			Payload: "pong",
		}
		res, err := s.Publish(ctx, PublishRequest{Pubsubmessage: msg})
		if err != nil {
			log.Error("publish err:", err.Error())
		}
		fmt.Println("ping: ", res.Result)
	}()
}

func (s *Server) Pong() {
	go func() {
		pong := make(chan interface{})
		ch, err := s.Subscribe(&SubscribeRequest{
			Topic: "ping:",
		})
		if err != nil {
			log.Error("[ERROR]Subscribe pign:", err.Error())
		}
		pong = ch

		for v := range pong {
			topic := v.(*PubSubMessage).GetTopic()
			payload := v.(*PubSubMessage).GetPayload()
			fmt.Printf("Receive Topic: %s, Receive Payload: %s \n", topic, payload)
		}
	}()
}
