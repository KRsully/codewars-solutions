package kata

func Solution(str string) []string {
  //Add an underscore if the length is odd
  if len(str)%2==1 {
    str+="_"
  }
  
  //Because the number of pairs is not known before runtime, use a Slice
  pairs := make([]string, 0)
  
  for i:=0; i<len(str); i+=2{
    pairs = append(pairs,string(str[i])+string(str[i+1]))
  }
  
  return pairs
  
}
