package conn

import (
	"fmt"
	"github.com/cebilon123/ElytraGo/packet"
	"os"
	"strconv"
)

var maxWorkers = os.Getenv("MaxWorkers")

// Worker is a struct responsible for handling packet jobs.
// Amount of workers can be set up easily through config
type Worker struct {
	id   int
	quit chan bool
}

func NewWorker(quit chan bool, id int) *Worker {
	return &Worker{quit: quit, id: id}
}

// Start makes worker starts listening for given channels
func (w *Worker) Start(clientPackets chan packet.IPacket, serverPackets chan packet.IPacket) {
	for {
		select {
		case cp := <-clientPackets:
			fmt.Printf("Worker(%v): Client-> PID: %v, Type: %v, Payload: %#x, String->: %s\n", w.id, cp.GetPid(), cp.GetType(), cp.GetPayload(), string(cp.GetPayload()))
		case sp := <-serverPackets:
			fmt.Printf("Worker(%v): Client-> PID: %v, Type: %v, Payload: %#x, String->: %s\n", w.id, sp.GetPid(), sp.GetType(), sp.GetPayload(), string(sp.GetPayload()))
		case <-w.quit:
			return
		}
	}
}

type WorkerDispatcher struct {
	ClientPackets chan packet.IPacket
	ServerPackets chan packet.IPacket
	quit          chan bool
}

func (wd *WorkerDispatcher) Close() error {
	wd.quit <- true
	return nil
}

func NewWorkerDispatcher(clientPackets chan packet.IPacket, serverPackets chan packet.IPacket) *WorkerDispatcher {
	return &WorkerDispatcher{ClientPackets: clientPackets, ServerPackets: serverPackets}
}

func (wd *WorkerDispatcher) SpawnWorkers() {
	maxW, err := strconv.Atoi(maxWorkers)
	//If somehow there is error while parsing max workers, just set it to default value
	if err != nil {
		maxW = 5
	}

	for i := 0; i < maxW; i++ {
		w := NewWorker(wd.quit, i)
		go w.Start(wd.ClientPackets, wd.ServerPackets)
	}
}
