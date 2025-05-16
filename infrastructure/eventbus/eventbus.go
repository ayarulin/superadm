package eventbus

type EventBus struct{}

func (b *EventBus) Publish(event any) error {
	return nil
}
