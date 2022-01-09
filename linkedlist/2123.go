package linkedlist

func pairSum(head *ListNode) int {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	max := 0
	for i := 0; i < len(nums)/2; i++ {
		if max < nums[i]+nums[len(nums)-i-1] {
			max = nums[i] + nums[len(nums)-i-1]
		}
	}
	return max
}
