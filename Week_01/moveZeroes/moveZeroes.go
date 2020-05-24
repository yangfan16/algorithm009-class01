package moveZeroes

func moveZeroes(nums []int) {
	zero := 0
	for index := range nums {
		if nums[index] != 0 {
			nums[index], nums[zero] = nums[zero], nums[index]
		}
		zero++
	}
}
