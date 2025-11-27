package main

import (
	stability_pattern "cloud-native-go/stability_patterns"
	"context"
	"time"
)

func showcaseChain() {
	stability_pattern.Breaker(stability_pattern.DebounceFirst(func(context context.Context) (string, error) {
		return "Placeholder for real API", nil
	}, time.Millisecond*2000), 5)
}

func main() {
	for {
		print("Hello World")
		time.Sleep(1 * time.Second)
	}
}
