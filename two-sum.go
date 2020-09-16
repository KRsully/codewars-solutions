package kata

import "sort"

func TwoSum(numbers []int, target int) [2]int {
    for i := range(numbers) {
      j := sort.Search(len(numbers), func(k int) bool {return numbers[k] == target-numbers[i]})
      if  j<len(numbers) && numbers[j] == target-numbers[i] && i!=j{
        indices := []int{i,j}
        sort.Ints(indices)
        var ret [2]int
        copy (ret[:], indices)
        return ret
      }
    }
    return [2]int{}
}
