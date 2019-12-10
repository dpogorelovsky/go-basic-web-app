package events

import (
	"log"
)

// FlushCache - handler wrapper
type FlushCache struct{}

// Handle -
func (f FlushCache) Handle(eventPayload interface{}) {
	log.Printf("handle flush cache. payload: %v", eventPayload)
}
