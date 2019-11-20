package runner

import "sync"

type Runner struct {
	sync.Mutex
}
