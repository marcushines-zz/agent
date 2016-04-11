// Package generator creates schedules of notifications to publish to clients.
package generator

import (
	ocpb "github.com/openconfig/reference/rpc/openconfig"
)

type Event struct {
	n *ocpb.Notification
	e error
}

// Generator emits Events to be processed by the agent.
type Generator interface {
	Start() <-chan Event
}

type RandomGenerator struct {
	c   *apb.Config
	out <-chan Event
}

// Start starts the generator
func (r *RandomGenerator) Start() <-chan Event {

}

// NewRandomGenerator returns a new initialized RandomGenerator ready to be started
// based on config.  If config is nil or invalid nil will be returned.
func NewRandomGenerator(config *apb.Config, outq <-chan Event) *RandomGenerator {
	if config == nil {
		return nil
	}
	r := &RandomGenerator{
		c:    config,
		outq: outq,
	}
	return r
}
