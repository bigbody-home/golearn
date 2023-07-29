package main

//横向扫描法，找到2个中的最长公共子串，再和后面的子串递归找到全部的最长子串

func main() {
	strs := []string{"ggower", "floller", "flppp"}
	println(longestCommonPrefix(strs))

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = pipeline(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func pipeline(s1 string, s2 string) string {
	length := min(s1, s2)
	index := 0
	for index < length && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}

func min(s1, s2 string) int {
	min := 0
	if len(s1) > len(s2) {
		min = len(s2)
	}
	min = len(s1)
	return min
}
