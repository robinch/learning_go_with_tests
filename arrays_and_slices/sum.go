package arrays_and_slices

func Sum(numbers []int) int {
  sum := 0
  
  for _, n := range numbers {
    sum += n
  }

  return sum
}

func SumAll(numbersToSum ...[]int) []int {
  var sums []int

  for _, numbers := range numbersToSum {
    sums = append(sums, Sum(numbers))
  }

  return sums
}

func SumOfAllTails(numbersToSum ...[]int) []int {
  var sums []int

  for _, numbers := range numbersToSum {
    if len(numbers) == 0 {
      sums = append(sums, 0)
    } else {
      tail := numbers[1:]
      sums = append(sums, Sum(tail))
    }
  }

  return sums
}
