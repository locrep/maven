package utils

import (
	"sync"

	. "github.com/onsi/ginkgo"
	"math"
	"time"
)

var BeforeAll = func(beforeAllFunc func()) {
	var once sync.Once

	BeforeEach(func() {
		once.Do(func() {
			beforeAllFunc()
		})
	})
}

var AfterAll = func(afterAllFunc func()) {
	var once sync.Once

	AfterEach(func() {
		once.Do(func() {
			afterAllFunc()
		})
	})
}

func WaitUntil(predicate func() bool, timeout time.Duration, interval time.Duration) bool {
	var totalWait time.Duration

	for predicate() == false && totalWait < timeout {
		time.Sleep(interval)
		totalWait = totalWait + interval
	}

	return totalWait < timeout
}

func IntContains(intSlice []int, searchInt int) bool {
	for _, value := range intSlice {
		if value == searchInt {
			return true
		}
	}
	return false
}

func StringContains(strSlice []string, searchStr string) bool {
	for _, value := range strSlice {
		if value == searchStr {
			return true
		}
	}
	return false
}

func IntMax(array []int) int {
	max := math.MaxInt32 * -1
	for _, v := range array {
		if max < v {
			max = v
		}
	}
	return max
}
