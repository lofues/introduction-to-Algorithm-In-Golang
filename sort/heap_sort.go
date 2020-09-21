package sort

func HeapSort(nums []int) {
	//	元素必须从1开始
	if len(nums) <= 1 {
		return
	}
	h := MakeBigHeap(nums)
	for i := Size(h) - 1; i > 1; i-- {
		h[i], h[1] = h[1], h[i]
		BigHeapify(h[:i], 1)
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = h[i+1]
	}
}

func MakeBigHeap(nums []int) []int {
	temp := []int{}
	temp = append(temp, -1)
	temp = append(temp, nums...)
	for i := len(temp) / 2; i >= 1; i-- {
		BigHeapify(temp, i)
	}
	return temp
}

func MakeSmallHeap(nums []int) []int {
	temp := []int{}
	temp = append(temp, -1)
	temp = append(temp, nums...)
	for i := len(temp) / 2; i >= 1; i-- {
		SmallHeapify(temp, i)
	}
	return temp
}

func BigHeapify(nums []int, i int) {
	var largest = i
	if Left(i) < Size(nums) && nums[Left(i)] > nums[i] {
		largest = Left(i)
	}
	if Right(i) < Size(nums) && nums[Right(i)] > nums[largest] {
		largest = Right(i)
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		BigHeapify(nums, largest)
	}
}

func SmallHeapify(nums []int, i int) {
	var smallest = i
	if Left(i) < Size(nums) && nums[Left(i)] < nums[i] {
		smallest = Left(i)
	}
	if Right(i) < Size(nums) && nums[Right(i)] < nums[smallest] {
		smallest = Right(i)
	}
	if smallest != i {
		nums[smallest], nums[i] = nums[i], nums[smallest]
		SmallHeapify(nums, smallest)
	}
}

func Left(i int) int {
	return 2 * i
}

func Right(i int) int {
	return 2*i + 1
}

func Size(nums []int) int {
	return len(nums)
}
