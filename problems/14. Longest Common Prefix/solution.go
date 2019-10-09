// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".

// Example 1:
// Input: ["flower","flow","flight"]
// Output: "fl"

// Example 2:
// Input: ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

package solution

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	if l == 1 {
		return strs[0]
	}

	i := 0
	for ; i < len(strs[0]); i++ {
		j := 1
		for ; j < l; j++ {
			if i >= len(strs[j]) {
				break
			}
			if strs[j][i] != strs[0][i] {
				break
			}
		}
		if j != l {
			break
		}
	}
	return strs[0][0:i]
}
