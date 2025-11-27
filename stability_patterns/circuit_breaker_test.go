package stability_pattern

import (
	"context"
	"errors"
	// "fmt"
	// "testing"
	"time"
)

// func TestMain(t *testing.T) {
// 	fmt.Println("test started")
// 	b := Breaker(buzz, 3)
// 	for i := 0; i < 3; i++ {
// 		resp, err := b(context.Background())
// 		fmt.Println(resp)
// 		fmt.Println(err.Error())
// 		time.Sleep(5000 * time.Millisecond)
// 	}

// }

func buzz(ctx context.Context) (string, error) {
	time.Sleep(300 * time.Millisecond)
	return "", errors.New("failed")
}
