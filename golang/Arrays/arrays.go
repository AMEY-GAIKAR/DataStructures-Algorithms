package arrays

import "math"

func maxInt(a, b int) int {
  if a > b {
    return a
  } else {
    return b
  }
}

func minInt(a, b int) int {
  if a < b {
    return a
  } else {
    return b
  }
}

func LinearSearch(nums []int, key int) int {
  for i, v := range nums {
    if v == key {
      return i
    }
  }
  return -1
}

func SearchRange(nums []int, target int) []int {
  answer := []int{-1, -1}
  for i:=0; i < len(nums); i++ {
    if nums[i] == target {
      answer[0] = i
      break
    }
  }
  for i:=len(nums)-1; i > -1; i-- {
    if nums[i] == target {
      answer[1] = i
      break
    }
  } 
  return answer
}

func CreateFrequencyTable(nums []int) map[int]int {
  frequencyTable := make(map[int]int)
  for _, v := range nums {
    if _, exist := frequencyTable[v]; exist {
      frequencyTable[v] ++
    } else {
      frequencyTable[v] = 1
    } 
  }
  return frequencyTable
}

func FindHighestFrequency(nums []int) int {
  frequencyTable := make(map[int]int) 
  var maxFrequency int = 0
  for _, v := range nums {
    if _, exist := frequencyTable[v]; exist {
      frequencyTable[v] ++
      maxFrequency = maxInt(frequencyTable[v], maxFrequency)
    } else {
      frequencyTable[v] = 1
    }
  }
  return maxFrequency
}

func FindLowestFrequency(nums []int) int {
  frequencyTable := make(map[int]int)
  var minFrequency int = math.MaxInt
  for _, v := range nums {
    if _, exist := frequencyTable[v]; exist {
      frequencyTable[v] ++
    } else {
      frequencyTable[v] = 1
    }
  }
  for _, v := range frequencyTable {
      minFrequency = minInt(minFrequency, v) 
    }
  return minFrequency
}

func FindLargestElement(nums []int) int {
  var largest int = nums[0]
  for _, v := range nums {
    if v > largest {
      largest = v
    }
  }
  return largest
}

func FindSmallestELement(nums []int) int {
  var smallest = nums[0]
  for _, v := range nums {
    if v < smallest {
      smallest = v
    }
  }
  return smallest
}

func FindSecondLargest(nums []int) int {
  var largest int = nums[0]
  var secondLargest int = nums[0]
  for i := range nums {
    if nums[i] > largest {
      secondLargest = largest
      largest = nums[i]
    } else if nums[i] > secondLargest && nums[i] < largest {
      secondLargest = nums[i]
    }
  }
  return secondLargest
}

func FindSingleI(nums []int) int {
  frequencyTable := make(map[int]int)
  for _, v := range nums {
    if _, exist := frequencyTable[v]; exist {
      frequencyTable[v] ++
    } else {
      frequencyTable[v] = 1
    }
  }
  for key, val := range frequencyTable {
    if val == 1 {
      return key
    }
  }
  return -1
}

func FindSingleII(nums []int) int {
  var xorSum int = 0
  for _, val := range nums {
    xorSum = xorSum ^ val
  }
  return xorSum
}

func CheckSorted(nums []int) bool {
  if len(nums) == 1  || len(nums) == 0 {
    return true
  }
  for i := 1; i <= len(nums) - 1; i++ {
    if nums[i] < nums[i-1] {
      return false
    }
  }
  return true
}

func FindEvenDigits(nums []int) int {
  var digitCount int = 0
  var answer int = 0
  for i := range(nums) {
    var digit int = nums[i]
    digitCount = 0
    for digit > 0 {
      digit = digit / 10
      digitCount += 1
    }
    if digitCount % 2  == 0{
      answer += 1
    }
  }
  return answer
}

func MaxConsecutiveOnes(nums []int) int {
  var localCount int = 0
  var globalCount int = 0
  for i := range(nums) {
    if nums[i] == 1 {
      localCount += 1
    } else {
      localCount = 0
    }
    globalCount = maxInt(localCount, globalCount)
  }
  return globalCount
}

func FindMissing(nums []int, n int) int {
  var sum int = 0
  for i := range nums {
    sum += nums[i]
  }
  return n*(n+1)/2 - sum
}

func TwoSumI(nums []int, target int) [2]int {
  for i := range nums {
    for j := range nums {
      if nums[j] + nums[i] == target {
        return [2]int{i, j}
      }
    }
  }
  return [2]int{-1, -1}
}

func TwoSumII(nums []int, target int) [2]int {
  search := make(map[int]int)
  for i, v := range nums {
    if idx, exist := search[target -v]; exist {
      return [2]int{idx, i}
    } else {
      search[v] = i
    }
  }
  return [2]int{-1, -1}
}

func LowerBound(nums []int, target int) int {
  var start int = 0
  var end int = len(nums) - 1
  var answer int = len(nums)
  for start <= end {
    var mid int = start + (end - start)/2
    if nums[mid] >= target {
      answer = mid
      end = mid - 1
    } else {
      start = mid + 1
    }
  }
  return answer
}

func RemoveDuplicates(nums []int) int {
  var i int = 0
  for j := 1; j < len(nums); j++ {
    if nums[i] != nums[j] {
      i ++
      nums[i] = nums[j]
    }
  }
  return i+1
}

func RemoveZeros(nums []int) []int {
  var i int = 0
  for j := 0; j < len(nums); j ++ {
    if nums[j] != 0 {
      nums[i], nums[j] = nums[j], nums[i]
      i ++
    }
  }
  return nums
}

