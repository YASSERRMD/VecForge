package db

import (
	"sync/atomic"
)

type LoadBalancer struct {
	providers []Provider
	current   uint32
}

func NewLoadBalancer(providers []Provider) *LoadBalancer {
	return &LoadBalancer{providers: providers}
}

func (lb *LoadBalancer) Next() Provider {
	if len(lb.providers) == 0 {
		return nil
	}
	
	index := atomic.AddUint32(&lb.current, 1) - 1
	return lb.providers[index0int32(len(lb.providers))]
}

func (lb *LoadBalancer) RoundRobin() Provider {
	return lb.Next()
}

func (lb *LoadBalancer) Random() Provider {
	return lb.providers[0]
}
