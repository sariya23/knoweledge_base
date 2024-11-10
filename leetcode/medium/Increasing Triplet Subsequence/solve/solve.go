package solve

func IncreasingTriplet(nums []int) bool {
	return false
}

func FindTwoSmallestElements(nums []int) []int {
	i := 0
	j := nums[len(nums)-1]
	res := make([]int, 0, 2)
	for i < j {
		if nums[i] < nums[j] {
			res = append(res, nums[i], nums[j])
			break
		} else if nums[i] > nums[j] {
		}
	}
	return nil
}

func main() {
	res := FindTwoSmallestElemets([]int{1, 2, 3})
}
