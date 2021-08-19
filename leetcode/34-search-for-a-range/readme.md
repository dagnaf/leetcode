# Search for a Range

## Description

```
Given an array of integers sorted in ascending order, find the starting and ending position of a given target value.

 Your algorithm's runtime complexity must be in the order of O(log n).

 If the target is not found in the array, return [-1, -1].


 For example,
 Given [5, 7, 7, 8, 8, 10] and target value 8,
 return [3, 4].
```

link: https://leetcode.com/problems/search-for-a-range/description/

difficulty: Medium

## Discussion

```
lower_bound(a, a+n, tar) 在数组中找到第一个大于等于>=tar的位置k，即a[k] >= tar，如果不存在返回end
upper_bound(a, a+n, tar)在数组中找到第一个大于tar的位置k，即a[k] > tar，如果不存在返回end
另一种做法求upper_bound就是利用lower_bound查询tar+1
```