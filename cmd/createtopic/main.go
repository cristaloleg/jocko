package main

import (
	"fmt"
	"net"

	"github.com/travisjeffery/jocko/protocol"
	"github.com/travisjeffery/jocko/testutil"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	brokerAddr = kingpin.Flag("brokeraddr", "Address for Broker to bind on").Default("0.0.0.0:9092").String()
	topic      = kingpin.Flag("topic", "Name of topic to create").String()
	partitions = kingpin.Flag("partitions", "Number of partitions").Int32()
)

func main() {
	kingpin.Parse()

	addr, err := net.ResolveTCPAddr("tcp", *brokerAddr)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}

	err = testutil.CreateTopic(conn, &protocol.CreateTopicRequest{
		Topic:             *topic,
		NumPartitions:     *partitions,
		ReplicationFactor: int16(1),
		ReplicaAssignment: nil,
		Configs:           nil,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("created topic successfully")
}
