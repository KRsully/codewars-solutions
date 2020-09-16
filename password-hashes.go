package kata

import (
  "crypto/md5"
  "encoding/hex"
)
func PassHash(str string) string {
  md5Bytes :=md5.Sum([]byte(str))
  //Since md5.Sum returns a fixed size array and hex.EncodeToString doesn't like that, take a slice of the whole array
  return hex.EncodeToString(md5Bytes[:])

}
