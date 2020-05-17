package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	m := make(map[int]int,len(nums))
	for i,k := range nums {
		if value,ok := m[target-k];ok {
			result = append(result,value)
			result = append(result,i)
		}
		m[k] = i
	}
	return result
}



func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums,target))
}