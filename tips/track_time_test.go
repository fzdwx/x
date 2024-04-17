package tips

import (
	"fmt"
	"testing"
	"time"
)

func TestTrackTimeFunc(t *testing.T) {
	_ = TrackTimeFunc(func() {
		fmt.Println("hello world")
		time.Sleep(1 * time.Second)
	})
}

func TestTrackTime(t *testing.T) {
	defer TrackTime(time.Now())

	time.Sleep(1 * time.Second)
}
