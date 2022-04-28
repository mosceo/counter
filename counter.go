package counter

import (
	"sync"
)

var mu sync.Mutex

var Counter int

func IncV2() {
	mu.Lock()
	Counter += 2
	mu.Unlock()
}

func Dec() {
	mu.Lock()
	Counter--
	mu.Unlock()
}

var Values []int

func AddValue(v int) {
	mu.Lock()
	Values = append(Values, v)
	mu.Unlock()
}
