// Generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_a(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, maximumLengthSubstring, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-390/problems/maximum-length-substring-with-two-occurrences/
// https://leetcode.cn/problems/maximum-length-substring-with-two-occurrences/