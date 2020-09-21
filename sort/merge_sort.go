package sort

func MergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	m := (l + r) >> 1
	MergeSort(nums, l, m)
	MergeSort(nums, m+1, r)
	temp := []int{}
	i, j := l, m+1
	for i <= m && j <= r {
		if nums[i] <= nums[j] {
			temp = append(temp, nums[i])
			i++
		} else {
			temp = append(temp, nums[j])
			j++
		}
	}
	temp = append(temp, nums[i:m+1]...)
	temp = append(temp, nums[j:r+1]...)
	for i := l; i <= r; i++ {
		nums[i] = temp[i-l]
	}
}
