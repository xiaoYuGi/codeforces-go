// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[1,2,0,3,0]`, `1`, 
			`3`,
		},
		{
			`[3,4,5,2,1,7,3,4,7]`, `3`, 
			`3`,
		},
		{
			`[1,2,4,1,2,5,1,2,6]`, `3`, 
			`3`,
		},
		// TODO 测试入参最小的情况
		{
			`[4,11,31,2,16,12,0,27,11,26,8,23,2,2,22,5,9]`, `3`,
			`13`,
		},
		{
			`[0,0,1,0,0,0]`, `3`,
			`1`,
		},
		{
			`[6,96,94,56,73,131,20,249,18,5,21,62,63,4,233,192,102,205,227,27,11,56,239,78,8,32,65,72,74,214,135,205,160,224,25,28,79,103,103,100,225,40,124,31,149,18,45,0,192,136,45,176,214,103,177,80,172,182,152,180,180,171,243,191,143,27,40,127,114,96,45,16,191,129,91,190,35,101,10,91,178,59,172,93,108]`, `15`,
			`69`,
		},
		{
			`[231,167,89,85,224,180,45,58,23,108,157,95,108,64,206,109,147,28,194,17,4,46,74,96,237,109,114,122,161,76,181,251,9,82,44,15,242,7,23,109,210,109,181,12,14,226,61,49,8,74,19,152,4,137,243,27,187,200,168,145,188,203,98,193,253,133,164,198,132,119,148,146,94,43,181,123,212,83,157]`, `2`,
			`75`,
		},
		{
			`[952,588,303,363,458,200,827,530,502,29,767,27,1016,63,644,948,703,986,638,542,392,615,671,350,1017,570,565,1012,377,864,323,50,148,624,555,63,239,309,100,473,667,618,669,850,755,257,890,295,112,32,381,971,146,35,890,577,200,453,548,563,140,964,279,497,404,525,93,923,489,885,533,647,463,1011,61,315,957,559,871,132,19,1019,260,607,103,7,443,766,327,479,155,876,986,686,733,240,998,634,926,816,117,522,341,822,954,874,200,768,218,605,540,844,853,858,920,847,691,643,622,838,214,496,16,819,38,578,213,446,961,709,71,980,264,50,329,248,913,607,664,597,966,526,410,902,812,406,133,155,48,692,8,956,826,1021,881,791,958,940,106,540,25,13,5,613,859,959,248,682,952,95,474,319,193,557,693,140,932,127,96,78,805,202,870,252,150,594,496,817,660,223,198,712,865,977,256,527,542,464,919,179,882,917,236,817,764,543,264,585,156,40,666,579,756,34,199,573,863,226,727,113,335,415,306,369,1,458,824,675,344,434,390,121,600,958,196,996,836,1020,918,509,125,177,794,596,122,23,252,843,1,588,55,768,163,796,709,753,98,637,791,714,784,112,1015,67,335,29,62,862,454,1020,469,406,317,212,323,310,355,868,153,214,180,809,286,145,743,446,369,61,85,456,589,693,538,602,755,241,879,729,623,987,649,552,670,234,900,182,355,251,562,884,655,574,754,574,975,127,892,508,740,562,124,191,364,162,396,246,387,862,434,773,582,239,240,212,995,509,782,377,351,461,72,8,687,64,1018,490,676,484,278,528,324,450,437,153,951,104,1015,333,147,299,689,166,897,908,557,881,824,427,733,513,485,273,126,715,940,628,211,810,1014,347,720,136,236,582,576,814,75,982,421,219,23,751,346,720,3,538,301,624,740,379,558,906,342,829,348,357,335,426,144,735,241,918,480,544,726,995,755,763,118,988,758,140,529,294,532,992,627,841,750,835,495,105,710,357,678,577,106,260,732,988,572,160,514,83,747,340,576,87,198,959,548,292,76,946,138,145,809,39,83,58,67,855,952,430,114,319,58,78,27,667,819,645,682,328,488,925,523,966,417,161,793,797,503,82,6,180,219,220,308,490,664,218,7,887,203,271,253,666,257,21,29,960,714,887,57,242,402,983,103,723,776,362,364,1020,427,372,1020,765,273,530,27,832,934,927,606,59,842,804,749,847,126,882,954,262,250,415,535,427,141,197,991,491,239,480,938,880,988,229,290,427,724,457,66,896,1,773,471,426,1005,194,62,695,112,836,187,221,114,965,443,985,1003,649,64,100,489,893,431,37,25,643,902,881,77,852,264,578,698,374,212,290,906,626,271,216,474,877,732,630,968,812,400,461,212,61,763,531,349,247,927,319,936,178,125,761,1018,936,933,663,714,96,206,794,144]`, `28`,
			`591`,
		},
	}
	targetCaseNum := -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minChanges, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-231/problems/make-the-xor-of-all-segments-equal-to-zero/
