package main

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//You can return the answer in any order.

//Example:
//
//Input: nums = [2,7,11,15], target = 9
//Output: [0,1]
//Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

//func twoSum(nums []int, target int) []int {
//	if len(nums) >= 2 {
//		m := make(map[int]int)
//
//		for i := 0; i < len(nums); i++ {
//			m[i] = int(nums[i])
//		}
//
//		var indexArr []int
//
//		for key, value := range m {
//
//			if value+m[key+1] == target {
//				indexArr = append(indexArr, key, key+1)
//			}
//
//		}
//
//		return indexArr
//	} else {
//		a := []int{0}
//		return a
//	}
//
//}

//func twoSum(nums []int, target int) []int {
//	if len(nums) >= 2 {
//		m := make(map[int]int)
//
//		for i, num := range nums {
//			difference := target - num
//
//			if index, isFound := m[difference]; isFound {
//				return []int{index, i}
//			}
//			m[num] = i
//
//		}
//
//		return []int{}
//	} else {
//		a := []int{0}
//		return a
//	}
//
//}

func twoSum(nums []int, target int) []int {
	if len(nums) >= 2 {
		m := make(map[int]int)

		for i, num := range nums {
			difference := target - num

			if index, isFound := m[difference]; isFound {
				return []int{index, i}
			}
			m[num] = i

		}

		return []int{}
	} else {
		a := []int{0}
		return a
	}

}
