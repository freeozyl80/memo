package main

import "fmt"

func max(int1, int2 int) int {
	if int1 > int2 {
		return int1
	}
	return int2
}

func maxSubArray(nums []int) int {
	cIdx := 0
	span := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[cIdx] > span {
			//说明梯度上升了，i 可以增加
			span = nums[i] - nums[cIdx]
		} else {
			// 说明梯度下降了， 所有 i 没有增加的意义
			if nums[i] < nums[cIdx] {
				cIdx = i
			}
		}
	}
	return span
}

// 0 100 -1 2
func main() {
	input := []int{0, 102, -1, 66}
	fmt.Println(maxSubArray(input))
}
