# Wildcard Matching

## Description

```
Implement wildcard pattern matching with support for '?' and '*'.
 
 
 '?' Matches any single character.
 '*' Matches any sequence of characters (including the empty sequence).
 
 The matching should cover the entire input string (not partial).
 
 The function prototype should be:
 bool isMatch(const char *s, const char *p)
 
 Some examples:
 isMatch("aa","a") ? false
 isMatch("aa","aa") ? true
 isMatch("aaa","aa") ? false
 isMatch("aa", "*") ? true
 isMatch("aa", "a*") ? true
 isMatch("ab", "?*") ? true
 isMatch("aab", "c*a*b") ? false
```

link: https://leetcode.com/problems/wildcard-matching/description/

difficulty: Hard

## Discussion

```
s 文本串

p 模式串

dp[i][j] 取文本s的前i个字符和模式p的前j个字符进行完整匹配的结果。
```