package problems

import (
	"bytes"
	"strings"
)

func AnSCaller() {
	//s := "abcabcbb"
	//s := "aab"
	s := "abcbbcad"
	lengthOfLongestSubstring(s)
	lengthOfLongestSubstringWithStringMap(s)
}

func numUniqueEmails(emails []string) int {
	m := map[string]bool{}
	for _, email := range emails {
		ld := strings.Split(email, "@")
		local := strings.Join(strings.Split(strings.Split(ld[0], "+")[0], "."), "")
		m[local+"@"+ld[1]] = true // assigning same key value for the map is simply discarded.
	}
	return len(m)
}

/*Edge cases - see Edge case func below to below one
Input: "2-5g-3-J"
2
Output: "25-G3J"
Expected: "2-5G-3J"*/
func licenseKeyFormatting(s string, k int) string {
	temp := strings.Join(strings.Split(s, "-"), "")
	for i, _ := range temp {
		if i-1 == k {
			temp = temp[:i-1] + "-" + temp[i-1:]
		}
	}
	return temp
}

func licenseKeyFormattingEdgeCases(s string, k int) string {
	s = strings.ToUpper(strings.Replace(s, "-", "", -1))
	if len(s) <= k {
		return s
	}
	f := len(s) % k
	var ans bytes.Buffer
	if f != 0 {
		ans.WriteString(s[:f] + "-")
	}

	for i := f + k; i < len(s); i += k {
		ans.WriteString(s[f:i])
		ans.WriteString("-")
		f = i
	}
	ans.WriteString(s[f:])
	return ans.String()

}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	//byte test
	var b byte
	b = 'a'
	println(b)

	res := -1 << 63
	l := 0
	m := make(map[byte]int)
	m[s[0]]++
	for i := 1; i < len(s); i++ {
		m[s[i]]++
		for m[s[i]] > 1 {
			m[s[l]]--
			if m[s[l]] <= 0 {
				delete(m, s[l])
			}
			l++
		}
		if len(m) > res {
			res = len(m)
		}
	}
	if len(m) > res {
		res = len(m)
	}
	return res
}

func lengthOfLongestSubstringWithStringMap(s string) int {
	if len(s) == 0 {
		return 0
	}
	//test map
	am := make(map[string]int)
	am["v"] = 1
	am["j"] = 1
	am["v"] = 2
	am["j"] = am["j"] + 1

	var res int
	l := 0 //sliding window 1, i below would be sliding window 2
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
		for m[string(s[i])] > 1 {
			m[string(s[l])]--
			if m[string(s[l])] == 0 {
				delete(m, string(s[l]))
			}
			l++ //moving sliding window1
		}
		if len(m) > res {
			res = len(m)
		}
	}
	if len(m) > res {
		res = len(m)
	}
	return res
}
