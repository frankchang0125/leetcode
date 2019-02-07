/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (32.85%)
 * Total Accepted:    396.7K
 * Total Submissions: 1.2M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 * 
 * If there is no common prefix, return an empty string "".
 * 
 * Example 1:
 * 
 * 
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 * 
 * 
 * Note:
 * 
 * All given inputs are in lowercase letters a-z.
 * 
 */
func longestCommonPrefix(strs []string) string {
    prefix := ""
    
    if len(strs) == 0 {
        return prefix
    }
    
    i := 0
    
    for {
        currentChar := ""
        for _, str := range strs {
            if i > len(str) - 1 {
                return prefix
            }
            
            if currentChar == "" {
                currentChar = string(str[i])
            } else if string(str[i]) != currentChar {
                return prefix
            }
        }
        
        prefix = fmt.Sprintf("%s%s", prefix, currentChar)
        
        i++
    }
    
    return prefix
}