func ContainsDuplicate(nums []int) bool {
  search := make(map[int]bool)
  for _, v := range nums {
    if _, exist := search[v]; exist {
      return true
    } else {
      search[v] = true
    }
  }
  return false
}

func LeftRotateK(nums []int, k int) []int {
  k = k % len(nums)
  if k != 0 {
    temp := make([]int, len(nums))
    temp = append(nums[k:], nums[:k]...)  
    copy(nums, temp) 
  }
  return nums
}

func RightRotateK(nums []int, k int) []int {
  k = k % len(nums)
  if k != 0 {
    temp := make([]int, len(nums))
    temp = append(nums[len(nums)-k:], nums[:len(nums)-k]...)  
    copy(nums, temp) 
  }
  return nums
}

func SortColors(nums []int)  {
  var left int = 0
  var mid int = 0
  var right int = len(nums) - 1
  for mid <= right {
    if nums[mid] == 0 {
      nums[mid], nums[left] = nums[left], nums[mid]
      left++
      mid++
    } else if nums[mid] == 2 {
      nums[right], nums[mid] = nums[mid], nums[right]
      right--
    } else {
      mid++
    }
  }
}

func CheckRotatedSorted(nums []int) bool {
  var count int = 0
  for i := 1; i < len(nums); i++ {
    if nums[i] < nums[i-1] {
       count++
    }
  }
  if count == 0 {
    return true
  }
  if count == 1 && nums[0] >= nums[len(nums)-1] {
    return true
  }
  return false
}

func MajorityElement(nums []int) int {
  var count int = 0
  var element int
  for _, v := range nums {
    if count == 0 {
      element = v
      count++
    } else if v == element {
      count++ 
    } else {
      count--
    }
  }
  return element
}

func MaxSubArraySumI(nums []int) int {
  var sum int = 0
  var max int = math.MinInt
  for _, v := range nums {
    sum = maxInt(sum+v, v)
    max = maxInt(sum, max)
  }
  return max
}

func MaxSubArraySumII(nums []int) []int {
  var sum int = 0
  var max int = math.MinInt
  var start int = 0
  var ansEnd, ansStart int = -1, -1
  for i := range nums {
    if sum == 0 {
      start = i
    }
    sum += nums[i]
    if sum > max {
      max = sum
      ansStart = start
      ansEnd = i
    }
    if sum < 0 {
      sum = 0
    }
  }
  return nums[ansStart:ansEnd+1]
}

func MaxProfit(prices []int) int {
  var min int = math.MaxInt
  var diff int = -1
  for _, v := range prices {
    min = minInt(min, v)
    diff = maxInt(diff, v-min)
  } 
  return diff
}

func RearrangeBySign(nums []int) []int {
  var even int = 0
  var odd int = 1
  new := make([]int, len(nums))
  for i := range nums {
    if nums[i] >= 0 {
      new[even] = nums[i]
      even += 2
    } else {
      new[odd] = nums[i]
      odd += 2
    }
  }
  return new
}

func Merge(nums1, nums2 []int, m, n int) []int {
  var i int = m - 1
  var j int = n - 1
  var k int = m + n - 1
  for j >= 0 && i >= 0 {
    if nums1[i] > nums2[j] {
      nums1[k] = nums1[i]
      i--
    } else {
      nums1[k] = nums2[j]
      j--
    }
    k--
  }
  for j >= 0 {
    nums1[k] = nums2[j]
    j--
    k--
  }
  return nums1
}

func LongestSubArrayWithSumI(nums []int, k int) int {
  var answer int = 0
  for i := range nums {
    var sum int = 0
    for j := i; j < len(nums); j++ {
      sum += nums[j]
      if sum == k {
        answer = maxInt(answer, j - i + 1)
      }
    }
  }
  return answer
}

func LongestSubArrayWithSumII(nums []int, k int) int {
  prefixSum := make(map[int]int)
  var sum int = 0
  var answer int = 0
  prefixSum[0] = 1
  for i := range nums {
    sum += nums[i]
    var remaining int = sum - k
    if _, ok := prefixSum[remaining]; ok {
      answer += prefixSum[remaining]
    }
    prefixSum[sum]++
  }
  return answer
}

func LongestSubArrayWithSumPositives(nums []int, k int) int {
  var i int = 0
  var answer int = 0
  var sum int = 0
  for j := range nums {
    sum += nums[j]
    if sum == k {
      answer = maxInt(answer, j - i + 1)
      i++
      sum -= nums[i]
    }
    for sum > k && i <= j {
      sum -= nums[i]
      i++
    }
  }
  return answer
}

func RotateMatrix(nums [][]int) [][]int {
  for i := range nums {
    for j := i; j < len(nums[0]); j++ {
      nums[i][j], nums[j][i] = nums[j][i], nums[i][j]
    }
  }
  for i := 0; i < len(nums); i++ {
    j, k := 0, len(nums[0]) - 1
    for j <= k {
      nums[i][j], nums[i][k] = nums[i][k], nums[i][j]
      j++
      k--
    }
  }
  return nums
}

func SetZeroes(matrix [][]int) [][]int {
  rows := make(map[int]bool)
  columns := make(map[int]bool)
  for i := 0; i < len(matrix); i++ {
    for j := 0; j < len(matrix[0]); j++ {
      if matrix[i][j] == 0 {
        rows[i] = true
        columns[j] = true
      }
    }
  }
  for i := range matrix {
    for j := range matrix[0] {
      if rows[i] || columns[j] {
      matrix[i][j] = 0
      }
    }
  }
  return matrix
}

func SpiralMatix(matrix [][]int) []int {

} 
