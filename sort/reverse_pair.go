package sort

func ReversePair(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	return _ReversePair(nums, 0, len(nums)-1)
}

func _ReversePair(nums []int, l, r int) int {
	if l >= r {
		return 0
	}
	m := (l + r) >> 1
	left := _ReversePair(nums, l, m)
	right := _ReversePair(nums, m+1, r)
	i, j := l, m+1
	temp := []int{}
	count := 0
	for i <= m && j <= r {
		if nums[i] <= nums[j] {
			temp = append(temp, nums[i])
			i++
		} else {
			temp = append(temp, nums[j])
			count += m - l + 1
			j++
		}
	}
	return count + left + right
}
