package jocko

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/hashicorp/raft"
	"github.com/travisjeffery/jocko/jocko/metadata"
)

type brokerLookup struct {
	lock            sync.RWMutex
	addressToBroker map[raft.ServerAddress]*metadata.Broker
	idToBroker      map[raft.ServerID]*metadata.Broker
}

func NewBrokerLookup() *brokerLookup {
	return &brokerLookup{
		addressToBroker: make(map[raft.ServerAddress]*metadata.Broker),
		idToBroker:      make(map[raft.ServerID]*metadata.Broker),
	}
}

func (b *brokerLookup) AddBroker(broker *metadata.Broker) {
	b.lock.Lock()
	b.addressToBroker[raft.ServerAddress(broker.RaftAddr)] = broker
	b.idToBroker[raft.ServerID(broker.ID.Int32())] = broker
	b.lock.Unlock()
}

func (b *brokerLookup) BrokerByAddr(addr raft.ServerAddress) *metadata.Broker {
	b.lock.RLock()
	svr, _ := b.addressToBroker[addr]
	b.lock.RUnlock()
	return svr
}

func (b *brokerLookup) BrokerByID(id raft.ServerID) *metadata.Broker {
	b.lock.RLock()
	svr, _ := b.idToBroker[id]
	b.lock.RUnlock()
	return svr
}

func (b *brokerLookup) BrokerAddr(id raft.ServerID) (raft.ServerAddress, error) {
	b.lock.RLock()
	svr, ok := b.idToBroker[id]
	b.lock.RUnlock()
	if !ok {
		return "", fmt.Errorf("no broker for id %v", id)
	}
	return raft.ServerAddress(svr.RaftAddr), nil
}

func (b *brokerLookup) RemoveBroker(broker *metadata.Broker) {
	b.lock.Lock()
	delete(b.addressToBroker, raft.ServerAddress(broker.RaftAddr))
	delete(b.idToBroker, raft.ServerID(broker.ID.Int32()))
	b.lock.Unlock()
}

func (b *brokerLookup) Brokers() []*metadata.Broker {
	b.lock.RLock()
	ret := make([]*metadata.Broker, 0, len(b.addressToBroker))
	for _, svr := range b.addressToBroker {
		ret = append(ret, svr)
	}
	b.lock.RUnlock()
	return ret
}

func (b *brokerLookup) RandomBroker() *metadata.Broker {
	brokers := b.Brokers()
	i := rand.Intn(len(brokers))
	return brokers[i]
}
