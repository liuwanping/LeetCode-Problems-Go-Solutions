package main

import (
	"context"
	"fmt"
	"time"
)

type A struct {
	M bool
	N int
	P bool
}

func main() {
	//res := longestPalindrome("aacabdkacaa")
	//nums1 := []int{1, 2}
	//nums2 := []int{3, 4, 5}
	//res := findMedianSortedArrays(nums1, nums2)
	//fmt.Println(res)
	startTime, endTime, _ :=  parseInputTimeString(context.TODO(), "20211201", "20211201")
	fmt.Println(startTime, endTime)

}

func parseInputTimeString(ctx context.Context, startDay, endDay string) (time.Time, time.Time, error) {
	var (
		startTime time.Time
		endTime   time.Time
	)
	for _, v := range map[string][]interface{}{
		"startDay": {startDay, &startTime},
		"endDay":   {endDay, &endTime},
	} {
		if dayTime, err := time.ParseInLocation("20060102", v[0].(string), time.UTC); err != nil {
			return time.Time{}, time.Time{}, nil
		} else {
			*(v[1].(*time.Time)) = dayTime
		}
	}
	return startTime, endTime, nil
}
