// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1971/problem/E
// https://codeforces.com/problemset/status/1971/problem/E
func Test_cf1971E(t *testing.T) {
	testCases := [][2]string{
		{
			`4
10 1 3
10
10
0
6
7
10 2 4
4 10
4 7
6
4
2
7
1000000000 1 1
1000000000
1000000000
99999999
6 1 3
6
5
2
6
5`,
			`0 6 7 
5 4 2 5 
99999999 
1 5 4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1971E)
}
