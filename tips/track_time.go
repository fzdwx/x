package tips

import (
	"fmt"
	"time"
)

// TrackTime is a function to track time
func TrackTime(pre time.Time) time.Duration {
	elapsed := time.Since(pre)
	fmt.Println("elapsed time: ", elapsed)
	return elapsed
}

func TrackTimeFuncE(f func() error) (time.Duration, error) {
	pre := time.Now()
	err := f()
	elapsed := TrackTime(pre)
	return elapsed, err
}

func TrackTimeFunc(f func()) time.Duration {
	elapsed, _ := TrackTimeFuncE(func() error {
		f()
		return nil
	})
	return elapsed
}
