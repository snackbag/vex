package vex

import "time"

func DoEvery(ms int, runnable func()) {
	DoEveryWithDelayAndIterations(ms, 0, -1, runnable)
}

func DoEveryWithDelay(ms int, delay int, runnable func()) {
	DoEveryWithDelayAndIterations(ms, delay, -1, runnable)
}

func DoEveryWithDelayAndIterations(ms int, delay int, iterations int, runnable func()) {
	go func() {
		if iterations == 0 {
			return
		}

		last := time.Now()
		timesDone := 0
		for last.UnixMilli()+int64(delay) > time.Now().UnixMilli() {
		}

		runnable()
		timesDone++
		last = time.Now()

		for iterations < 0 || timesDone < iterations {
			if last.UnixMilli()+int64(ms) < time.Now().UnixMilli() {
				runnable()
				timesDone++
				last = time.Now()
			}
		}
	}()
}
