func twoSum_2ForLoopsSolution(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
            if nums[i] + nums[j] == target && i != j {
                return []int{i,j};
            }
	    }
	}
    return []int{};
}

func twoSum_TwoPassSolution(nums []int, target int) []int {
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
    
    for i := 0; i < len(nums); i++ {
        complement := target - nums[i]
        
        index, containsComplement := m[complement]
        if containsComplement && index != i {
            return []int{i,index};
        }
	}
    
    return []int{};
}


func twoSum_OnePassSolution(nums []int, target int) []int {
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        complement := target - nums[i]
        index, containsComplement := m[complement]
        
        if containsComplement {
            return []int{i,index};
        }
        
		m[nums[i]] = i
	}
    return []int{};
}

