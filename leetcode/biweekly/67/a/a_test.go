// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`[2,1,3,3]`, `2`, 
			`[3,3]`,
		},
		{
			`[-1,-2,3,4]`, `3`, 
			`[-1,3,4]`,
		},
		{
			`[3,4,3,3]`, `2`, 
			`[3,4]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxSubsequence, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-67/problems/find-subsequence-of-length-k-with-the-largest-sum/
