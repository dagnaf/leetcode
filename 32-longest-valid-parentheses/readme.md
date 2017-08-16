# Longest Valid Parentheses

## Description

```
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.


 For "(()", the longest valid parentheses substring is "()", which has length = 2.


 Another example is ")()())", where the longest valid parentheses substring is "()()", which has length = 4.
```

link: https://leetcode.com/problems/longest-valid-parentheses/description/

difficulty: Hard

## Discussion

```
先匹配每个括号对应的位置，然后遍历相邻的括号匹配。
可以不需要存储匹配：
  匹配用到的栈存放的是下标而不是字符'(', ')'
  在对括号匹配时，匹配到的括号将全部弹出栈，而未匹配到的则被留在栈中
  再次遍历栈的下标，这些下标是未匹配到的括号，分割了连续匹配的括号，所以可以直接从这些栈中剩余的值得到结果。
```