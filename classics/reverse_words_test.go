package classics

import (
	"testing"
)

//ReverseWords 反转单词
func TestReverseWords(t *testing.T) {
	str := " I am a student!"
	reverseStr := ReverseWords(str)
	t.Log(reverseStr)
}
