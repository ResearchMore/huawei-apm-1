package pool

import (
	"sync"
	"time"
)

const (
	defaultPool   = 1024
	startPool     = 128
	defauleTicker = time.Second * 60
)

var pool chan *Pool

// Pool
type Pool struct {
	Mutex    *sync.RWMutex
	PoolChan chan func()
}

// newPool return new Pool default Pool length is 1024
// start pool is 128
func newPool() (*Pool, error) {

	return nil, nil
}

// ClosePool when chan use the end use this func to close chan.
// and to much idle chan will close idle chan.
func ClosePool(p Pool) {
	close(p.PoolChan)
}

func init() {
	pool = make(chan *Pool, defaultPool)
}
