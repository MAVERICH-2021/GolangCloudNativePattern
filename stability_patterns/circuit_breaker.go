package stability_pattern

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//use closure to keep state for all func invokation
//If have to keep state for multi instance -> use redis | other midware health check and service discovery
func Breaker(circuit Circuit, failureThresh uint) Circuit {
	failureCount := 0
	lastAttempt := time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {
		m.RLock() // narrow to read lock
		d := failureCount - int(failureThresh)

		if d >= 0 {
			possibleRetryAt := lastAttempt.Add(200 * time.Second << d) //That is huge amount if d is big TODO: improve backoff algrithem
			if time.Now().Before(possibleRetryAt) {
				m.RUnlock()
				return "", fmt.Errorf("Service is currenly unreachable. Recovering... next try at %v", possibleRetryAt)
			}
		}

		m.RUnlock()

		resp, err := circuit(ctx)

		m.Lock() //avoid race condition
		defer m.Unlock()
		lastAttempt = time.Now()
		if err != nil {
			failureCount++
			return resp, err
		} else {
			failureCount = 0
			return resp, err
		}

	}
}
