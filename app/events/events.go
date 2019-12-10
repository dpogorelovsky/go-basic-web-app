package events

import "github.com/dpogorelovsky/pubsub"

const (
	DBModelChanged = "DB_MODEL_CHANGED"
)

type EventEngine interface {
}

// SetupEvents - setting up possible events, initilization of whole event system
func SetupEvents() pubsub.Emitter {
	list := []string{DBModelChanged}

	emitter := pubsub.NewEmitter(list)
	flushCache := FlushCache{}
	flushCacheSub := pubsub.NewSubscriber(flushCache)
	emitter.AddSubscriber(DBModelChanged, flushCacheSub)

	return emitter
}
