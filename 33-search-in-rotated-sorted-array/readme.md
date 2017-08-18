# Search in Rotated Sorted Array

## Description

```
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

 (i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

 You are given a target value to search. If found in the array return its index, otherwise return -1.

 You may assume no duplicate exists in the array.
```

link: https://leetcode.com/problems/search-in-rotated-sorted-array/description/

difficulty: Medium

## Discussion

```
查找目标t和数组a的第一个元素比较a[0]，可知t是属于翻转前点所在下标之前还是之后，
同样根据这一点，在二分查找获取中间数后，也可以判断中间数是在之前还是之后。
如果在同一段中，则按照常规的区间更新法则；否则区间在更新时只需考虑t所在的段即可。
```

favorite: 1
