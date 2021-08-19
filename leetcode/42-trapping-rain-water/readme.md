# Trapping Rain Water

## Description

```
 Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.



 For example,
 Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.




 The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!
```

link: https://leetcode.com/problems/trapping-rain-water/description/

difficulty: Hard

## Discussion

```
1、DP, l[i] 记录包含i这段可以到达的最左端点，r[i]类似。
2、用栈记录高度，遇到比当前高的就弹出计算。
```

favorite: 2