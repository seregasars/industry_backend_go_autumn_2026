package main

import(
	"math"
)

type Stats struct {
	Count         int
	Sum, Min, Max int64
}


func Calc(nums []int64) Stats {
	if len(nums) < 2{
		return Stats{}
	}

	s := Stats{}
	s.Count = len(nums) - 1
	s.Sum = 0
	s.Min = math.MaxInt64
	s.Max = math.MinInt64


	for i := 0; i < len(nums) - 1;i++{
		diff := nums[i+1] - nums[i]
		s.Sum += diff
		if diff < s.Min {
			s.Min = diff
		}
		if diff > s.Max {
			s.Max = diff
		}

	}
	return s
}
