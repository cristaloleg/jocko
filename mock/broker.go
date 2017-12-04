// Code generated by mocker; DO NOT EDIT
// github.com/travisjeffery/mocker
package mock

import (
	"context"
	"github.com/travisjeffery/jocko"
	"github.com/travisjeffery/jocko/protocol"
	"sync"
)

var (
	lockBrokerJoin     sync.RWMutex
	lockBrokerRun      sync.RWMutex
	lockBrokerShutdown sync.RWMutex
)

// Broker is a mock implementation of Broker.
//
//     func TestSomethingThatUsesBroker(t *testing.T) {
//
//         // make and configure a mocked Broker
//         mockedBroker := &Broker{
//             JoinFunc: func(addr ...string) protocol.Error {
// 	               panic("TODO: mock out the Join method")
//             },
//             RunFunc: func(in1 context.Context,in2 <-chan jocko.Request,in3 chan<- jocko.Response)  {
// 	               panic("TODO: mock out the Run method")
//             },
//             ShutdownFunc: func() error {
// 	               panic("TODO: mock out the Shutdown method")
//             },
//         }
//
//         // TODO: use mockedBroker in code that requires Broker
//         //       and then make assertions.
//
//     }
type Broker struct {
	// JoinFunc mocks the Join method.
	JoinFunc func(addr ...string) protocol.Error

	// RunFunc mocks the Run method.
	RunFunc func(in1 context.Context, in2 <-chan jocko.Request, in3 chan<- jocko.Response)

	// ShutdownFunc mocks the Shutdown method.
	ShutdownFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// Join holds details about calls to the Join method.
		Join []struct {
			// Addr is the addr argument value.
			Addr []string
		}
		// Run holds details about calls to the Run method.
		Run []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 <-chan jocko.Request
			// In3 is the in3 argument value.
			In3 chan<- jocko.Response
		}
		// Shutdown holds details about calls to the Shutdown method.
		Shutdown []struct {
		}
	}
}

// Reset resets the calls made to the mocked APIs.
func (mock *Broker) Reset() {
	lockBrokerJoin.Lock()
	mock.calls.Join = nil
	lockBrokerJoin.Unlock()
	lockBrokerRun.Lock()
	mock.calls.Run = nil
	lockBrokerRun.Unlock()
	lockBrokerShutdown.Lock()
	mock.calls.Shutdown = nil
	lockBrokerShutdown.Unlock()
}

// Join calls JoinFunc.
func (mock *Broker) Join(addr ...string) protocol.Error {
	if mock.JoinFunc == nil {
		panic("moq: Broker.JoinFunc is nil but Broker.Join was just called")
	}
	callInfo := struct {
		Addr []string
	}{
		Addr: addr,
	}
	lockBrokerJoin.Lock()
	mock.calls.Join = append(mock.calls.Join, callInfo)
	lockBrokerJoin.Unlock()
	return mock.JoinFunc(addr...)
}

// JoinCalled returns true if at least one call was made to Join.
func (mock *Broker) JoinCalled() bool {
	lockBrokerJoin.RLock()
	defer lockBrokerJoin.RUnlock()
	return len(mock.calls.Join) > 0
}

// JoinCalls gets all the calls that were made to Join.
// Check the length with:
//     len(mockedBroker.JoinCalls())
func (mock *Broker) JoinCalls() []struct {
	Addr []string
} {
	var calls []struct {
		Addr []string
	}
	lockBrokerJoin.RLock()
	calls = mock.calls.Join
	lockBrokerJoin.RUnlock()
	return calls
}

// Run calls RunFunc.
func (mock *Broker) Run(in1 context.Context, in2 <-chan jocko.Request, in3 chan<- jocko.Response) {
	if mock.RunFunc == nil {
		panic("moq: Broker.RunFunc is nil but Broker.Run was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 <-chan jocko.Request
		In3 chan<- jocko.Response
	}{
		In1: in1,
		In2: in2,
		In3: in3,
	}
	lockBrokerRun.Lock()
	mock.calls.Run = append(mock.calls.Run, callInfo)
	lockBrokerRun.Unlock()
	mock.RunFunc(in1, in2, in3)
}

// RunCalled returns true if at least one call was made to Run.
func (mock *Broker) RunCalled() bool {
	lockBrokerRun.RLock()
	defer lockBrokerRun.RUnlock()
	return len(mock.calls.Run) > 0
}

// RunCalls gets all the calls that were made to Run.
// Check the length with:
//     len(mockedBroker.RunCalls())
func (mock *Broker) RunCalls() []struct {
	In1 context.Context
	In2 <-chan jocko.Request
	In3 chan<- jocko.Response
} {
	var calls []struct {
		In1 context.Context
		In2 <-chan jocko.Request
		In3 chan<- jocko.Response
	}
	lockBrokerRun.RLock()
	calls = mock.calls.Run
	lockBrokerRun.RUnlock()
	return calls
}

// Shutdown calls ShutdownFunc.
func (mock *Broker) Shutdown() error {
	if mock.ShutdownFunc == nil {
		panic("moq: Broker.ShutdownFunc is nil but Broker.Shutdown was just called")
	}
	callInfo := struct {
	}{}
	lockBrokerShutdown.Lock()
	mock.calls.Shutdown = append(mock.calls.Shutdown, callInfo)
	lockBrokerShutdown.Unlock()
	return mock.ShutdownFunc()
}

// ShutdownCalled returns true if at least one call was made to Shutdown.
func (mock *Broker) ShutdownCalled() bool {
	lockBrokerShutdown.RLock()
	defer lockBrokerShutdown.RUnlock()
	return len(mock.calls.Shutdown) > 0
}

// ShutdownCalls gets all the calls that were made to Shutdown.
// Check the length with:
//     len(mockedBroker.ShutdownCalls())
func (mock *Broker) ShutdownCalls() []struct {
} {
	var calls []struct {
	}
	lockBrokerShutdown.RLock()
	calls = mock.calls.Shutdown
	lockBrokerShutdown.RUnlock()
	return calls
}