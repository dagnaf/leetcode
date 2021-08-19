# First Missing Positive

## Description

```
 Given an unsorted integer array, find the first missing positive integer.



 For example,
 Given [1,2,0] return 3,
 and [3,4,-1,1] return 2.



 Your algorithm should run in O(n) time and uses constant space.
```

link: https://leetcode.com/problems/first-missing-positive/description/

difficulty: Hard

## Discussion

```
1、题目中数组是无序的，由于限制复杂度，不能直接排序查找。
2、数组中的数字是原来有序数组的一个置换。
3、在求解缺失数时只需要考虑1, 2, 3, ... , n - 1, n，其他数无用，其中n是数组长度。
4、所以把输入数组中的需要考虑的数字放回原来的位置就可以得到有序数组。
```

favorite: 2
