func minSteps(s string, t string) int {
  count := 0
  m1 := make(map[rune]int)
  m2 := make(map[rune]int)
  for _, char := range s{
      m1[char]+=1
  }
  for _, char := range t{
      m2[char]+=1
  }
  for _, char := range s{
      if m2[char] > 0 {
          m2[char]-=1
      } else if m2[char] == 0{
          count+=1
      }
  }
  return count

}
