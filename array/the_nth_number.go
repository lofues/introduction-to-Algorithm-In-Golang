package array

import (
	"math/rand"
	"time"
)

// 从nums中找出第n小的数
func TheNthNumber(nums []int, l, r, n int) int {
	if l == r {
		return nums[l]
	}
	m := randomPartition(nums, l, r)
	leftNum := m - l + 1
	if n == leftNum {
		return nums[m]
	} else if n < leftNum {
		return TheNthNumber(nums, l, m-1, n)
	} else {
		return TheNthNumber(nums, m+1, r, n-leftNum)
	}
}

func randomPartition(nums []int, l, r int) int {
	rand.Seed(time.Now().UnixNano())
	m := rand.Intn(r-l+1) + l
	i, j, x := l, r, nums[m]
	for i < j {
		for nums[j] > x && i < j {
			j--
		}
		for nums[i] <= x && i < j {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}
