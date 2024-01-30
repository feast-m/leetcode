package longestsubstring_test

import (
	"fmt"
	"testing"

	longestsubstring "leetcode/longest-substring"
)

func TestDo(t *testing.T) {
	x := longestsubstring.Do("pwwkew")
	fmt.Println(x)
}
