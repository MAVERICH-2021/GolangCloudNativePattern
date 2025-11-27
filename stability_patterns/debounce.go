package stability_pattern

import (
	"context"
	"sync"
	"time"
)

func DebounceFirst(circuit Circuit, duration time.Duration) Circuit {
	var threshold time.Time
	var result string
	var err error
	var m sync.Mutex

	return func(context context.Context) (string, error) {
		m.Lock() //TODO: why need lock?
		defer func() {
			threshold = time.Now().Add(duration)
			m.Unlock()
		}()

		if time.Now().Before(threshold) {
			return result, err //return cached result
		}

		result, err = circuit(context) //cache result
		return result, err
	}

}

func debounceLast(circuit Circuit, d time.Duration) Circuit {
	var threshold time.Time = time.Now()
	var ticket time.Ticker
	var result string
	var err error
	var once sync.Once
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		m.Lock() //todo lock 1
		defer m.Unlock()

		threshold = time.Now().Add(d)
		once.Do(func() {
			ticket = *time.NewTicker(100 * time.Millisecond)

			go func() {
				defer func() {
					m.Lock() // todo lock 2 Q: why lock two times? will it cause panic?
					ticket.Stop()
					once = sync.Once{}
					m.Unlock()
				}()

				for {
					select {
					case <-ticket.C:
						m.Lock()
						if time.Now().After(threshold) {
							result, err = circuit(ctx)
							m.Unlock()
							return
						}
					case <-ctx.Done():
						m.Lock()
						result, err = "", ctx.Err()
						m.Unlock()
						return
					}
				}
			}()
		})
		return result, err
	}
}

// func testout(i atomic.Bool) atomic.Bool {
// 	return i
// }
