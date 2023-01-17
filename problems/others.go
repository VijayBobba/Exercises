package problems

import "strings"

func numUniqueEmails(emails []string) int {
	m := map[string]bool{}
	for _, email := range emails {
		ld := strings.Split(email, "@")
		local := strings.Join(strings.Split(strings.Split(ld[0], "+")[0], "."), "")
		m[local+"@"+ld[1]] = true // assigning same key value for the map is simply discarded.
	}
	return len(m)
}

/*Input: "2-5g-3-J"
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
