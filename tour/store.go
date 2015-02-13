package tour

import (
	"github.com/xebia/microgen/events"
)

type Store interface {
	GetTourOnYear(year int) (*Tour, bool)
	StoreEvent(envelope *events.Envelope)
}
