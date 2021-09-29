// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`5`, `[[2,5,4],[1,5,1]]`, 
			`7`,
		},
		{
			`20`, `[[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]`, 
			`20`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxTaxiEarnings, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-61/problems/maximum-earnings-from-taxi/