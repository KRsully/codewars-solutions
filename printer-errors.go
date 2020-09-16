package kata

import (
  "fmt"
  "strings"
  "sort"
)

func PrinterError(s string) string {
   letters := strings.Split(s, "")
  sort.Strings(letters)
  s = strings.Join(letters,"")
  numerator := 0
  denominator := len(s)
  if index := strings.IndexAny(s, "nopqrstuvwxyz");index != -1 {
    numerator = len(s)-index
  }
  return fmt.Sprintf("%d/%d", numerator, denominator)
}
