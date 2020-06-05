package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func commonPrefix(str1, str2 string) string {
	i := 0
	for ; i < len(str1) && i < len(str2); i++ {
		if str1[i] != str2[i] {
			break
		}
	}
	return str1[:i]
}

func longestCommonPrefix(strArr []string) string {
	if len(strArr) == 0 {
		return ""
	}
	prefix := strArr[0]
	for i := 1; i < len(strArr); i++ {
		prefix = commonPrefix(prefix, strArr[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func main() {
	strArr := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strArr))
}
