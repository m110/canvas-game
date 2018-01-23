package broker

import (
	"reflect"

	"github.com/m110/canvas-game/pkg/game/events"
)

type SubscribeRequest struct {
	EventType reflect.Type
	Channel   chan interface{}
}

type DisconnectRequest struct {
	Channel chan interface{}
}

type Broker struct {
	subscribers map[reflect.Type][]chan interface{}

	subscribe  chan SubscribeRequest
	disconnect chan DisconnectRequest
	publish    chan events.Event
}

func NewBroker() *Broker {
	broker := &Broker{
		subscribers: make(map[reflect.Type][]chan interface{}),

		subscribe:  make(chan SubscribeRequest),
		disconnect: make(chan DisconnectRequest),
		publish:    make(chan events.Event),
	}

	go broker.start()

	return broker
}

func (b *Broker) Subscribe(event events.Event) chan interface{} {
	channel := make(chan interface{})
	b.subscribe <- SubscribeRequest{
		EventType: reflect.TypeOf(event),
		Channel:   channel,
	}
	return channel
}

func (b *Broker) Disconnect(channel chan interface{}) {
	b.disconnect <- DisconnectRequest{
		Channel: channel,
	}
}

func (b *Broker) Publish(event events.Event) {
	b.publish <- event
}

func (b *Broker) start() {
	for {
		select {
		case request := <-b.subscribe:
			b.subscribeClient(request)
		case request := <-b.disconnect:
			b.disconnectClient(request)
		case event := <-b.publish:
			b.publishEvent(event)
		}
	}
}

func (b *Broker) subscribeClient(request SubscribeRequest) {
	b.subscribers[request.EventType] = append(b.subscribers[request.EventType], request.Channel)
}

func (b *Broker) disconnectClient(request DisconnectRequest) {
	for event, subscribers := range b.subscribers {
		var channels []chan interface{}
		for _, channel := range subscribers {
			if channel == request.Channel {
				close(channel)
			} else {
				channels = append(channels, channel)
			}
		}

		b.subscribers[event] = channels
	}
}

func (b *Broker) publishEvent(event events.Event) {
	for _, subscriber := range b.subscribers[reflect.TypeOf(event)] {
		subscriber <- event
	}
}
