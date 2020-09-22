package sort

// 只能排序大于等于0的数组
// O(n)
func CountingSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	temp := make([]int, max+1)
	res := make([]int, len(nums))
	// temp临时数组计算小于等于当前数组下标的元素的个数
	for _, num := range nums {
		temp[num]++
	}
	// 计算前缀和
	for i := 1; i < len(temp); i++ {
		temp[i] += temp[i-1]
	}
	// 将数组元素按下标填回数组
	for j := len(nums) - 1; j >= 0; j-- {
		res[temp[nums[j]]-1] = nums[j]
		temp[nums[j]]--
	}
	return res
}
