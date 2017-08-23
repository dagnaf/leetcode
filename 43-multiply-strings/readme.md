# Multiply Strings

## Description

```
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2.

 Note:

 The length of both num1 and num2 is < 110.
 Both num1 and num2 contains only digits 0-9.
 Both num1 and num2 does not contain any leading zero.
 You must not use any built-in BigInteger library or convert the inputs to integer directly.

```

link: https://leetcode.com/problems/multiply-strings/description/

difficulty: Medium

## Discussion

```
模拟乘法，对应每一位的输出需要移位。
n位数乘以m位数结果最多为n+m+2位。
```